package meyerhof

import (
	"testing"

	dt "github.com/geoport/GeoGo/data"
	"github.com/geoport/GeoGo/internal"
)

func TestCalcBearingCapacity(t *testing.T) {
	soilProfile := dt.SoilProfile.Copy()
	foundationData := dt.FoundationData
	loadData := dt.LoadData

	soilProfile.CalcLayerDepths()
	expected := 62.62
	output := CalcBearingCapacity(
		soilProfile, foundationData, loadData, "short",
	)
	bearingCapacity := output.UltimateBearingCapacity
	if !internal.AssertFloat(expected, bearingCapacity, 0.1) {
		t.Errorf("Got %v, want %v", bearingCapacity, expected)
	}
}
