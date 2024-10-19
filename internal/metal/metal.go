package metal

/*
#cgo darwin LDFLAGS: -framework Metal

#include "metal.h"
*/
import "C"

import "github.com/csnewman/go-gfx/hal"

type Graphics struct {
	device C.id
	queue  C.id
}

func NewGraphics() hal.Graphics {
	return &Graphics{}
}

func (g *Graphics) Init(cfg hal.GPUConfig) error {
	r := C.gfx_mtl_open(&g.device, &g.queue)

	switch r {
	case C.GFX_SUCCESS:
		return nil

	default:
		panic("unexpected response")
	}
}

func (g *Graphics) CreateSurface(rawWH hal.WindowHandle) (hal.Surface, error) {
	wh, ok := rawWH.(hal.MetalWindowHandle)
	if !ok {
		return nil, hal.ErrUnsupportedWindowHandle
	}

	layer := C.id(wh.Layer)

	r := C.gfx_mtl_configure_surface(g.device, layer)

	switch r {
	case C.GFX_SUCCESS:
		return &Surface{
			graphics: g,
			layer:    layer,
		}, nil

	default:
		panic("unexpected response")
	}
}

type Surface struct {
	graphics *Graphics
	layer    C.id
}

func (s *Surface) AcquireTexture() (hal.SurfaceTexture, error) {
	var text C.id

	C.gfx_mtl_acquire_texture(s.layer, &text)

	return &SurfaceTexture{
		graphics: s.graphics,
		texture:  text,
	}, nil
}

type SurfaceTexture struct {
	graphics *Graphics
	texture  C.id
}

func (s *SurfaceTexture) Present() error {
	C.gfx_mtl_present_texture(s.graphics.queue, s.texture)

	return nil
}

func (s *SurfaceTexture) Discard() {
	C.gfx_mtl_discard_surface_texture(s.texture)
}
