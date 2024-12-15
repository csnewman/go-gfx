package sdl2

import (
	"fmt"
	"os"
	"runtime"
	"unsafe"
)

/*
#include "sdl_wrapper.h"
#include "SDL_vulkan.h"
*/
import "C"

func (p *Platform) vkInit() error {
	// TODO: remove
	if runtime.GOOS == "darwin" {
		err := os.Setenv("SDL_VULKAN_LIBRARY", "/usr/local/lib/libvulkan.dylib")
		if err != nil {
			panic(err)
		}
	}

	p.logger.Debug("Loading Vulkan")

	if v := C.SDL_Vulkan_LoadLibrary(nil); v != 0 {
		return fmt.Errorf("failed to load vulkan: %w", getError())
	}

	p.vkAddr = C.SDL_Vulkan_GetVkGetInstanceProcAddr()
	if p.vkAddr == nil {
		return fmt.Errorf("failed to get vk addr: %w", getError())
	}

	var count C.uint32_t

	if C.SDL_Vulkan_GetInstanceExtensions(nil, &count, nil) == 0 {
		return fmt.Errorf("failed to get vk count: %w", getError())
	}

	res := make([]unsafe.Pointer, count)

	if C.SDL_Vulkan_GetInstanceExtensions(nil, &count, (**C.char)(unsafe.Pointer(unsafe.SliceData(res)))) == 0 {
		return fmt.Errorf("failed to get vk extensions: %w", getError())
	}

	p.vkExts = make([]string, count)

	for i := 0; i < int(count); i++ {
		p.vkExts[i] = C.GoString((*C.char)(res[i]))
	}

	p.logger.Debug("Required Vulkan extensions", "exts", p.vkExts)

	return nil
}

func (p *Platform) VKGetInstanceProcAddr() unsafe.Pointer {
	return p.vkAddr
}

func (p *Platform) RequiredVKExtensions() []string {
	return p.vkExts
}
