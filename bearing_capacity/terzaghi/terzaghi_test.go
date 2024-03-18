package terzaghi

import (
	"testing"

	dt "github.com/geoport/GeoGo/data"
	pkg "github.com/geoport/GeoGo/internal"
)

func TestCalcBearingCapacityFactors(t *testing.T) {
	expectedNc := 172.3
	expectedNq := 173.3
	expectedNg := 305.9
	Nc, Nq, Ng := calcBearingCapacityFactors(45)

	if !pkg.AssertFloat(Nc, expectedNc, 0.1) {
		t.Errorf("Got %v, want %v for Nc", Nc, expectedNc)
	}
	if !pkg.AssertFloat(Nq, expectedNq, 0.1) {
		t.Errorf("Got %v, want %v for Nq", Nq, expectedNq)
	}
	if !pkg.AssertFloat(Ng, expectedNg, 0.1) {
		t.Errorf("Got %v, want %v for Ng", Ng, expectedNg)
	}
}

func TestCalcKp(t *testing.T) {
	expectedKp := 306.4
	Kp := calcKp(45)

	if !pkg.AssertFloat(Kp, expectedKp, 0.1) {
		t.Errorf("Got %v, want %v for Kp", Kp, expectedKp)
	}
}

func TestCalcBearingCapacity(t *testing.T) {
	soilProfile := dt.SoilProfile.Copy()
	foundationData := dt.FoundationData

	soilProfile.CalcLayerDepths()
	expected := 84.44
	output := CalcBearingCapacity(
		soilProfile, foundationData, 50, "short",
	)
	bearingCapacity := output.UltimateBearingCapacity
	if !pkg.AssertFloat(expected, bearingCapacity, 0.1) {
		t.Errorf("Got %v, want %v", bearingCapacity, expected)
	}
}
