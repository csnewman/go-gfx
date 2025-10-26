package vulkan

import (
	"fmt"
	"log/slog"
	"math"
	"unsafe"

	"github.com/csnewman/go-gfx/ffi"
	"github.com/csnewman/go-gfx/gfx"
	"github.com/csnewman/go-gfx/vk"
)

type Surface struct {
	graphics      *Graphics
	logger        *slog.Logger
	surface       vk.SurfaceKHR
	format        vk.Format
	colorSpace    vk.ColorSpaceKHR
	minImageCount int
	transform     vk.SurfaceTransformFlagsKHR
	swapchain     vk.SwapchainKHR
	images        []*Image
	entries       []*SurfaceEntry
	frameCount    int
	currentEntry  int
	width         int
	height        int
}

type SurfaceEntry struct {
	commandPool vk.CommandPool
	buffers     []vk.CommandBuffer
	bufferPos   int
	imgSem      vk.Semaphore
	completeSem vk.Semaphore
	fence       vk.Fence
}

func (g *Graphics) CreateSurface(handle gfx.SurfaceHandle, size gfx.PhysicalSize) (gfx.Surface, error) {
	arena := ffi.NewArena()
	defer arena.Close()

	var surface vk.SurfaceKHR

	switch handle.SurfaceHandleType() {
	case gfx.VulkanSurfaceHandleType:
		sh := handle.(gfx.VulkanSurfaceHandle)

		rawSurf, err := sh.CreateVkSurface(unsafe.Pointer(g.instance))
		if err != nil {
			return nil, err
		}

		surface = vk.SurfaceKHR(uintptr(rawSurf))
	case gfx.MetalSurfaceHandleType:
		sh := handle.(gfx.MetalSurfaceHandle)

		createInfo := vk.MetalSurfaceCreateInfoEXTAlloc(arena, 1)
		createInfo.SetSType(vk.STRUCTURE_TYPE_METAL_SURFACE_CREATE_INFO_EXT)
		createInfo.SetPLayer(sh.MetalLayer())

		surfaceRef := ffi.RefAlloc[vk.SurfaceKHR](arena, 1)

		if err := mapError(vk.CreateMetalSurfaceEXT(g.instance, createInfo, vk.AllocationCallbacksNil, surfaceRef)); err != nil {
			return nil, err
		}

		surface = surfaceRef.Get()
	case gfx.Win32SurfaceHandleType:
		sh := handle.(gfx.Win32SurfaceHandle)

		createInfo := vk.Win32SurfaceCreateInfoKHRAlloc(arena, 1)
		createInfo.SetSType(vk.STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR)
		createInfo.SetHinstance(sh.Win32Instance())
		createInfo.SetHwnd(sh.Win32Handle())

		surfaceRef := ffi.RefAlloc[vk.SurfaceKHR](arena, 1)

		if err := mapError(vk.CreateWin32SurfaceKHR(g.instance, createInfo, vk.AllocationCallbacksNil, surfaceRef)); err != nil {
			return nil, err
		}

		surface = surfaceRef.Get()
	default:
		return nil, fmt.Errorf("%w: %v", gfx.ErrUnsupportedSurfaceHandle, handle.SurfaceHandleType())
	}

	capabilities := vk.SurfaceCapabilitiesKHRAlloc(arena, 1)

	if err := mapError(vk.GetPhysicalDeviceSurfaceCapabilitiesKHR(g.physicalDevice, surface, capabilities)); err != nil {
		return nil, err
	}

	// TODO: min & max width height

	g.logger.Debug("Surface capabilities", "capabilities", capabilities.Raw())

	formatCountRef := ffi.RefAlloc[uint32](arena, 1)

	if err := mapError(vk.GetPhysicalDeviceSurfaceFormatsKHR(g.physicalDevice, surface, formatCountRef, vk.SurfaceFormatKHRNil)); err != nil {
		return nil, err
	}

	formats := vk.SurfaceFormatKHRAlloc(arena, int(formatCountRef.Get()))

	if err := mapError(vk.GetPhysicalDeviceSurfaceFormatsKHR(g.physicalDevice, surface, formatCountRef, formats)); err != nil {
		return nil, err
	}

	var format vk.SurfaceFormatKHR
	foundFormat := false

	for i := 0; i < int(formatCountRef.Get()); i++ {
		entry := formats.Offset(i)

		if entry.GetFormat() == vk.FORMAT_B8G8R8A8_UNORM {
			foundFormat = true
			format = entry

			break
		}
	}

	if !foundFormat {
		return nil, gfx.ErrIncompatibleSurface
	}

	s := &Surface{
		logger:        g.logger,
		graphics:      g,
		surface:       surface,
		format:        format.GetFormat(),
		colorSpace:    format.GetColorSpace(),
		minImageCount: int(capabilities.GetMinImageCount()),
		transform:     capabilities.GetCurrentTransform(),
		frameCount:    3,
	}

	if size.Width == 0 && size.Height == 0 {
		currentExtent := capabilities.RefCurrentExtent()

		size.Width = int(currentExtent.GetWidth())
		size.Height = int(currentExtent.GetHeight())
	}

	if err := s.Resize(size); err != nil {
		return nil, err
	}

	for i := 0; i < s.frameCount; i++ {
		commandInfo := vk.CommandPoolCreateInfoAlloc(arena, 1)
		commandInfo.SetSType(vk.STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO)
		commandInfo.SetQueueFamilyIndex(uint32(s.graphics.graphicsFamily))
		commandInfo.SetFlags(vk.COMMAND_POOL_CREATE_TRANSIENT_BIT)

		poolRef := ffi.RefAlloc[vk.CommandPool](arena, 1)

		if err := mapError(vk.CreateCommandPool(s.graphics.device, commandInfo, vk.AllocationCallbacksNil, poolRef)); err != nil {
			return nil, err
		}

		semInfo := vk.SemaphoreCreateInfoAlloc(arena, 1)
		semInfo.SetSType(vk.STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO)

		imgSemRef := ffi.RefAlloc[vk.Semaphore](arena, 1)

		if err := mapError(vk.CreateSemaphore(s.graphics.device, semInfo, vk.AllocationCallbacksNil, imgSemRef)); err != nil {
			return nil, err
		}

		completeSemRef := ffi.RefAlloc[vk.Semaphore](arena, 1)

		if err := mapError(vk.CreateSemaphore(s.graphics.device, semInfo, vk.AllocationCallbacksNil, completeSemRef)); err != nil {
			return nil, err
		}

		fenceInfo := vk.FenceCreateInfoAlloc(arena, 1)
		fenceInfo.SetSType(vk.STRUCTURE_TYPE_FENCE_CREATE_INFO)
		fenceInfo.SetFlags(vk.FENCE_CREATE_SIGNALED_BIT)

		fenceRef := ffi.RefAlloc[vk.Fence](arena, 1)

		if err := mapError(vk.CreateFence(s.graphics.device, fenceInfo, vk.AllocationCallbacksNil, fenceRef)); err != nil {
			return nil, err
		}

		s.entries = append(s.entries, &SurfaceEntry{
			commandPool: poolRef.Get(),
			imgSem:      imgSemRef.Get(),
			completeSem: completeSemRef.Get(),
			fence:       fenceRef.Get(),
		})
	}

	return s, nil
}

