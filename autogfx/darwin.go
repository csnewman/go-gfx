package autogfx

import (
	"github.com/csnewman/go-gfx/appkit"
	"github.com/csnewman/go-gfx/gfx"
	"github.com/csnewman/go-gfx/vulkan"
	"sync/atomic"
)

func platformRun(cfg Config) error {
	d := &darwin{
		cfg: cfg,
	}

	plt, err := appkit.Init(appkit.Config{
		LoadVulkan: true,
		Init:       d.init,
	})
	if err != nil {
		return err
	}

	d.plt = plt

	return plt.Run()
}

type resizeData struct {
	Width  float64
	Height float64
	Scale  float64
}

type darwin struct {
	cfg       Config
	plt       *appkit.Platform
	window    *appkit.Window
	resizeEvt atomic.Pointer[resizeData]
	closing   atomic.Bool
}

func (d *darwin) init() error {
	window, err := d.plt.NewWindow(appkit.WindowConfig{
		Title:            d.cfg.Title,
		Width:            d.cfg.Width,
		Height:           d.cfg.Height,
		OnResize:         d.resized,
		OnCloseRequested: d.closeRequested,
	})
	if err != nil {
		return err
	}

	d.window = window

	size := window.Size()

	vk, err := vulkan.Init(vulkan.Config{
		Logger:   d.cfg.Logger,
		Platform: d.plt,
	})
	if err != nil {
		return err
	}

	surface, err := vk.CreateSurface(window, size.PhysicalSize())
	if err != nil {
		return err
	}

	if err := d.cfg.Init(vk, surface, size); err != nil {
		return err
	}

	go func() {
		for {
			if d.closing.Swap(false) {
				if err := d.cfg.OnExiting(); err != nil {
					panic(err)
				}

				// TODO: free gfx

				d.plt.Exit()
				break
			}

			evt := d.resizeEvt.Swap(nil)
			if evt != nil {
				if err := d.cfg.OnResize(gfx.LogicalSize{
					Width:  evt.Width,
					Height: evt.Height,
					Scale:  evt.Scale,
				}); err != nil {
					panic(err)
				}
			}

			sf, err := surface.Acquire()
			if err != nil {
				panic(err)
			}

			if err := d.cfg.OnRender(sf); err != nil {
				panic(err)
			}
		}
	}()

	return nil
}

func (d *darwin) resized(width float64, height float64, scale float64) {
	d.resizeEvt.Store(&resizeData{
		Width:  width,
		Height: height,
		Scale:  scale,
	})
}

func (d *darwin) closeRequested() {
	d.closing.Store(true)
}
