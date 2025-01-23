package appkit

import (
	"fmt"
	"github.com/csnewman/go-gfx/gfx"
	"sync/atomic"
	"unsafe"
)

/*
#cgo darwin LDFLAGS: -framework AppKit -framework QuartzCore

#include "appkit.h"
*/
import "C"

var (
	runCounter atomic.Uint32
	halCfg     Config
)

type Config struct {
	LoadVulkan bool
	Init       func() error
}

func Init(cfg Config) (*Platform, error) {
	if !runCounter.CompareAndSwap(0, 1) {
		return nil, gfx.ErrAlreadyRunning
	}

	halCfg = cfg

	p := &Platform{}

	if cfg.LoadVulkan {
		if err := p.vkInit(); err != nil {
			return nil, fmt.Errorf("failed to init Vulkan: %v", err)
		}
	}

	return p, nil
}

type Platform struct {
	vkAddr unsafe.Pointer
}

func (p *Platform) Run() error {
	r := C.gfx_ak_run()

	switch r {
	case C.GFX_SUCCESS:
		return nil

	case C.GFX_NOT_MAIN_THREAD:
		return gfx.ErrNotMainThread

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
		window := value.(*Window)

		C.gfx_ak_free_context(window.ptr)

		return true
	})

	windows.Clear()
}
