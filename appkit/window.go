package appkit

import (
	"github.com/csnewman/go-gfx/gfx"
	"log"
	"sync"
	"sync/atomic"
	"unsafe"
)

/*
#include "appkit.h"
*/
import "C"

var (
	windowCounter atomic.Uint64
	windows       sync.Map
)

type Window struct {
	id    uint64
	cfg   WindowConfig
	ptr   C.id
	layer C.id
}

type WindowConfig struct {
	Title            string
	Width            int
	Height           int
	OnResize         func(width float64, height float64, scale float64)
	OnCloseRequested func()
}

func (p *Platform) NewWindow(cfg WindowConfig) (*Window, error) {
	window := &Window{
		id:  windowCounter.Add(1),
		cfg: cfg,
	}

	r := C.gfx_ak_new_window(
		C.uint64_t(window.id),
		unsafe.Pointer(unsafe.StringData(cfg.Title)),
		C.int(len(cfg.Title)),
		C.int(cfg.Width),
		C.int(cfg.Height),
		&window.ptr,
		&window.layer,
	)

	switch r {
	case C.GFX_SUCCESS:
		windows.Store(window.id, window)

		return window, nil

	default:
		panic("unexpected response")
	}
}

func (w *Window) SurfaceHandleType() gfx.SurfaceHandleType {
	return gfx.MetalSurfaceHandleType
}

func (w *Window) MetalLayer() unsafe.Pointer {
	return unsafe.Pointer(w.layer)
}

func (w *Window) Size() gfx.LogicalSize {
	var (
		width  C.double
		height C.double
		scale  C.double
	)

	C.gfx_ak_size(w.ptr, &width, &height, &scale)

	return gfx.LogicalSize{
		Width:  float64(width),
		Height: float64(height),
		Scale:  float64(scale),
	}
}

//func (p *Platform) CloseWindow(id hal.Window) {
//	raw, ok := windows.Load(id)
//	if !ok {
//		return
//	}
//
//	C.gfx_ak_close_window(raw.(C.id))
//}

//export gfx_ak_close_requested_callback
func gfx_ak_close_requested_callback(id uint64) {
	raw, ok := windows.Load(id)
	if !ok {
		return
	}

	window := raw.(*Window)

	window.cfg.OnCloseRequested()
}

//export gfx_ak_window_closed_callback
func gfx_ak_window_closed_callback(id uint64) {
	raw, ok := windows.LoadAndDelete(id)
	if !ok {
		return
	}

	C.gfx_ak_free_context(raw.(*Window).ptr)
}

//export gfx_ak_draw_callback
func gfx_ak_draw_callback(id uint64, drawable unsafe.Pointer) {
	log.Println("draw")

	//halCfg.WindowRender(hal.Window(id), hal.MetalRenderToken{Drawable: drawable})
}

//export gfx_ak_resize_callback
func gfx_ak_resize_callback(id uint64, width float64, height float64, scale float64) {
	raw, ok := windows.Load(id)
	if !ok {
		return
	}

	window := raw.(*Window)

	window.cfg.OnResize(width, height, scale)
}
