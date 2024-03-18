package tezcan_ozdemir

type Result struct {
	VS                       float64 `json:"VS"`
	UnitWeight               float64 `json:"unitWeight"`
	UltimateBearingCapacity  float64 `json:"ultimateBearingCapacity"`
	AllowableBearingCapacity float64 `json:"allowableBearingCapacity"`
	IsSafe                   bool    `json:"isSafe"`
	SafetyFactor             float64 `json:"safetyFactor"`
}
