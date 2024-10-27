package vulkan

/*
#cgo pkg-config: vulkan

#include <stdlib.h>
#include <vulkan/vulkan.h>
#include <vulkan/vulkan_metal.h>
#include "vulkan.h"

const char* GFX_VK_KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME = VK_KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME;
const char* GFX_VK_KHR_SURFACE_EXTENSION_NAME = VK_KHR_SURFACE_EXTENSION_NAME;
const char* GFX_VK_EXT_DEBUG_UTILS_EXTENSION_NAME = VK_EXT_DEBUG_UTILS_EXTENSION_NAME;
const char* GFX_VK_LAYER_KHRONOS_validation = "VK_LAYER_KHRONOS_validation";
const char* GFX_VK_KHR_portability_subset = "VK_KHR_portability_subset";
const char* GFX_VK_KHR_DYNAMIC_RENDERING_EXTENSION_NAME = VK_KHR_DYNAMIC_RENDERING_EXTENSION_NAME;
const char* GFX_VK_EXT_METAL_SURFACE_EXTENSION_NAME = VK_EXT_METAL_SURFACE_EXTENSION_NAME;

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
	"runtime"
	"unsafe"
)

type Graphics struct {
	instance       C.VkInstance
	debugMessenger C.VkDebugUtilsMessengerEXT
	device         C.VkDevice
	graphicsQueue  C.VkQueue
}

func NewGraphics() hal.Graphics {
	return &Graphics{}
}

func (g *Graphics) Init(cfg hal.GPUConfig) error {
	if err := g.createInstance(); err != nil {
		return err
	}

	device, err := g.selectDevice()
	if err != nil {
		return err
	}

	if err := g.createDevice(device); err != nil {
		return err
	}

	return nil
}

func (g *Graphics) createInstance() error {
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
	exts = append(exts, C.GFX_VK_EXT_METAL_SURFACE_EXTENSION_NAME)

	exts = append(exts, C.GFX_VK_EXT_DEBUG_UTILS_EXTENSION_NAME)

	instInfo.enabledExtensionCount = C.uint32_t(len(exts))
	instInfo.ppEnabledExtensionNames = unsafe.SliceData(exts)
	pinner.Pin(instInfo.ppEnabledExtensionNames)

	var layers []*C.char

	layers = append(layers, C.GFX_VK_LAYER_KHRONOS_validation)

	instInfo.enabledLayerCount = C.uint32_t(len(layers))
	instInfo.ppEnabledLayerNames = unsafe.SliceData(layers)
	pinner.Pin(instInfo.ppEnabledLayerNames)

	if err := mapError(C.vkCreateInstance(&instInfo, nil, &g.instance)); err != nil {
		return err
	}

	var debugInfo C.VkDebugUtilsMessengerCreateInfoEXT
	debugInfo.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT
	debugInfo.messageSeverity = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT | C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT | C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT
	debugInfo.messageType = C.VK_DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT | C.VK_DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT | C.VK_DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT
	debugInfo.pfnUserCallback = C.PFN_vkDebugUtilsMessengerCallbackEXT(C.gfx_vk_debug_callback)

	if err := mapError(C.gfx_vkCreateDebugUtilsMessengerEXT(g.instance, &debugInfo, nil, &g.debugMessenger)); err != nil {
		return err
	}

	return nil
}

type selectedDevice struct {
	device         C.VkPhysicalDevice
	graphicsFamily int
}

func (g *Graphics) selectDevice() (*selectedDevice, error) {
	var physicalDeviceCount C.uint32_t

	if err := mapError(C.vkEnumeratePhysicalDevices(g.instance, &physicalDeviceCount, nil)); err != nil {
		return nil, err
	}

	physicalDevices := make([]C.VkPhysicalDevice, physicalDeviceCount)

	if err := mapError(C.vkEnumeratePhysicalDevices(
		g.instance,
		&physicalDeviceCount,
		unsafe.SliceData(physicalDevices)),
	); err != nil {
		return nil, err
	}

	physicalDevices = physicalDevices[:physicalDeviceCount]

	currentScore := -1

	var bestDevice *selectedDevice

	for _, device := range physicalDevices {
		var props C.VkPhysicalDeviceProperties
		C.vkGetPhysicalDeviceProperties(device, &props)

		// TODO: check if device can present

		score := 0

		switch props.deviceType {
		case C.VK_PHYSICAL_DEVICE_TYPE_OTHER:
			score = 1
		case C.VK_PHYSICAL_DEVICE_TYPE_CPU:
			score = 2
		case C.VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU:
			score = 3
		case C.VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU:
			score = 4
		case C.VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU:
			score = 5
		default:
			continue
		}

		var queueFamilyCount C.uint32_t
		C.vkGetPhysicalDeviceQueueFamilyProperties(device, &queueFamilyCount, nil)

		queueFamilies := make([]C.VkQueueFamilyProperties, queueFamilyCount)

		C.vkGetPhysicalDeviceQueueFamilyProperties(device, &queueFamilyCount, unsafe.SliceData(queueFamilies))

		queueFamilies = queueFamilies[:queueFamilyCount]

		graphicsQueue := -1

		for i, family := range queueFamilies {
			if family.queueFlags&C.VK_QUEUE_GRAPHICS_BIT != 0 {
				graphicsQueue = i

				break
			}
		}

		if graphicsQueue == -1 {
			continue
		}

		// TODO: other scoring tie breakers

		if score > currentScore {
			currentScore = score
			bestDevice = &selectedDevice{
				device:         device,
				graphicsFamily: graphicsQueue,
			}
		}
	}

	if currentScore < 0 {
		return nil, hal.ErrNoSuitableDevice
	}

	return bestDevice, nil
}

func (g *Graphics) createDevice(sel *selectedDevice) error {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	priority := C.float(1.0)

	var queueCreateInfo C.VkDeviceQueueCreateInfo
	queueCreateInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
	queueCreateInfo.queueFamilyIndex = C.uint32_t(sel.graphicsFamily)
	queueCreateInfo.queueCount = 1
	queueCreateInfo.pQueuePriorities = &priority
	pinner.Pin(queueCreateInfo.pQueuePriorities)

	var dynamicRenderingFeatures C.VkPhysicalDeviceDynamicRenderingFeatures
	dynamicRenderingFeatures.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_FEATURES
	dynamicRenderingFeatures.dynamicRendering = C.VkBool32(1)

	var createInfo C.VkDeviceCreateInfo
	createInfo.pNext = unsafe.Pointer(&dynamicRenderingFeatures)
	pinner.Pin(createInfo.pNext)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
	createInfo.queueCreateInfoCount = 1
	createInfo.pQueueCreateInfos = &queueCreateInfo
	pinner.Pin(createInfo.pQueueCreateInfos)

	var exts []*C.char

	exts = append(exts, C.GFX_VK_KHR_portability_subset)
	exts = append(exts, C.GFX_VK_KHR_DYNAMIC_RENDERING_EXTENSION_NAME)

	createInfo.enabledExtensionCount = C.uint32_t(len(exts))
	createInfo.ppEnabledExtensionNames = unsafe.SliceData(exts)
	pinner.Pin(createInfo.ppEnabledExtensionNames)

	if err := mapError(C.vkCreateDevice(sel.device, &createInfo, nil, &g.device)); err != nil {
		return err
	}

	C.vkGetDeviceQueue(g.device, C.uint32_t(sel.graphicsFamily), 0, &g.graphicsQueue)

	return nil
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
