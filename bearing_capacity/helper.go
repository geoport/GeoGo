package bearingcapacity

import (
	"github.com/geoport/GeoGo/models"

	np "github.com/geoport/numpy4go/vectors"
)

// CalcEffectiveUnitWeight calculates the effective unit weight of a soil profile between surface and Df + B depth.
//
// Parameters:
//
// -Df (float64) : Depth of the foundation in meters.
//
// -B (float64) : Width of the foundation in meters.
//
// -soilProfile (SoilProfile)
//
// -term (string) : Defines the type of analysis (short || long)
//
// Returns:
//
// -effectiveUnitWeight (float64) : Effective unit weight to use in bearing capacity analysis.
func CalcEffectiveUnitWeight(Df, B float64, soilProfile models.SoilProfile, term string) float64 {
	var gwt float64
	if term == "short" {
		gwt = soilProfile.Gwt
	} else {
		gwt = Df + B
	}
	profileView := soilProfile.SliceProfile(0, Df+B)

	gwtLayerIndex := profileView.GetLayerIndex(gwt)

	var weightedUnitWeights []float64

	for i, layer := range profileView.Layers {
		if i < gwtLayerIndex || gwt >= Df+B {
			unitWeight := layer.DryUnitWeight
			thickness := layer.Thickness
			weightedUnitWeights = append(weightedUnitWeights, thickness*unitWeight)
		} else if i > gwtLayerIndex {
			unitWeight := layer.SaturatedUnitWeight
			thickness := layer.Thickness
			weightedUnitWeights = append(weightedUnitWeights, thickness*unitWeight)
		} else {
			unitWeightUpper := layer.DryUnitWeight
			var thicknessUpper float64
			if i == 0 {
				thicknessUpper = gwt
			} else {
				thicknessUpper = gwt - profileView.Layers[i-1].Depth
			}
			weightedUnitWeights = append(weightedUnitWeights, thicknessUpper*unitWeightUpper)

			unitWeightLower := layer.SaturatedUnitWeight
			thicknessLower := layer.Depth - gwt
			weightedUnitWeights = append(weightedUnitWeights, thicknessLower*unitWeightLower)
		}
	}

	viewLayerDepths := profileView.GetLayerDepths()
	maxDepth := viewLayerDepths[len(viewLayerDepths)-1]
	averageUnitWeight := np.Sum(weightedUnitWeights) / maxDepth

	var effectiveUnitWeight float64
	if gwt <= Df {
		effectiveUnitWeight = averageUnitWeight - 1
	} else if gwt > Df+B {
		effectiveUnitWeight = averageUnitWeight
	} else {
		effectiveUnitWeight = averageUnitWeight - (1 - (gwt-Df)/B)
	}

	return effectiveUnitWeight
}

// GetSoilParams returns long term or short term soil parameters to use in bearing capacity analysis.
//
// Parameters:
//
// -layerDepth (float64) : The depth of the soil layer where the soil parameters will be obtained.
//
// -soilProfile (SoilProfile)
//
// -term (string) : Defines the type of analysis (short || long)
//
// Returns:
//
// -cohesion (float64) : Cohesion value to use in bearing capacity analysis.
//
// -frictionAngle (float64) : Friction angle value to use in bearing capacity analysis.
func GetSoilParams(layerDepth float64, soilProfile models.SoilProfile, term string) (float64, float64) {
	gwt := soilProfile.Gwt
	layerIndex := soilProfile.GetLayerIndex(layerDepth)
	layer := soilProfile.Layers[layerIndex]

	var cohesion, frictionAngle float64

	if gwt <= layerDepth && term == "short" {
		cohesion = layer.UndrainedShearStrength
		frictionAngle = layer.FrictionAngle
	} else {
		cohesion = layer.Cohesion
		frictionAngle = layer.EffectiveFrictionAngle
	}

	return cohesion, frictionAngle
}

// CalcStress calculates the normal or effective stress at a given depth depending on the term.
//
// Parameters:
//
// -soilProfile (SoilProfile)
//
// -Df (float64) : Depth of the foundation in meters.
//
// -term (string) : Defines the type of analysis (short || long)
//
// Returns:
//
// -stress (float64) : Normal stress in short term or effective stress in long term at a given depth.
func CalcStress(soilProfile models.SoilProfile, Df float64, term string) float64 {
	var stress float64

	if term == "short" {
		stress = soilProfile.CalcNormalStress(Df)
	} else {
		stress = soilProfile.CalcEffectiveStress(Df)
	}

	return stress
}

// CalcEccentricity calculates the eccentricity of the vertical load.
//
// Parameters:
//
// -moment (float64) : Moment acting on the foundation in ton.m.
//
// -verticalLoad (float64) : Vertical load acting on the foundation in ton.
//
// Returns:
//
// -eccentricity (float64) : Eccentricity of the vertical load in meters.
func CalcEccentricity(moment, verticalLoad float64) float64 {
	return moment / verticalLoad
}

// CalcEffectiveDimensions calculates the effective dimensions of the foundation.
//
// Parameters:
//
// -foundation (Foundation)
//
// -loads (Load)
//
// Returns:
//
// -B_ (float64) : Effective width of the foundation in meters.
//
// -L_ (float64) : Effective length of the foundation in meters.
func CalcEffectiveDimensions(foundation models.Foundation, loads models.Load) (float64, float64) {
	B := foundation.FoundationWidth
	L := foundation.FoundationLength

	Mx := loads.MomentLoadX
	My := loads.MomentLoadY

	P := loads.VerticalLoad

	ex := CalcEccentricity(Mx, P)
	ey := CalcEccentricity(My, P)

	B_ := B - 2*ex
	L_ := L - 2*ey

	if B_ > L_ {
		return L_, B_
	} else {
		return B_, L_
	}
}
