package vulkan

import "github.com/csnewman/go-gfx/hal"

/*
#include <stdlib.h>
#include <vulkan/vulkan.h>
#include <vulkan/vulkan_metal.h>
#include "vulkan.h"
*/
import "C"

type Surface struct {
	surface C.VkSurfaceKHR
}

func (g *Graphics) CreateSurface(rawWH hal.WindowHandle) (hal.Surface, error) {
	// TODO: support other handles
	wh, ok := rawWH.(hal.MetalWindowHandle)
	if !ok {
		return nil, hal.ErrUnsupportedWindowHandle
	}

	var surface C.VkSurfaceKHR

	_ = wh

	var createInfo C.VkMetalSurfaceCreateInfoEXT
	createInfo.sType = C.VK_STRUCTURE_TYPE_METAL_SURFACE_CREATE_INFO_EXT
	createInfo.pLayer = wh.Layer

	if err := mapError(C.vkCreateMetalSurfaceEXT(g.instance, &createInfo, nil, &surface)); err != nil {
		return nil, err
	}

	return &Surface{
		surface: surface,
	}, nil
}

func (s Surface) TextureFormat() hal.TextureFormat {
	//TODO implement me
	panic("implement me")
}

func (s Surface) AcquireTexture() (hal.SurfaceTexture, error) {
	//TODO implement me
	panic("implement me")
}
