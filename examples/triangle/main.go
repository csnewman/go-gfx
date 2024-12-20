package main

import (
	_ "embed"
	"fmt"
	"github.com/csnewman/go-gfx/gfx"
	"github.com/csnewman/go-gfx/sdl2"
	"log"
	"log/slog"
	"os"
	"runtime"
	"time"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

var lastPrint time.Time
var count int

//go:embed shader.spv
var mainShader []byte

type Example struct {
	logger           *slog.Logger
	platform         *sdl2.Platform
	window           *sdl2.Window
	graphics         *gfx.Graphics
	surface          *gfx.Surface
	shader           *gfx.Shader
	vertexFunction   *gfx.ShaderFunction
	fragmentFunction *gfx.ShaderFunction
	trianglePipeline *gfx.RenderPipeline
	vertData         *gfx.Buffer
}

func (e *Example) init() error {
	var err error

	e.shader, err = e.graphics.CreateShader(gfx.ShaderConfig{
		SPIRV: mainShader,
	})
	if err != nil {
		return fmt.Errorf("error loading shader: %w", err)
	}

	e.vertexFunction, err = e.shader.Function("vertexShader")
	if err != nil {
		return fmt.Errorf("error creating vertex function: %w", err)
	}

	e.fragmentFunction, err = e.shader.Function("fragmentShader")
	if err != nil {
		return fmt.Errorf("error creating fragment function: %w", err)
	}

	e.trianglePipeline, err = e.graphics.CreateRenderPipeline(gfx.RenderPipelineDescriptor{
		VertexFunction: e.vertexFunction,
		VertexBindings: []gfx.VertexBinding{
			{
				Binding: 0,
				Stride:  16,
				Rate:    gfx.VertexRateVertex,
				Attributes: []gfx.VertexAttribute{
					{
						Location: 0,
						Offset:   0,
					},
				},
			},
		},
		FragmentFunction: e.fragmentFunction,
		ColorAttachments: []gfx.RenderPipelineColorAttachment{
			{
				Format: e.surface.Format(),
			},
		},
	})
	if err != nil {
		return fmt.Errorf("error creating rendering pipeline: %w", err)
	}

	floatData := []float32{
		0.0, -0.5, 0.0, 0,
		0.5, 0.5, 0.0, 0,
		-0.5, 0.5, 0.0, 0,
	}
	byteData := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(floatData))), len(floatData)*4)
	e.vertData = e.graphics.CreateBuffer(byteData)

	return nil
}

func (e *Example) Run() error {
	var err error

	if e.platform, err = sdl2.Init(sdl2.Config{
		Logger: e.logger,
		Init:   e.init,
	}); err != nil {
		return fmt.Errorf("failed to initialise platform: %w", err)
	}

	if e.graphics, err = gfx.Init(gfx.Config{
		Logger:   e.logger,
		Platform: e.platform,
	}); err != nil {
		return err
	}

	e.logger.Info("Creating main window")

	if e.window, err = e.platform.NewWindow(sdl2.WindowConfig{
		Title:  "Triangle | go-gfx",
		Width:  800,
		Height: 600,
		//OnClosed: e.closed,
		OnRender: e.render,
		OnResize: e.resize,
	}); err != nil {
		return fmt.Errorf("failed to create main window: %w", err)
	}

	if e.surface, err = e.graphics.CreateSurface(e.window); err != nil {
		return fmt.Errorf("failed to create surface: %w", err)
	}

	return e.platform.Run()
}

func (e *Example) closed() {
	log.Println("Main window closed")

	e.platform.Exit()
}

func (e *Example) resize(width float64, height float64) {
	log.Println("Main window resized", width, height)
}

func (e *Example) render() {
	frame, err := e.surface.Acquire()
	if err != nil {
		panic(err)
	}

	count++

	if time.Since(lastPrint) >= time.Second {
		lastPrint = time.Now()

		e.logger.Info("FPS", "fps", count)

		count = 0
	}

	buffer := frame.CreateCommandBuffer()

	buffer.Barrier(gfx.Barrier{
		Images: []gfx.ImageBarrier{
			{
				Image:     frame.Image(),
				SrcLayout: gfx.ImageLayoutUndefined,
				DstLayout: gfx.ImageLayoutAttachment,
			},
		},
	})

	buffer.BeginRenderPass(gfx.RenderPassDescriptor{
		ColorAttachments: []gfx.RenderPassColorAttachment{
			{
				Target:     frame,
				Load:       false,
				ClearColor: gfx.NewColor(0, 1, 0, 1),
				Discard:    false,
			},
		},
	})

	buffer.SetRenderPipeline(e.trianglePipeline)
	buffer.SetVertexBuffer(0, e.vertData, 0)
	buffer.Draw(0, 3)

	buffer.EndRenderPass()

	buffer.Barrier(gfx.Barrier{
		Images: []gfx.ImageBarrier{
			{
				Image:     frame.Image(),
				SrcLayout: gfx.ImageLayoutAttachment,
				DstLayout: gfx.ImageLayoutPresent,
			},
		},
	})

	buffer.Submit()

	if err := frame.Present(); err != nil {
		panic(err)
	}
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}))

	logger.Info("Triangle example")

	ex := &Example{
		logger: logger,
	}

	if err := ex.Run(); err != nil {
		panic(err)
	}

	logger.Info("App exited")
}
