package backends

import (
	"github.com/csnewman/go-gfx/backends/vulkan"
)

var defaults []func()

func init() {
	defaults = append(
		defaults,
		vulkan.Register,
	)
}

func RegisterDefaults() {
	for _, f := range defaults {
		f()
	}
}
