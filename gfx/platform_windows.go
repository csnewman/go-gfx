//go:build windows

package gfx

import (
	"github.com/csnewman/go-gfx/hal"
	"github.com/csnewman/go-gfx/internal/windows"
)

func DefaultPlatform() hal.Platform {
	return windows.NewPlatform()
}
