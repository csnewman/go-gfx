package gfx

import (
	"fmt"
	"github.com/csnewman/go-gfx/hal"
)

type WindowConfig struct {
	Title            string
	Width            int
	Height           int
	OnCloseRequested func()
	OnClosed         func()
	OnRender         func(f *Frame)
	OnResize         func(width float64, height float64)
}

type Window struct {
	app     *Application
	id      hal.Window
	cfg     WindowConfig
	surface hal.Surface

	surfaceWidth  float64
	width         float64
	surfaceHeight float64
	height        float64
}

func (a *Application) NewWindow(cfg WindowConfig) (*Window, error) {
	w := &Window{
		app: a,
		cfg: cfg,
	}

	id, wh, err := a.platform.NewWindow(hal.WindowConfig{
		Title:  cfg.Title,
		Width:  cfg.Width,
		Height: cfg.Height,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create window: %w", err)
	}

	w.id = id

	a.windows.Set(id, w)

	w.surface, err = a.graphics.CreateSurface(wh)
	if err != nil {
		return nil, fmt.Errorf("failed to create surface: %w", err)
	}

	w.width = float64(cfg.Width)
	w.surfaceWidth = float64(cfg.Width)
	w.height = float64(cfg.Height)
	w.surfaceHeight = float64(cfg.Height)

	return w, nil
}

func (w *Window) TextureFormat() TextureFormat {
	return w.surface.TextureFormat()
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

func (w *Window) Start() {
	go w.thread()
}

func (w *Window) thread() {
	for {
		if w.width != w.surfaceWidth || w.height != w.surfaceHeight {
			if err := w.surface.Resize(int(w.width), int(w.height)); err != nil {
				panic(err)
			}

			w.surfaceWidth = w.width
			w.surfaceHeight = w.height

			if w.cfg.OnResize != nil {
				w.cfg.OnResize(w.width, w.height)
			}
		}

		frame, err := w.surface.Acquire()
		if err != nil {
			// TODO: handle error
			panic(err)
		}

		w.cfg.OnRender(&Frame{
			app:   w.app,
			frame: frame,
		})
	}
}

func (a *Application) windowResized(id hal.Window, width float64, height float64) {
	w, ok := a.windows.Get(id)
	if !ok {
		return
	}

	// TODO: fix race
	w.width = width
	w.height = height
}
