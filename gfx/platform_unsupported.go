//go:build !darwin

package gfx

import "github.com/csnewman/go-gfx/hal"

func DefaultPlatform() hal.Platform {
	panic("unsupported platform")
}
