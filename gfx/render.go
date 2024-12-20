package gfx

/*
#include "vulkan.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

type RenderPipelineColorAttachment struct {
	Format Format
}

type VertexRate int

const VertexRateVertex VertexRate = 0
const VertexRateInstance VertexRate = 1

type VertexBinding struct {
	Binding    int
	Stride     int
	Rate       VertexRate
	Attributes []VertexAttribute
}

type VertexAttribute struct {
	Location int
	// TODO: format
	Offset int
}

type RenderPipelineDescriptor struct {
	VertexFunction   *ShaderFunction
	VertexBindings   []VertexBinding
	FragmentFunction *ShaderFunction
	ColorAttachments []RenderPipelineColorAttachment
}

type RenderPipeline struct {
	pipeline C.VkPipeline
}

func (g *Graphics) CreateRenderPipeline(des RenderPipelineDescriptor) (*RenderPipeline, error) {
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
		var stage C.VkPipelineShaderStageCreateInfo
		stage.sType = C.VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
		stage.stage = C.VK_SHADER_STAGE_VERTEX_BIT
		stage.module = des.VertexFunction.shader.shader
		stage.pName = C.CString(des.VertexFunction.function)
		defer C.free(unsafe.Pointer(stage.pName))

		shaderStages = append(shaderStages, stage)
	}

	if des.FragmentFunction != nil {
		var stage C.VkPipelineShaderStageCreateInfo
		stage.sType = C.VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
		stage.stage = C.VK_SHADER_STAGE_FRAGMENT_BIT
		stage.module = des.FragmentFunction.shader.shader
		stage.pName = C.CString(des.FragmentFunction.function)
		defer C.free(unsafe.Pointer(stage.pName))

		shaderStages = append(shaderStages, stage)
	}

	var vertexInputInfo C.VkPipelineVertexInputStateCreateInfo
	vertexInputInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO
	vertexInputInfo.vertexBindingDescriptionCount = 0
	vertexInputInfo.vertexAttributeDescriptionCount = 0

	if len(des.VertexBindings) > 0 {
		var bindingDes []C.VkVertexInputBindingDescription
		var attrDes []C.VkVertexInputAttributeDescription

		for _, binding := range des.VertexBindings {
			var des C.VkVertexInputBindingDescription
			des.binding = C.uint32_t(binding.Binding)
			des.stride = C.uint32_t(binding.Stride)

			// des.inputRate TODO

			bindingDes = append(bindingDes, des)

			for _, attribute := range binding.Attributes {
				var attr C.VkVertexInputAttributeDescription

				attr.binding = C.uint32_t(binding.Binding)
				attr.location = C.uint32_t(attribute.Location)

				//attr.format = // TODO
				attr.format = C.VK_FORMAT_R32G32B32_SFLOAT

				attr.offset = C.uint32_t(attribute.Offset)

				attrDes = append(attrDes, attr)
			}
		}

		vertexInputInfo.vertexBindingDescriptionCount = C.uint32_t(len(bindingDes))
		vertexInputInfo.pVertexBindingDescriptions = unsafe.SliceData(bindingDes)
		pinner.Pin(vertexInputInfo.pVertexBindingDescriptions)

		vertexInputInfo.vertexAttributeDescriptionCount = C.uint32_t(len(attrDes))
		vertexInputInfo.pVertexAttributeDescriptions = unsafe.SliceData(attrDes)
		pinner.Pin(vertexInputInfo.pVertexAttributeDescriptions)
	}

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

	var viewportState C.VkPipelineViewportStateCreateInfo
	viewportState.sType = C.VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO

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

	pipelineInfo.pViewportState = &viewportState
	pinner.Pin(pipelineInfo.pViewportState)

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

func ToFormat(format Format) C.VkFormat {
	switch format {
	case FormatBGRA8UNorm:
		return C.VK_FORMAT_B8G8R8A8_UNORM
	default:
		panic("unknown format")
	}
}
