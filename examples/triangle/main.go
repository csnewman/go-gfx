package main

import (
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"
	"unsafe"

	"github.com/csnewman/go-gfx/backends"
	"github.com/csnewman/go-gfx/gfx"
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
	app              *gfx.Application
	shader           gfx.Shader
	vertexFunction   gfx.ShaderFunction
	fragmentFunction gfx.ShaderFunction
	trianglePipeline gfx.RenderPipeline
	vertData         gfx.Buffer
}

func (e *Example) init(app *gfx.Application) error {
	e.app = app

	var err error

	e.shader, err = app.CreateShader(gfx.ShaderConfig{
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

	e.trianglePipeline, err = app.CreateRenderPipeline(gfx.RenderPipelineDescriptor{
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
						Format:   gfx.FormatRGB32SFloat,
					},
				},
			},
		},
		FragmentFunction: e.fragmentFunction,
		ColorAttachments: []gfx.RenderPipelineColorAttachment{
			{
				Format: app.SurfaceFormat(),
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

	e.vertData, err = app.CreateBuffer(gfx.BufferDescriptor{
		Size:  len(byteData),
		Usage: gfx.BufferUsageHostUpload | gfx.BufferUsagePersistentMap,
	})
	if err != nil {
		return fmt.Errorf("error creating buffer: %w", err)
	}

	if err := e.vertData.CopyFrom(byteData); err != nil {
		return fmt.Errorf("error creating buffer: %w", err)
	}

	if err := e.vertData.Flush(); err != nil {
		return fmt.Errorf("error creating buffer: %w", err)
	}

	return nil
}

func (e *Example) resize(size gfx.LogicalSize) error {
	e.logger.Info("Window resized", "size", size)

	return nil
}

func (e *Example) render(frame *gfx.Frame) error {
	count++

	if time.Since(lastPrint) >= time.Second {
		lastPrint = time.Now()

		e.logger.Info("FPS", "fps", count)

		count = 0
	}

	encoder := frame.BeginRenderPass(gfx.RenderPassDescriptor{
		ColorAttachments: []gfx.RenderPassColorAttachment{
			{
				Target:     frame.View(),
				Load:       false,
				ClearColor: gfx.NewColor(0, 1, 0, 1),
				Discard:    false,
			},
		},
	})

	encoder.SetRenderPipeline(e.trianglePipeline)
	encoder.SetVertexBuffer(0, e.vertData, 0)
	encoder.Draw(0, 3)

	encoder.End()

	return nil
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

	backends.RegisterDefaults()

	if err := gfx.Run(gfx.Config{
		Logger: logger,
		Title:  "Triangle | go-gfx",
		Width:  800,
		Height: 600,

		Init:     ex.init,
		OnRender: ex.render,
		OnResize: ex.resize,
	}); err != nil {
		panic(err)
	}

	logger.Info("App exited")
}
