package gfx

import (
	"cmp"
	"log/slog"
	"slices"
)

type Config struct {
	Logger *slog.Logger

	Title  string
	Width  int
	Height int

	Init     func(app *Application) error
	OnRender func(frame *Frame) error
	OnResize func(size LogicalSize) error
}

type Application struct {
	logger   *slog.Logger
	cfg      Config
	platform Platform
	graphics Graphics
	surface  Surface
	size     LogicalSize
}

func Run(cfg Config) error {
	if len(platforms) == 0 {
		return ErrNoPlatforms
	}

	slices.SortFunc(platforms, func(a, b Platform) int {
		return cmp.Compare(b.Priority(), a.Priority())
	})

	var platform Platform

	for _, p := range platforms {
		if !p.Available() {
			cfg.Logger.Debug("Platform unavailable", "name", p.Name())

			continue
		}

		platform = p

		break
	}

	if platform == nil {
		return ErrNoPlatformAvailable
	}

	cfg.Logger.Info("Initialising platform", "name", platform.Name())

	app := &Application{
		logger:   cfg.Logger,
		cfg:      cfg,
		platform: platform,
	}

	return platform.Run(PlatformInit{
		Logger:  cfg.Logger,
		Cfg:     cfg,
		Init:    app.init,
		Render:  app.render,
		Resized: app.resized,
	})
}

func (a *Application) init() error {
	if len(graphics) == 0 {
		return ErrNoGraphics
	}

	slices.SortFunc(graphics, func(a, b Graphics) int {
		return cmp.Compare(b.Priority(), a.Priority())
	})

	var g Graphics

	for _, ga := range graphics {
		if !ga.Available(a.platform) {
			a.logger.Debug("Graphics unavailable", "name", ga.Name())

			continue
		}

		g = ga

		break
	}

	if g == nil {
		return ErrNoGraphicsAvailable
	}

	a.graphics = g

	a.logger.Info("Initialising graphics", "name", g.Name())

	if err := g.Init(GraphicsInit{
		Logger:   a.logger,
		Platform: a.platform,
	}); err != nil {
		return err
	}

	surf := a.platform.PrimarySurface()

	a.size = surf.Size()

	surface, err := g.CreateSurface(surf, a.size.PhysicalSize())
	if err != nil {
		return err
	}

	a.surface = surface

	if err := a.cfg.Init(a); err != nil {
		return err
	}

	return nil
}

type Frame struct {
	sf   SurfaceFrame
	size LogicalSize
}

func (f *Frame) BeginRenderPass(descriptor RenderPassDescriptor) RenderPassEncoder {
	return f.sf.BeginRenderPass(descriptor)
}

func (f *Frame) Size() LogicalSize {
	return f.size
}

func (f *Frame) Index() int {
	return f.sf.Index()
}

func (f *Frame) View() ImageView {
	return f.sf.View()
}

func (a *Application) render() error {
	sf, err := a.surface.Acquire()
	if err != nil {
		panic(err)
	}

	if err := a.cfg.OnRender(&Frame{
		sf:   sf,
		size: a.size,
	}); err != nil {
		panic(err)
	}

	if err := sf.Present(); err != nil {

		return err
	}

	return nil
}

func (a *Application) resized(size LogicalSize) error {
	a.size = size

	if err := a.surface.Resize(size.PhysicalSize()); err != nil {
		return err
	}

	if err := a.cfg.OnResize(size); err != nil {
		return err
	}

	return nil
}

func (a *Application) SurfaceFormat() Format {
	return a.surface.Format()
}

func (a *Application) FrameCount() int {
	return a.surface.FrameCount()
}

func (a *Application) CreateShader(config ShaderConfig) (Shader, error) {
	return a.graphics.CreateShader(config)
}

func (a *Application) CreateRenderPipeline(descriptor RenderPipelineDescriptor) (RenderPipeline, error) {
	return a.graphics.CreateRenderPipeline(descriptor)
}

func (a *Application) CreateBuffer(descriptor BufferDescriptor) (Buffer, error) {
	return a.graphics.CreateBuffer(descriptor)
}

func (a *Application) CreateSampler(descriptor SamplerDescriptor) (Sampler, error) {
	return a.graphics.CreateSampler(descriptor)
}

func (a *Application) CreateImage(descriptor ImageDescriptor) (Image, error) {
	return a.graphics.CreateImage(descriptor)
}

func (a *Application) CreateCommandEncoder() (CommandEncoder, error) {
	return a.graphics.CreateCommandEncoder()
}
