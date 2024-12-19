package gfx

import "unsafe"

type VulkanPlatform interface {
	VKGetInstanceProcAddr() unsafe.Pointer

	RequiredVKExtensions() []string
}

type SurfaceHandleType string

const (
	VulkanSurfaceHandleType = "vulkan"
	MetalSurfaceHandleType  = "metal"
	Win32SurfaceHandleType  = "win32"
)

type SurfaceHandle interface {
	SurfaceHandleType() SurfaceHandleType
}

type VulkanSurfaceHandle interface {
	SurfaceHandle
	CreateVkSurface(instance unsafe.Pointer) (unsafe.Pointer, error)
}

type MetalSurfaceHandle interface {
	SurfaceHandle
	MetalLayer() unsafe.Pointer
}

type Win32SurfaceHandle interface {
	SurfaceHandle
	Win32Instance() unsafe.Pointer
	Win32Handle() unsafe.Pointer
}
