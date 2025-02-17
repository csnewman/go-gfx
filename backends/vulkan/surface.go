package vulkan

import (
	"fmt"
	"github.com/csnewman/go-gfx/gfx"
	"log/slog"
	"math"
	"runtime"
	"unsafe"
)

/*
#include "vulkan.h"
*/
import "C"

type Surface struct {
	graphics      *Graphics
	logger        *slog.Logger
	surface       C.VkSurfaceKHR
	format        C.VkFormat
	colorSpace    C.VkColorSpaceKHR
	minImageCount int
	transform     C.VkSurfaceTransformFlagBitsKHR
	swapchain     C.VkSwapchainKHR
	images        []*Image
	entries       []*SurfaceEntry
	frameCount    int
	currentEntry  int
	width         int
	height        int
}

type SurfaceEntry struct {
	commandPool C.VkCommandPool
	buffers     []C.VkCommandBuffer
	bufferPos   int
	imgSem      C.VkSemaphore
	completeSem C.VkSemaphore
	fence       C.VkFence
}

func (g *Graphics) CreateSurface(handle gfx.SurfaceHandle, size gfx.PhysicalSize) (gfx.Surface, error) {
	var surface C.VkSurfaceKHR

	switch handle.SurfaceHandleType() {
	case gfx.VulkanSurfaceHandleType:
		sh := handle.(gfx.VulkanSurfaceHandle)

		rawSurf, err := sh.CreateVkSurface(unsafe.Pointer(g.instance))
		if err != nil {
			return nil, err
		}

		surface = C.VkSurfaceKHR(rawSurf)
	case gfx.MetalSurfaceHandleType:
		sh := handle.(gfx.MetalSurfaceHandle)

		var createInfo C.VkMetalSurfaceCreateInfoEXT
		createInfo.sType = C.VK_STRUCTURE_TYPE_METAL_SURFACE_CREATE_INFO_EXT
		createInfo.pLayer = sh.MetalLayer()

		if err := mapError(C.vkCreateMetalSurfaceEXT(g.instance, &createInfo, nil, &surface)); err != nil {
			return nil, err
		}
	case gfx.Win32SurfaceHandleType:
		sh := handle.(gfx.Win32SurfaceHandle)

		var createInfo C.VkWin32SurfaceCreateInfoKHR
		createInfo.sType = C.VK_STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR
		createInfo.hinstance = C.HINSTANCE(sh.Win32Instance())
		createInfo.hwnd = C.HWND(sh.Win32Handle())

		if err := mapError(C.vkCreateWin32SurfaceKHR(g.instance, &createInfo, nil, &surface)); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("%w: %v", gfx.ErrUnsupportedSurfaceHandle, handle.SurfaceHandleType())
	}

	var capabilities C.VkSurfaceCapabilitiesKHR

	if err := mapError(C.vkGetPhysicalDeviceSurfaceCapabilitiesKHR(g.physicalDevice, surface, &capabilities)); err != nil {
		return nil, err
	}

	// TODO: min & max width height

	g.logger.Debug("Surface capabilities", "capabilities", capabilities)

	var formatCount C.uint32_t

	if err := mapError(C.vkGetPhysicalDeviceSurfaceFormatsKHR(g.physicalDevice, surface, &formatCount, nil)); err != nil {
		return nil, err
	}

	formats := make([]C.VkSurfaceFormatKHR, formatCount)

	if err := mapError(C.vkGetPhysicalDeviceSurfaceFormatsKHR(g.physicalDevice, surface, &formatCount, unsafe.SliceData(formats))); err != nil {
		return nil, err
	}

	formats = formats[:formatCount]

	var format C.VkSurfaceFormatKHR
	foundFormat := false

	for _, fmt := range formats {
		if fmt.format == C.VK_FORMAT_B8G8R8A8_UNORM {
			foundFormat = true
			format = fmt

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
		format:        format.format,
		colorSpace:    format.colorSpace,
		minImageCount: int(capabilities.minImageCount),
		transform:     capabilities.currentTransform,
		frameCount:    3,
	}

	if size.Width == 0 && size.Height == 0 {
		size.Width = int(capabilities.currentExtent.width)
		size.Height = int(capabilities.currentExtent.height)
	}

	if err := s.Resize(size); err != nil {
		return nil, err
	}

	for i := 0; i < s.frameCount; i++ {
		var commandInfo C.VkCommandPoolCreateInfo

		commandInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
		commandInfo.queueFamilyIndex = C.uint32_t(s.graphics.graphicsFamily)
		commandInfo.flags = C.VK_COMMAND_POOL_CREATE_TRANSIENT_BIT

		var commandPool C.VkCommandPool

		if err := mapError(C.vkCreateCommandPool(s.graphics.device, &commandInfo, nil, &commandPool)); err != nil {
			return nil, err
		}

		var semInfo C.VkSemaphoreCreateInfo
		semInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO

		var imgSem C.VkSemaphore

		if err := mapError(C.vkCreateSemaphore(s.graphics.device, &semInfo, nil, &imgSem)); err != nil {
			return nil, err
		}

		var completeSem C.VkSemaphore

		if err := mapError(C.vkCreateSemaphore(s.graphics.device, &semInfo, nil, &completeSem)); err != nil {
			return nil, err
		}

		var fenceInfo C.VkFenceCreateInfo
		fenceInfo.sType = C.VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
		fenceInfo.flags = C.VK_FENCE_CREATE_SIGNALED_BIT

		var fence C.VkFence

		if err := mapError(C.vkCreateFence(s.graphics.device, &fenceInfo, nil, &fence)); err != nil {
			return nil, err
		}

		s.entries = append(s.entries, &SurfaceEntry{
			commandPool: commandPool,
			imgSem:      imgSem,
			completeSem: completeSem,
			fence:       fence,
		})
	}

	return s, nil
}

func (s *Surface) destroy() {
	for _, entry := range s.entries {
		C.vkDestroyFence(s.graphics.device, entry.fence, nil)
		C.vkDestroySemaphore(s.graphics.device, entry.imgSem, nil)
		C.vkDestroySemaphore(s.graphics.device, entry.completeSem, nil)
		C.vkDestroyCommandPool(s.graphics.device, entry.commandPool, nil)
	}

	s.entries = nil

	s.destroySwapchain()
}

func (s *Surface) createSwapchain() error {
	var createInfo C.VkSwapchainCreateInfoKHR
	createInfo.sType = C.VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR
	createInfo.surface = s.surface
	createInfo.minImageCount = C.uint32_t(s.minImageCount)
	createInfo.imageFormat = s.format
	createInfo.imageColorSpace = s.colorSpace
	createInfo.imageExtent.width = C.uint32_t(s.width)
	createInfo.imageExtent.height = C.uint32_t(s.height)
	createInfo.imageArrayLayers = 1
	createInfo.imageUsage = C.VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
	createInfo.imageSharingMode = C.VK_SHARING_MODE_EXCLUSIVE
	createInfo.preTransform = s.transform
	createInfo.compositeAlpha = C.VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR
	createInfo.presentMode = C.VK_PRESENT_MODE_FIFO_KHR
	createInfo.clipped = C.VkBool32(1)

	if err := mapError(C.vkCreateSwapchainKHR(s.graphics.device, &createInfo, nil, &s.swapchain)); err != nil {
		return err
	}

	var imageCount C.uint32_t

	if err := mapError(C.vkGetSwapchainImagesKHR(s.graphics.device, s.swapchain, &imageCount, nil)); err != nil {
		return err
	}

	images := make([]C.VkImage, imageCount)

	if err := mapError(C.vkGetSwapchainImagesKHR(s.graphics.device, s.swapchain, &imageCount, unsafe.SliceData(images))); err != nil {
		return err
	}

	images = images[:imageCount]

	for _, image := range images {
		var createInfo C.VkImageViewCreateInfo
		createInfo.sType = C.VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
		createInfo.viewType = C.VK_IMAGE_VIEW_TYPE_2D
		createInfo.components.r = C.VK_COMPONENT_SWIZZLE_IDENTITY
		createInfo.components.g = C.VK_COMPONENT_SWIZZLE_IDENTITY
		createInfo.components.b = C.VK_COMPONENT_SWIZZLE_IDENTITY
		createInfo.components.a = C.VK_COMPONENT_SWIZZLE_IDENTITY
		createInfo.subresourceRange.baseMipLevel = 0
		createInfo.subresourceRange.levelCount = C.VK_REMAINING_MIP_LEVELS
		createInfo.subresourceRange.baseArrayLayer = 0
		createInfo.subresourceRange.layerCount = C.VK_REMAINING_ARRAY_LAYERS
		createInfo.image = image
		createInfo.format = s.format
		createInfo.subresourceRange.aspectMask = C.VK_IMAGE_ASPECT_COLOR_BIT

		var view C.VkImageView

		if err := mapError(C.vkCreateImageView(s.graphics.device, &createInfo, nil, &view)); err != nil {
			return err
		}

		s.images = append(s.images, &Image{
			image:  image,
			view:   view,
			width:  s.width,
			height: s.height,
		})
	}

	return nil
}

func (s *Surface) destroySwapchain() {
	for _, image := range s.images {
		C.vkDestroyImageView(s.graphics.device, image.view, nil)
	}

	s.images = nil

	C.vkDestroySwapchainKHR(s.graphics.device, s.swapchain, nil)
}

func (s *Surface) Resize(size gfx.PhysicalSize) error {
	s.logger.Debug("Surface resize", "width", size.Width, "height", size.Height)

	s.width = size.Width
	s.height = size.Height

	var capabilities C.VkSurfaceCapabilitiesKHR

	if err := mapError(C.vkGetPhysicalDeviceSurfaceCapabilitiesKHR(s.graphics.physicalDevice, s.surface, &capabilities)); err != nil {
		return err
	}

	// TODO: min & max width height

	s.logger.Debug("Surface capabilities", "capabilities", capabilities)

	if err := mapError(C.vkDeviceWaitIdle(s.graphics.device)); err != nil {
		return err
	}

	s.destroySwapchain()

	return s.createSwapchain()
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
	entry := s.entries[s.currentEntry]

	if err := mapError(C.vkWaitForFences(
		s.graphics.device,
		1,
		&entry.fence,
		C.VkBool32(1),
		C.uint64_t(math.MaxUint64),
	)); err != nil {
		return nil, err
	}

	if err := mapError(C.vkResetFences(s.graphics.device, 1, &entry.fence)); err != nil {
		return nil, err
	}

	var imgIndex C.uint32_t

	// TODO: handle outdated & suboptimal
	if err := mapError(C.vkAcquireNextImageKHR(
		s.graphics.device,
		s.swapchain,
		C.uint64_t(math.MaxUint64),
		entry.imgSem,
		nil,
		&imgIndex,
	)); err != nil {
		return nil, err
	}

	s.currentEntry = (s.currentEntry + 1) % len(s.entries)

	sf := &SurfaceFrame{
		graphics: s.graphics,
		surface:  s,
		entry:    entry,
		img:      s.images[imgIndex],
		imgIndex: int(imgIndex),
		index:    s.currentEntry,
	}

	if err := mapError(C.vkResetCommandPool(s.graphics.device, sf.entry.commandPool, 0)); err != nil {
		return nil, err
	}

	sf.entry.bufferPos = 0

	sf.buffer = sf.CreateCommandBuffer()

	sf.buffer.Barrier(Barrier{
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
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	f.buffer.Barrier(Barrier{
		Images: []ImageBarrier{
			{
				Image:     f.img,
				SrcLayout: ImageLayoutAttachment,
				DstLayout: ImageLayoutPresent,
			},
		},
	})

	if err := f.buffer.SubmitFrame(f); err != nil {
		return err
	}

	var presentInfo C.VkPresentInfoKHR
	presentInfo.sType = C.VK_STRUCTURE_TYPE_PRESENT_INFO_KHR
	presentInfo.pNext = nil
	presentInfo.swapchainCount = 1

	swapchain := f.surface.swapchain
	presentInfo.pSwapchains = &swapchain
	pinner.Pin(presentInfo.pSwapchains)

	ind := C.uint32_t(f.imgIndex)
	presentInfo.pImageIndices = &ind
	pinner.Pin(presentInfo.pImageIndices)

	complete := f.entry.completeSem
	presentInfo.pWaitSemaphores = &complete
	pinner.Pin(presentInfo.pWaitSemaphores)

	presentInfo.waitSemaphoreCount = 1

	if err := mapError(C.vkQueuePresentKHR(f.graphics.graphicsQueue, &presentInfo)); err != nil {
		return err
	}

	return nil
}

//
//func (f *SurfaceFrame) Discard() {
//	//TODO implement me
//	panic("implement me")
//}
