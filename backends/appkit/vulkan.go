package appkit

import (
	"github.com/csnewman/go-gfx/internal/ffi"
	"os"
	"unsafe"
)

func (p *Platform) LoadVulkan() error {
	_ = os.Setenv("MVK_CONFIG_USE_METAL_ARGUMENT_BUFFERS", "1")

	lib := ffi.Open(p.logger, []string{
		"libvulkan.dylib",
		"libvulkan.1.dylib",
		"libMoltenVK.dylib",
		"/usr/local/lib/libvulkan.dylib",
	})
	if lib == nil {
		panic("failed to open libvulkan.dylib")
	}

	p.vkAddr = lib.Symbol("vkGetInstanceProcAddr")

	return nil
}

func (p *Platform) VKInstanceProcAddr() unsafe.Pointer {
	return p.vkAddr
}

func (p *Platform) RequiredVKExtensions() []string {
	return []string{"VK_KHR_surface", "VK_EXT_metal_surface"}
}
