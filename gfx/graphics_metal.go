//go:build darwin

package gfx

import (
	"github.com/csnewman/go-gfx/hal"
	"github.com/csnewman/go-gfx/internal/metal"
)

func DefaultGraphics() hal.Graphics {
	return metal.NewGraphics()
}
