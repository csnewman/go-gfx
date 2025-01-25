package vulkan

/*
#include "vulkan.h"
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/csnewman/go-gfx/gfx"
)

type RenderPipeline struct {
	pipeline C.VkPipeline
}

func (g *Graphics) CreateRenderPipeline(des gfx.RenderPipelineDescriptor) (gfx.RenderPipeline, error) {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

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
		vf := (des.VertexFunction).(*ShaderFunction)

		var stage C.VkPipelineShaderStageCreateInfo
		stage.sType = C.VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
		stage.stage = C.VK_SHADER_STAGE_VERTEX_BIT
		stage.module = vf.shader.shader
		stage.pName = C.CString(vf.function)
		defer C.free(unsafe.Pointer(stage.pName))

		shaderStages = append(shaderStages, stage)
	}

	if des.FragmentFunction != nil {
		vf := (des.FragmentFunction).(*ShaderFunction)

		var stage C.VkPipelineShaderStageCreateInfo
		stage.sType = C.VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
		stage.stage = C.VK_SHADER_STAGE_FRAGMENT_BIT
		stage.module = vf.shader.shader
		stage.pName = C.CString(vf.function)
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
				attr.format = ToFormat(attribute.Format)
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
	rasterizer.depthBiasEnable = C.VkBool32(0)
	rasterizer.lineWidth = 1.0

	switch des.CullMode {
	case gfx.CullModeNone:
		rasterizer.cullMode = C.VK_CULL_MODE_NONE
	case gfx.CullModeFront:
		rasterizer.cullMode = C.VK_CULL_MODE_FRONT_BIT
	case gfx.CullModeBack:
		rasterizer.cullMode = C.VK_CULL_MODE_BACK_BIT
	default:
		return nil, fmt.Errorf("%w: invalid cull mode %d", gfx.ErrInvalidDescriptor, des.CullMode)
	}

	if des.FrontFaceClockwise {
		rasterizer.frontFace = C.VK_FRONT_FACE_CLOCKWISE
	} else {
		rasterizer.frontFace = C.VK_FRONT_FACE_COUNTER_CLOCKWISE
	}

	var multisampling C.VkPipelineMultisampleStateCreateInfo
	multisampling.sType = C.VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO
	multisampling.sampleShadingEnable = C.VkBool32(0)
	multisampling.rasterizationSamples = C.VK_SAMPLE_COUNT_1_BIT

	var colorBlendAttachment C.VkPipelineColorBlendAttachmentState
	colorBlendAttachment.colorWriteMask = C.VK_COLOR_COMPONENT_R_BIT | C.VK_COLOR_COMPONENT_G_BIT | C.VK_COLOR_COMPONENT_B_BIT | C.VK_COLOR_COMPONENT_A_BIT
	//colorBlendAttachment.blendEnable = C.VkBool32(0)

	colorBlendAttachment.blendEnable = C.VK_TRUE
	colorBlendAttachment.srcColorBlendFactor = C.VK_BLEND_FACTOR_SRC_ALPHA
	colorBlendAttachment.dstColorBlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
	colorBlendAttachment.colorBlendOp = C.VK_BLEND_OP_ADD
	colorBlendAttachment.srcAlphaBlendFactor = C.VK_BLEND_FACTOR_ONE
	colorBlendAttachment.dstAlphaBlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
	colorBlendAttachment.alphaBlendOp = C.VK_BLEND_OP_ADD

	var colorBlending C.VkPipelineColorBlendStateCreateInfo
	colorBlending.sType = C.VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO
	colorBlending.logicOpEnable = C.VkBool32(0)
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

	pipelineInfo.layout = g.pipelineLayout

	var pipeline C.VkPipeline

	if err := mapError(C.vkCreateGraphicsPipelines(g.device, nil, 1, &pipelineInfo, nil, &pipeline)); err != nil {
		return nil, err
	}

	return &RenderPipeline{
		pipeline: pipeline,
	}, nil
}

type RenderPassEncoder struct {
	buffer *CommandBuffer
}

func (r *RenderPassEncoder) SetRenderPipeline(pipeline gfx.RenderPipeline) {
	r.buffer.SetRenderPipeline(pipeline)
}

func (r *RenderPassEncoder) SetVertexBuffer(binding int, buffer gfx.Buffer, offset int) {
	r.buffer.SetVertexBuffer(binding, buffer, offset)
}

func (r *RenderPassEncoder) SetIndexBuffer(buffer gfx.Buffer, offset int) {
	r.buffer.SetIndexBuffer(buffer, offset)
}

func (r *RenderPassEncoder) SetPushConstants(offset int, size int, data unsafe.Pointer) {
	r.buffer.SetPushConstants(offset, size, data)
}

func (r *RenderPassEncoder) Draw(start int, count int) {
	r.buffer.Draw(start, count)
}

func (r *RenderPassEncoder) DrawIndexed(start int, count int, vertexOffset int) {
	r.buffer.DrawIndexed(start, count, vertexOffset)
}

func (r *RenderPassEncoder) End() {
	r.buffer.EndRenderPass()
}
