//go:build windows

package gfx

import (
	"github.com/csnewman/go-gfx/hal"
	"github.com/csnewman/go-gfx/internal/windows"
)

func WindowsPlatformEnabled() bool {
	return true
}

func WindowsPlatform() hal.Platform {
	return windows.NewPlatform()
}
