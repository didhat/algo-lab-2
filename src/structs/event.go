package structs

type Event struct {
	ZippedX      int
	IsStart      bool
	ZippedYStart int
	ZippedYEnd   int
}

func NewEvent(zippedX int, isStart bool, zippedYStart int, zippedYEnd int) Event {
	return Event{ZippedX: zippedX, IsStart: isStart, ZippedYStart: zippedYStart, ZippedYEnd: zippedYEnd}
}
