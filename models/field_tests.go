package models

type PressuremeterExp struct {
	Depth            float64 `json:"depth"`
	LimitPressure    float64 `json:"limitPressure"`
	NetLimitPressure float64 `json:"netLimitPressure"`
}

type Pressuremeter struct {
	Exps []PressuremeterExp `json:"exps"`
}

type MaswExp struct {
	Thickness               float64 `json:"thickness"`
	ShearWaveVelocity       float64 `json:"shearWaveVelocity"`
	CompressionWaveVelocity float64 `json:"compressionWaveVelocity"`
}

type MASW struct {
	Exps []MaswExp `json:"exps"`
}

type CptExp struct {
	Depth          float64 `json:"depth"`
	ConeResistance float64 `json:"coneResistance"`
	PorePressure   float64 `json:"porePressure"`
}

type CPT struct {
	Exps []CptExp `json:"exps"`
}

type SptExp struct {
	Depth       float64 `json:"depth"`
	N           int     `json:"N"`
	N60         int     `json:"N60"`
	N160        int     `json:"N160"`
	N160F       int     `json:"N160F"`
	Cr          float64 `json:"Cr"`
	Cn          float64 `json:"Cn"`
	FineContent float64 `json:"fineContent"`
	Alpha       float64 `json:"alpha"`
	Beta        float64 `json:"beta"`
}

type SPT struct {
	Exps                     []SptExp `json:"exps"`
	EnergyCorrectionFactor   float64  `json:"energyCorrectionFactor"`
	DiameterCorrectionFactor float64  `json:"diameterCorrectionFactor"`
	SamplerCorrectionFactor  float64  `json:"samplerCorrectionFactor"`
	MakeCorrection           bool     `json:"makeCorrection"`
	AverageN                 int64    `json:"averageN"`
}
