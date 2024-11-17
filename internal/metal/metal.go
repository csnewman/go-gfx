package metal

/*
#cgo darwin LDFLAGS: -framework Metal

#include <stdlib.h>
#include "metal.h"
*/
import "C"

import (
	"errors"
	"github.com/csnewman/go-gfx/hal"
	"github.com/csnewman/go-gfx/internal/spvc"
	"unsafe"
)

type Graphics struct {
	device C.id
	queue  C.id
}

func NewGraphics() hal.Graphics {
	return &Graphics{}
}

func (g *Graphics) Init(cfg hal.GPUConfig) error {
	r := C.gfx_mtl_open(&g.device, &g.queue)

	switch r {
	case C.GFX_SUCCESS:
		return nil

	default:
		panic("unexpected response")
	}
}

func (g *Graphics) CreateSurface(rawWH hal.WindowHandle) (hal.Surface, error) {
	wh, ok := rawWH.(hal.MetalWindowHandle)
	if !ok {
		return nil, hal.ErrUnsupportedWindowHandle
	}

	layer := C.id(wh.Layer)

	format := hal.TextureFormatBGRA8UNorm

	var (
		width  C.double
		height C.double
	)

	r := C.gfx_mtl_configure_surface(g.device, layer, ToPixelFormat(format), &width, &height)

	switch r {
	case C.GFX_SUCCESS:
		return &Surface{
			graphics: g,
			layer:    layer,
			format:   format,
			width:    int(width),
			height:   int(height),
		}, nil

	default:
		panic("unexpected response")
	}
}

func ToPixelFormat(tf hal.TextureFormat) C.int {
	switch tf {
	case hal.TextureFormatBGRA8UNorm:
		return 80
	default:
		panic("unknown format")
	}
}

type Surface struct {
	graphics *Graphics
	layer    C.id
	format   hal.TextureFormat
	width    int
	height   int
}

func (s *Surface) Resize(width int, height int) error {
	s.width = width
	s.height = height

	C.gfx_mtl_resize_surface(s.layer, C.int(width), C.int(height))

	return nil
}

func (s *Surface) TextureFormat() hal.TextureFormat {
	return s.format
}

func (s *Surface) Acquire() (hal.SurfaceFrame, error) {
	var (
		draw C.id
		text C.id
	)

	C.gfx_mtl_acquire_surface(s.layer, &draw, &text)

	return &SurfaceFrame{
		graphics: s.graphics,
		drawable: draw,
		texture:  text,
		width:    s.width,
		height:   s.height,
	}, nil
}

type SurfaceFrame struct {
	graphics *Graphics
	drawable C.id
	texture  C.id
	width    int
	height   int
}

func (f *SurfaceFrame) Texture() hal.Texture {
	//TODO implement me
	//panic("implement me")
	return nil
}

func (f *SurfaceFrame) Present() error {
	C.gfx_mtl_present_texture(f.graphics.queue, f.drawable)

	return nil
}

func (f *SurfaceFrame) Discard() {
	C.gfx_mtl_discard_surface_texture(f.drawable)
}

func (f *SurfaceFrame) View() hal.TextureView {
	// TODO: ownership

	return &TextureView{
		texture: f.texture,
	}
}

type TextureView struct {
	texture C.id
}

type Shader struct {
	funcs map[string]*ShaderFunction
}

func (g *Graphics) CreateShader(cfg hal.ShaderConfig) (hal.Shader, error) {
	ctx, err := spvc.NewContext()
	if err != nil {
		return nil, err
	}

	defer ctx.Close()

	ir, err := ctx.ParseSPIRV(cfg.SPIRV)
	if err != nil {
		return nil, err
	}

	cc, err := ir.CreateCompiler(spvc.BackendMSL)
	if err != nil {
		return nil, err
	}

	eps, err := cc.EntryPoints()
	if err != nil {
		return nil, err
	}

	funcs := make(map[string]*ShaderFunction)

	// TODO: cleanup on error
	for _, e := range eps {
		com, err := ir.CreateCompiler(spvc.BackendMSL)
		if err != nil {
			return nil, err
		}

		if err := com.SetEntryPoint(e); err != nil {
			return nil, err
		}

		com.SetMSLPlatform(spvc.MSLPlatformMacOS)

		src, err := com.Compile()
		if err != nil {
			return nil, err
		}

		cleanedName := com.GetCleanedEntryPoint(e)

		var (
			lib      C.id
			errorStr *C.char
		)

		r := C.gfx_mtl_create_shader(
			g.device,
			unsafe.Pointer(unsafe.StringData(src)),
			C.int(len(src)),
			&lib,
			&errorStr,
		)

		switch r {
		case C.GFX_SUCCESS:
			var fun C.id

			C.gfx_mtl_get_shader_function(
				lib,
				unsafe.Pointer(unsafe.StringData(cleanedName)),
				C.int(len(cleanedName)),
				&fun,
			)

			if fun == nil {
				// TODO: handle better
				return nil, hal.ErrFunctionNotFound
			}

			funcs[e.Name] = &ShaderFunction{
				lib:      lib,
				function: fun,
			}
		case C.GFX_SEE_ERROR:
			defer C.free(unsafe.Pointer(errorStr))

			return nil, errors.New(C.GoString(errorStr))
		default:
			panic("unexpected response")
		}

	}

	return &Shader{
		funcs: funcs,
	}, nil
}

