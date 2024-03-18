package meyerhof

import (
	"math"

	pkg "github.com/geoport/GeoGo/internal"
)

func calcKp(phi float64) float64 {
	return math.Pow(math.Tan(pkg.Radian(45+phi/2)), 2)
}

// calcShapeFactors calculates the shape factors for foundation design in geotechnical engineering.
//
// Parameters:
//
// - B (float64): Width of the foundation (in meters).
//
// - L (float64): Length of the foundation (in meters).
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
// Sc, Sq, Sg := calcShapeFactors(2.0, 4.0, 35.0)
func calcShapeFactors(B, L, phi float64) (float64, float64, float64) {
	var Sq, Sg float64

	kp := calcKp(phi)

	Sc := 1 + 0.2*kp*(B/L)
	if phi == 0 {
		Sq = 1
		Sg = 1
	} else {
		Sq = 1 + 0.1*kp*(B/L)
		Sg = 1 + 0.1*kp*(B/L)
	}

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
	Ng := (Nq - 1) * math.Tan(pkg.Radian(phi*1.4))

	return Nc, Nq, Ng
}

// calcLoadInclinationFactors calculates the load inclination factors for a foundation based on various parameters including the soil's angle of internal friction, cohesion, foundation dimensions, and applied pressure. These factors adjust the bearing capacity of the foundation considering the effect of inclined loads.
//
// Parameters:
//
// - phi (float64): Angle of internal friction of the soil (in degrees), affecting the soil's shear strength.
//
// - B (float64): Width of the foundation (in meters).
//
// - L (float64): Length of the foundation (in meters).
//
// - Vmax (float64): Maximum horizontal load applied on the foundation (in tons).
//
// - foundationPressure (float64): Pressure applied by the foundation on the soil (in tons per square meter), calculated as the vertical load divided by the foundation area.
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
// ic, iq, ig := calcLoadInclinationFactors(30, 20, 50, 100, 150)
func calcLoadInclinationFactors(
	phi, B, L, Vmax, foundationPressure float64,
) (float64, float64, float64) {
	verticalLoad := foundationPressure * B * L

	delta := math.Atan(Vmax / verticalLoad)

	ig := math.Pow(1-delta/math.Pi, 2)
	ic := math.Pow(1-2*delta/math.Pi, 2)
	iq := ic

	return ic, iq, ig
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
	var dc, dq, dg float64

	kp := calcKp(phi)

	db := Df / B

	dc = 1 + 0.2*db*math.Sqrt(kp)

	if phi == 0 {
		dq = 1
	} else {
		dq = 1 + 0.1*db*math.Sqrt(kp)
	}

	dg = dq

	return dc, dq, dg
}
