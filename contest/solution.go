package main

import (
	"fmt"
	"sort"
)

type Event struct {
	ZippedX      int
	IsStart      bool
	ZippedYStart int
	ZippedYEnd   int
}

func NewEvent(zippedX int, isStart bool, zippedXStart int, zippedYEnd int) Event {
	return Event{ZippedX: zippedX, IsStart: isStart, ZippedYStart: zippedXStart, ZippedYEnd: zippedYEnd}
}

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{X: x, Y: y}
}

type Rectangle struct {
	LeftDown Point
	RightTop Point
}

func NewRectangleFromPrimitives(x1 int, y1 int, x2 int, y2 int) Rectangle {
	left := NewPoint(x1, y1)
	right := NewPoint(x2, y2)
	return Rectangle{LeftDown: left, RightTop: right}
}

type SegTreeNode struct {
	left  *SegTreeNode
	right *SegTreeNode
	sum   int
}

func NewSegTreeNode(left, right *SegTreeNode, sum int) SegTreeNode {
	return SegTreeNode{left: left, right: right, sum: sum}
}

func NewEmptySegTreeNode() SegTreeNode {
	return SegTreeNode{left: nil, right: nil, sum: 0}
}

func AddToSegTree(root SegTreeNode, left, right, rangeStart, rangeEnd, value int) SegTreeNode {
	if left >= rangeEnd || right <= rangeStart {
		return root
	}

	if rangeStart <= left && right <= rangeEnd {
		newRoot := NewSegTreeNode(root.left, root.right, root.sum)
		newRoot.sum += value
		return newRoot
	}

	mid := (left + right) / 2
	newRoot := NewSegTreeNode(root.left, root.right, root.sum)

	if newRoot.left == nil {
		_left := NewEmptySegTreeNode()
		newRoot.left = &_left
	}

	newLeft := AddToSegTree(*newRoot.left, left, mid, rangeStart, rangeEnd, value)
	newRoot.left = &newLeft

	if newRoot.right == nil {
		_right := NewEmptySegTreeNode()
		newRoot.right = &_right
	}
	newRight := AddToSegTree(*newRoot.right, mid, right, rangeStart, rangeEnd, value)
	newRoot.right = &newRight

	return newRoot
}

func GetSum(root SegTreeNode, left, right, targetZippedX int) int {
	if right-left == 1 {
		return root.sum
	}

	mid := (left + right) / 2

	if targetZippedX < mid {
		if root.left == nil {
			return root.sum
		}
		return root.sum + GetSum(*root.left, left, mid, targetZippedX)
	} else {
		if root.right == nil {
			return root.sum
		}
		return root.sum + GetSum(*root.right, mid, right, targetZippedX)
	}

}

type ZippedCordsImp struct {
	zippedXCords []int
	zippedYCords []int
}

func createZippedCordsFromRecs(recs []Rectangle) ZippedCordsImp {
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

func (zp ZippedCordsImp) GetZippedPoint(unzippedPoint Point) Point {
	return NewPoint(findPointPosition(zp.zippedXCords, unzippedPoint.X), findPointPosition(zp.zippedYCords, unzippedPoint.Y))
}

func (zp ZippedCordsImp) IsPointBeyondZippedField(point Point) bool {
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

type PersistentTreeAlgo struct {
	recs         []Rectangle
	zipCords     ZippedCordsImp
	roots        []SegTreeNode
	rootsZippedX []int
}

func NewPersistentTreeAlgo(recs []Rectangle) PersistentTreeAlgo {
	zipped := createZippedCordsFromRecs(recs)
	return PersistentTreeAlgo{recs: recs, zipCords: zipped}
}

func (pta *PersistentTreeAlgo) QueryPoint(point Point) int {
	if pta.zipCords.IsPointBeyondZippedField(point) {
		return 0
	}
	zippedPoint := pta.zipCords.GetZippedPoint(point)

	rootForAnswer := pta.roots[findPointPosition(pta.rootsZippedX, zippedPoint.X)]

	return GetSum(rootForAnswer, 0, pta.zipCords.YSegmentsNumber(), zippedPoint.Y)

}

func (pta *PersistentTreeAlgo) Prepare() {
	events := pta.createEventsForPersistentSegTree()
	pta.createPersistentSegmentTree(events)
}

func (pta *PersistentTreeAlgo) createPersistentSegmentTree(events []Event) {
	root := NewEmptySegTreeNode()

	prevZippedX := events[0].ZippedX
	var val int
	for _, ev := range events {
		if ev.ZippedX != prevZippedX {
			pta.roots = append(pta.roots, root)
			pta.rootsZippedX = append(pta.rootsZippedX, prevZippedX)
			prevZippedX = ev.ZippedX
		}
		if ev.IsStart {
			val = 1
		} else {
			val = -1
		}
		root = AddToSegTree(root, 0, pta.zipCords.YSegmentsNumber(), ev.ZippedYStart, ev.ZippedYEnd, val)
	}

	pta.roots = append(pta.roots, root)
	pta.rootsZippedX = append(pta.rootsZippedX, prevZippedX)
}

func (pta *PersistentTreeAlgo) createEventsForPersistentSegTree() []Event {
	events := make([]Event, 0, len(pta.recs)*2)

	for _, rec := range pta.recs {
		event1 := NewEvent(
			pta.zipCords.GetZippedX(rec.LeftDown.X),
			true,
			pta.zipCords.GetZippedY(rec.LeftDown.Y),
			pta.zipCords.GetZippedY(rec.RightTop.Y+1))

		event2 := NewEvent(
			pta.zipCords.GetZippedX(rec.RightTop.X+1),
			false,
			pta.zipCords.GetZippedY(rec.LeftDown.Y),
			pta.zipCords.GetZippedY(rec.RightTop.Y+1))
		events = append(events, event1, event2)
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].ZippedX < events[j].ZippedX
	})

	return events
}

func main() {
	var n, m, x1, x2, y1, y2, xForCheck, yForCheck int
	fmt.Scanf("%d", &n)
	recs := make([]Rectangle, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2)
		recs = append(recs, NewRectangleFromPrimitives(x1, y1, x2, y2))
	}
	fmt.Scanf("%d", &m)
	points := make([]Point, 0, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &xForCheck, &yForCheck)
		points = append(points, NewPoint(xForCheck, yForCheck))
	}
	if n > 0 {
		algo := NewPersistentTreeAlgo(recs)
		algo.Prepare()
		for _, p := range points {
			result := algo.QueryPoint(p)
			fmt.Printf("%d ", result)
		}
	} else {
		fmt.Printf("")
	}

}
