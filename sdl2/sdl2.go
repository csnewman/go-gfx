package sdl2

/*
#cgo CFLAGS: -I${SRCDIR}/../thirdparty/SDL/include -fobjc-arc -Wno-deprecated-declarations

#cgo darwin LDFLAGS: -framework AppKit -framework QuartzCore -framework IOKit -framework Carbon
#cgo windows LDFLAGS: -l Gdi32 -l Imm32 -l OleAut32 -l Ole32 -l Version -l Api-ms-win-core-version-l1-1-0 -lm -lwinmm

#include "sdl_wrapper.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"log/slog"
	"unsafe"
)

var ErrSDL = errors.New("sdl error")

func getError() error {
	return fmt.Errorf("%w: %v", ErrSDL, C.GoString(C.SDL_GetError()))
}

type Config struct {
	Logger *slog.Logger
	Init   func() error
}

func Init(c Config) (*Platform, error) {
	c.Logger.Debug("Initialising")

	C.SDL_Init(C.SDL_INIT_VIDEO)

	p := &Platform{
		logger:  c.Logger,
		init:    c.Init,
		windows: make(map[uint]*Window),
	}

	// TODO: expose alternatives
	if err := p.vkInit(); err != nil {
		return nil, err
	}

	return p, nil
}

type Platform struct {
	logger  *slog.Logger
	init    func() error
	windows map[uint]*Window
	vkAddr  unsafe.Pointer
	vkExts  []string
}

func (p *Platform) Run() error {
	if err := p.init(); err != nil {
		return err
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
