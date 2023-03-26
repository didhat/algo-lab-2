package algo

import (
	"testing"
)

func TestMapAlgo_Prepare(t *testing.T) {
	recs := getBasicRecs()
	algo := NewMapAlgo(recs)

	algo.Prepare()

}

func TestMapAlgo_QueryPoint(t *testing.T) {
	recs := getBasicRecs()
	algo := NewMapAlgo(recs)
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
