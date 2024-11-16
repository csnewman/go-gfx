//go:build darwin || windows

package gfx

import (
	"github.com/csnewman/go-gfx/hal"
	"github.com/csnewman/go-gfx/internal/vulkan"
)

func VulkanGraphicsEnabled() bool {
	return true
}

func VulkanGraphics() hal.Graphics {
	return vulkan.NewGraphics()
}
