package gfx

import "github.com/csnewman/go-gfx/hal"

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

type RenderPassEncoder struct {
	buffer hal.CommandBuffer
}

func (e *RenderPassEncoder) SetPipeline(pipeline *RenderPipeline) {
	e.buffer.SetRenderPipeline(pipeline.pipeline)
}

func (e *RenderPassEncoder) SetVertexBuffer(data *Buffer) {
	e.buffer.SetVertexBuffer(data.buffer)
}

func (e *RenderPassEncoder) Draw(start int, count int) {
	e.buffer.Draw(start, count)
}
