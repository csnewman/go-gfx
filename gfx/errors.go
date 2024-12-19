package gfx

/*
#include "vulkan.h"
*/
import "C"

import (
	"errors"
	"fmt"
)

var (
	ErrUnexpectedSystemResponse = errors.New("unexpected system response")
	ErrUnexpectedStatus         = errors.New("unexpected status")
	ErrAlreadyRunning           = errors.New("already running")
	ErrNotMainThread            = errors.New("not on main thread")
	ErrUnsupportedSurfaceHandle = errors.New("unsupported surface handle")
	ErrIncompatibleSurface      = errors.New("incompatible surface")
	ErrFunctionNotFound         = errors.New("function not found")
	ErrIncompatibleDriver       = errors.New("incompatible driver")
	ErrNoSuitableDevice         = errors.New("could not find a suitable device")
	ErrMissingFeature           = errors.New("feature missing")
	ErrInitializationFailed     = errors.New("initialization failed")
)

func mapError(err C.VkResult) error {
	if err >= 0 {
		return nil
	}

	switch err {
	case C.VK_ERROR_INITIALIZATION_FAILED:
		return ErrInitializationFailed

	case C.VK_ERROR_LAYER_NOT_PRESENT:
		return fmt.Errorf("%w: layer", ErrMissingFeature)

	case C.VK_ERROR_EXTENSION_NOT_PRESENT:
		return fmt.Errorf("%w: extension", ErrMissingFeature)

	case C.VK_ERROR_INCOMPATIBLE_DRIVER:
		return ErrIncompatibleDriver

	default:
		return fmt.Errorf("%w: vulkan %v", ErrUnexpectedStatus, err)
	}
}
