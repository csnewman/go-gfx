package windows

import (
	"github.com/csnewman/go-gfx/hal"
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

func (p *Platform) WindowType() hal.WindowHandleType {
	return hal.Win32WindowHandleType
}

func (p *Platform) NewWindow(cfg hal.WindowConfig) (hal.Window, hal.WindowHandle, error) {
	id := hal.Window(windowCounter.Add(1))

	titleStr, err := syscall.UTF16PtrFromString(cfg.Title)
	if err != nil {
		return 0, nil, err
	}

	var handle C.HWND

	res := C.gfx_windows_new_window(
		C.uint64_t(id),
		C.LPCWSTR(unsafe.Pointer(titleStr)),
		C.int(cfg.Width),
		C.int(cfg.Height),
		&handle,
	)

	switch res {
	case C.GFX_SUCCESS:
		return id, hal.Win32WindowHandle{
			Instance: unsafe.Pointer(p.instance),
			Handle:   unsafe.Pointer(handle),
		}, nil
	case C.GFX_CALL_ERROR:
		return 0, 0, hal.ErrUnexpectedSystemResponse
	default:
		panic("unexpected result")
	}
}

//export gfx_windows_resize_callback
func gfx_windows_resize_callback(id uint64, width int64, height int64) {
	halCfg.WindowResized(hal.Window(id), float64(width), float64(height))
}

func (p *Platform) CloseWindow(id hal.Window) {
	//TODO implement me
	panic("implement me")
}
