package vulkan

/*
#include "vulkan.h"
*/
import "C"

import (
	"fmt"

	"github.com/csnewman/go-gfx/gfx"
)

func mapError(err C.VkResult) error {
	if err >= 0 {
		return nil
	}

	switch err {
	case C.VK_ERROR_LAYER_NOT_PRESENT:
		return fmt.Errorf("%w: layer", gfx.ErrMissingFeature)

	case C.VK_ERROR_EXTENSION_NOT_PRESENT:
		return fmt.Errorf("%w: extension", gfx.ErrMissingFeature)

	case C.VK_ERROR_INCOMPATIBLE_DRIVER:
		return gfx.ErrIncompatibleDriver

	default:
		return fmt.Errorf("%w: vulkan %v", gfx.ErrUnexpectedStatus, err)
	}
}
