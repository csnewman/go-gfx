package appkit

import (
	"sync"
	"sync/atomic"
)

/*
#include "appkit.h"
*/
import "C"

var (
	windowCounter atomic.Uint32
	windows       sync.Map
)

type Window uint32

func NewWindow(width int, height int) (Window, error) {
	var res C.id

	id := Window(windowCounter.Add(1))

	r := C.gfx_ak_new_window(C.uint32_t(id), C.int(width), C.int(height), &res)

	switch r {
	case C.GFX_SUCCESS:
		windows.Store(id, res)
		return id, nil

	default:
		panic("unexpected response")
	}
}

//export gfx_ak_close_requested_callback
func gfx_ak_close_requested_callback(id uint32) {
	wid := Window(id)

	if _, ok := windows.Load(wid); !ok {
		return
	}

	callbacks.CloseRequested(wid)
}

//export gfx_ak_window_closed_callback
func gfx_ak_window_closed_callback(id uint32) {
	wid := Window(id)

	raw, ok := windows.LoadAndDelete(wid)
	if !ok {
		return
	}

	callbacks.Closed(wid)

	ptr := raw.(C.id)
	C.gfx_ak_free_context(ptr)
}

func (wid Window) Close() {
	raw, ok := windows.Load(wid)
	if !ok {
		return
	}

	ptr := raw.(C.id)

	C.gfx_ak_close_window(ptr)
}

//export gfx_ak_draw_callback
func gfx_ak_draw_callback(id uint32) {
	wid := Window(id)

	if _, ok := windows.Load(wid); !ok {
		return
	}

	callbacks.Render(wid)
}
