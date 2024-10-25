package vulkan

/*
#cgo pkg-config: vulkan

#include <stdlib.h>
#include <vulkan/vulkan.h>
#include "vulkan.h"

const char* GFX_VK_KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME = VK_KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME;
const char* GFX_VK_KHR_SURFACE_EXTENSION_NAME = VK_KHR_SURFACE_EXTENSION_NAME;
const char* GFX_VK_EXT_DEBUG_UTILS_EXTENSION_NAME = VK_EXT_DEBUG_UTILS_EXTENSION_NAME;
const char* GFX_VK_LAYER_KHRONOS_validation = "VK_LAYER_KHRONOS_validation";

VkBool32 gfx_vk_debug_callback(
	VkDebugUtilsMessageSeverityFlagBitsEXT           messageSeverity,
    VkDebugUtilsMessageTypeFlagsEXT                  messageTypes,
    const VkDebugUtilsMessengerCallbackDataEXT*      pCallbackData,
    void*                                            pUserData
);

VkResult gfx_vkCreateDebugUtilsMessengerEXT(
    VkInstance                                  instance,
    const VkDebugUtilsMessengerCreateInfoEXT*   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDebugUtilsMessengerEXT*                   pMessenger
) {
	return GFX_VK_EXT_FUNC(vkCreateDebugUtilsMessengerEXT, instance, pCreateInfo, pAllocator, pMessenger);
}

*/
import "C"

import (
	"github.com/csnewman/go-gfx/hal"
	"log"
	"runtime"
	"unsafe"
)

type Graphics struct {
	instance       C.VkInstance
	debugMessenger C.VkDebugUtilsMessengerEXT
}

func NewGraphics() hal.Graphics {
	return &Graphics{}
}

func (g *Graphics) Init(cfg hal.GPUConfig) error {
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

	exts = append(exts, C.GFX_VK_KHR_SURFACE_EXTENSION_NAME)
	exts = append(exts, C.GFX_VK_EXT_DEBUG_UTILS_EXTENSION_NAME)

	instInfo.enabledExtensionCount = C.uint32_t(len(exts))
	instInfo.ppEnabledExtensionNames = unsafe.SliceData(exts)
	pinner.Pin(instInfo.ppEnabledExtensionNames)

	var layers []*C.char

	layers = append(layers, C.GFX_VK_LAYER_KHRONOS_validation)

	instInfo.enabledLayerCount = C.uint32_t(len(layers))
	instInfo.ppEnabledLayerNames = unsafe.SliceData(layers)
	pinner.Pin(instInfo.ppEnabledLayerNames)

	err := mapError(C.vkCreateInstance(&instInfo, nil, &g.instance))
	if err != nil {
		return err
	}

	var debugInfo C.VkDebugUtilsMessengerCreateInfoEXT
	debugInfo.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT
	debugInfo.messageSeverity = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT | C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT | C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT
	debugInfo.messageType = C.VK_DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT | C.VK_DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT | C.VK_DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT
	debugInfo.pfnUserCallback = C.PFN_vkDebugUtilsMessengerCallbackEXT(C.gfx_vk_debug_callback)

	err = mapError(C.gfx_vkCreateDebugUtilsMessengerEXT(g.instance, &debugInfo, nil, &g.debugMessenger))
	if err != nil {
		return err
	}

	var physicalDeviceCount C.uint32_t
	err = mapError(C.vkEnumeratePhysicalDevices(g.instance, &physicalDeviceCount, nil))
	if err != nil {
		return err
	}

	log.Println("count", physicalDeviceCount)

	return nil
}

func (g *Graphics) CreateSurface(windowHandle hal.WindowHandle) (hal.Surface, error) {
	//TODO implement me
	panic("implement me")
}

func (g *Graphics) CreateShader(cfg hal.ShaderConfig) (hal.Shader, error) {
	//TODO implement me
	panic("implement me")
}

func (g *Graphics) CreateBuffer(data []byte) hal.Buffer {
	//TODO implement me
	panic("implement me")
}

func (g *Graphics) CreateRenderPipeline(des hal.RenderPipelineDescriptor) (hal.RenderPipeline, error) {
	//TODO implement me
	panic("implement me")
}

func (g *Graphics) CreateCommandBuffer() hal.CommandBuffer {
	//TODO implement me
	panic("implement me")
}
