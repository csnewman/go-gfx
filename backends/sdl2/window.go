package sdl2

import (
	"unsafe"

	"github.com/csnewman/go-gfx/gfx"
)

/*
#include "sdl_wrapper.h"
*/
import "C"

type Window struct {
	id         uint
	window     *C.SDL_Window
	render     func() error
	resize     func(size gfx.LogicalSize) error
	width      int
	height     int
	nextWidth  int
	nextHeight int
}

type WindowConfig struct {
	Title    string
	Width    int
	Height   int
	OnResize func(size gfx.LogicalSize) error
	OnRender func() error
}

func (p *Platform) NewWindow(cfg WindowConfig) (*Window, error) {
	cTitle := C.CString(cfg.Title)
	defer C.free(unsafe.Pointer(cTitle))

	window := C.SDL_CreateWindow(
		cTitle,
		C.SDL_WINDOWPOS_CENTERED,
		C.SDL_WINDOWPOS_CENTERED,
		C.int(cfg.Width),
		C.int(cfg.Height),
		C.SDL_WINDOW_SHOWN|C.SDL_WINDOW_RESIZABLE|C.SDL_WINDOW_VULKAN,
	)
	if window == nil {
		return nil, getError()
	}

	id := uint(C.SDL_GetWindowID(window))

	w := &Window{
		id:         id,
		window:     window,
		render:     cfg.OnRender,
		resize:     cfg.OnResize,
		width:      cfg.Width,
		nextWidth:  cfg.Width,
		height:     cfg.Height,
		nextHeight: cfg.Height,
	}

	p.windows[id] = w

	return w, nil
}

func (w *Window) SurfaceHandleType() gfx.SurfaceHandleType {
	return gfx.VulkanSurfaceHandleType
}

func (w *Window) doRender() error {
	if w.width != w.nextWidth || w.height != w.nextHeight {
		w.width, w.height = w.nextWidth, w.nextHeight

		if w.resize != nil {
			if err := w.resize(gfx.LogicalSize{
				Width:  float64(w.width),
				Height: float64(w.height),
				Scale:  1.0,
			}); err != nil {
				return err
			}
		}
	}

	return w.render()
}

func (w *Window) Size() gfx.LogicalSize {
	return gfx.LogicalSize{
		Width:  float64(w.width),
		Height: float64(w.height),
		Scale:  1.0,
	}
}
