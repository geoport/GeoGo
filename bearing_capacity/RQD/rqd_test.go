package RQD

import (
	"testing"

	dt "github.com/geoport/GeoGo/data"

	np "github.com/geoport/numpy4go/vectors"
)

func TestCalcBearingCapacity(t *testing.T) {
	soilProfile := dt.SoilProfile.Copy()
	foundationData := dt.FoundationData
	soilProfile.CalcLayerDepths()

	expectedQfRatio := 0.2
	expectedUltimateBearingCapacity := 342.64
	expectedQLab := 1713.18
	output := CalcBearingCapacity(soilProfile, foundationData, 50)

	qfRatio := output.QFRatio
	ultimateBearingCapacity := output.UltimateBearingCapacity
	qLab := output.QLab
	if np.Round(qfRatio, 2) != expectedQfRatio {
		t.Errorf("Expected qfRatio to be %f, got %f", expectedQfRatio, qfRatio)
	}
	if np.Round(ultimateBearingCapacity, 2) != expectedUltimateBearingCapacity {
		t.Errorf(
			"Expected ultimateBearingCapacity to be %f, got %f", expectedUltimateBearingCapacity,
			ultimateBearingCapacity,
		)
	}
	if np.Round(qLab, 2) != expectedQLab {
		t.Errorf("Expected qLab to be %f, got %f", expectedQLab, qLab)
	}
}
