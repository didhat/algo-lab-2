package algo

import (
	"lab2/src/structs"
	"testing"
)

func TestPersistentTreeAlgo_Prepare(t *testing.T) {
	recs := getBasicRecs()
	algo := NewPersistentTreeAlgo(recs)

	algo.Prepare()

	answ := algo.QueryPoint(structs.NewPoint(2, 2))

	if answ != 1 {
		t.Error("error")
	}

}

func TestPersistentTreeAlgo_QueryPoint(t *testing.T) {
	recs := getBasicRecs()
	algo := NewPersistentTreeAlgo(recs)
	algo.Prepare()

	for _, d := range testDataForBasicRecs {
		t.Run(d.name, func(t *testing.T) {
			result := algo.QueryPoint(d.pointForCheck)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
		})
	}
}

func TestPersistentTreeAlgo_QueryPointWithRandomTestCases(t *testing.T) {
	recs, testCases := generateRandomTestCase(100, 1000)
	algo := NewPersistentTreeAlgo(recs)
	algo.Prepare()

	for _, d := range testCases {
		t.Run(d.name, func(t *testing.T) {
			result := algo.QueryPoint(d.pointForCheck)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
		})
	}

}
