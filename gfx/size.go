package gfx

type LogicalSize struct {
	Width  float64
	Height float64
	Scale  float64
}

func (l LogicalSize) PhysicalSize() PhysicalSize {
	return PhysicalSize{
		Width:  int(l.Width * l.Scale),
		Height: int(l.Height * l.Scale),
	}
}

type PhysicalSize struct {
	Width  int
	Height int
}
