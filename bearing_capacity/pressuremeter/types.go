package Pressuremeter

type Result struct {
	Kp                       float64 `json:"kp"`
	AllowableBearingCapacity float64 `json:"allowableBearingCapacity"`
	IsSafe                   bool    `json:"isSafe"`
	EffectivePressure        float64 `json:"effectivePressure"`
	NetEffectivePressure     float64 `json:"netEffectivePressure"`
}
