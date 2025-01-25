package backends

import (
	"github.com/csnewman/go-gfx/backends/appkit"
	"github.com/csnewman/go-gfx/backends/vulkan"
)

func RegisterDefaults() {
	appkit.Register()
	vulkan.Register()
}
