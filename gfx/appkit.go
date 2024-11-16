//go:build darwin

package gfx

import (
	"github.com/csnewman/go-gfx/hal"
	"github.com/csnewman/go-gfx/internal/appkit"
)

func AppKitPlatformEnabled() bool {
	return true
}

func AppKitPlatform() hal.Platform {
	return appkit.NewPlatform()
}
