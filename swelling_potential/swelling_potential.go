package swelling_potential

import (
	"math"

	"github.com/geoport/GeoGo/models"
)

// KayabaliYaldiz2014 is the function to calculate the swelling pressure of the soil by Kayabali and Yaldiz 2014 method
func KayabaliYaldiz2014(PL float64, LL float64, dryUnitWeight float64, waterContent float64) float64 {
	// water content is in %
	// dryUnitWeight is in t/m^3
	// output is in t/m^2
	return -3.08*waterContent + 102.5*dryUnitWeight + 0.635*LL + 4.24*PL - 220.8
}

// CalcSwellingPotential calculates the swelling potential of the soil
func CalcSwellingPotential(
	soilProfile models.SoilProfile, foundationData models.Foundation, foundationPressure float64,
) models.SwellingPotential {
	Df := foundationData.FoundationDepth
	B := foundationData.FoundationWidth
	L := foundationData.FoundationLength
	Qnet := foundationPressure - soilProfile.CalcNormalStress(Df)
	F := Qnet * B * L
	layerCenters := soilProfile.GetLayerCenters()

	var effectiveStress float64
	var deltaSigma float64
	var safety bool
	var swellingPressure float64
	var pressures, effectiveStresses, deltaSigmas []float64
	var checkSafety []bool

	for _, layer := range soilProfile.Layers {
		z := layer.Center
		PL := layer.PlasticLimit
		LL := layer.LiquidLimit
		dryUnitWeight := layer.DryUnitWeight
		waterContent := layer.WaterContent
		if z < Df {
			effectiveStress = 0
			deltaSigma = 0
		} else {
			effectiveStress = soilProfile.CalcEffectiveStress(z)
			deltaSigma = F / ((B + z - Df) * (L + z - Df))
		}
		// safety is true if swelling pressure is smaller than effective stress + delta sigma
		if layer.IsCohesive {
			swellingPressure = KayabaliYaldiz2014(PL, LL, dryUnitWeight, waterContent)
			safety = swellingPressure <= (effectiveStress + deltaSigma)
		} else {
			swellingPressure = 0
			safety = true
		}
		effectiveStresses = append(effectiveStresses, effectiveStress)
		deltaSigmas = append(deltaSigmas, deltaSigma)
		pressures = append(pressures, math.Max(0, swellingPressure))
		checkSafety = append(checkSafety, safety)
	}

	result := models.SwellingPotential{
		SwellingPressures:     pressures,
		IsSafe:                checkSafety,
		LayerCenters:          layerCenters,
		EffectiveStresses:     effectiveStresses,
		DeltaSigmas:           deltaSigmas,
		NetFoundationPressure: Qnet,
	}

	return result
}
