package autogfx

import (
	"github.com/csnewman/go-gfx/gfx"
	"log/slog"
)

type Config struct {
	Logger *slog.Logger

	Title  string
	Width  int
	Height int

	Init      func(g gfx.Graphics, s gfx.Surface, size gfx.LogicalSize) error
	OnRender  func(frame gfx.SurfaceFrame) error
	OnResize  func(size gfx.LogicalSize) error
	OnExiting func() error
}

func Run(cfg Config) error {
	return platformRun(cfg)
}
