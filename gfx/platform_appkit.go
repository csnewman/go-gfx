//go:build darwin

package gfx

import (
	"github.com/csnewman/go-gfx/hal"
	"github.com/csnewman/go-gfx/internal/appkit"
)

func DefaultPlatform() hal.Platform {
	return appkit.NewPlatform()
}
