package gfx

import "github.com/csnewman/go-gfx/hal"

type RenderPipelineColorAttachment struct {
	Format TextureFormat
}

type RenderPipelineDescriptor struct {
	VertexFunction   *ShaderFunction
	FragmentFunction *ShaderFunction
	ColorAttachments []RenderPipelineColorAttachment
}

type RenderPipeline struct {
	pipeline hal.RenderPipeline
}

func (a *Application) NewRenderPipeline(descriptor RenderPipelineDescriptor) (*RenderPipeline, error) {
	halDes := hal.RenderPipelineDescriptor{
		ColorAttachments: make([]hal.RenderPipelineColorAttachment, len(descriptor.ColorAttachments)),
	}

	if descriptor.VertexFunction != nil {
		halDes.VertexFunction = descriptor.VertexFunction.function
	}

	if descriptor.FragmentFunction != nil {
		halDes.FragmentFunction = descriptor.FragmentFunction.function
	}

	for i, attachment := range descriptor.ColorAttachments {
		halDes.ColorAttachments[i] = hal.RenderPipelineColorAttachment{
			Format: attachment.Format,
		}
	}

	pipeline, err := a.graphics.CreateRenderPipeline(halDes)
	if err != nil {
		return nil, err
	}

	return &RenderPipeline{
		pipeline: pipeline,
	}, nil
}
