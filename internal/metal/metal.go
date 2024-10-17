package metal

/*
#cgo darwin LDFLAGS: -framework Metal

#include "metal.h"
*/
import "C"

import "github.com/csnewman/go-gfx/hal"

type graphics struct {
	device C.id
}

func NewGraphics() hal.Graphics {
	return &graphics{}
}

func (g *graphics) Init(cfg hal.GPUConfig) error {
	r := C.gfx_mtl_open(&g.device)

	switch r {
	case C.GFX_SUCCESS:
		return nil

	default:
		panic("unexpected response")
	}
}

func (g *graphics) CreateSurface(wh hal.WindowHandle) (hal.Surface, error) {
	return 0, nil
}
