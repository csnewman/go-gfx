//go:build !darwin

package gfx

import (
	"github.com/csnewman/go-gfx/hal"
)

func MetalGraphicsEnabled() bool {
	return false
}

func MetalGraphics() hal.Graphics {
	panic("unsupported platform")
}
