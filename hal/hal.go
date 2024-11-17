package hal

import (
	"unsafe"
)

type PlatformConfig struct {
	Init                 func() error
	WindowCloseRequested func(w Window)
	WindowClosed         func(w Window)
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

	WindowType() WindowHandleType
	NewWindow(cfg WindowConfig) (Window, WindowHandle, error)
	CloseWindow(id Window)
}

type GPUConfig struct {
	WindowType WindowHandleType
}

type Graphics interface {
	Init(cfg GPUConfig) error

	CreateSurface(windowHandle WindowHandle) (Surface, error)
	CreateShader(cfg ShaderConfig) (Shader, error)
	CreateBuffer(data []byte) Buffer
	CreateRenderPipeline(des RenderPipelineDescriptor) (RenderPipeline, error)
}

type WindowHandle interface{}

type WindowHandleType string

type MetalWindowHandle struct {
	Layer unsafe.Pointer
}

const MetalWindowHandleType WindowHandleType = "MetalWindowHandle"

type Win32WindowHandle struct {
	Instance unsafe.Pointer
	Handle   unsafe.Pointer
}

const Win32WindowHandleType WindowHandleType = "Win32WindowHandle"

type Surface interface {
	TextureFormat() TextureFormat

	Resize(width int, height int) error

	Acquire() (SurfaceFrame, error)
}

type Texture interface {
}

type SurfaceFrame interface {
	Texture() Texture
	View() TextureView

	Present() error
	Discard()

	CreateCommandBuffer() CommandBuffer
}

type ShaderConfig struct {
	Source string
	Code   []byte
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
	Barrier(barrier Barrier)

	BeginRenderPass(description RenderPassDescriptor)
	SetRenderPipeline(pipeline RenderPipeline)
	SetVertexBuffer(data Buffer)
	Draw(start int, count int)
	EndRenderPass()

	Submit()
}

type TextureLayout int

const (
	TextureLayoutUndefined TextureLayout = iota
	TextureLayoutAttachment
	TextureLayoutRead
	TextureLayoutPresent
)

type TextureBarrier struct {
	Texture   Texture
	SrcLayout TextureLayout
	DstLayout TextureLayout
}

type Barrier struct {
	Textures []TextureBarrier
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
