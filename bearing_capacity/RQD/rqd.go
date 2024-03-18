package RQD

import (
	"github.com/geoport/GeoGo/models"

	"math"
)

func getQfRatio(rqd float64) float64 {
	if rqd < 0.25 {
		return 0.15
	} else if 0.25 <= rqd && rqd < 0.50 {
		return 0.2
	} else if 0.50 <= rqd && rqd < 0.75 {
		return 0.25

	} else if 0.75 <= rqd && rqd < 0.90 {
		return 0.3 + (rqd-0.75)*0.4/0.15

	} else {
		return math.Max(1, 0.7+(rqd-0.9)*0.3)
	}
}

// CalcBearingCapacity is a function that returns ultimate and allowable bearing capacity, qf and qLab.
func CalcBearingCapacity(
	soilProfile models.SoilProfile, foundationData models.Foundation, foundationPressure float64,
) models.RQD {
	Df := foundationData.FoundationDepth
	layerIndex := soilProfile.GetLayerIndex(Df)
	layer := soilProfile.Layers[layerIndex]

	rqd := layer.RQD / 100
	kp := layer.Kp
	IS50 := layer.IS50 // t/m2

	qfRatio := getQfRatio(rqd)

	qLab := kp * IS50
	ultimateBearingCapacity := qLab * qfRatio
	allowableBearingCapacity := ultimateBearingCapacity / 7

	result := models.RQD{
		UltimateBearingCapacity:  ultimateBearingCapacity,
		AllowableBearingCapacity: allowableBearingCapacity,
		QFRatio:                  qfRatio,
		QLab:                     qLab,
		Kp:                       kp,
		Is50:                     IS50,
		IsSafe:                   allowableBearingCapacity >= foundationPressure,
		RQD:                      rqd,
	}

	return result
}
