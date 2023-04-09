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
			treeAlgo.Prepare()
		})

	}

}

func BenchmarkAllAlgoQueryPoint(b *testing.B) {
	testCases := generator.GenerateManyTestsForBenchMark()

	for _, v := range testCases {
		basicAlgo := NewBasicAlgo(v.Rectangles)
		mapAlgo := NewMapAlgo(v.Rectangles)
		treeAlgo := NewPersistentTreeAlgo(v.Rectangles)

		basicAlgo.Prepare()
		mapAlgo.Prepare()
		treeAlgo.Prepare()

		b.Run(fmt.Sprintf("basicAlgo:Query:%d", len(v.Rectangles)), func(b *testing.B) {
			for _, point := range v.Points {
				basicAlgo.QueryPoint(point)
			}
		})

		b.Run(fmt.Sprintf("mapAlgo:Query:%d", len(v.Rectangles)), func(b *testing.B) {
			for _, point := range v.Points {
				mapAlgo.QueryPoint(point)
			}
		})

		b.Run(fmt.Sprintf("treeAlgo:Query:%d", len(v.Rectangles)), func(b *testing.B) {
			for _, point := range v.Points {
				treeAlgo.QueryPoint(point)
			}
		})

	}

}
