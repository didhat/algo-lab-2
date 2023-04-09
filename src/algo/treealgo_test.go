package algo

import (
	"fmt"
	"lab2/src/generator"
	"lab2/src/structs"
	"testing"
)

func TestPersistentTreeAlgo_Prepare(t *testing.T) {
	recs := getBasicRecs()
	algo := NewPersistentTreeAlgo(recs)

	algo.Prepare()

	answ := algo.QueryPoint(structs.NewPoint(2, 2))

	if answ != 1 {
		t.Error("error")
	}

}

func TestPersistentTreeAlgo_QueryPoint(t *testing.T) {
	recs := getBasicRecs()
	algo := NewPersistentTreeAlgo(recs)
	algo.Prepare()

	for _, d := range testDataForBasicRecs {
		t.Run(d.name, func(t *testing.T) {
			result := algo.QueryPoint(d.pointForCheck)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
		})
	}
}

func TestPersistentTreeAlgo_QueryPointWithRandomTestCases(t *testing.T) {
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

func BenchmarkPersistentTreeAlgo_Prepare(b *testing.B) {
	benchmarksTests := generator.GenerateManyTestsForBenchMark()

	for _, v := range benchmarksTests {
		b.Run(fmt.Sprintf("SegTreeAlgo:%d", len(v.Rectangles)), func(b *testing.B) {
			algo := NewPersistentTreeAlgo(v.Rectangles)
			algo.Prepare()
		})
	}

}

func BenchmarkPersistentTreeAlgo_QueryPoint(b *testing.B) {
	benchmarksTests := generator.GenerateManyTestsForBenchMark()

	for _, v := range benchmarksTests {
		algo := NewPersistentTreeAlgo(v.Rectangles)
		algo.Prepare()
		b.Run(fmt.Sprintf("SegTreeAlgo_Query:%d", len(v.Rectangles)), func(b *testing.B) {
			for _, point := range v.Points {
				algo.QueryPoint(point)
			}
		})
	}

}
