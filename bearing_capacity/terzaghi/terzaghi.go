package terzaghi

import (
	"math"

	helper "github.com/geoport/GeoGo/bearing_capacity"
	pkg "github.com/geoport/GeoGo/internal"
	"github.com/geoport/GeoGo/models"
)

// calcBearingCapacityFactors is a function that returns bearing capacity factors (Nc,Nq & Ng).
//
// Parameters:
//
// - phi (float64): Angle of internal friction of the soil (in degrees).
//
// Returns:
//
// - Nc (float64): Cohesion factor, which is a function of the angle of internal friction. For phi = 0, a default value of 5.14 is used, based on empirical evidence.
//
// - Nq (float64): Bearing capacity factor related to the depth of the foundation. It is calculated using the exponential and tangent functions, reflecting the influence of soil friction.
//
// - Ng (float64): Bearing capacity factor related to the weight of the soil above the failure wedge. It is derived from Nq and phi, indicating the effect of soil weight on bearing capacity.
//
// Usage:
//
// Nc, Nq, Ng := calcBearingCapacityFactors(35.0)
func calcBearingCapacityFactors(phi float64) (float64, float64, float64) {
	var Nc, Nq, Ng float64

	if phi == 0 {
		Nc = 5.7
		Nq = 1
		Ng = 0

		return Nc, Nq, Ng
	}
	phiRad := pkg.Radian(phi)
	Nq = math.Exp(2*(0.75*math.Pi-phiRad*0.5)*math.Tan(phiRad)) / (2 * math.Pow(math.Cos(pkg.Radian(45+phi/2)), 2))
	Nc = (Nq - 1) / math.Tan(phiRad)

	kp := calcKp(phi)

	Ng = 0.5 * math.Tan(phiRad) * ((kp / math.Pow(math.Cos(phiRad), 2)) - 1)

	return Nc, Nq, Ng
}

func calcKp(phi float64) float64 {
	lnKp := 7.60635563e-08*math.Pow(phi, 5) - 7.48182055e-06*math.Pow(phi, 4) + 2.53533753e-04*math.Pow(phi, 3) - 2.29013627e-03*math.Pow(phi, 2) + 3.41309817e-02*phi + 2.36753692

	return math.Exp(lnKp)
}

func CalcBearingCapacity(soilProfile models.SoilProfile, foundationData models.Foundation,
	foundationPressure float64, term string) Result {

	Df := foundationData.FoundationDepth
	B_ := foundationData.FoundationWidth
	foundationType := foundationData.FoundationType

	effectiveUnitWeight := helper.CalcEffectiveUnitWeight(Df, B_, soilProfile, term)
	stress := helper.CalcStress(soilProfile, Df, term)

	cohesion, phi := helper.GetSoilParams(Df, soilProfile, term)

	Nc, Nq, Ng := calcBearingCapacityFactors(phi)

	var sc, sg float64

	if foundationType == "square" {
		sc = 1.3
		sg = 0.8
	} else if foundationType == "round" {
		sc = 1.3
		sg = 0.6
	} else {
		sc = 1
		sg = 1
	}

	ultimateBearingCapacity := cohesion*Nc*sc + stress*Nq + 0.5*effectiveUnitWeight*B_*Ng*sg

	return Result{
		UltimateBearingCapacity: ultimateBearingCapacity,
		BearingCapacityFactors: BearingCapacityFactors{
			Nc: Nc,
			Nq: Nq,
			Ng: Ng,
		},
		SoilParams: BCSoilParams{
			UnitWeight:    effectiveUnitWeight,
			Cohesion:      cohesion,
			FrictionAngle: phi,
		},
	}

}
