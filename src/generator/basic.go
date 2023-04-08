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
