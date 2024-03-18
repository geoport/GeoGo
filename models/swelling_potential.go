package models

type SwellingPotential struct {
	LayerCenters          []float64 `json:"layerCenters"`
	EffectiveStresses     []float64 `json:"effectiveStresses"`
	DeltaSigmas           []float64 `json:"deltaSigmas"`
	SwellingPressures     []float64 `json:"swellingPressures"`
	IsSafe                []bool    `json:"isSafe"`
	NetFoundationPressure float64   `json:"netFoundationPressure"`
}
