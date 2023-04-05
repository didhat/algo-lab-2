package algo

import "lab2/src/structs"

type Algo interface {
	Prepare()
	QueryPoint(point structs.Point)
}

type ZippedCords interface {
	GetZippedPoint(point structs.Point) structs.Point
	IsPointBeyondZippedField(point structs.Point) bool
}
