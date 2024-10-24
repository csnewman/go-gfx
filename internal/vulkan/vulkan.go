package vulkan

/*
#cgo pkg-config: vulkan

#include <stdlib.h>
#include <vulkan/vulkan.h>
#include "vulkan.h"

const char* GFX_VK_KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME = VK_KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME;
*/
import "C"

import (
	"github.com/csnewman/go-gfx/hal"
	"runtime"
	"unsafe"
)

type Graphics struct {
}

func NewGraphics() hal.Graphics {
	return &Graphics{}
}

func (g Graphics) Init(cfg hal.GPUConfig) error {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	var instInfo C.VkInstanceCreateInfo
	instInfo.sType = C.VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO

	var appInfo C.VkApplicationInfo

	appInfo.sType = C.VK_STRUCTURE_TYPE_APPLICATION_INFO

	appInfo.pApplicationName = C.CString("TODO")
	defer C.free(unsafe.Pointer(appInfo.pApplicationName))

	appInfo.applicationVersion = C.gfx_vk_version(1, 0, 0)

	appInfo.pEngineName = C.CString("go-gfx")
	defer C.free(unsafe.Pointer(appInfo.pEngineName))

	appInfo.engineVersion = C.gfx_vk_version(1, 0, 0)
	appInfo.apiVersion = C.VK_API_VERSION_1_3

	instInfo.pApplicationInfo = &appInfo
	pinner.Pin(instInfo.pApplicationInfo)

	var exts []*C.char

	exts = append(exts, C.GFX_VK_KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME)
	instInfo.flags |= C.VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT_KHR

	instInfo.enabledExtensionCount = C.uint32_t(len(exts))
	instInfo.ppEnabledExtensionNames = unsafe.SliceData(exts)
	pinner.Pin(instInfo.ppEnabledExtensionNames)

	var inst C.VkInstance

	err := mapError(C.vkCreateInstance(&instInfo, nil, &inst))
	if err != nil {
		return err
	}

	return nil
}

func (g Graphics) CreateSurface(windowHandle hal.WindowHandle) (hal.Surface, error) {
	//TODO implement me
	panic("implement me")
}

func (g Graphics) CreateShader(cfg hal.ShaderConfig) (hal.Shader, error) {
	//TODO implement me
	panic("implement me")
}

func (g Graphics) CreateBuffer(data []byte) hal.Buffer {
	//TODO implement me
	panic("implement me")
}

func (g Graphics) CreateRenderPipeline(des hal.RenderPipelineDescriptor) (hal.RenderPipeline, error) {
	//TODO implement me
	panic("implement me")
}

func (g Graphics) CreateCommandBuffer() hal.CommandBuffer {
	//TODO implement me
	panic("implement me")
}
