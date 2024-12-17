package utils

var Up = MakePoint(0, 1)
var Down = MakePoint(0, -1)
var Left = MakePoint(-1, 0)
var Right = MakePoint(1, 0)

var All = []Point{
	Up,
	Down,
	Left,
	Right,
}

func Manhattan(a, b Point) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

type Point struct {
	X, Y int
}

func MakePoint(x, y int) Point {
	return Point{x, y}
}

func (p Point) Add(o Point) Point {
	return Point{
		X: p.X + o.X,
		Y: p.Y + o.Y,
	}
}

func (p Point) Sub(o Point) Point {
	return Point{
		X: p.X - o.X,
		Y: p.Y - o.Y,
	}
}

// Rotate does a clockwise rotation of the point around the origin.
func (p Point) Rotate(dir rune) Point {
	q := complex(float64(p.X), float64(p.Y))
	switch dir {
	case 'c':
		q *= -1i
	case 'a':
		q *= 1i
	default:
		panic("can only rotate 'c' or 'a'")
	}

	return Point{
		X: int(real(q)),
		Y: int(imag(q)),
	}
}
