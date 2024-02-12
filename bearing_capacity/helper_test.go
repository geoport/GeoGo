package bearingcapacity

import (
	"GeoGo/internal"
	"GeoGo/models"
	"testing"

	np "github.com/geoport/numpy4go/vectors"
)

var soilProfile = models.SoilProfile{
	Layers: []models.SoilLayer{
		{
			DryUnitWeight:          1.8,
			SaturatedUnitWeight:    1.9,
			Thickness:              3,
			Cohesion:               1,
			UndrainedShearStrength: 4,
			FrictionAngle:          30,
			EffectiveFrictionAngle: 20,
		},
		{
			DryUnitWeight:          1.9,
			SaturatedUnitWeight:    2,
			Thickness:              5,
			Cohesion:               2,
			UndrainedShearStrength: 5,
			FrictionAngle:          32,
			EffectiveFrictionAngle: 22,
		},
		{
			DryUnitWeight:          2,
			SaturatedUnitWeight:    2.1,
			Thickness:              50,
			Cohesion:               3,
			UndrainedShearStrength: 7,
			FrictionAngle:          33,
			EffectiveFrictionAngle: 23,
		},
	},
	Gwt: 1,
}

func TestCalcEffectiveUnitWeight(t *testing.T) {
	soilProfile_ := soilProfile.Copy()
	var effectiveUnitWeight, expected float64
	// case 1
	effectiveUnitWeight = CalcEffectiveUnitWeight(10, 0, soilProfile_, "short")
	expected = 0.98
	if np.Round(effectiveUnitWeight, 2) != expected {
		t.Errorf("Got %v, want %v for effectiveUnitWeight for case 1", effectiveUnitWeight, expected)
	}

	// case 2
	soilProfile_.Gwt = 12
	effectiveUnitWeight = CalcEffectiveUnitWeight(10, 5, soilProfile_, "short")
	expected = 1.35
	if np.Round(effectiveUnitWeight, 2) != expected {
		t.Errorf("Got %v, want %v for effectiveUnitWeight for case 2", effectiveUnitWeight, expected)
	}

	// case 3
	soilProfile_.Gwt = 30
	effectiveUnitWeight = CalcEffectiveUnitWeight(10, 0, soilProfile_, "short")
	expected = 1.89
	if np.Round(effectiveUnitWeight, 2) != expected {
		t.Errorf("Got %v, want %v for effectiveUnitWeight for case 3", effectiveUnitWeight, expected)
	}
}

func TestGetSoilParams(t *testing.T) {
	soilProfile.CalcLayerDepths()
	cohesion1, phi1 := GetSoilParams(4, soilProfile, "short")
	cohesion2, phi2 := GetSoilParams(10, soilProfile, "long")

	output1 := []float64{cohesion1, phi1}
	output2 := []float64{cohesion2, phi2}

	expected1 := []float64{5, 32}
	expected2 := []float64{3, 23}

	if !internal.AssertFloatArray(output1, expected1, 0.1) {
		t.Errorf("Got %v, want %v for case 1", output1, expected1)
	}

	if !internal.AssertFloatArray(output2, expected2, 0.1) {
		t.Errorf("Got %v, want %v for case 2", output2, expected2)
	}
}
