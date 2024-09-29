package appkit

/*
#cgo darwin LDFLAGS: -framework AppKit

#include "appkit.h"
*/
import "C"

import (
	"errors"
	"sync/atomic"
)

var (
	runCounter        atomic.Uint32
	callbacks         Callbacks
	ErrAlreadyRunning = errors.New("already running")
	ErrNotMainThread  = errors.New("not on main thread")
)

type Callbacks struct {
	Init func()
}

func Run(cb Callbacks) error {
	if !runCounter.CompareAndSwap(0, 1) {
		return ErrAlreadyRunning
	}

	callbacks = cb

	r := C.gfx_ak_run()

	switch r {
	case C.GFX_SUCCESS:
		return nil

	case C.GFX_NOT_MAIN_THREAD:
		return ErrNotMainThread

	default:
		panic("unexpected response")
	}
}

func NewWindow(width int, height int) error {
	r := C.gfx_ak_new_window(C.int(width), C.int(height))

	switch r {
	case C.GFX_SUCCESS:
		return nil

	default:
		panic("unexpected response")
	}
}

//export gfx_ak_init_callback
func gfx_ak_init_callback() {
	callbacks.Init()
}
