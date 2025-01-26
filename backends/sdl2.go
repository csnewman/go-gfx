//go:build windows || linux

package backends

import "github.com/csnewman/go-gfx/backends/sdl2"

func init() {
	defaults = append(
		defaults,
		sdl2.Register,
	)
}
