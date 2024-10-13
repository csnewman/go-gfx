package hal

type PlatformConfig struct {
	Init                 func() error
	WindowCloseRequested func(w Window)
	WindowClosed         func(w Window)
	WindowRender         func(w Window)
}

type WindowConfig struct {
	Width  int
	Height int
}

type Window uint64

type Platform interface {
	Run(cfg PlatformConfig) error
	Exit()

	NewWindow(cfg WindowConfig) (Window, error)
	CloseWindow(id Window)
}
