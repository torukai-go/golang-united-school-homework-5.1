package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (sq *Square) End() Point {
	// implement me
	var endPoint Point = Point{x: sq.start.x + int(sq.a), y: sq.start.y + int(sq.a)}
	return endPoint
}

func (sq *Square) Area() uint {
	// implement me
	var area uint = sq.a * sq.a
	return area
}

func (sq *Square) Perimeter() uint {
	// implement me
	var perim uint = sq.a * 4
	return perim
}