type ShaderFunction struct {
	lib      C.id
	function C.id
}

func (s *Shader) ResolveFunction(name string) (hal.ShaderFunction, error) {
	fun, ok := s.funcs[name]
	if !ok {

		return nil, hal.ErrFunctionNotFound
	}

	return fun, nil
}

type Buffer struct {
	buffer C.id
}

func (g *Graphics) CreateBuffer(data []byte) hal.Buffer {
	var buf C.id

	C.gfx_mtl_buffer_from_bytes(
		g.device,
		unsafe.Pointer(unsafe.SliceData(data)),
		C.int(len(data)),
		&buf,
	)

	return &Buffer{
		buffer: buf,
	}
}

type RenderPipeline struct {
	pipeline C.id
}

func (g *Graphics) CreateRenderPipeline(des hal.RenderPipelineDescriptor) (hal.RenderPipeline, error) {
	var (
		vert C.id
		frag C.id
	)

	if des.VertexFunction != nil {
		vf, ok := des.VertexFunction.(*ShaderFunction)
		if !ok {
			panic("unexpected type")
		}

		vert = vf.function
	}

	if des.FragmentFunction != nil {
		vf, ok := des.FragmentFunction.(*ShaderFunction)
		if !ok {
			panic("unexpected type")
		}

		frag = vf.function
	}

	cAttachs := make([]C.PipelineColorAttachment, len(des.ColorAttachments))

	for i, c := range des.ColorAttachments {
		cAttachs[i] = C.PipelineColorAttachment{
			format: ToPixelFormat(c.Format),
		}
	}

	cAttachsPtr := unsafe.SliceData(cAttachs)

	var (
		pipeline C.id
		errorStr *C.char
	)

	r := C.gfx_mtl_create_render_pipeline(
		g.device,
		vert,
		frag,
		cAttachsPtr,
		C.uint64_t(len(cAttachs)),
		&pipeline,
		&errorStr,
	)

	switch r {
	case C.GFX_SUCCESS:
		return &RenderPipeline{
			pipeline: pipeline,
		}, nil

	case C.GFX_SEE_ERROR:
		defer C.free(unsafe.Pointer(errorStr))

		return nil, errors.New(C.GoString(errorStr))
	default:
		panic("unexpected response")
	}
}

type CommandBuffer struct {
	frame         *SurfaceFrame
	buffer        C.id
	renderEncoder C.id
}

func (f *SurfaceFrame) CreateCommandBuffer() hal.CommandBuffer {
	var buf C.id

	// TODO: synchronise
	C.gfx_mtl_create_command_buf(f.graphics.queue, &buf)

	return &CommandBuffer{
		frame:  f,
		buffer: buf,
	}
}

func (c *CommandBuffer) Barrier(barrier hal.Barrier) {
	// Unneeded on Metal
}

func (c *CommandBuffer) BeginRenderPass(description hal.RenderPassDescriptor) {
	cAttachs := make([]C.ColorAttachment, len(description.ColorAttachments))

	for i, c := range description.ColorAttachments {
		tv, ok := c.View.(*TextureView)
		if !ok {
			panic("unexpected view type")
		}

		cAttachs[i] = C.ColorAttachment{
			view:  tv.texture,
			load:  C.bool(c.Load),
			store: C.bool(!c.Discard),
			r:     C.double(c.ClearColor.R),
			g:     C.double(c.ClearColor.G),
			b:     C.double(c.ClearColor.B),
			a:     C.double(c.ClearColor.A),
		}
	}

	cAttachsPtr := unsafe.SliceData(cAttachs)

	C.gfx_mtl_begin_rpass(
		c.buffer,
		cAttachsPtr,
		C.uint64_t(len(cAttachs)),
		&c.renderEncoder,
	)

	// Flipped-y viewport to match Vulkan
	C.gfx_mtl_set_viewport(
		c.renderEncoder,
		C.double(0.0),
		C.double(float64(c.frame.height)),
		C.double(float64(c.frame.width)),
		C.double(float64(-c.frame.height)),
		C.double(0.0),
		C.double(1.0),
	)
}

func (c *CommandBuffer) SetRenderPipeline(pipeline hal.RenderPipeline) {
	p, ok := pipeline.(*RenderPipeline)
	if !ok {
		panic("unexpected type")
	}

	C.gfx_mtl_set_render_pipeline(c.renderEncoder, p.pipeline)
}

func (c *CommandBuffer) SetVertexBuffer(data hal.Buffer) {
	d, ok := data.(*Buffer)
	if !ok {
		panic("unexpected type")
	}

	C.gfx_mtl_set_vertex_buffer(c.renderEncoder, d.buffer)
}

func (c *CommandBuffer) Draw(start int, count int) {
	C.gfx_mtl_draw(c.renderEncoder, C.int(start), C.int(count))
}

func (c *CommandBuffer) EndRenderPass() {
	C.gfx_mtl_end_rpass(c.renderEncoder)
}

func (c *CommandBuffer) Submit() {
	C.gfx_mtl_cbuf_submit(c.buffer)
}
