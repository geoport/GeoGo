package vesic

import (
	"math"

	pkg "github.com/geoport/GeoGo/internal"
)

// calcShapeFactors calculates the shape factors for foundation design in geotechnical engineering.
//
// Parameters:
//
// - B (float64): Width of the foundation (in meters).
//
// - L (float64): Length of the foundation (in meters).
//
// - Nq (float64): Bearing capacity factor related to the angle of internal friction of the soil.
//
// - Nc (float64): Bearing capacity factor related to cohesion.
//
// - phi (float64): Angle of internal friction of the soil (in degrees).
//
// Returns:
//
// - Sc (float64): Shape factor for cohesion, accounting for the foundation's dimensions and soil properties.
//
// - Sq (float64): Shape factor for the angle of internal friction, modifying the bearing capacity factor Nq based on the foundation's width and length, and the soil's angle of internal friction.
//
// - Sg (float64): Shape factor for foundation's geometry, specifically affecting the gradient of the soil's shear strength with depth.
//
// Usage:
//
// Sc, Sq, Sg := calcShapeFactors(2.0, 4.0, 30.0, 17.5, 35.0)
func calcShapeFactors(B, L, Nq, Nc, phi float64) (float64, float64, float64) {
	Sc := 1 + (B/L)*(Nq/Nc)
	Sq := 1 + (B/L)*math.Tan(pkg.Radian(phi))
	Sg := 1 - 0.4*B/L

	Sg = math.Max(Sg, 0.6)

	return Sc, Sq, Sg
}

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
	var Nc float64
	Nq := math.Exp(math.Pi*math.Tan(pkg.Radian(phi))) * math.Pow(math.Tan(pkg.Radian(45+phi/2)), 2)
	if phi == 0 {
		Nc = 5.14
	} else {
		Nc = (Nq - 1) / math.Tan(pkg.Radian(phi))
	}
	Ng := 2 * (Nq - 1) * math.Tan(pkg.Radian(phi))

	return Nc, Nq, Ng
}

// calcLoadInclinationFactors calculates the load inclination factors for a foundation based on various parameters including the soil's angle of internal friction, cohesion, foundation dimensions, and applied pressure. These factors adjust the bearing capacity of the foundation considering the effect of inclined loads.
//
// Parameters:
//
// - phi (float64): Angle of internal friction of the soil (in degrees), affecting the soil's shear strength.
//
// - cohesion (float64): Cohesion value of the soil (in tons per square meter), representing the soil's inherent shear strength independent of its internal friction.
//
// - B (float64): Width of the foundation (in meters).
//
// - L (float64): Length of the foundation (in meters).
//
// - baseAngle (float64) : Angle of the foundation to the soil base (in degrees).
//
// - Vmax (float64): Maximum horizontal load applied on the foundation (in tons).
//
// - verticalLoad (float64): Load applied by the foundation on the soil (in tons).
//
// Returns:
//
// - ic (float64): Load inclination factor for cohesion, modifying the cohesion bearing capacity factor (Nc) based on the load's inclination.
//
// - iq (float64): Load inclination factor for the depth (related to Nq), adjusting the bearing capacity factor Nq considering the inclined load.
//
// - ig (float64): Load inclination factor for the weight of the soil (related to Ng), affecting the bearing capacity related to the soil's weight above the failure wedge.
//
// Usage Example:
//
// ic, iq, ig := calcLoadInclinationFactors(30, 0.5, 20, 50, 100, 150)
func calcLoadInclinationFactors(
	phi, cohesion, B, L, baseAngle, Vmax, verticalLoad float64,
) (float64, float64, float64) {
	Nc, Nq, _ := calcBearingCapacityFactors(phi)

	var ic, iq, ig float64

	if baseAngle > 0 {
		A := B * L
		Ca := cohesion * 0.75
		m := (2 + (B / L)) / (1 + (B / L))
		if phi == 0 {
			ic = 1 - m*Vmax/(A*Ca*Nc)
			iq = 1
			ig = 1
		} else {
			iq = math.Pow(1-(Vmax/(verticalLoad+A*Ca*(1/math.Tan(pkg.Radian(phi))))), m)
			ig = math.Pow(1-(Vmax/(verticalLoad+A*Ca*(1/math.Tan(pkg.Radian(phi))))), m+1)
			ic = iq - (1-iq)/(Nq-1)
		}
	} else {
		ic = 1
		iq = 1
		ig = 1
	}

	return ic, iq, ig
}

