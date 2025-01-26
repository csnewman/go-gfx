//go:build windows

package backends

import (
	"github.com/csnewman/go-gfx/backends/windows"
)

func init() {
	defaults = append(
		defaults,
		windows.Register,
	)
}
