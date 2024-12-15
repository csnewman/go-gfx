package sdl2

import (
	"unsafe"
)

/*
#include "sdl_wrapper.h"
*/
import "C"

type Window struct {
	window *C.SDL_Window
}

type WindowConfig struct {
	Title  string
	Width  int
	Height int
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

	return nil, nil
}