func (s *Surface) destroy() {
	for _, entry := range s.entries {
		vk.DestroyFence(s.graphics.device, entry.fence, vk.AllocationCallbacksNil)
		vk.DestroySemaphore(s.graphics.device, entry.imgSem, vk.AllocationCallbacksNil)
		vk.DestroySemaphore(s.graphics.device, entry.completeSem, vk.AllocationCallbacksNil)
		vk.DestroyCommandPool(s.graphics.device, entry.commandPool, vk.AllocationCallbacksNil)
	}

	s.entries = nil

	s.destroySwapchain()
}

func (s *Surface) createSwapchain(arena *ffi.Arena) error {
	s.logger.Debug("Creating swapchain")

	createInfo := vk.SwapchainCreateInfoKHRAlloc(arena, 1)
	createInfo.SetSType(vk.STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR)
	createInfo.SetSurface(s.surface)
	createInfo.SetMinImageCount(uint32(s.minImageCount))
	createInfo.SetImageFormat(s.format)
	createInfo.SetImageColorSpace(s.colorSpace)

	extent := createInfo.RefImageExtent()
	extent.SetWidth(uint32(s.width))
	extent.SetHeight(uint32(s.height))

	createInfo.SetImageArrayLayers(1)
	createInfo.SetImageUsage(vk.IMAGE_USAGE_COLOR_ATTACHMENT_BIT)
	createInfo.SetImageSharingMode(vk.SHARING_MODE_EXCLUSIVE)
	createInfo.SetPreTransform(s.transform)
	createInfo.SetCompositeAlpha(vk.COMPOSITE_ALPHA_OPAQUE_BIT_KHR)
	createInfo.SetPresentMode(vk.PRESENT_MODE_FIFO_KHR)
	createInfo.SetClipped(true)

	swapchainRef := ffi.RefAlloc[vk.SwapchainKHR](arena, 1)

	if err := mapError(vk.CreateSwapchainKHR(s.graphics.device, createInfo, vk.AllocationCallbacksNil, swapchainRef)); err != nil {
		return err
	}

	s.swapchain = swapchainRef.Get()

	imageCountRef := ffi.RefAlloc[uint32](arena, 1)

	if err := mapError(vk.GetSwapchainImagesKHR(s.graphics.device, s.swapchain, imageCountRef, ffi.RefNil[vk.Image]())); err != nil {
		return err
	}

	images := ffi.RefAlloc[vk.Image](arena, int(imageCountRef.Get()))

	if err := mapError(vk.GetSwapchainImagesKHR(s.graphics.device, s.swapchain, imageCountRef, images)); err != nil {
		return err
	}

	for i := 0; i < int(imageCountRef.Get()); i++ {
		image := images.Offset(i).Get()

		createInfo := vk.ImageViewCreateInfoAlloc(arena, 1)
		createInfo.SetSType(vk.STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO)
		createInfo.SetViewType(vk.IMAGE_VIEW_TYPE_2D)

		components := createInfo.RefComponents()
		components.SetR(vk.COMPONENT_SWIZZLE_IDENTITY)
		components.SetG(vk.COMPONENT_SWIZZLE_IDENTITY)
		components.SetB(vk.COMPONENT_SWIZZLE_IDENTITY)
		components.SetA(vk.COMPONENT_SWIZZLE_IDENTITY)

		subresourceRange := createInfo.RefSubresourceRange()
		subresourceRange.SetBaseMipLevel(0)
		subresourceRange.SetLevelCount(vk.REMAINING_MIP_LEVELS)
		subresourceRange.SetBaseArrayLayer(0)
		subresourceRange.SetLayerCount(vk.REMAINING_ARRAY_LAYERS)
		subresourceRange.SetAspectMask(vk.IMAGE_ASPECT_COLOR_BIT)

		createInfo.SetImage(image)
		createInfo.SetFormat(s.format)

		viewRef := ffi.RefAlloc[vk.ImageView](arena, 1)

		if err := mapError(vk.CreateImageView(s.graphics.device, createInfo, vk.AllocationCallbacksNil, viewRef)); err != nil {
			return err
		}

		s.images = append(s.images, &Image{
			image:  image,
			view:   viewRef.Get(),
			width:  s.width,
			height: s.height,
		})
	}

	return nil
}

