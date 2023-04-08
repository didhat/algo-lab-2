package generator

import (
	"testing"
)

func TestGenerateRectangles(t *testing.T) {
	numberRectangles := 10
	recs := GenerateRectangles(numberRectangles)

	if len(recs) != numberRectangles {
		t.Errorf("Excepted %d rectangles, but got %d", numberRectangles, len(recs))
	}

}

func TestGeneratePoints(t *testing.T) {
	numberPoints := 10
	minX, maxX, minY, maxY := 0, 10, 0, 10

	points := GeneratePoints(numberPoints, minX, maxX, minY, maxY)

	if len(points) != numberPoints {
		t.Errorf("Excpeted %d points, but got %d", numberPoints, len(points))
	}

	for _, point := range points {
		if point.X > maxX || point.X < minX {
			t.Error("point X greater or less than must be!")
		}
		if point.Y > maxY || point.Y < minY {
			t.Error("point Y greater or less than must be")
		}
	}

}

