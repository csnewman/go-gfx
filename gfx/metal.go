//go:build darwin

package gfx

import (
	"github.com/csnewman/go-gfx/hal"
	"github.com/csnewman/go-gfx/internal/metal"
)

func MetalGraphicsEnabled() bool {
	return true
}

func MetalGraphics() hal.Graphics {
	return metal.NewGraphics()
}
