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

	r := C.gfx_mtl_configure_surface(g.device, layer)

	switch r {
	case C.GFX_SUCCESS:
		return &Surface{
			graphics: g,
			layer:    layer,
		}, nil

	default:
		panic("unexpected response")
	}
}

type Surface struct {
	graphics *Graphics
	layer    C.id
}

func (s *Surface) AcquireTexture(rawTK hal.RenderToken) (hal.SurfaceTexture, error) {
	tk, ok := rawTK.(hal.MetalRenderToken)
	if !ok {
		panic("unexpected render token")
	}

	return &SurfaceTexture{
		graphics: s.graphics,
		drawable: C.id(tk.Drawable),
		texture:  C.gfx_mtl_get_drawable_texture(C.id(tk.Drawable)),
	}, nil
}

type SurfaceTexture struct {
	graphics *Graphics
	drawable C.id
	texture  C.id
}

func (s *SurfaceTexture) Present() error {
	C.gfx_mtl_present_texture(s.graphics.queue, s.drawable)

	return nil
}

func (s *SurfaceTexture) Discard() {
	C.gfx_mtl_discard_surface_texture(s.drawable)
}

func (s *SurfaceTexture) View() hal.TextureView {
	// TODO: ownership

	return &TextureView{
		texture: s.texture,
	}
}

type TextureView struct {
	texture C.id
}

type Shader struct {
	shader C.id
}

func (g *Graphics) CreateShader(cfg hal.ShaderConfig) (hal.Shader, error) {
	var (
		lib      C.id
		errorStr *C.char
	)

	r := C.gfx_mtl_create_shader(
		g.device,
		unsafe.Pointer(unsafe.StringData(cfg.Source)),
		C.int(len(cfg.Source)),
		&lib,
		&errorStr,
	)

	switch r {
	case C.GFX_SUCCESS:
		return &Shader{
			shader: lib,
		}, nil

	case C.GFX_SEE_ERROR:
		defer C.free(unsafe.Pointer(errorStr))

		return nil, errors.New(C.GoString(errorStr))
	default:
		panic("unexpected response")
	}
}

type ShaderFunction struct {
	function C.id
}

func (s *Shader) ResolveFunction(name string) (hal.ShaderFunction, error) {
	var fun C.id

	C.gfx_mtl_get_shader_function(
		s.shader,
		unsafe.Pointer(unsafe.StringData(name)),
		C.int(len(name)),
		&fun,
	)

	if fun == nil {
		return nil, hal.ErrFunctionNotFound
	}

	return &ShaderFunction{
		function: fun,
	}, nil
}

type CommandBuffer struct {
	buffer        C.id
	renderEncoder C.id
}

func (g *Graphics) CreateCommandBuffer() hal.CommandBuffer {
	var buf C.id

	// TODO: synchronise
	C.gfx_mtl_create_command_buf(g.queue, &buf)

	return &CommandBuffer{
		buffer: buf,
	}
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
}

func (c *CommandBuffer) EndRenderPass() {
	C.gfx_mtl_end_rpass(c.renderEncoder)
}

func (c *CommandBuffer) Submit() {
	C.gfx_mtl_cbuf_submit(c.buffer)
}
