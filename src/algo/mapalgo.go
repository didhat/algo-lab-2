package algo

import (
	"lab2/src/structs"
	"sort"
)

type MapAlgo struct {
	recs          []structs.Rectangle
	preparedMap   [][]int
	xZippedCoords []int
	yZippedCoords []int
}

func NewMapAlgo(recs []structs.Rectangle) MapAlgo {
	mapForPrepare := make([][]int, len(recs)*4)
	for index := range mapForPrepare {
		mapForPrepare[index] = make([]int, len(recs)*4)
	}
	xCoords := make([]int, 0, len(recs)*4)
	yCoords := make([]int, 0, len(recs)*4)
	return MapAlgo{recs: recs, preparedMap: mapForPrepare, xZippedCoords: xCoords, yZippedCoords: yCoords}
}

func (ma *MapAlgo) QueryPoint(point structs.Point) int {
	zippedPoint := structs.NewPoint(zipCord(ma.xZippedCoords, point.X), zipCord(ma.yZippedCoords, point.Y))
	if ma.isPointBeyondMap(point) {
		return 0
	}
	return ma.preparedMap[zippedPoint.X][zippedPoint.Y]
}

func (ma *MapAlgo) Prepare() {

	for _, rec := range ma.recs {
		ma.xZippedCoords = append(ma.xZippedCoords, rec.LeftDown.X, rec.RightTop.X, rec.RightTop.X+1, rec.LeftDown.X+1)
		ma.yZippedCoords = append(ma.yZippedCoords, rec.RightTop.Y, rec.LeftDown.Y, rec.LeftDown.Y+1, rec.RightTop.Y+1)
	}

	sort.Sort(sort.IntSlice(ma.xZippedCoords))
	sort.Sort(sort.IntSlice(ma.yZippedCoords))
	ma.xZippedCoords = ma.removeDuplicates(ma.xZippedCoords)
	ma.yZippedCoords = ma.removeDuplicates(ma.yZippedCoords)

	for _, rec := range ma.recs {
		zippedLeft := structs.NewPoint(zipCord(ma.xZippedCoords, rec.LeftDown.X), zipCord(ma.yZippedCoords, rec.LeftDown.Y))
		zippedRight := structs.NewPoint(zipCord(ma.xZippedCoords, rec.RightTop.X), zipCord(ma.yZippedCoords, rec.RightTop.Y))
		for i := zippedLeft.X; i <= zippedRight.X; i++ {
			for j := zippedLeft.Y; j <= zippedRight.Y; j++ {
				ma.preparedMap[i][j]++
			}
		}
	}

}

func (ma *MapAlgo) removeDuplicates(sortedCoords []int) []int {
	filteredCoords := make([]int, 0, 2*len(ma.recs))
	prev := sortedCoords[len(sortedCoords)-1]

	for _, val := range sortedCoords {
		if val != prev {
			prev = val
			filteredCoords = append(filteredCoords, val)
		} else if val == prev {
			continue
		}
	}
	return filteredCoords

}

func (ma *MapAlgo) isPointBeyondMap(point structs.Point) bool {
	if point.X < ma.xZippedCoords[0] || point.Y < ma.yZippedCoords[0] {
		return true
	}
	return false
}

func zipCord(sortedCoords []int, target int) int {
	low, high, mid := 0, len(sortedCoords)-1, 0

	for low <= high {
		mid = (low + high) / 2
		if sortedCoords[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low - 1
}
