package HorizontalSliding

import (
	"math"

	"github.com/geoport/GeoGo/models"
)

// getParams returns the effective unit weight, cohesion and friction angle of the layer at foundation depth
func getParams(soilProfile models.SoilProfile, Df float64) (float64, float64, float64) {
	foundationLayerIndex := soilProfile.GetLayerIndex(Df)
	layer := soilProfile.Layers[foundationLayerIndex]
	cohesion := layer.Cohesion
	Cu := layer.UndrainedShearStrength
	phi := layer.FrictionAngle
	dryUnitWeight := layer.DryUnitWeight
	saturatedUnitWeight := layer.SaturatedUnitWeight

	var selectedUnitWeight float64
	var selectedCohesion float64
	var selectedPhi float64
	if soilProfile.Gwt <= Df {
		selectedUnitWeight = saturatedUnitWeight - 1
		if Cu > 0 {
			selectedCohesion = Cu
			selectedPhi = 0
		} else {
			selectedCohesion = cohesion
			selectedPhi = phi
		}
	} else {
		selectedUnitWeight = dryUnitWeight
		selectedCohesion = cohesion
		selectedPhi = phi
	}
	return selectedCohesion, selectedPhi, selectedUnitWeight
}

// CalcHorizontalSliding calculates the horizontal sliding force between soil and foundation
func CalcHorizontalSliding(
	soilProfile models.SoilProfile, foundationData models.Foundation, foundationPressure, horizontalLoadX, horizontalLoadY float64,
) models.HorizontalSliding {
	var Rth float64

	Df := foundationData.FoundationDepth
	B := foundationData.FoundationWidth
	L := foundationData.FoundationLength

	Vx := horizontalLoadX
	Vy := horizontalLoadY

	surfaceFriction := foundationData.SurfaceFrictionCoefficient
	Ptv := foundationPressure * B * L

	cohesion, phi, unitWeight := getParams(soilProfile, Df)
	kp := math.Pow(math.Tan((45+phi/2)*math.Pi/180), 2)
	if soilProfile.Gwt > Df {
		Rth = Ptv * surfaceFriction / 1.1
	} else {
		Rth = L * B * cohesion / 1.1
	}

	rpkX := B * (0.5 * math.Pow(Df, 2) * unitWeight) * kp
	rpkY := L * (0.5 * math.Pow(Df, 2) * unitWeight) * kp
	rptX := rpkX / 1.4
	rptY := rpkY / 1.4

	sumX := Rth + 0.3*rptX
	sumY := Rth + 0.3*rptY

	isSafeY := Vy <= sumY
	isSafeX := Vx <= sumX
	result := models.HorizontalSliding{
		Rth:     Rth,
		Ptv:     Ptv,
		RpkX:    rpkX,
		RpkY:    rpkY,
		RptX:    rptX,
		RptY:    rptY,
		SumX:    sumX,
		SumY:    sumY,
		IsSafeX: isSafeX,
		IsSafeY: isSafeY,
		Ac:      L * B,
		VthX:    Vx,
		VthY:    Vy,
	}

	return result
}
