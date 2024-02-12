package vesic

import (
	"testing"

	"github.com/geoport/GeoGo/internal"
	"github.com/geoport/GeoGo/models"
)

var foundationData = models.Foundation{
	FoundationDepth:            2.0,
	FoundationWidth:            10.0,
	FoundationLength:           20.,
	SurfaceFrictionCoefficient: 0.6,
	FoundationBaseAngle:        0,
	SlopeAngle:                 0,
}

var soilProfile = models.SoilProfile{
	Gwt: 5,
	Layers: []models.SoilLayer{
		{
			DryUnitWeight:          1.8,
			SaturatedUnitWeight:    1.9,
			Thickness:              3,
			Cohesion:               1,
			UndrainedShearStrength: 3,
			FrictionAngle:          21,
			EffectiveFrictionAngle: 21,
		},
		{
			DryUnitWeight:          1.9,
			SaturatedUnitWeight:    2,
			Thickness:              5,
			Cohesion:               0.5,
			UndrainedShearStrength: 0,
			FrictionAngle:          28,
			EffectiveFrictionAngle: 28,
		},
		{
			DryUnitWeight:          2,
			SaturatedUnitWeight:    2.1,
			Thickness:              50,
			Cohesion:               1,
			UndrainedShearStrength: 5,
			FrictionAngle:          24,
			EffectiveFrictionAngle: 24,
		},
	},
}

func TestCalcBearingCapacityByVesic(t *testing.T) {
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
