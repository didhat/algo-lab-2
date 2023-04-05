package algo

import (
	"lab2/src/structs"
)

type MapAlgo struct {
	recs        []structs.Rectangle
	preparedMap [][]int
	zipCords    ZippedCords
}

func NewMapAlgo(recs []structs.Rectangle) MapAlgo {
	mapForPrepare := make([][]int, len(recs)*4)
	for index := range mapForPrepare {
		mapForPrepare[index] = make([]int, len(recs)*4)
	}
	zipped := createZippedCordsFromRecs(recs)
	return MapAlgo{recs: recs, preparedMap: mapForPrepare, zipCords: zipped}
}

func (ma *MapAlgo) QueryPoint(point structs.Point) int {
	zippedPoint := ma.zipCords.GetZippedPoint(point)
	if ma.zipCords.IsPointBeyondZippedField(point) {
		return 0
	}
	return ma.preparedMap[zippedPoint.X][zippedPoint.Y]
}

func (ma *MapAlgo) Prepare() {

	for _, rec := range ma.recs {
		zippedLeft := ma.zipCords.GetZippedPoint(rec.LeftDown)
		zippedRight := ma.zipCords.GetZippedPoint(rec.RightTop)
		for i := zippedLeft.X; i <= zippedRight.X; i++ {
			for j := zippedLeft.Y; j <= zippedRight.Y; j++ {
				ma.preparedMap[i][j]++
			}
		}
	}

}