func (s *Surface) destroySwapchain() {
	s.logger.Debug("Destroying swapchain", "dev", s.graphics.device, "swapchain", s.swapchain)
	
	for _, image := range s.images {
		vk.DestroyImageView(s.graphics.device, image.view, vk.AllocationCallbacksNil)
	}

	s.images = nil

	if s.swapchain != vk.SwapchainKHRNil {
		vk.DestroySwapchainKHR(s.graphics.device, s.swapchain, vk.AllocationCallbacksNil)
	}
}

func (s *Surface) Resize(size gfx.PhysicalSize) error {
	s.logger.Debug("Surface resize", "width", size.Width, "height", size.Height)

	arena := ffi.NewArena()
	defer arena.Close()

	s.width = size.Width
	s.height = size.Height

	capabilities := vk.SurfaceCapabilitiesKHRAlloc(arena, 1)

	if err := mapError(vk.GetPhysicalDeviceSurfaceCapabilitiesKHR(s.graphics.physicalDevice, s.surface, capabilities)); err != nil {
		return err
	}

	// TODO: min & max width height

	s.logger.Debug("Surface capabilities", "capabilities", capabilities.Raw())

	if err := mapError(vk.DeviceWaitIdle(s.graphics.device)); err != nil {
		return err
	}

	s.destroySwapchain()

	return s.createSwapchain(arena)
}

