package main

import (
	"math/rand"
	"testing"
)

type testCaseForAlgo struct {
	name          string
	pointForCheck Point
	expected      int
}

func newTestCase(name string, pointForCheck Point, expected int) testCaseForAlgo {
	return testCaseForAlgo{name: name, pointForCheck: pointForCheck, expected: expected}
}

func GenerateRectangles(rectanglesNumber int) []Rectangle {
	rectangles := make([]Rectangle, 0, rectanglesNumber)

	for i := 1; i <= rectanglesNumber; i++ {
		rectangles = append(rectangles, NewRectangleFromPrimitives(10*i, 10*i, 10*(2*rectanglesNumber-i), 10*(2*rectanglesNumber-i)))
	}
	return rectangles
}

func GeneratePoints(pointsNumber, minX, maxX, minY, maxY int) []Point {
	points := make([]Point, 0, pointsNumber)

	for i := 0; i < pointsNumber; i++ {
		points = append(points, NewPoint(rand.Intn(maxX-minX)+minX, rand.Intn(maxY-minY)+minY))
	}
	return points
}

type BenchmarkTestCase struct {
	Rectangles []Rectangle
	Points     []Point
}

func NewBenchmarkTestCase(recs []Rectangle, points []Point) BenchmarkTestCase {
	return BenchmarkTestCase{Rectangles: recs, Points: points}
}

func generateTestCaseForBenchmarks(rectanglesNumber, pointsNumber int) BenchmarkTestCase {
	recs := GenerateRectangles(rectanglesNumber)
	points := GeneratePoints(pointsNumber, 0, 10*(2*rectanglesNumber), 0, 10*(2*rectanglesNumber))

	return NewBenchmarkTestCase(recs, points)

}

type BasicAlgo struct {
	rects []Rectangle
}

func NewBasicAlgo(rects []Rectangle) BasicAlgo {
	return BasicAlgo{rects: rects}
}

func (ba BasicAlgo) Prepare() {}

func (ba BasicAlgo) QueryPoint(point Point) int {
	answer := 0

	for _, rec := range ba.rects {
		if rec.LeftDown.X <= point.X && point.X <= rec.RightTop.X && rec.LeftDown.Y <= point.Y && point.Y <= rec.RightTop.Y {
			answer++
		}
	}
	return answer
}

func GenerateRandomTestCase(recsNumber, pointNumberForCheck int) ([]Rectangle, []testCaseForAlgo) {
	rectangles := GenerateRectangles(recsNumber)
	minX, maxX, minY, maxY := -1000, -10, 901, 10*2*recsNumber

	pointsForCheck := GeneratePoints(pointNumberForCheck, minX, maxX, minY, maxY)

	basicAlgoForGetRightAnswer := NewBasicAlgo(rectangles)
	basicAlgoForGetRightAnswer.Prepare()
	testCases := make([]testCaseForAlgo, 0, pointNumberForCheck)

	for _, point := range pointsForCheck {
		expected := basicAlgoForGetRightAnswer.QueryPoint(point)
		testCases = append(testCases, newTestCase("random", point, expected))
	}

	return rectangles, testCases
}

func TestPersistentTreeAlgo_QueryPoint(t *testing.T) {

	recs, testCases := GenerateRandomTestCase(100, 1000)
	algo := NewPersistentTreeAlgo(recs)
	algo.Prepare()

	for _, d := range testCases {
		t.Run(d.name, func(t *testing.T) {
			result := algo.QueryPoint(d.pointForCheck)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
		})
	}

}
