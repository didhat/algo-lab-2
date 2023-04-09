package generator

import (
	"lab2/src/structs"
	"math/rand"
)

func GenerateRectangles(rectanglesNumber int) []structs.Rectangle {
	rectangles := make([]structs.Rectangle, 0, rectanglesNumber)

	for i := 1; i <= rectanglesNumber; i++ {
		rectangles = append(rectangles, structs.NewRectangleFromPrimitives(10*i, 10*i, 10*(2*rectanglesNumber-i), 10*(2*rectanglesNumber-i)))
	}
	return rectangles
}

func GeneratePoints(pointsNumber, minX, maxX, minY, maxY int) []structs.Point {
	points := make([]structs.Point, 0, pointsNumber)

	for i := 0; i < pointsNumber; i++ {
		points = append(points, structs.NewPoint(rand.Intn(maxX-minX)+minX, rand.Intn(maxY-minY)+minY))
	}
	return points
}

type BenchmarkTestCase struct {
	Rectangles []structs.Rectangle
	Points     []structs.Point
}

func NewBenchmarkTestCase(recs []structs.Rectangle, points []structs.Point) BenchmarkTestCase {
	return BenchmarkTestCase{Rectangles: recs, Points: points}
}

func generateTestCaseForBenchmarks(rectanglesNumber, pointsNumber int) BenchmarkTestCase {
	recs := GenerateRectangles(rectanglesNumber)
	points := GeneratePoints(pointsNumber, 0, 10*(2*rectanglesNumber), 0, 10*(2*rectanglesNumber))

	return NewBenchmarkTestCase(recs, points)

}

func GenerateManyTestsForBenchMark() []BenchmarkTestCase {
	pointsNumber := 100
	rectanglesNumbers := []int{100, 500, 1000, 2000}

	testCases := make([]BenchmarkTestCase, 0, len(rectanglesNumbers))

	for _, recsNumber := range rectanglesNumbers {
		testCases = append(testCases, generateTestCaseForBenchmarks(recsNumber, pointsNumber))
	}

	return testCases
}
