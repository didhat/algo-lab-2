package structs

type Rectangle struct {
	LeftDown Point
	RightTop Point
}

func NewRectangleFromPrimitives(x1 int, y1 int, x2 int, y2 int) Rectangle {
	left := NewPoint(x1, y1)
	right := NewPoint(x2, y2)
	return Rectangle{LeftDown: left, RightTop: right}
}

func NewRectangleFromPoints(leftDown Point, rightTop Point) Rectangle {
	return Rectangle{LeftDown: leftDown, RightTop: rightTop}
}
