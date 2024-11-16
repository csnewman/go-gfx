package windows

import (
	"fmt"
	"github.com/csnewman/go-gfx/hal"
	"sync/atomic"
)

/*
#include "helper.h"
*/
import "C"

var (
	runCounter atomic.Uint32
	halCfg     hal.PlatformConfig
)

type Platform struct {
}

func NewPlatform() hal.Platform {
	return &Platform{}
}

func (p *Platform) Run(cfg hal.PlatformConfig) error {
	if !runCounter.CompareAndSwap(0, 1) {
		return hal.ErrAlreadyRunning
	}

	halCfg = cfg

	res := C.gfx_windows_init()

	switch res {
	case C.GFX_SUCCESS:
		return nil
	case C.GFX_MODULE_ERROR:
		return fmt.Errorf("%w: failed to get module", hal.ErrUnexpectedSystemResponse)
	case C.GFX_CLASS_ERROR:
		return fmt.Errorf("%w: failed to create window class", hal.ErrUnexpectedSystemResponse)
	default:
		panic("unexpected result")
	}

	return nil
}

//export gfx_windows_init_callback
func gfx_windows_init_callback() {
	if err := halCfg.Init(); err != nil {
		panic(err)
	}
}

func (p *Platform) Exit() {
	//TODO implement me
	panic("implement me")
}
