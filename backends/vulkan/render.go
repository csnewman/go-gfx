package vulkan

import (
	"fmt"
	"unsafe"

	"github.com/csnewman/go-gfx/ffi"
	"github.com/csnewman/go-gfx/gfx"
	"github.com/csnewman/go-gfx/vk"
)

type RenderPipeline struct {
	pipeline vk.Pipeline
}

func (g *Graphics) CreateRenderPipeline(des gfx.RenderPipelineDescriptor) (gfx.RenderPipeline, error) {
	arena := ffi.NewArena()
	defer arena.Close()

	colorFmts := make([]vk.Format, len(des.ColorAttachments))

	for i, c := range des.ColorAttachments {
		colorFmts[i] = ToFormat(c.Format)
	}

	renderingInfo := vk.PipelineRenderingCreateInfoAlloc(arena, 1)
	renderingInfo.SetSType(vk.STRUCTURE_TYPE_PIPELINE_RENDERING_CREATE_INFO)
	renderingInfo.SetColorAttachmentCount(uint32(len(colorFmts)))
	renderingInfo.SetPColorAttachmentFormats(ffi.RefFromValues(arena, colorFmts...))

	stageCount := 0

	if des.VertexFunction != nil {
		stageCount++
	}

	if des.FragmentFunction != nil {
		stageCount++
	}

	shaderStages := vk.PipelineShaderStageCreateInfoAlloc(arena, stageCount)
	shaderStagesOffset := 0

	if des.VertexFunction != nil {
		vf := (des.VertexFunction).(*ShaderFunction)

		stage := shaderStages.Offset(shaderStagesOffset)
		stage.SetSType(vk.STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO)
		stage.SetStage(vk.SHADER_STAGE_VERTEX_BIT)
		stage.SetModule(vf.shader.shader)
		stage.SetPName(ffi.CStringAlloc(arena, vf.function))

		shaderStagesOffset++
	}

	if des.FragmentFunction != nil {
		vf := (des.FragmentFunction).(*ShaderFunction)

		stage := shaderStages.Offset(shaderStagesOffset)
		stage.SetSType(vk.STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO)
		stage.SetStage(vk.SHADER_STAGE_FRAGMENT_BIT)
		stage.SetModule(vf.shader.shader)
		stage.SetPName(ffi.CStringAlloc(arena, vf.function))

		shaderStagesOffset++
	}

	vertexInputInfo := vk.PipelineVertexInputStateCreateInfoAlloc(arena, 1)
	vertexInputInfo.SetSType(vk.STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO)
	vertexInputInfo.SetVertexBindingDescriptionCount(0)
	vertexInputInfo.SetVertexAttributeDescriptionCount(0)

	if len(des.VertexBindings) > 0 {
		attrCount := 0

		for _, binding := range des.VertexBindings {
			attrCount += len(binding.Attributes)
		}

		bindingDes := vk.VertexInputBindingDescriptionAlloc(arena, len(des.VertexBindings))
		attrDes := vk.VertexInputAttributeDescriptionAlloc(arena, attrCount)
		attrOffset := 0

		for i, binding := range des.VertexBindings {
			des := bindingDes.Offset(i)
			des.SetBinding(uint32(binding.Binding))
			des.SetStride(uint32(binding.Stride))

			// des.inputRate TODO

			for _, attribute := range binding.Attributes {
				attr := attrDes.Offset(attrOffset)
				attr.SetBinding(uint32(binding.Binding))
				attr.SetLocation(uint32(attribute.Location))
				attr.SetFormat(ToFormat(attribute.Format))
				attr.SetOffset(uint32(attribute.Offset))

				attrOffset++
			}
		}

		vertexInputInfo.SetVertexBindingDescriptionCount(uint32(len(des.VertexBindings)))
		vertexInputInfo.SetPVertexBindingDescriptions(bindingDes)
		vertexInputInfo.SetVertexAttributeDescriptionCount(uint32(attrCount))
		vertexInputInfo.SetPVertexAttributeDescriptions(attrDes)
	}

	inputAssembly := vk.PipelineInputAssemblyStateCreateInfoAlloc(arena, 1)
	inputAssembly.SetSType(vk.STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO)
	inputAssembly.SetTopology(vk.PRIMITIVE_TOPOLOGY_TRIANGLE_LIST)
	inputAssembly.SetPrimitiveRestartEnable(false)

	rasterizer := vk.PipelineRasterizationStateCreateInfoAlloc(arena, 1)
	rasterizer.SetSType(vk.STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO)
	rasterizer.SetDepthClampEnable(false)
	rasterizer.SetRasterizerDiscardEnable(false)
	rasterizer.SetPolygonMode(vk.POLYGON_MODE_FILL)
	rasterizer.SetDepthBiasEnable(false)
	rasterizer.SetLineWidth(1.0)

	switch des.CullMode {
	case gfx.CullModeNone:
		rasterizer.SetCullMode(vk.CULL_MODE_NONE)
	case gfx.CullModeFront:
		rasterizer.SetCullMode(vk.CULL_MODE_FRONT_BIT)
	case gfx.CullModeBack:
		rasterizer.SetCullMode(vk.CULL_MODE_BACK_BIT)
	default:
		return nil, fmt.Errorf("%w: invalid cull mode %d", gfx.ErrInvalidDescriptor, des.CullMode)
	}

	if des.FrontFaceClockwise {
		rasterizer.SetFrontFace(vk.FRONT_FACE_CLOCKWISE)
	} else {
		rasterizer.SetFrontFace(vk.FRONT_FACE_COUNTER_CLOCKWISE)
	}

	multisampling := vk.PipelineMultisampleStateCreateInfoAlloc(arena, 1)
	multisampling.SetSType(vk.STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO)
	multisampling.SetSampleShadingEnable(false)
	multisampling.SetRasterizationSamples(vk.SAMPLE_COUNT_1_BIT)

	colorBlendAttachment := vk.PipelineColorBlendAttachmentStateAlloc(arena, 1)
	colorBlendAttachment.SetColorWriteMask(vk.COLOR_COMPONENT_R_BIT | vk.COLOR_COMPONENT_G_BIT | vk.COLOR_COMPONENT_B_BIT | vk.COLOR_COMPONENT_A_BIT)
	//colorBlendAttachment.blendEnable = C.VkBool32(0)

	colorBlendAttachment.SetBlendEnable(true)
	colorBlendAttachment.SetSrcColorBlendFactor(vk.BLEND_FACTOR_SRC_ALPHA)
	colorBlendAttachment.SetDstColorBlendFactor(vk.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA)
	colorBlendAttachment.SetColorBlendOp(vk.BLEND_OP_ADD)
	colorBlendAttachment.SetSrcAlphaBlendFactor(vk.BLEND_FACTOR_ONE)
	colorBlendAttachment.SetDstAlphaBlendFactor(vk.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA)
	colorBlendAttachment.SetAlphaBlendOp(vk.BLEND_OP_ADD)

	colorBlending := vk.PipelineColorBlendStateCreateInfoAlloc(arena, 1)
	colorBlending.SetSType(vk.STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO)
	colorBlending.SetLogicOpEnable(false)
	colorBlending.SetAttachmentCount(1)
	colorBlending.SetPAttachments(colorBlendAttachment)

	// TODO: restore, correct by calloc for now!
	//colorBlending.blendConstants[0] = 0.0
	//colorBlending.blendConstants[1] = 0.0
	//colorBlending.blendConstants[2] = 0.0
	//colorBlending.blendConstants[3] = 0.0

	var dynamicStates []vk.DynamicState

	dynamicStates = append(dynamicStates, vk.DYNAMIC_STATE_VIEWPORT_WITH_COUNT)
	dynamicStates = append(dynamicStates, vk.DYNAMIC_STATE_SCISSOR_WITH_COUNT)

	dynamicState := vk.PipelineDynamicStateCreateInfoAlloc(arena, 1)
	dynamicState.SetSType(vk.STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO)
	dynamicState.SetDynamicStateCount(uint32(len(dynamicStates)))
	dynamicState.SetPDynamicStates(ffi.RefFromValues(arena, dynamicStates...))

	viewportState := vk.PipelineViewportStateCreateInfoAlloc(arena, 1)
	viewportState.SetSType(vk.STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO)

	pipelineInfo := vk.GraphicsPipelineCreateInfoAlloc(arena, 1)
	pipelineInfo.SetSType(vk.STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO)
	pipelineInfo.SetPNext(uintptr(renderingInfo))
	pipelineInfo.SetStageCount(uint32(stageCount))
	pipelineInfo.SetPStages(shaderStages)
	pipelineInfo.SetPVertexInputState(vertexInputInfo)
	pipelineInfo.SetPInputAssemblyState(inputAssembly)
	pipelineInfo.SetPViewportState(viewportState)
	pipelineInfo.SetPRasterizationState(rasterizer)
	pipelineInfo.SetPMultisampleState(multisampling)
	pipelineInfo.SetPColorBlendState(colorBlending)
	pipelineInfo.SetPDynamicState(dynamicState)
	pipelineInfo.SetLayout(g.pipelineLayout)

	pipelineRef := ffi.RefAlloc[vk.Pipeline](arena, 1)

	if err := mapError(vk.CreateGraphicsPipelines(g.device, vk.PipelineCacheNil, 1, pipelineInfo, vk.AllocationCallbacksNil, pipelineRef)); err != nil {
		return nil, err
	}

	return &RenderPipeline{
		pipeline: pipelineRef.Get(),
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
