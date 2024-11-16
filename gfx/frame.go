package gfx

import "github.com/csnewman/go-gfx/hal"

type Frame struct {
	app       *Application
	frame     hal.SurfaceFrame
	presented bool

	passes []RenderPassDescriptor
}

func (f *Frame) Close() {
	if f.presented {
		return
	}

	f.frame.Discard()
}

func (f *Frame) TextureView() *TextureView {
	return viewFromHal(f.frame.View())
}

type RenderPassColorAttachment struct {
	Target     TextureViewable
	Load       bool
	ClearColor Color
	Discard    bool
}

type RenderPassDescriptor struct {
	ColorAttachments []RenderPassColorAttachment

	Body func(enc *RenderPassEncoder)
}

func (f *Frame) QueueRenderPass(descriptor RenderPassDescriptor) {
	f.passes = append(f.passes, descriptor)
}

func (f *Frame) Present() error {
	f.presented = true

	buffer := f.frame.CreateCommandBuffer()

	buffer.Barrier(hal.Barrier{
		Textures: []hal.TextureBarrier{
			{
				Texture:   f.frame.Texture(),
				SrcLayout: hal.TextureLayoutUndefined,
				DstLayout: hal.TextureLayoutAttachment,
			},
		},
	})

	for _, pass := range f.passes {
		halDes := hal.RenderPassDescriptor{
			ColorAttachments: make([]hal.RenderPassColorAttachment, len(pass.ColorAttachments)),
		}

		for i, attachment := range pass.ColorAttachments {
			halDes.ColorAttachments[i] = hal.RenderPassColorAttachment{
				View:       attachment.Target.TextureView().view,
				Load:       attachment.Load,
				ClearColor: attachment.ClearColor,
				Discard:    attachment.Discard,
			}
		}

		buffer.BeginRenderPass(halDes)

		pass.Body(&RenderPassEncoder{
			buffer: buffer,
		})

		buffer.EndRenderPass()
	}

	buffer.Barrier(hal.Barrier{
		Textures: []hal.TextureBarrier{
			{
				Texture:   f.frame.Texture(),
				SrcLayout: hal.TextureLayoutAttachment,
				DstLayout: hal.TextureLayoutPresent,
			},
		},
	})

	buffer.Submit()

	return f.frame.Present()
}
