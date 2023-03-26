package algo

import (
	"lab2/src/structs"
	"testing"
)

var testDataForBasicRecs = []struct {
	name          string
	pointForCheck structs.Point
	expected      int
}{
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
