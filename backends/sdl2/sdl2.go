package sdl2

/*
#cgo CFLAGS: -I${SRCDIR}/../../thirdparty/SDL/include -fobjc-arc -Wno-deprecated-declarations

#cgo darwin LDFLAGS: -framework AppKit -framework QuartzCore -framework IOKit -framework Carbon
#cgo windows LDFLAGS: -l Gdi32 -l Imm32 -l OleAut32 -l Ole32 -l Version -l Api-ms-win-core-version-l1-1-0 -lm -lwinmm

#include "sdl_wrapper.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"github.com/csnewman/go-gfx/gfx"
	"log/slog"
	"sync/atomic"
	"unsafe"
)

var ErrSDL = errors.New("sdl error")

func getError() error {
	return fmt.Errorf("%w: %v", ErrSDL, C.GoString(C.SDL_GetError()))
}

var (
	instance   *Platform
	runCounter atomic.Uint32
)

func Register() {
	if instance != nil {
		return
	}

	instance = &Platform{
		windows: make(map[uint]*Window),
	}

	gfx.RegisterPlatform(instance)
}

type Platform struct {
	logger  *slog.Logger
	cfg     gfx.Config
	windows map[uint]*Window
	vkAddr  unsafe.Pointer
	vkExts  []string

	initCallback   func() error
	renderCallback func() error
	resizeCallback func(size gfx.LogicalSize) error
	mainWindow     *Window
	exitRequested  atomic.Bool
}

func (p *Platform) Name() string {
	return "sdl2"
}

func (p *Platform) Priority() int {
	return -1
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

	C.SDL_Init(C.SDL_INIT_VIDEO)

	if err := p.vkInit(); err != nil {
		return err
	}

	var err error

	if p.mainWindow, err = p.NewWindow(WindowConfig{
		Title:    p.cfg.Title,
		Width:    p.cfg.Width,
		Height:   p.cfg.Height,
		OnResize: p.onResize,
		OnRender: p.onRender,
		//OnCloseRequested: p.OnCloseRequested,
	}); err != nil {
		panic(err)
	}

	if err := p.initCallback(); err != nil {
		panic(err)
	}

	for {
		var event C.SDL_Event

		for C.SDL_PollEvent(&event) != 0 {
			ptr := &event
			ty := *(*C.Uint32)(unsafe.Pointer(ptr))

			switch ty {
			case C.SDL_QUIT:
				p.logger.Info("Quiting")
				return nil

			case C.SDL_WINDOWEVENT:

				evt := (*C.SDL_WindowEvent)(unsafe.Pointer(ptr))
				wid := uint(evt.windowID)

				w, ok := p.windows[wid]
				if !ok {
					p.logger.Debug("Ignored window event", "wid", wid, "evt", evt.event)

					continue
				}

				switch evt.event {
				case C.SDL_WINDOWEVENT_RESIZED, C.SDL_WINDOWEVENT_SIZE_CHANGED:
					w.nextWidth = int(int32(evt.data1))
					w.nextHeight = int(int32(evt.data2))
				default:
					p.logger.Info("Window event received", "time", evt.timestamp, "wid", evt.windowID, "evt", evt.event)
				}

			case C.SDL_MOUSEMOTION:
				// Ignore

			default:
				p.logger.Debug("Unhandled event", "ty", ty)
			}
		}

		for _, window := range p.windows {
			if err := window.doRender(); err != nil {
				return err
			}
		}
	}
}

func (p *Platform) Exit() {
	//TODO implement me
	panic("implement me")
}

func (p *Platform) onResize(size gfx.LogicalSize) error {
	if err := p.resizeCallback(size); err != nil {
		panic(err)
	}

	if err := p.renderCallback(); err != nil {
		panic(err)
	}

	return nil
}

func (p *Platform) onRender() error {
	return p.renderCallback()
}
