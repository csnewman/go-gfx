package main

import (
	_ "embed"
	"github.com/csnewman/go-gfx/gfx"
	"log"
	"runtime"
	"time"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

var lastPrint time.Time
var count int

//go:embed shader.metal
var metalShader string

//go:embed shader.spv
var mainShader []byte

type Example struct {
	app              *gfx.Application
	window           *gfx.Window
	shader           *gfx.Shader
	vertexFunction   *gfx.ShaderFunction
	fragmentFunction *gfx.ShaderFunction
	vertData         *gfx.Buffer
	trianglePipeline *gfx.RenderPipeline
}

func (e *Example) init(app *gfx.Application) error {
	e.app = app

	log.Println("Creating main window")

	window, err := e.app.NewWindow(gfx.WindowConfig{
		Title:    "Triangle",
		Width:    800,
		Height:   600,
		OnClosed: e.closed,
		OnRender: e.render,
		OnResize: e.resize,
	})
	if err != nil {
		return err
	}

	e.window = window

	log.Println("init complete")

	e.shader, err = e.app.LoadShader(gfx.ShaderConfig{
		Source: metalShader,
		Code:   mainShader,
	})
	if err != nil {
		return err
	}

	e.vertexFunction, err = e.shader.Function("vertexShader")
	if err != nil {
		return err
	}

	e.fragmentFunction, err = e.shader.Function("fragmentShader")
	if err != nil {
		return err
	}

	e.trianglePipeline, err = e.app.NewRenderPipeline(gfx.RenderPipelineDescriptor{
		VertexFunction:   e.vertexFunction,
		FragmentFunction: e.fragmentFunction,
		ColorAttachments: []gfx.RenderPipelineColorAttachment{
			{
				Format: e.window.TextureFormat(),
			},
		},
	})
	if err != nil {
		return err
	}

	floatData := []float32{
		-0.5, -0.5, 0.0, 0,
		0.5, -0.5, 0.0, 0,
		0.0, 0.5, 0.0, 0,
	}
	byteData := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(floatData))), len(floatData)*4)
	e.vertData = e.app.NewBuffer(byteData)

	e.window.Start()

	return nil
}

func (e *Example) Run() error {
	return gfx.Run(gfx.ApplicationConfig{
		Init: e.init,
	})
}

func (e *Example) closed() {
	log.Println("Main window closed")

	e.app.Exit()
}

func (e *Example) resize(width float64, height float64) {
	log.Println("Main window resized", width, height)
}

func (e *Example) render(frame *gfx.Frame) {
	defer frame.Close()

	count++

	if time.Since(lastPrint) >= time.Second {
		lastPrint = time.Now()

		log.Println("FPS", count)
		count = 0
	}

	frame.QueueRenderPass(gfx.RenderPassDescriptor{
		ColorAttachments: []gfx.RenderPassColorAttachment{
			{
				Target:     frame,
				Load:       false,
				ClearColor: gfx.NewColor(1, 0, 0, 1),
				Discard:    false,
			},
		},
		Body: func(enc *gfx.RenderPassEncoder) {
			enc.SetPipeline(e.trianglePipeline)
			//enc.SetVertexBuffer(e.vertData)
			enc.Draw(0, 3)
		},
	})

	if err := frame.Present(); err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Triangle Example")

	ex := &Example{}

	if err := ex.Run(); err != nil {
		panic(err)
	}

	log.Println("App exited")
}
