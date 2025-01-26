package windows

import (
	"syscall"
	"unsafe"
)

func (p *Platform) LoadVulkan() error {

	lib, err := syscall.LoadDLL("vulkan-1.dll")
	if err != nil {
		panic("failed to open vulkan-1.dll")
	}

	proc, err := lib.FindProc("vkGetInstanceProcAddr")
	if err != nil {
		panic("failed to find proc")
	}

	p.vkAddr = unsafe.Pointer(proc.Addr())

	return nil
}

func (p *Platform) VKInstanceProcAddr() unsafe.Pointer {
	return p.vkAddr
}

func (p *Platform) RequiredVKExtensions() []string {
	return []string{"VK_KHR_surface", "VK_KHR_win32_surface"}
}
