package gfx

import (
	"errors"
)

var (
	ErrInvalidDescriptor        = errors.New("invalid descriptor")
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
