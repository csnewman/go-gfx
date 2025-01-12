package gfx

import "unsafe"

type Graphics interface {
	CreateShader(cfg ShaderConfig) (Shader, error)

	CreateRenderPipeline(descriptor RenderPipelineDescriptor) (RenderPipeline, error)

	CreateImage(descriptor ImageDescriptor) (Image, error)

	CreateBuffer(descriptor BufferDescriptor) (Buffer, error)
}

type Surface interface {
	Resize(width int, height int) error

	Acquire() (SurfaceFrame, error)

	Format() Format
}

type SurfaceFrame interface {
	ImageViewer

	Index() int

	BeginRenderPass(descriptor RenderPassDescriptor) RenderPassEncoder

	Present() error
}

type ShaderConfig struct {
	SPIRV []byte
}

type Shader interface {
	Function(name string) (ShaderFunction, error)
}

type ShaderFunction interface{}

type ImageType int

const (
	ImageType1D ImageType = iota
	ImageType2D
	ImageType3D
)

type ImageUsage int

const (
	ImageUsageAttachment ImageUsage = 1 << iota
	ImageUsageSampled
	ImageUsageCopySrc
	ImageUsageCopyDst
)

type ImageDescriptor struct {
	Type   ImageType
	Width  int
	Height int
	Depth  int
	Format Format
	Usage  ImageUsage
}

type Image interface{}

type Format int

func (f Format) Depth() bool {
	// TODO
	return false
}

const (
	FormatBGRA8UNorm Format = iota
	FormatRGBA8UNorm
	FormatRGBA16SFloat
	FormatRGB32SFloat
	FormatRG32SFloat
)

type RenderPipelineColorAttachment struct {
	Format Format
}

type VertexRate int

const VertexRateVertex VertexRate = 0
const VertexRateInstance VertexRate = 1

type VertexBinding struct {
	Binding    int
	Stride     int
	Rate       VertexRate
	Attributes []VertexAttribute
}

type VertexAttribute struct {
	Location int
	Offset   int
	Format   Format
}

type CullMode int

const (
	CullModeNone CullMode = iota
	CullModeFront
	CullModeBack
)

type RenderPipelineDescriptor struct {
	VertexFunction     ShaderFunction
	VertexBindings     []VertexBinding
	FragmentFunction   ShaderFunction
	ColorAttachments   []RenderPipelineColorAttachment
	CullMode           CullMode
	FrontFaceClockwise bool
}

type RenderPipeline interface{}

type BufferUsage int

const (
	BufferUsageHostRandomAccess BufferUsage = 1 << iota
	BufferUsageHostUpload
	BufferUsagePersistentMap
)

type BufferDescriptor struct {
	Size  int
	Usage BufferUsage
}

type Buffer interface {
	Close()

	Size() int

	MappedPtr() (unsafe.Pointer, bool)

	CopyFrom(slice []byte) error

	Flush() error
}

type ImageViewer interface {
	ImageView() ImageView
}

type ImageView interface {
	Width() int
	Height() int
}

type RenderPassDescriptor struct {
	ColorAttachments []RenderPassColorAttachment
}

type RenderPassColorAttachment struct {
	Target     ImageViewer
	Load       bool
	ClearColor Color
	Discard    bool
}

type RenderPassEncoder interface {
	SetRenderPipeline(pipeline RenderPipeline)

	SetVertexBuffer(binding int, buffer Buffer, offset int)

	SetIndexBuffer(buffer Buffer, offset int)

	SetPushConstants(offset int, size int, data unsafe.Pointer)

	Draw(start int, count int)

	DrawIndexed(start int, count int, vertexOffset int)

	End()
}
