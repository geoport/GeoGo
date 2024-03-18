package RQD

type Result struct {
	Kp                       float64 `json:"kp"`
	Is50                     float64 `json:"Is50"`
	UltimateBearingCapacity  float64 `json:"ultimateBearingCapacity"`
	AllowableBearingCapacity float64 `json:"allowableBearingCapacity"`
	IsSafe                   bool    `json:"isSafe"`
	QLab                     float64 `json:"qLab"`
	QFRatio                  float64 `json:"qfRatio"`
	RQD                      float64 `json:"RQD"`
}
