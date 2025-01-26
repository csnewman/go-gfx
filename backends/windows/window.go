package windows

import (
	"github.com/csnewman/go-gfx/gfx"
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"
)

/*
#include "helper.h"
*/
import "C"

var (
	windowCounter atomic.Uint64
	windows       sync.Map
)

type Window struct {
	plt    *Platform
	id     uint64
	cfg    WindowConfig
	handle C.HWND
}

func (w *Window) Size() gfx.LogicalSize {
	var (
		width  C.int
		height C.int
	)

	C.gfx_windows_size(w.handle, &width, &height)

	return gfx.LogicalSize{
		Width:  float64(width),
		Height: float64(height),
		Scale:  1.0,
	}
}

type WindowConfig struct {
	Title            string
	Width            int
	Height           int
	OnResize         func(size gfx.LogicalSize)
	OnCloseRequested func()
}

func (p *Platform) NewWindow(cfg WindowConfig) (*Window, error) {
	window := &Window{
		plt: p,
		id:  windowCounter.Add(1),
		cfg: cfg,
	}

	titleStr, err := syscall.UTF16PtrFromString(cfg.Title)
	if err != nil {
		return nil, err
	}

	res := C.gfx_windows_new_window(
		C.uint64_t(window.id),
		C.LPCWSTR(unsafe.Pointer(titleStr)),
		C.int(cfg.Width),
		C.int(cfg.Height),
		&window.handle,
	)

	switch res {
	case C.GFX_SUCCESS:
		windows.Store(window.id, window)

		return window, nil
	case C.GFX_CALL_ERROR:
		return nil, gfx.ErrUnexpectedSystemResponse
	default:
		panic("unexpected result")
	}
}

func (w *Window) SurfaceHandleType() gfx.SurfaceHandleType {
	return gfx.Win32SurfaceHandleType
}

func (w *Window) Win32Instance() unsafe.Pointer {
	return unsafe.Pointer(w.plt.instance)
}

func (w *Window) Win32Handle() unsafe.Pointer {
	return unsafe.Pointer(w.handle)
}

//export gfx_windows_resize_callback
func gfx_windows_resize_callback(id uint64, width int64, height int64) {
	raw, ok := windows.Load(id)
	if !ok {
		return
	}

	window := raw.(*Window)

	window.cfg.OnResize(gfx.LogicalSize{
		Width:  float64(width),
		Height: float64(height),
		Scale:  1.0, // TODO: fix
	})
}

//export gfx_windows_close_requested_callback
func gfx_windows_close_requested_callback(id uint64) {
	raw, ok := windows.Load(id)
	if !ok {
		return
	}

	window := raw.(*Window)

	window.cfg.OnCloseRequested()
}

//export gfx_windows_window_closed_callback
func gfx_windows_window_closed_callback(id uint64) {
	//raw, ok := windows.LoadAndDelete(id)
	//if !ok {
	//	return
	//}
	//
	//window := raw.(*Window)
	//
	//window.cfg.OnCloseRequested()

	// TODO: cleanup?
}

//func (p *Platform) CloseWindow(id hal.Window) {
//	raw, ok := windows.Load(id)
//	if !ok {
//		return
//	}
//
//	C.gfx_windows_close_window(raw.(C.HWND))
//}
