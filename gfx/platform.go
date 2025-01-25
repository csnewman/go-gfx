package gfx

import (
	"errors"
	"log/slog"
	"unsafe"
)

var (
	platforms []Platform

	ErrNoPlatforms         = errors.New("no platforms registered")
	ErrNoPlatformAvailable = errors.New("no platforms available")
	ErrUnsupportedPlatform = errors.New("unsupported platform")
)

func RegisterPlatform(platform Platform) {
	platforms = append(platforms, platform)
}

type PlatformInit struct {
	Logger  *slog.Logger
	Cfg     Config
	Init    func() error
	Render  func() error
	Resized func(size LogicalSize) error
}

type Platform interface {
	Name() string

	Priority() int

	Available() bool

	Run(init PlatformInit) error

	PrimarySurface() SurfaceHandle
}

type VulkanPlatform interface {
	Platform

	LoadVulkan() error

	VKInstanceProcAddr() unsafe.Pointer

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

	Size() LogicalSize
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
