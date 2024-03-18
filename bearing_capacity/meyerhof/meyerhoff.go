package meyerhof

import (
	helper "github.com/geoport/GeoGo/bearing_capacity"
	"github.com/geoport/GeoGo/models"
)

// CalcBearingCapacity is a function that calculates the ultimate and allowable bearing capacity of a foundation using the Meyerhofs method.
//
// Parameters:
//
// - soilProfile (models.SoilProfile): The soil profile to be analyzed.
//
// - foundationData (models.Foundation): The foundation data to be analyzed.
//
// - horizontalLoadX (float64): The horizontal load in the X direction (in t/m).
//
// - horizontalLoadY (float64): The horizontal load in the Y direction (in t/m).
//
// - foundationPressure (float64): The foundation pressure (in t/m2).
//
// - term (string): Defines the type of analysis (short || long).
//
// Returns:
//
// - result (Result): The result of the bearing capacity analysis.
func CalcBearingCapacity(
	soilProfile models.SoilProfile, foundationData models.Foundation,
	horizontalLoadX, horizontalLoadY, foundationPressure float64, term string,
) Result {
	//unitWeight is in t/m3
	//cohesion is in t/m2
	//stress is in t/m2
	//bearing capacity is in t/m2
	Vmax := max(horizontalLoadX, horizontalLoadY)
	Df := foundationData.FoundationDepth
	B_ := foundationData.FoundationWidth
	L_ := foundationData.FoundationLength

	effectiveUnitWeight := helper.CalcEffectiveUnitWeight(Df, B_, soilProfile, term)
	stress := helper.CalcStress(soilProfile, Df, term)

	cohesion, phi := helper.GetSoilParams(Df, soilProfile, term)

	Nc, Nq, Ng := calcBearingCapacityFactors(phi)
	Sc, Sq, Sg := calcShapeFactors(B_, L_, phi)
	Dc, Dq, Dg := calcDepthFactors(Df, B_, phi)
	Ic, Iq, Ig := calcLoadInclinationFactors(phi, B_, L_, Vmax, foundationPressure)

	var fc, fq, fg float64

	if horizontalLoadX+horizontalLoadY > 0 {
		fc = Ic
		fq = Iq
		fg = Ig
	} else {
		fc = Sc
		fq = Sq
		fg = Sg
	}

	partC := cohesion * Nc * fc * Dc
	partQ := stress * Nq * fq * Dq
	partG := 0.5 * effectiveUnitWeight * B_ * Ng * fg * Dg

	ultimateBearingCapacity := partC + partQ + partG

	result := Result{
		UltimateBearingCapacity: ultimateBearingCapacity,
		BearingCapacityFactors: BearingCapacityFactors{
			Nc: Nc,
			Nq: Nq,
			Ng: Ng,
		},
		ShapeFactors: ShapeFactors{
			Sc: Sc,
			Sq: Sq,
			Sg: Sg,
		},
		DepthFactors: DepthFactors{
			Dc: Dc,
			Dq: Dq,
			Dg: Dg,
		},
		LoadInclinationFactors: LoadInclinationFactors{
			Ic: Ic,
			Iq: Iq,
			Ig: Ig,
		},
		SoilParams: BCSoilParams{Cohesion: cohesion, FrictionAngle: phi, UnitWeight: effectiveUnitWeight},
	}

	return result
}