func (s *Surface) Format() gfx.Format {
	// TODO: fix

	return gfx.FormatBGRA8UNorm
}

func (s *Surface) FrameCount() int {
	return s.frameCount
}

type SurfaceFrame struct {
	graphics *Graphics
	surface  *Surface
	entry    *SurfaceEntry
	img      *Image
	imgIndex int
	index    int
	buffer   *CommandBuffer
}

func (s *Surface) Acquire() (gfx.SurfaceFrame, error) {
	arena := ffi.NewArena()
	defer arena.Close()

	entry := s.entries[s.currentEntry]
	fenceRef := ffi.RefFromValues(arena, entry.fence)

	if err := mapError(vk.WaitForFences(
		s.graphics.device,
		1,
		fenceRef,
		true,
		math.MaxUint64,
	)); err != nil {
		return nil, err
	}

	if err := mapError(vk.ResetFences(s.graphics.device, 1, fenceRef)); err != nil {
		return nil, err
	}

	imgIndexRef := ffi.RefAlloc[uint32](arena, 1)

	// TODO: handle outdated & suboptimal
	if err := mapError(vk.AcquireNextImageKHR(
		s.graphics.device,
		s.swapchain,
		math.MaxUint64,
		entry.imgSem,
		vk.FenceNil,
		imgIndexRef,
	)); err != nil {
		return nil, err
	}

	imgIndex := imgIndexRef.Get()

	s.currentEntry = (s.currentEntry + 1) % len(s.entries)

	sf := &SurfaceFrame{
		graphics: s.graphics,
		surface:  s,
		entry:    entry,
		img:      s.images[imgIndex],
		imgIndex: int(imgIndex),
		index:    s.currentEntry,
	}

	if err := mapError(vk.ResetCommandPool(s.graphics.device, sf.entry.commandPool, 0)); err != nil {
		return nil, err
	}

	sf.entry.bufferPos = 0

	sf.buffer = sf.CreateCommandBuffer(arena)

	sf.buffer.Barrier(arena, Barrier{
		Images: []ImageBarrier{
			{
				Image:     sf.img,
				SrcLayout: ImageLayoutUndefined,
				DstLayout: ImageLayoutAttachment,
			},
		},
	})

	return sf, nil
}

func (f *SurfaceFrame) Index() int {
	return f.index
}

func (f *SurfaceFrame) View() gfx.ImageView {
	return f.img.DefaultView()
}

func (f *SurfaceFrame) Image() *Image {
	return f.img
}

func (f *SurfaceFrame) BeginRenderPass(descriptor gfx.RenderPassDescriptor) gfx.RenderPassEncoder {
	f.buffer.BeginRenderPass(descriptor)

	return &RenderPassEncoder{
		buffer: f.buffer,
	}
}

func (f *SurfaceFrame) Present() error {
	arena := ffi.NewArena()
	defer arena.Close()

	f.buffer.Barrier(arena, Barrier{
		Images: []ImageBarrier{
			{
				Image:     f.img,
				SrcLayout: ImageLayoutAttachment,
				DstLayout: ImageLayoutPresent,
			},
		},
	})

	if err := f.buffer.SubmitFrame(arena, f); err != nil {
		return err
	}

	presentInfo := vk.PresentInfoKHRAlloc(arena, 1)
	presentInfo.SetSType(vk.STRUCTURE_TYPE_PRESENT_INFO_KHR)
	presentInfo.SetSwapchainCount(1)

	presentInfo.SetPSwapchains(ffi.RefFromValues(arena, f.surface.swapchain))
	presentInfo.SetPImageIndices(ffi.RefFromValues[uint32](arena, uint32(f.imgIndex)))

	presentInfo.SetPWaitSemaphores(ffi.RefFromValues(arena, f.entry.completeSem))
	presentInfo.SetWaitSemaphoreCount(1)

	if err := mapError(vk.QueuePresentKHR(f.graphics.graphicsQueue, presentInfo)); err != nil {
		return err
	}

	return nil
}

//
//func (f *SurfaceFrame) Discard() {
//	//TODO implement me
//	panic("implement me")
//}
