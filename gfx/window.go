package gfx

import (
	"github.com/csnewman/go-gfx/hal"
)

type WindowConfig struct {
	Title            string
	Width            int
	Height           int
	OnCloseRequested func()
	OnClosed         func()
	OnRender         func()
}

type Window struct {
	app *Application
	id  hal.Window
	cfg WindowConfig
}

func (a *Application) NewWindow(cfg WindowConfig) (*Window, error) {
	w := &Window{
		app: a,
		cfg: cfg,
	}

	id, err := a.platform.NewWindow(hal.WindowConfig{
		Width:  cfg.Width,
		Height: cfg.Height,
	})
	if err != nil {
		return nil, err
	}

	w.id = id

	a.windows.Set(id, w)

	return w, nil
}

func (a *Application) windowCloseRequested(id hal.Window) {
	w, ok := a.windows.Get(id)
	if !ok {
		return
	}

	if w.cfg.OnCloseRequested != nil {
		w.cfg.OnCloseRequested()

		return
	}

	w.Close()
}

func (a *Application) windowClosed(id hal.Window) {
	w, ok := a.windows.Remove(id)
	if !ok {
		return
	}

	if w.cfg.OnClosed != nil {
		w.cfg.OnClosed()
	}

	// TODO: auto exit app?
}

func (w *Window) Close() {
	w.app.platform.CloseWindow(w.id)
}

func (a *Application) windowRender(id hal.Window) {
	w, ok := a.windows.Get(id)
	if !ok {
		return
	}

	if w.cfg.OnRender != nil {
		w.cfg.OnRender()
	}
}
