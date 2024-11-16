//go:build !darwin && !windows

package gfx

import (
	"github.com/csnewman/go-gfx/hal"
)

func VulkanGraphicsEnabled() bool {
	return false
}

func VulkanGraphics() hal.Graphics {
	panic("unsupported platform")
}
