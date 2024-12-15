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
		logger: c.Logger,
		init:   c.Init,
	}

	// TODO: expose alternatives
	if err := p.vkInit(); err != nil {
		return nil, err
	}

	return p, nil
}

type Platform struct {
	logger *slog.Logger
	init   func() error
	vkAddr unsafe.Pointer
	vkExts []string
}

func (p *Platform) Run() error {
	if err := p.init(); err != nil {
		return err
	}

	for {
		var event C.SDL_Event

		for C.SDL_WaitEvent(&event) != 0 {
			ptr := &event
			ty := *(*C.Uint32)(unsafe.Pointer(ptr))

			switch ty {
			case C.SDL_QUIT:
				p.logger.Info("Quiting")
				return nil

			case C.SDL_MOUSEMOTION:
				// Ignore

			default:
				p.logger.Debug("Unhandled event", "ty", ty)
			}
		}
	}
}

func (p *Platform) Exit() {
	//TODO implement me
	panic("implement me")
}

//func (p *Platform) GetInstanceExtensions() ([]unsafe.Pointer, error) {
//	var count C.uint32_t
//
//	if C.SDL_Vulkan_GetInstanceExtensions(p.window, &count, nil) == 0 {
//		return nil, getError()
//	}
//
//	res := make([]unsafe.Pointer, count)
//
//	if C.SDL_Vulkan_GetInstanceExtensions(p.window, &count, (**C.char)(unsafe.Pointer(unsafe.SliceData(res)))) == 0 {
//		return nil, getError()
//	}
//
//	return res[:count], nil
//}
//
//func (p *Platform) GetLoader() unsafe.Pointer {
//	return C.SDL_Vulkan_GetVkGetInstanceProcAddr()
//}
