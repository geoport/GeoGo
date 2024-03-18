package Pressuremeter

import (
	"github.com/geoport/GeoGo/models"

	"math"

	np "github.com/geoport/numpy4go/vectors"
)

func calcEffectivePressure(psData models.Pressuremeter) (float64, float64) {
	effectivePressure := 1.
	netEffectivePressure := 1.

	numOfData := float64(len(psData.Exps))
	for _, exp := range psData.Exps {
		effectivePressure *= exp.LimitPressure
		netEffectivePressure *= exp.NetLimitPressure
	}

	effectivePressure = math.Pow(effectivePressure, 1/numOfData)
	netEffectivePressure = math.Pow(netEffectivePressure, 1/numOfData)

	return effectivePressure, netEffectivePressure

}

func getKp(foundationData models.Foundation, soilProfile models.SoilProfile, effectivePressure float64) float64 {
	Df := foundationData.FoundationDepth
	B := foundationData.FoundationWidth
	L := foundationData.FoundationLength

	layerIndex := soilProfile.GetLayerIndex(Df)
	layer := soilProfile.Layers[layerIndex]

	soilClass := layer.SoilClass
	BL := B / L
	DL := Df / B

	var c, kp float64

	if np.Contains([]string{"CH", "CL", "MH", "ML"}, soilClass) {
		if effectivePressure < 122.365 {
			c = 0.25
		} else if 122.365 <= effectivePressure && effectivePressure < 203.94 {
			c = 0.35
		} else {
			c = 0.5
		}
		kp = 0.8 * math.Pow(1+c*(0.6+math.Pow(0.4, BL)), DL)
	} else {
		if effectivePressure < 101.97 {
			c = 0.35
		} else if 101.97 <= effectivePressure && effectivePressure < 203.94 {
			c = 0.5
		} else {
			c = 0.8
		}
		kp = math.Pow(1+1+c*(0.6+math.Pow(0.4, BL)), DL)
	}
	return kp
}

func CalcBearingCapacity(
	soilProfile models.SoilProfile, foundationData models.Foundation, foundationPressure float64, psData models.Pressuremeter,
) Result {
	effectivePressure, netEffectivePressure := calcEffectivePressure(psData)
	kp := getKp(foundationData, soilProfile, effectivePressure)

	quNet := kp * netEffectivePressure

	result := Result{
		EffectivePressure:        effectivePressure,
		NetEffectivePressure:     netEffectivePressure,
		AllowableBearingCapacity: quNet,
		IsSafe:                   quNet >= foundationPressure,
	}

	return result
}
