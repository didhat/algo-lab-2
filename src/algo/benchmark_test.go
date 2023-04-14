package algo

import (
	"fmt"
	"lab2/src/generator"
	"testing"
)

func BenchmarkAllAlgoPreparing(b *testing.B) {
	testCases := generator.GenerateManyTestsForBenchMark()

	for _, v := range testCases {
		mapAlgo := NewMapAlgo(v.Rectangles)
		treeAlgo := NewPersistentTreeAlgo(v.Rectangles)

		b.Run(fmt.Sprintf("mapAlgo:Prepare:%d", len(v.Rectangles)), func(b *testing.B) {
			mapAlgo.Prepare()
		})

		b.Run(fmt.Sprintf("treeAlgo:Prepare:%d", len(v.Rectangles)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				treeAlgo.Prepare()
			}
		})

	}

}

var blackhole int

func BenchmarkAllAlgoQueryPoint(b *testing.B) {
	testCases := generator.GenerateManyTestsForBenchMark()

	for _, v := range testCases {
		basicAlgo := NewBasicAlgo(v.Rectangles)
		mapAlgo := NewMapAlgo(v.Rectangles)
		treeAlgo := NewPersistentTreeAlgo(v.Rectangles)

		basicAlgo.Prepare()
		treeAlgo.Prepare()

		//if len(v.Rectangles) <= 500 {
		mapAlgo.Prepare()
		//}

		b.Run(fmt.Sprintf("basicAlgo:Query:%d", len(v.Rectangles)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, point := range v.Points {
					blackhole = basicAlgo.QueryPoint(point)
				}
			}
		})

		//if len(v.Rectangles) <= 500 {
		b.Run(fmt.Sprintf("mapAlgo:Query:%d", len(v.Rectangles)), func(b *testing.B) {
			for _, point := range v.Points {
				blackhole = mapAlgo.QueryPoint(point)
			}
		})
		//}

		b.Run(fmt.Sprintf("treeAlgo:Query:%d", len(v.Rectangles)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, point := range v.Points {
					blackhole = treeAlgo.QueryPoint(point)
				}
			}
		})

	}

}
