//go:build !darwin && !windows

package gfx

import "github.com/csnewman/go-gfx/hal"

func DefaultPlatform() hal.Platform {
	panic("unsupported platform")
}
