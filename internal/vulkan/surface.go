package vulkan

import (
	"log/slog"
	"math"
	"runtime"
	"unsafe"

	"github.com/csnewman/go-gfx/hal"
)

/*
#include <stdlib.h>
#include <vulkan/vulkan.h>
#include <vulkan/vulkan_metal.h>
#include "vulkan.h"
*/
import "C"

type Surface struct {
	graphics      *Graphics
	surface       C.VkSurfaceKHR
	format        C.VkFormat
	colorSpace    C.VkColorSpaceKHR
	minImageCount int
	transform     C.VkSurfaceTransformFlagBitsKHR
	swapchain     C.VkSwapchainKHR
	images        []*SurfaceImage
	entries       []*SurfaceEntry
	currentEntry  int
	width         int
	height        int
}

type SurfaceImage struct {
	image  C.VkImage
	view   C.VkImageView
	width  int
	height int
}

type SurfaceEntry struct {
	commandPool C.VkCommandPool
	imgSem      C.VkSemaphore
	completeSem C.VkSemaphore
	fence       C.VkFence
}

func (g *Graphics) CreateSurface(rawWH hal.WindowHandle) (hal.Surface, error) {
	// TODO: support other handles
	wh, ok := rawWH.(hal.MetalWindowHandle)
	if !ok {
		return nil, hal.ErrUnsupportedWindowHandle
	}

	var surface C.VkSurfaceKHR

	var createInfo C.VkMetalSurfaceCreateInfoEXT
	createInfo.sType = C.VK_STRUCTURE_TYPE_METAL_SURFACE_CREATE_INFO_EXT
	createInfo.pLayer = wh.Layer

	if err := mapError(C.vkCreateMetalSurfaceEXT(g.instance, &createInfo, nil, &surface)); err != nil {
		return nil, err
	}

	var capabilities C.VkSurfaceCapabilitiesKHR

	if err := mapError(C.vkGetPhysicalDeviceSurfaceCapabilitiesKHR(g.physicalDevice, surface, &capabilities)); err != nil {
		return nil, err
	}

	// TODO: min & max width height

	slog.Info("cap", "cap", capabilities)

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
		return nil, hal.ErrIncompatibleSurface
	}

	s := &Surface{
		graphics:      g,
		surface:       surface,
		format:        format.format,
		colorSpace:    format.colorSpace,
		minImageCount: int(capabilities.minImageCount),
		transform:     capabilities.currentTransform,
	}

	if err := s.Resize(int(capabilities.currentExtent.width), int(capabilities.currentExtent.height)); err != nil {
		return nil, err
	}

	for i := 0; i < 3; i++ {
		var commandInfo C.VkCommandPoolCreateInfo

		commandInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
		commandInfo.queueFamilyIndex = C.uint32_t(g.graphicsFamily)
		commandInfo.flags = C.VK_COMMAND_POOL_CREATE_TRANSIENT_BIT

		var commandPool C.VkCommandPool

		if err := mapError(C.vkCreateCommandPool(g.device, &commandInfo, nil, &commandPool)); err != nil {
			return nil, err
		}

		var semInfo C.VkSemaphoreCreateInfo
		semInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO

		var imgSem C.VkSemaphore

		if err := mapError(C.vkCreateSemaphore(g.device, &semInfo, nil, &imgSem)); err != nil {
			return nil, err
		}

		var completeSem C.VkSemaphore

		if err := mapError(C.vkCreateSemaphore(g.device, &semInfo, nil, &completeSem)); err != nil {
			return nil, err
		}

		var fenceInfo C.VkFenceCreateInfo
		fenceInfo.sType = C.VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
		fenceInfo.flags = C.VK_FENCE_CREATE_SIGNALED_BIT

		var fence C.VkFence

		if err := mapError(C.vkCreateFence(g.device, &fenceInfo, nil, &fence)); err != nil {
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

func (s *Surface) Resize(width int, height int) error {
	slog.Info("resize", "width", width, "height", height)

	s.width = width
	s.height = height

	var createInfo C.VkSwapchainCreateInfoKHR
	createInfo.sType = C.VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR
	createInfo.surface = s.surface
	createInfo.minImageCount = C.uint32_t(s.minImageCount)
	createInfo.imageFormat = s.format
	createInfo.imageColorSpace = s.colorSpace
	createInfo.imageExtent.width = C.uint32_t(width)
	createInfo.imageExtent.height = C.uint32_t(height)
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

		s.images = append(s.images, &SurfaceImage{
			image:  image,
			view:   view,
			width:  width,
			height: height,
		})
	}

	return nil
}

func (s *Surface) TextureFormat() hal.TextureFormat {
	// TODO: fix

	return hal.TextureFormatBGRA8UNorm
}

type SurfaceFrame struct {
	graphics *Graphics
	surface  *Surface
	entry    *SurfaceEntry
	img      *SurfaceImage
	index    int
}

func (s *Surface) Acquire() (hal.SurfaceFrame, error) {
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

	return &SurfaceFrame{
		graphics: s.graphics,
		surface:  s,
		entry:    entry,
		img:      s.images[imgIndex],
		index:    int(imgIndex),
	}, nil
}

func (f *SurfaceFrame) Texture() hal.Texture {
	return &Texture{
		img: f.img.image,
	}
}

func (f *SurfaceFrame) View() hal.TextureView {
	return &TextureView{
		view: f.img.view,
	}
}

func (f *SurfaceFrame) Present() error {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	var presentInfo C.VkPresentInfoKHR
	presentInfo.sType = C.VK_STRUCTURE_TYPE_PRESENT_INFO_KHR
	presentInfo.pNext = nil
	presentInfo.swapchainCount = 1

	swapchain := f.surface.swapchain
	presentInfo.pSwapchains = &swapchain
	pinner.Pin(presentInfo.pSwapchains)

	ind := C.uint32_t(f.index)
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

func (f *SurfaceFrame) Discard() {
	//TODO implement me
	panic("implement me")
}
