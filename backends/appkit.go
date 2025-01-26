//go:build darwin

package backends

import (
	"github.com/csnewman/go-gfx/backends/appkit"
)

func init() {
	defaults = append(
		defaults,
		appkit.Register,
	)
}
