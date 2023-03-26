package algo

import "lab2/src/structs"

type Algo interface {
	Prepare()
	QueryPoint(point structs.Point)
}
