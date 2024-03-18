package tezcan_ozdemir

import (
	"testing"

	dt "github.com/geoport/GeoGo/data"

	np "github.com/geoport/numpy4go/vectors"
)

func TestCalcBearingCapacity(t *testing.T) {
	soilProfile := dt.SoilProfile.Copy()
	foundationData := dt.FoundationData

	soilProfile.CalcLayerDepths()
	expectedSafetyFactor := 4.
	expectedBearingCapacity := 14.63

	output := CalcBearingCapacity(soilProfile, foundationData, 50)
	safetyFactor := output.SafetyFactor
	allowableBearingCapacity := output.AllowableBearingCapacity
	if np.Round(safetyFactor, 2) != expectedSafetyFactor {
		t.Errorf("Expected safety factor to be %f, got %f", expectedSafetyFactor, safetyFactor)
	}
	if np.Round(allowableBearingCapacity, 2) != expectedBearingCapacity {
		t.Errorf("Expected bearing capacity to be %f, got %f", expectedBearingCapacity, allowableBearingCapacity)
	}
}
