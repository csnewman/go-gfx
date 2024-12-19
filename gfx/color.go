package gfx

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

func NewColor(r float64, g float64, b float64, a float64) Color {
	return Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
