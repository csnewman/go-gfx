//go:build !windows

package gfx

import (
	"github.com/csnewman/go-gfx/hal"
)

func WindowsPlatformEnabled() bool {
	return false
}

func WindowsPlatform() hal.Platform {
	panic("unsupported platform")
}
