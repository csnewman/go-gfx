package gfx

import "github.com/csnewman/go-gfx/hal"

type ApplicationConfig struct {
	Init func(app *Application) error
}

type Application struct {
	platform hal.Platform
	cfg      ApplicationConfig
	windows  tmap[hal.Window, *Window]
}

func Run(cfg ApplicationConfig) error {
	return RunWith(cfg, DefaultPlatform())
}

func RunWith(cfg ApplicationConfig, platform hal.Platform) error {
	app := &Application{
		platform: platform,
		cfg:      cfg,
	}

	return app.run()
}

func (a *Application) run() error {
	return a.platform.Run(hal.PlatformConfig{
		Init:                 a.init,
		WindowCloseRequested: a.windowCloseRequested,
		WindowClosed:         a.windowClosed,
		WindowRender:         a.windowRender,
	})
}

func (a *Application) init() error {
	if a.cfg.Init == nil {
		return nil
	}

	return a.cfg.Init(a)
}

func (a *Application) Exit() {
	a.platform.Exit()

	// TODO: cleanup?
}
