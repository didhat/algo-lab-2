package structs

type Event struct {
	ZippedX      int
	IsOpen       bool
	ZippedYStart int
	ZippedYEnd   int
}

func NewEvent(zippedX int, isOpen bool, zippedXStart int, zippedYStart int) Event {
	return Event{ZippedX: zippedX, IsOpen: isOpen, ZippedYStart: zippedXStart, ZippedYEnd: zippedYStart}
}
