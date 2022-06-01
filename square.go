package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (square Square) End() Point {
	// implement me
	var p Point = Point{x: square.start.x + int(square.a), y: square.start.y + int(square.a)}
	return p
}

func (square Square) Area() uint {
	// implement me
	return square.a * square.a
}

func (square Square) Perimeter() uint {
	return square.a * 4

}
type Point struct {
	x, y int
}
