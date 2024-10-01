package appkit

import (
	"errors"
	"sync/atomic"
)

/*

#cgo darwin LDFLAGS: -framework AppKit

#include "appkit.h"
*/
import "C"

var (
	runCounter        atomic.Uint32
	callbacks         Callbacks
	ErrAlreadyRunning = errors.New("already running")
	ErrNotMainThread  = errors.New("not on main thread")
)

type Callbacks struct {
	Init func()

	CloseRequested func(w Window)

	Closed func(w Window)
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

//export gfx_ak_init_callback
func gfx_ak_init_callback() {
	callbacks.Init()
}

func Stop() {
	C.gfx_ak_stop()

	windows.Range(func(key, value any) bool {
		ptr := value.(C.id)
		C.gfx_ak_free_context(ptr)

		return true
	})

	windows.Clear()
}
