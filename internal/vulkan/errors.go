package vulkan

/*
#include "vulkan.h"
*/
import "C"
import (
	"fmt"
	"github.com/csnewman/go-gfx/hal"
)

func mapError(err C.VkResult) error {
	if err >= 0 {
		return nil
	}

	switch err {
	case C.VK_ERROR_EXTENSION_NOT_PRESENT:
		return hal.ErrMissingFeature

	case C.VK_ERROR_INCOMPATIBLE_DRIVER:
		return hal.ErrIncompatibleDriver

	default:
		return fmt.Errorf("%w: vulkan %v", hal.ErrUnexpectedStatus, err)
	}
}
