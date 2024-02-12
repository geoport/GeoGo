package vesic

import (
	helper "github.com/geoport/GeoGo/bearing_capacity"
	"github.com/geoport/GeoGo/models"
)

func CalcBearingCapacity(
	soilProfile models.SoilProfile, foundationData models.Foundation,
	horizontalLoadX, horizontalLoadY, foundationPressure float64, term string,
) models.VesicBearingCapacity {
	//unitWeight is in ton/m3
	//cohesion is in ton/m2
	//normalStress is in t/m2
	//bearing capacity is in ton/m2
	Vmax := max(horizontalLoadX, horizontalLoadY)
	Df := foundationData.FoundationDepth
	B_ := foundationData.FoundationWidth
	L_ := foundationData.FoundationLength
	slopeAngle := foundationData.SlopeAngle
	baseAngle := foundationData.FoundationBaseAngle
	effectiveUnitWeight := helper.CalcEffectiveUnitWeight(Df, B_, soilProfile, term)
	normalStress := soilProfile.CalcNormalStress(Df)

	cohesion, phi := helper.GetSoilParams(Df, soilProfile, term)

	Nc, Nq, Ng := calcBearingCapacityFactors(phi)
	Sc, Sq, Sg := calcShapeFactors(B_, L_, Nq, Nc, phi)
	Dc, Dq, Dg := calcDepthFactors(Df, B_, phi)
	Ic, Iq, Ig := calcLoadInclinationFactors(phi, cohesion, B_, L_, baseAngle, Vmax, foundationPressure)
	Bc, Bq, Bg := calcBaseFactors(phi, slopeAngle, baseAngle)
	Gc, Gq, Gg := calcGroundFactors(Iq, slopeAngle, phi)

	partC := cohesion * Nc * Sc * Dc * Bc * Gc
	partQ := normalStress * Nq * Sq * Dq * Bq * Gq
	partG := 0.5 * effectiveUnitWeight * B_ * Ng * Sg * Dg * Bg * Gg

	ultimateBearingCapacity := partC + partQ + partG
	allowableBearingCapacity := ultimateBearingCapacity / 1.5

	result := models.VesicBearingCapacity{
		UltimateBearingCapacity:  ultimateBearingCapacity,
		AllowableBearingCapacity: allowableBearingCapacity,
		Nc:                       Nc,
		Nq:                       Nq,
		Ng:                       Ng,
		Sc:                       Sc,
		Sq:                       Sq,
		Sg:                       Sg,
		Dc:                       Dc,
		Dq:                       Dq,
		Dg:                       Dg,
		Ic:                       Ic,
		Iq:                       Iq,
		Ig:                       Ig,
		Bc:                       Bc,
		Bq:                       Bq,
		Bg:                       Bg,
		Gc:                       Gc,
		Gq:                       Gq,
		Gg:                       Gg,
		UnitWeight:               effectiveUnitWeight,
		Cohesion:                 cohesion,
		FrictionAngle:            phi,
		IsSafe:                   allowableBearingCapacity >= foundationPressure,
	}

	return result
}
