package hal

import "errors"

var (
	ErrAlreadyRunning          = errors.New("already running")
	ErrNotMainThread           = errors.New("not on main thread")
	ErrUnsupportedWindowHandle = errors.New("unsupported window handle")
	ErrFunctionNotFound        = errors.New("function not found")
)
