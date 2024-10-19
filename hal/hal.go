package hal

import "unsafe"

type PlatformConfig struct {
	Init                 func() error
	WindowCloseRequested func(w Window)
	WindowClosed         func(w Window)
	WindowRender         func(w Window)
}

type WindowConfig struct {
	Width  int
	Height int
}

type Window uint64

type Platform interface {
	Run(cfg PlatformConfig) error
	Exit()

	NewWindow(cfg WindowConfig) (Window, WindowHandle, error)
	CloseWindow(id Window)
}

type GPUConfig struct {
}

type Graphics interface {
	Init(cfg GPUConfig) error

	CreateSurface(windowHandle WindowHandle) (Surface, error)

	CreateCommandBuffer() CommandBuffer
}

type WindowHandle interface{}

type MetalWindowHandle struct {
	Layer unsafe.Pointer
}

type Surface interface {
	AcquireTexture() (SurfaceTexture, error)
}

type SurfaceTexture interface {
	View() TextureView

	Present() error
	Discard()
}

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

type ColorAttachment struct {
	View       TextureView
	Load       bool
	ClearColor Color
	Discard    bool
}

type RenderPassDescriptor struct {
	ColorAttachments []ColorAttachment
}

type CommandBuffer interface {
	BeginRenderPass(description RenderPassDescriptor)
	EndRenderPass()

	Submit()
}

type TextureView interface {
}
