package effective_depth

import (
	"github.com/geoport/GeoGo/models"

	"math"
)

// getDifferance returns the difference between delta sigma and %10 of effective stress at given depth
func getDifferance(z float64, F float64, B float64, Df float64, L float64, sp models.SoilProfile) float64 {
	DG := F / ((B + z - Df) * (L + z - Df))
	effectiveStress := sp.CalcEffectiveStress(z)
	return DG - 0.1*effectiveStress
}

// findEffectiveDepth finds the effective depth of the stress by using bisection method
func findEffectiveDepth(F float64, B float64, Df float64, L float64, sp models.SoilProfile) float64 {
	boundary1 := Df
	boundary2 := Df + 1.5*B
	middle := (boundary1 + boundary2) / 2

	n := 0

	// if the stress difference is positive for both boundaries, increase the boundary2
	if getDifferance(boundary1, F, B, Df, L, sp)*getDifferance(boundary2, F, B, Df, L, sp) > 0 {
		boundary2 = 100 * B
	}

	// keep iterating till the difference is less than 0.001 or number of steps is greater than 100
	for math.Abs(getDifferance(middle, F, B, Df, L, sp)) > 0.01 && n < 100 {
		n = n + 1
		if boundary1 == boundary2 && boundary1 == middle && n > 10 {
			return 0
		}
		if getDifferance(middle, F, B, Df, L, sp) > 0 {
			boundary1 = middle
		} else {
			boundary2 = middle
		}
		middle = (boundary1 + boundary2) / 2
	}
	return middle
}

// CalcEffectiveDepth calculates the effective depth of the stress by bisection method and TBDY method
func CalcEffectiveDepth(
	soilProfile models.SoilProfile, foundationData models.Foundation, foundationPressure float64,
) float64 {
	Df := foundationData.FoundationDepth
	B := foundationData.FoundationWidth
	L := foundationData.FoundationLength

	Qnet := foundationPressure - soilProfile.CalcNormalStress(Df)
	F := Qnet * B * L

	effectiveDepth := findEffectiveDepth(F, B, Df, L, soilProfile)

	return effectiveDepth
}
