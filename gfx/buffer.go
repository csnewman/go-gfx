package gfx

import "github.com/csnewman/go-gfx/hal"

type Buffer struct {
	buffer hal.Buffer
}

func (a *Application) NewBuffer(data []byte) *Buffer {
	buffer := a.graphics.CreateBuffer(data)

	return &Buffer{
		buffer: buffer,
	}
}
