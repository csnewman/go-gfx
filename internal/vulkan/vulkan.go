package vulkan

/*
#cgo pkg-config: vulkan
#cgo CXXFLAGS: -std=c++20

#include <stdlib.h>
#include <vulkan/vulkan.h>
#include <vulkan/vulkan_metal.h>
#include "vulkan.h"
#include "vk_mem_alloc.h"

const char* GFX_VK_KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME = VK_KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME;
const char* GFX_VK_KHR_SURFACE_EXTENSION_NAME = VK_KHR_SURFACE_EXTENSION_NAME;
const char* GFX_VK_EXT_DEBUG_UTILS_EXTENSION_NAME = VK_EXT_DEBUG_UTILS_EXTENSION_NAME;
const char* GFX_VK_LAYER_KHRONOS_validation = "VK_LAYER_KHRONOS_validation";
const char* GFX_VK_KHR_portability_subset = "VK_KHR_portability_subset";
const char* GFX_VK_KHR_DYNAMIC_RENDERING_EXTENSION_NAME = VK_KHR_DYNAMIC_RENDERING_EXTENSION_NAME;
const char* GFX_VK_EXT_METAL_SURFACE_EXTENSION_NAME = VK_EXT_METAL_SURFACE_EXTENSION_NAME;
const char* GFX_VK_KHR_SWAPCHAIN_EXTENSION_NAME = VK_KHR_SWAPCHAIN_EXTENSION_NAME;
const char* GFX_VK_EXT_EXTENDED_DYNAMIC_STATE_EXTENSION_NAME = VK_EXT_EXTENDED_DYNAMIC_STATE_EXTENSION_NAME;
const char* GFX_VK_EXT_EXTENDED_DYNAMIC_STATE_2_EXTENSION_NAME = VK_EXT_EXTENDED_DYNAMIC_STATE_2_EXTENSION_NAME;
const char* GFX_VK_EXT_EXTENDED_DYNAMIC_STATE_3_EXTENSION_NAME = VK_EXT_EXTENDED_DYNAMIC_STATE_3_EXTENSION_NAME;
const char* GFX_VK_KHR_SYNCHRONIZATION_2_EXTENSION_NAME = VK_KHR_SYNCHRONIZATION_2_EXTENSION_NAME;

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
	instance        C.VkInstance
	debugMessenger  C.VkDebugUtilsMessengerEXT
	physicalDevice  C.VkPhysicalDevice
	graphicsFamily  int
	device          C.VkDevice
	graphicsQueue   C.VkQueue
	memoryAllocator C.VmaAllocator
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

	g.physicalDevice = sel.device
	g.graphicsFamily = sel.graphicsFamily

	priority := C.float(1.0)

	var queueCreateInfo C.VkDeviceQueueCreateInfo
	queueCreateInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
	queueCreateInfo.queueFamilyIndex = C.uint32_t(sel.graphicsFamily)
	queueCreateInfo.queueCount = 1
	queueCreateInfo.pQueuePriorities = &priority
	pinner.Pin(queueCreateInfo.pQueuePriorities)

	var extendedDynamicState C.VkPhysicalDeviceExtendedDynamicStateFeaturesEXT
	extendedDynamicState.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_FEATURES_EXT
	extendedDynamicState.extendedDynamicState = C.VkBool32(1)

	var extendedDynamicState3 C.VkPhysicalDeviceExtendedDynamicState3FeaturesEXT
	extendedDynamicState3.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_3_FEATURES_EXT
	extendedDynamicState3.pNext = unsafe.Pointer(&extendedDynamicState)
	pinner.Pin(extendedDynamicState3.pNext)

	var dynamicRenderingFeatures C.VkPhysicalDeviceDynamicRenderingFeatures
	dynamicRenderingFeatures.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_FEATURES
	dynamicRenderingFeatures.pNext = unsafe.Pointer(&extendedDynamicState3)
	pinner.Pin(dynamicRenderingFeatures.pNext)
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

	// TODO: switch to vk1.3
	exts = append(exts, C.GFX_VK_KHR_DYNAMIC_RENDERING_EXTENSION_NAME)
	exts = append(exts, C.GFX_VK_EXT_EXTENDED_DYNAMIC_STATE_EXTENSION_NAME)
	exts = append(exts, C.GFX_VK_EXT_EXTENDED_DYNAMIC_STATE_2_EXTENSION_NAME)
	exts = append(exts, C.GFX_VK_EXT_EXTENDED_DYNAMIC_STATE_3_EXTENSION_NAME)
	exts = append(exts, C.GFX_VK_KHR_SYNCHRONIZATION_2_EXTENSION_NAME)
	exts = append(exts, C.GFX_VK_KHR_SWAPCHAIN_EXTENSION_NAME)

	createInfo.enabledExtensionCount = C.uint32_t(len(exts))
	createInfo.ppEnabledExtensionNames = unsafe.SliceData(exts)
	pinner.Pin(createInfo.ppEnabledExtensionNames)

	if err := mapError(C.vkCreateDevice(sel.device, &createInfo, nil, &g.device)); err != nil {
		return err
	}

	C.vkGetDeviceQueue(g.device, C.uint32_t(sel.graphicsFamily), 0, &g.graphicsQueue)

	var vmaInfo C.VmaAllocatorCreateInfo
	vmaInfo.vulkanApiVersion = C.VK_API_VERSION_1_3
	vmaInfo.physicalDevice = sel.device
	vmaInfo.device = g.device
	vmaInfo.instance = g.instance

	if err := mapError(C.vmaCreateAllocator(&vmaInfo, &g.memoryAllocator)); err != nil {
		return err
	}

	return nil
}

type Shader struct {
	shader C.VkShaderModule
}

func (g *Graphics) CreateShader(cfg hal.ShaderConfig) (hal.Shader, error) {

	var createInfo C.VkShaderModuleCreateInfo
	createInfo.sType = C.VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO
	createInfo.codeSize = C.size_t(len(cfg.Code))
	createInfo.pCode = (*C.uint32_t)(unsafe.Pointer(unsafe.SliceData(cfg.Code)))

	var shaderModule C.VkShaderModule

	if err := mapError(C.vkCreateShaderModule(g.device, &createInfo, nil, &shaderModule)); err != nil {
		return nil, err
	}

	return &Shader{
		shader: shaderModule,
	}, nil
}

type ShaderFunction struct {
	shader   *Shader
	function string
}

func (s *Shader) ResolveFunction(name string) (hal.ShaderFunction, error) {
	return &ShaderFunction{
		shader:   s,
		function: name,
	}, nil
}

type RenderPipeline struct {
	pipeline C.VkPipeline
}

func (g *Graphics) CreateRenderPipeline(des hal.RenderPipelineDescriptor) (hal.RenderPipeline, error) {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	var pipelineLayoutInfo C.VkPipelineLayoutCreateInfo
	pipelineLayoutInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO
	pipelineLayoutInfo.setLayoutCount = 0
	pipelineLayoutInfo.pushConstantRangeCount = 0

	var pipelineLayout C.VkPipelineLayout

	if err := mapError(C.vkCreatePipelineLayout(g.device, &pipelineLayoutInfo, nil, &pipelineLayout)); err != nil {
		return nil, err
	}

	var renderingInfo C.VkPipelineRenderingCreateInfoKHR
	renderingInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_RENDERING_CREATE_INFO_KHR

	colorFmts := make([]C.VkFormat, len(des.ColorAttachments))

	for i, c := range des.ColorAttachments {
		colorFmts[i] = ToFormat(c.Format)
	}

	renderingInfo.colorAttachmentCount = C.uint32_t(len(colorFmts))
	renderingInfo.pColorAttachmentFormats = unsafe.SliceData(colorFmts)
	pinner.Pin(renderingInfo.pColorAttachmentFormats)

	var shaderStages []C.VkPipelineShaderStageCreateInfo

	if des.VertexFunction != nil {
		vf, ok := des.VertexFunction.(*ShaderFunction)
		if !ok {
			panic("unexpected type")
		}

		var stage C.VkPipelineShaderStageCreateInfo
		stage.sType = C.VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
		stage.stage = C.VK_SHADER_STAGE_VERTEX_BIT
		stage.module = vf.shader.shader
		stage.pName = C.CString(vf.function)
		defer C.free(unsafe.Pointer(stage.pName))

		shaderStages = append(shaderStages, stage)
	}

	if des.FragmentFunction != nil {
		ff, ok := des.FragmentFunction.(*ShaderFunction)
		if !ok {
			panic("unexpected type")
		}

		var stage C.VkPipelineShaderStageCreateInfo
		stage.sType = C.VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
		stage.stage = C.VK_SHADER_STAGE_FRAGMENT_BIT
		stage.module = ff.shader.shader
		stage.pName = C.CString(ff.function)
		defer C.free(unsafe.Pointer(stage.pName))

		shaderStages = append(shaderStages, stage)
	}

	var vertexInputInfo C.VkPipelineVertexInputStateCreateInfo
	vertexInputInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO
	vertexInputInfo.vertexBindingDescriptionCount = 0
	vertexInputInfo.vertexAttributeDescriptionCount = 0

	var inputAssembly C.VkPipelineInputAssemblyStateCreateInfo
	inputAssembly.sType = C.VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO
	inputAssembly.topology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST
	inputAssembly.primitiveRestartEnable = C.VkBool32(0)

	var rasterizer C.VkPipelineRasterizationStateCreateInfo
	rasterizer.sType = C.VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO
	rasterizer.depthClampEnable = C.VkBool32(0)
	rasterizer.rasterizerDiscardEnable = C.VkBool32(0)
	rasterizer.polygonMode = C.VK_POLYGON_MODE_FILL
	rasterizer.cullMode = C.VK_CULL_MODE_BACK_BIT
	rasterizer.frontFace = C.VK_FRONT_FACE_CLOCKWISE
	rasterizer.depthBiasEnable = C.VkBool32(0)
	rasterizer.lineWidth = 1.0

	var multisampling C.VkPipelineMultisampleStateCreateInfo
	multisampling.sType = C.VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO
	multisampling.sampleShadingEnable = C.VkBool32(0)
	multisampling.rasterizationSamples = C.VK_SAMPLE_COUNT_1_BIT

	var colorBlendAttachment C.VkPipelineColorBlendAttachmentState
	colorBlendAttachment.colorWriteMask = C.VK_COLOR_COMPONENT_R_BIT | C.VK_COLOR_COMPONENT_G_BIT | C.VK_COLOR_COMPONENT_B_BIT | C.VK_COLOR_COMPONENT_A_BIT
	colorBlendAttachment.blendEnable = C.VkBool32(0)

	var colorBlending C.VkPipelineColorBlendStateCreateInfo
	colorBlending.sType = C.VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO
	colorBlending.logicOpEnable = C.VkBool32(0)
	colorBlending.logicOp = C.VK_LOGIC_OP_COPY
	colorBlending.attachmentCount = 1
	colorBlending.pAttachments = &colorBlendAttachment
	pinner.Pin(colorBlending.pAttachments)
	colorBlending.blendConstants[0] = 0.0
	colorBlending.blendConstants[1] = 0.0
	colorBlending.blendConstants[2] = 0.0
	colorBlending.blendConstants[3] = 0.0

	var dynamicStates []C.VkDynamicState

	dynamicStates = append(dynamicStates, C.VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT)
	dynamicStates = append(dynamicStates, C.VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT)

	var dynamicState C.VkPipelineDynamicStateCreateInfo
	dynamicState.sType = C.VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO
	dynamicState.dynamicStateCount = C.uint32_t(len(dynamicStates))
	dynamicState.pDynamicStates = unsafe.SliceData(dynamicStates)
	pinner.Pin(dynamicState.pDynamicStates)

	var pipelineInfo C.VkGraphicsPipelineCreateInfo

	pipelineInfo.sType = C.VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO

	pipelineInfo.pNext = unsafe.Pointer(&renderingInfo)
	pinner.Pin(pipelineInfo.pNext)

	pipelineInfo.stageCount = C.uint32_t(len(shaderStages))
	pipelineInfo.pStages = unsafe.SliceData(shaderStages)
	pinner.Pin(pipelineInfo.pStages)

	pipelineInfo.pVertexInputState = &vertexInputInfo
	pinner.Pin(pipelineInfo.pVertexInputState)

	pipelineInfo.pInputAssemblyState = &inputAssembly
	pinner.Pin(pipelineInfo.pInputAssemblyState)

	pipelineInfo.pRasterizationState = &rasterizer
	pinner.Pin(pipelineInfo.pRasterizationState)

	pipelineInfo.pMultisampleState = &multisampling
	pinner.Pin(pipelineInfo.pMultisampleState)

	pipelineInfo.pColorBlendState = &colorBlending
	pinner.Pin(pipelineInfo.pColorBlendState)

	pipelineInfo.pDynamicState = &dynamicState
	pinner.Pin(pipelineInfo.pDynamicState)

	pipelineInfo.layout = pipelineLayout

	var pipeline C.VkPipeline

	if err := mapError(C.vkCreateGraphicsPipelines(g.device, nil, 1, &pipelineInfo, nil, &pipeline)); err != nil {
		return nil, err
	}

	return &RenderPipeline{
		pipeline: pipeline,
	}, nil
}

func ToFormat(format hal.TextureFormat) C.VkFormat {
	switch format {
	case hal.TextureFormatBGRA8UNorm:
		return C.VK_FORMAT_B8G8R8A8_UNORM
	default:
		panic("unknown format")
	}
}
