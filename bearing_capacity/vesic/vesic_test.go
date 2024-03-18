package vesic

import (
	"testing"

	dt "github.com/geoport/GeoGo/data"
	"github.com/geoport/GeoGo/internal"
)

func TestCalcBearingCapacityByVesic(t *testing.T) {
	soilProfile := dt.SoilProfile.Copy()
	foundationData := dt.FoundationData

	soilProfile.CalcLayerDepths()
	expected := 76.77
	output := CalcBearingCapacity(
		soilProfile, foundationData, 150, 150, 50, "short",
	)
	bearingCapacity := output.UltimateBearingCapacity
	if !internal.AssertFloat(expected, bearingCapacity, 0.1) {
		t.Errorf("Got %v, want %v", bearingCapacity, expected)
	}
}
