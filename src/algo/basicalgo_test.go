package algo

import (
	"fmt"
	"lab2/src/generator"
	"lab2/src/structs"
	"testing"
)

type testCaseForAlgo struct {
	name          string
	pointForCheck structs.Point
	expected      int
}

func newTestCase(name string, pointForCheck structs.Point, expected int) testCaseForAlgo {
	return testCaseForAlgo{name: name, pointForCheck: pointForCheck, expected: expected}
}

var testDataForBasicRecs = []testCaseForAlgo{
	{"first", structs.NewPoint(2, 2), 1},
	{"second", structs.NewPoint(12, 12), 1},
	{"third", structs.NewPoint(10, 4), 2},
	{"fourth", structs.NewPoint(5, 5), 3},
	{"fifth", structs.NewPoint(2, 10), 0},
}

func getBasicRecs() []structs.Rectangle {
	rec1 := structs.NewRectangleFromPrimitives(2, 2, 6, 8)
	rec2 := structs.NewRectangleFromPrimitives(5, 4, 9, 10)
	rec3 := structs.NewRectangleFromPrimitives(4, 0, 11, 6)
	rec4 := structs.NewRectangleFromPrimitives(8, 2, 12, 12)
	recs := []structs.Rectangle{rec1, rec2, rec3, rec4}
	return recs
}

func TestBasicAlgo_Prepare(t *testing.T) {
	var mock []structs.Rectangle
	algo := NewBasicAlgo(mock)

	algo.Prepare()
}

func TestBasicAlgo_QueryPoint(t *testing.T) {
	recs := getBasicRecs()

	algo := NewBasicAlgo(recs)

	for _, d := range testDataForBasicRecs {
		t.Run(d.name, func(t *testing.T) {
			result := algo.QueryPoint(d.pointForCheck)
			if result != d.expected {
				t.Errorf("Excepted %d, got %d", d.expected, result)
			}
		})
	}

	res := algo.QueryPoint(structs.Point{X: 2, Y: 2})
	if res != 1 {
		t.Error("error: excepted: 1, get:", res)
	}
}

func GenerateRandomTestCase(recsNumber, pointNumberForCheck int) ([]structs.Rectangle, []testCaseForAlgo) {
	rectangles := generator.GenerateRectangles(recsNumber)
	minX, maxX, minY, maxY := 0, 10*2*recsNumber, 0, 10*2*recsNumber

	pointsForCheck := generator.GeneratePoints(pointNumberForCheck, minX, maxX, minY, maxY)

	basicAlgoForGetRightAnswer := NewBasicAlgo(rectangles)
	basicAlgoForGetRightAnswer.Prepare()
	testCases := make([]testCaseForAlgo, 0, pointNumberForCheck)

	for _, point := range pointsForCheck {
		expected := basicAlgoForGetRightAnswer.QueryPoint(point)
		testCases = append(testCases, newTestCase("random", point, expected))
	}

	return rectangles, testCases
}

func BenchmarkBasicAlgo_QueryPoint(b *testing.B) {
	benchMarksTests := generator.GenerateManyTestsForBenchMark()

	for _, v := range benchMarksTests {
		algo := NewBasicAlgo(v.Rectangles)
		b.Run(fmt.Sprintf("Recs:%d", len(v.Rectangles)), func(b *testing.B) {
			for _, point := range v.Points {
				algo.QueryPoint(point)
			}
		})
	}

}
