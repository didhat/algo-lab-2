package algo

import (
	"testing"
)

func TestMapAlgo_Prepare(t *testing.T) {
	recs := getBasicRecs()
	algo := NewMapAlgo(recs)

	algo.Prepare()

}

func TestMapAlgo_QueryPoint(t *testing.T) {
	recs := getBasicRecs()
	algo := NewMapAlgo(recs)
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

func TestMapAlgo_QueryPointWithRandomTestCase(t *testing.T) {
	recs, testCases := GenerateRandomTestCase(100, 100)
	algo := NewMapAlgo(recs)
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

//func BenchmarkMapAlgo_Prepare(b *testing.B) {
//	benchMarkTests := generator.GenerateManyTestsForBenchMark()
//
//	for _, v := range benchMarkTests {
//		b.Run(fmt.Sprintf("MapAlgo:%d", len(v.Rectangles)), func(b *testing.B) {
//			algo := NewMapAlgo(v.Rectangles)
//			algo.Prepare()
//		})
//	}
//
//}
//
//func BenchmarkMapAlgo_QueryPoint(b *testing.B) {
//	benchmarkTests := generator.GenerateManyTestsForBenchMark()
//
//	for _, v := range benchmarkTests {
//		algo := NewMapAlgo(v.Rectangles)
//		algo.Prepare()
//		b.Run(fmt.Sprintf("MapAlgo:%d", len(v.Rectangles)), func(b *testing.B) {
//			for _, point := range v.Points {
//				algo.QueryPoint(point)
//			}
//		})
//	}

//}
