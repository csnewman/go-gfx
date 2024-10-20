package main

import (
	"github.com/csnewman/go-gfx/gfx"
	"log"
	"runtime"
	"time"

	_ "embed"
)

func init() {
	runtime.LockOSThread()
}

var lastPrint time.Time
var count int

//go:embed shader.metal
var metalShader string

type Example struct {
	app              *gfx.Application
	window           *gfx.Window
	shader           *gfx.Shader
	vertexFunction   *gfx.ShaderFunction
	fragmentFunction *gfx.ShaderFunction
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

	if time.Since(lastPrint) > time.Second {
		lastPrint = time.Now()

		log.Println("FPS", count)
		count = 0
	}

	buf := frame.NewCommandBuffer()

	rp := buf.BeginRenderPass(gfx.RenderPassDescriptor{
		ColorAttachments: []gfx.ColorAttachment{
			{
				Target:     frame,
				Load:       false,
				ClearColor: gfx.NewColor(1, 0, 1, 1),
				Discard:    false,
			},
		},
	})

	rp.End()

	buf.Submit()

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
