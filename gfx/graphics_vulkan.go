//go:build windows

package gfx

import (
	"github.com/csnewman/go-gfx/hal"
	"github.com/csnewman/go-gfx/internal/vulkan"
)

func DefaultGraphics() hal.Graphics {
	return vulkan.NewGraphics()
}
