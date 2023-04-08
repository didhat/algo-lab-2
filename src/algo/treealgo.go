package algo

import (
	"lab2/src/structs"
	"sort"
)

type PersistentTreeAlgo struct {
	recs         []structs.Rectangle
	zipCords     ZippedCords
	roots        []structs.SegTreeNode
	rootsZippedX []int
}

func NewPersistentTreeAlgo(recs []structs.Rectangle) PersistentTreeAlgo {
	zipped := createZippedCordsFromRecs(recs)
	return PersistentTreeAlgo{recs: recs, zipCords: zipped}
}

func (pta *PersistentTreeAlgo) QueryPoint(point structs.Point) int {
	if pta.zipCords.IsPointBeyondZippedField(point) {
		return 0
	}
	zippedPoint := pta.zipCords.GetZippedPoint(point)

	rootForAnswer := pta.roots[findPointPosition(pta.rootsZippedX, zippedPoint.X)]

	return structs.GetSum(rootForAnswer, 0, pta.zipCords.YSegmentsNumber(), zippedPoint.Y)

}

func (pta *PersistentTreeAlgo) Prepare() {
	events := make([]structs.Event, 0, len(pta.recs)*2)

	for _, rec := range pta.recs {
		event1 := structs.NewEvent(
			pta.zipCords.GetZippedX(rec.LeftDown.X),
			true,
			pta.zipCords.GetZippedY(rec.LeftDown.Y),
			pta.zipCords.GetZippedY(rec.RightTop.Y+1))

		event2 := structs.NewEvent(
			pta.zipCords.GetZippedX(rec.RightTop.X+1),
			false,
			pta.zipCords.GetZippedY(rec.LeftDown.Y),
			pta.zipCords.GetZippedY(rec.RightTop.Y+1))
		events = append(events, event1, event2)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].ZippedX < events[j].ZippedX
	})

	root := structs.NewEmptySegTreeNode()

	prevZippedX := events[0].ZippedX
	var val int
	for _, ev := range events {
		if ev.ZippedX != prevZippedX {
			pta.roots = append(pta.roots, root)
			pta.rootsZippedX = append(pta.rootsZippedX, prevZippedX)
			prevZippedX = ev.ZippedX
		}
		if ev.IsOpen {
			val = 1
		} else {
			val = -1
		}
		root = structs.AddToSegTree(root, 0, pta.zipCords.YSegmentsNumber(), ev.ZippedYStart, ev.ZippedYEnd, val)
	}

	pta.roots = append(pta.roots, root)
	pta.rootsZippedX = append(pta.rootsZippedX, prevZippedX)
}
