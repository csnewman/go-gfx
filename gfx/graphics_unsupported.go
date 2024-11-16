//go:build !darwin && !windows

package gfx

import "github.com/csnewman/go-gfx/hal"

func DefaultGraphics() hal.Graphics {
	panic("unsupported platform")
}
