package appkit

import (
	"log/slog"
	"sync/atomic"
	"unsafe"

	"github.com/csnewman/go-gfx/gfx"
)

/*
#cgo darwin LDFLAGS: -framework AppKit -framework QuartzCore

#include "appkit.h"
*/
import "C"

var (
	instance   *Platform
	runCounter atomic.Uint32
)

func Register() {
	if instance != nil {
		return
	}

	instance = &Platform{}

	gfx.RegisterPlatform(instance)
}

type Platform struct {
	logger         *slog.Logger
	vkAddr         unsafe.Pointer
	cfg            gfx.Config
	initCallback   func() error
	renderCallback func() error
	resizeCallback func(size gfx.LogicalSize) error
	mainWindow     *Window
	exitRequested  atomic.Bool
}

func (p *Platform) Name() string {
	return "appkit"
}

func (p *Platform) Priority() int {
	return 0
}

func (p *Platform) Available() bool {
	return true
}

func (p *Platform) PrimarySurface() gfx.SurfaceHandle {
	return p.mainWindow
}

func (p *Platform) Run(init gfx.PlatformInit) error {
	if !runCounter.CompareAndSwap(0, 1) {
		return gfx.ErrAlreadyRunning
	}

	p.logger = init.Logger
	p.cfg = init.Cfg
	p.initCallback = init.Init
	p.renderCallback = init.Render
	p.resizeCallback = init.Resized

	r := C.gfx_ak_run()

	switch r {
	case C.GFX_SUCCESS:
	case C.GFX_NOT_MAIN_THREAD:
		return gfx.ErrNotMainThread

	default:
		panic("unexpected response")
	}

	var err error

	if p.mainWindow, err = p.NewWindow(WindowConfig{
		Title:            p.cfg.Title,
		Width:            p.cfg.Width,
		Height:           p.cfg.Height,
		OnResize:         p.onResize,
		OnCloseRequested: p.OnCloseRequested,
	}); err != nil {
		panic(err)
	}

	if err := p.initCallback(); err != nil {
		panic(err)
	}

	for {
		if p.exitRequested.Load() {
			break
		}

		C.gfx_ak_process_events()

		if err := init.Render(); err != nil {
			panic(err)
		}

	}

	//if err := ; err != nil {
	//	return err
	//}

	return nil
}

//export gfx_ak_init_callback
func gfx_ak_init_callback() {
	instance.init()
}

func (p *Platform) init() {
	var err error

	if p.mainWindow, err = p.NewWindow(WindowConfig{
		Title:            p.cfg.Title,
		Width:            p.cfg.Width,
		Height:           p.cfg.Height,
		OnResize:         p.onResize,
		OnCloseRequested: p.OnCloseRequested,
	}); err != nil {
		panic(err)
	}

	if err := p.initCallback(); err != nil {
		panic(err)
	}
}

func (p *Platform) onResize(size gfx.LogicalSize) {
	if err := p.resizeCallback(size); err != nil {
		panic(err)
	}

	if err := p.renderCallback(); err != nil {
		panic(err)
	}

}

func (p *Platform) OnCloseRequested() {
	p.exitRequested.Store(true)
}

func (p *Platform) Exit() {
	//C.gfx_ak_stop()

	windows.Range(func(key, value any) bool {
		window := value.(*Window)

		C.gfx_ak_free_context(window.ptr)

		return true
	})

	windows.Clear()
}
