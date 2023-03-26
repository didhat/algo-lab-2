package algo

import "lab2/src/structs"

type BasicAlgo struct {
	rects []structs.Rectangle
}

func NewBasicAlgo(rects []structs.Rectangle) BasicAlgo {
	return BasicAlgo{rects: rects}
}

func (ba BasicAlgo) Prepare() {}

func (ba BasicAlgo) QueryPoint(point structs.Point) int {
	answer := 0

	for _, rec := range ba.rects {
		if rec.LeftDown.X <= point.X && point.X <= rec.RightTop.X && rec.LeftDown.Y <= point.Y && point.Y <= rec.RightTop.Y {
			answer++
		}
	}
	return answer
}
