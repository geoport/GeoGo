package tezcan_ozdemir

import (
	"github.com/geoport/GeoGo/models"
)

// GetSoilParams returns unit weight and shear wave velocity at foundation layer.
func getSoilParams(Df float64, soilProfile models.SoilProfile) (float64, float64) {
	layerIndex := soilProfile.GetLayerIndex(Df)
	layer := soilProfile.Layers[layerIndex]
	VS := layer.ShearWaveVelocity
	gwt := soilProfile.Gwt

	var unitWeight float64

	if gwt < Df {
		unitWeight = layer.SaturatedUnitWeight
	} else {
		unitWeight = layer.DryUnitWeight
	}

	return VS, unitWeight
}

// CalcBearingCapacity is a function that returns allowable bearing capacity and the safety factor.
func CalcBearingCapacity(
	soilProfile models.SoilProfile, foundationData models.Foundation, foundationPressure float64,
) Result {
	Df := foundationData.FoundationDepth
	VS, unitWeight := getSoilParams(Df, soilProfile)
	//VS is in m/s
	//unitWeight is in t/m3
	//bearing capacity is in t/m3

	var safetyFactor, bearingCapacity float64
	if VS >= 4000 {
		safetyFactor = 1.4
		bearingCapacity = 0.071 * unitWeight * VS
	} else if VS >= 750 && VS < 4000 {
		safetyFactor = 4.6 - VS*8e-4
		bearingCapacity = 0.1 * unitWeight * VS / safetyFactor
	} else {
		safetyFactor = 4
		bearingCapacity = 0.025 * unitWeight * VS
	}

	result := Result{
		AllowableBearingCapacity: bearingCapacity,
		SafetyFactor:             safetyFactor,
		IsSafe:                   bearingCapacity >= foundationPressure,
		UnitWeight:               unitWeight,
		VS:                       VS,
	}
	return result
}
