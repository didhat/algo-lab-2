package algo

import (
	"lab2/src/structs"
	"sort"
)

type ZippedCordsImp struct {
	zippedXCords []int
	zippedYCords []int
}

func createZippedCordsFromRecs(recs []structs.Rectangle) ZippedCordsImp {
	xCords := make([]int, 0, len(recs)*4)
	yCords := make([]int, 0, len(recs)*4)

	for _, rec := range recs {
		xCords = append(xCords, rec.LeftDown.X, rec.RightTop.X, rec.RightTop.X+1)
		yCords = append(yCords, rec.LeftDown.Y, rec.RightTop.Y, rec.RightTop.Y+1)
	}

	sort.Sort(sort.IntSlice(xCords))
	sort.Sort(sort.IntSlice(yCords))

	xCords = removeDuplicates(xCords)
	yCords = removeDuplicates(yCords)

	return ZippedCordsImp{zippedYCords: yCords, zippedXCords: xCords}

}

func (zp ZippedCordsImp) GetZippedPoint(unzippedPoint structs.Point) structs.Point {
	return structs.NewPoint(findPointPosition(zp.zippedXCords, unzippedPoint.X), findPointPosition(zp.zippedYCords, unzippedPoint.Y))
}

func (zp ZippedCordsImp) IsPointBeyondZippedField(point structs.Point) bool {
	if point.X < zp.zippedXCords[0] || point.Y < zp.zippedYCords[0] {
		return true
	}
	return false
}

func (zp ZippedCordsImp) GetZippedX(x int) int {
	return findPointPosition(zp.zippedXCords, x)
}

func (zp ZippedCordsImp) GetZippedY(y int) int {
	return findPointPosition(zp.zippedYCords, y)
}

func (zp ZippedCordsImp) YSegmentsNumber() int {
	return len(zp.zippedYCords)
}

func removeDuplicates(sortedCoords []int) []int {
	filteredCoords := make([]int, 0, len(sortedCoords))
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

func findPointPosition(sortedCoords []int, target int) int {
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
