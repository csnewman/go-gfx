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

type RenderPassColorAttachment struct {
	Target     TextureViewable
	Load       bool
	ClearColor Color
	Discard    bool
}

type RenderPassDescriptor struct {
	ColorAttachments []RenderPassColorAttachment
}

func (b *CommandBuffer) BeginRenderPass(descriptor RenderPassDescriptor) *RenderPass {
	halDes := hal.RenderPassDescriptor{
		ColorAttachments: make([]hal.RenderPassColorAttachment, len(descriptor.ColorAttachments)),
	}

	for i, attachment := range descriptor.ColorAttachments {
		halDes.ColorAttachments[i] = hal.RenderPassColorAttachment{
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

func (p *RenderPass) SetPipeline(pipeline *RenderPipeline) {
	p.buf.buffer.SetRenderPipeline(pipeline.pipeline)
}

func (p *RenderPass) SetVertexBuffer(data *Buffer) {
	p.buf.buffer.SetVertexBuffer(data.buffer)
}

func (p *RenderPass) Draw(start int, count int) {
	p.buf.buffer.Draw(start, count)
}
