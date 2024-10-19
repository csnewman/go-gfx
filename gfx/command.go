package gfx

import "github.com/csnewman/go-gfx/hal"

func (f *Frame) NewCommandBuffer() *CommandBuffer {
	buffer := f.app.graphics.CreateCommandBuffer()

	return &CommandBuffer{
		buffer: buffer,
	}
}

type CommandBuffer struct {
	buffer hal.CommandBuffer
}

func (b *CommandBuffer) Submit() {
	b.buffer.Submit()
}

type RenderPass struct {
	buf *CommandBuffer
}

type Color = hal.Color

func NewColor(r float64, g float64, b float64, a float64) Color {
	return Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

var Black = Color{
	A: 1,
}

type ColorAttachment struct {
	Target     TextureViewable
	Load       bool
	ClearColor Color
	Discard    bool
}

type RenderPassDescriptor struct {
	ColorAttachments []ColorAttachment
}

func (b *CommandBuffer) BeginRenderPass(descriptor RenderPassDescriptor) *RenderPass {
	halDes := hal.RenderPassDescriptor{
		ColorAttachments: make([]hal.ColorAttachment, len(descriptor.ColorAttachments)),
	}

	for i, attachment := range descriptor.ColorAttachments {
		halDes.ColorAttachments[i] = hal.ColorAttachment{
			View:       attachment.Target.TextureView().view,
			Load:       attachment.Load,
			ClearColor: attachment.ClearColor,
			Discard:    attachment.Discard,
		}
	}

	b.buffer.BeginRenderPass(halDes)

	return &RenderPass{
		buf: b,
	}
}

func (p *RenderPass) End() {
	p.buf.buffer.EndRenderPass()
}
