//go:build !darwin

package gfx

import (
	"github.com/csnewman/go-gfx/hal"
)

func AppKitPlatformEnabled() bool {
	return false
}

func AppKitPlatform() hal.Platform {
	panic("unsupported platform")
}
