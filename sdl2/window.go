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
	id     uint
	window *C.SDL_Window
	render func()
}

type WindowConfig struct {
	Title    string
	Width    int
	Height   int
	OnResize func(w float64, h float64)
	OnRender func()
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
		id:     id,
		window: window,
		render: cfg.OnRender,
	}

	p.windows[id] = w

	return w, nil
}

func (w *Window) SurfaceHandleType() gfx.SurfaceHandleType {
	return gfx.VulkanSurfaceHandleType
}
