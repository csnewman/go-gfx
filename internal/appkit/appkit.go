package appkit

import (
	"github.com/csnewman/go-gfx/hal"
	"sync/atomic"
)

/*
#cgo darwin LDFLAGS: -framework AppKit -framework QuartzCore

#include "appkit.h"
*/
import "C"

var (
	runCounter atomic.Uint32
	halCfg     hal.PlatformConfig
)

func NewPlatform() hal.Platform {
	return &Platform{}
}

type Platform struct {
}

func (p *Platform) Run(cfg hal.PlatformConfig) error {
	if !runCounter.CompareAndSwap(0, 1) {
		return hal.ErrAlreadyRunning
	}

	halCfg = cfg

	r := C.gfx_ak_run()

	switch r {
	case C.GFX_SUCCESS:
		return nil

	case C.GFX_NOT_MAIN_THREAD:
		return hal.ErrNotMainThread

	default:
		panic("unexpected response")
	}
}

//export gfx_ak_init_callback
func gfx_ak_init_callback() {
	if err := halCfg.Init(); err != nil {
		panic(err)
	}
}

func (p *Platform) Exit() {
	C.gfx_ak_stop()

	windows.Range(func(key, value any) bool {
		ptr := value.(C.id)
		C.gfx_ak_free_context(ptr)

		return true
	})

	windows.Clear()
}