// calcBaseFactors calculates the base correction factors for foundation design in geotechnical engineering, taking into account the effects of slope angle and base angle on the foundation's bearing capacity.
//
// Parameters:
//
// - phi (float64): Angle of internal friction of the soil (in degrees), indicative of the soil's shear strength.
//
// - slopeAngle (float64): Angle of the slope (in degrees) relative to the horizontal on which the foundation rests.
//
// - baseAngle (float64): Angle of the foundation base (in degrees) relative to the horizontal.
//
// Returns:
//
// - bc (float64): Base correction factor for cohesion, adjusting the cohesion component of bearing capacity based on the slope angle. For soils with zero angle of internal friction (phi = 0), a specific formula is applied.
//
// - bq (float64): Base correction factor for the depth, modifying the bearing capacity factor related to the depth based on the base angle.
//
// - bg (float64): Base correction factor for the weight of the soil, identical to bq in this context, indicating the adjustment to the bearing capacity related to the soil's weight above the failure wedge.
//
// Usage Example:
//
// bc, bq, bg := calcBaseFactors(30, 10, 5)
func calcBaseFactors(phi, slopeAngle, baseAngle float64) (float64, float64, float64) {
	var bc, bg, bq float64

	if phi == 0 {
		bc = 1 - pkg.Radian(slopeAngle)/5.14
	} else {
		bc = 1 - 2*pkg.Radian(slopeAngle)/(5.14*math.Tan(pkg.Radian(phi)))
	}

	bq = math.Pow(1-pkg.Radian(baseAngle)*math.Tan(pkg.Radian(phi)), 2)
	bg = bq

	return bc, bq, bg
}

// calcGroundFactors calculates the ground correction factors for evaluating the stability and bearing capacity of foundations on sloped surfaces. These factors adjust the bearing capacity equations to account for the effects of slope on the foundation's performance.
//
// Parameters:
//
// - iq (float64): Load inclination factor for the depth (related to Nq), which adjusts the bearing capacity factor Nq considering the inclined load.
//
// - slopeAngle (float64): Angle of the slope (in degrees) on which the foundation is situated.
//
// - phi (float64): Angle of internal friction of the soil (in degrees).
//
// Returns:
//
// - gc (float64): Ground correction factor for cohesion.
//
// - gq (float64): Ground correction factor for the depth.
//
// - gg (float64): Ground correction factor for the weight of the soil.
//
// Usage Example:
//
// gc, gq, gg := calcGroundFactors(1.2, 15, 30)
func calcGroundFactors(iq, slopeAngle, phi float64) (float64, float64, float64) {
	var gc, gq, gg float64

	if phi == 0 {
		gc = 1 - pkg.Radian(slopeAngle)/5.14
	} else {
		gc = iq - (1-iq)/(5.14*math.Tan(pkg.Radian(phi)))
	}

	gq = math.Pow(1-math.Tan(pkg.Radian(slopeAngle)), 2)
	gg = gq

	return gc, gq, gg
}

// calcDepthFactors calculates the depth correction factors for assessing the bearing capacity of foundations, considering the effect of foundation depth. These factors are crucial for the design and analysis of deep foundations where the depth influences the soil's bearing capacity.
//
// Parameters:
//
// - Df (float64): Depth of the foundation (in meters).
//
// - B (float64): Width of the foundation (in meters).
//
// - phi (float64): Angle of internal friction of the soil (in degrees).
//
// Returns:
//
// - dc (float64): Depth correction factor for cohesion.
//
// - dq (float64): Depth correction factor for the bearing capacity related to the depth of the foundation.
//
// - dg (float64): Depth correction factor for the weight of the soil.
//
// Usage Example:
//
// dc, dq, dg := calcDepthFactors(3.0, 2.0, 30)
func calcDepthFactors(Df, B, phi float64) (float64, float64, float64) {
	var dc, dq, dg, DB float64

	if Df/B <= 1 {
		DB = Df / B
	} else {
		DB = math.Atan(Df / B)
	}

	dc = 1 + 0.4*DB
	dq = 1 + 2*math.Tan(pkg.Radian(phi))*math.Pow(1-math.Sin(pkg.Radian(phi)), 2)*DB
	dg = 1

	return dc, dq, dg
}
