package vulkan

import (
	"log/slog"
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

	return s, nil
}

func (s *Surface) Resize(width int, height int) error {
	slog.Info("resize", "width", width, "height", height)

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

	return nil
}

func (s *Surface) TextureFormat() hal.TextureFormat {
	// TODO: fix

	return hal.TextureFormatBGRA8UNorm
}

func (s *Surface) AcquireTexture() (hal.SurfaceTexture, error) {
	//TODO implement me
	panic("implement me")
}
