package hal

import (
	"unsafe"
)

type PlatformConfig struct {
	Init                 func() error
	WindowCloseRequested func(w Window)
	WindowClosed         func(w Window)
	WindowRender         func(w Window, token RenderToken)
	WindowResized        func(w Window, width float64, height float64)
}

type WindowConfig struct {
	Title  string
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
	CreateShader(cfg ShaderConfig) (Shader, error)
	CreateBuffer(data []byte) Buffer
	CreateRenderPipeline(des RenderPipelineDescriptor) (RenderPipeline, error)
	CreateCommandBuffer() CommandBuffer
}

type RenderToken any

type MetalRenderToken struct {
	Drawable unsafe.Pointer
}

type WindowHandle interface{}

type MetalWindowHandle struct {
	Layer unsafe.Pointer
}

type Surface interface {
	TextureFormat() TextureFormat

	AcquireTexture(token RenderToken) (SurfaceTexture, error)
}

type SurfaceTexture interface {
	View() TextureView

	Present() error
	Discard()
}

type ShaderConfig struct {
	Source string
}

type Shader interface {
	ResolveFunction(name string) (ShaderFunction, error)
}

type ShaderFunction interface {
}

type Buffer interface {
}

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

type RenderPassColorAttachment struct {
	View       TextureView
	Load       bool
	ClearColor Color
	Discard    bool
}

type RenderPassDescriptor struct {
	ColorAttachments []RenderPassColorAttachment
}

type CommandBuffer interface {
	BeginRenderPass(description RenderPassDescriptor)
	SetRenderPipeline(pipeline RenderPipeline)
	SetVertexBuffer(data Buffer)
	Draw(start int, count int)
	EndRenderPass()

	Submit()
}

type TextureView interface {
}

type TextureFormat int

const (
	TextureFormatBGRA8UNorm TextureFormat = iota
)

type RenderPipelineColorAttachment struct {
	Format TextureFormat
}

type RenderPipelineDescriptor struct {
	VertexFunction   ShaderFunction
	FragmentFunction ShaderFunction
	ColorAttachments []RenderPipelineColorAttachment
}

type RenderPipeline interface {
}
