package models

type VesicBearingCapacity struct {
	Nq                       float64 `json:"Nq"`
	Nc                       float64 `json:"Nc"`
	Ng                       float64 `json:"Ng"`
	Sq                       float64 `json:"Sq"`
	Sc                       float64 `json:"Sc"`
	Sg                       float64 `json:"Sg"`
	Dq                       float64 `json:"Dq"`
	Dc                       float64 `json:"Dc"`
	Dg                       float64 `json:"Dg"`
	Iq                       float64 `json:"Iq"`
	Ic                       float64 `json:"Ic"`
	Ig                       float64 `json:"Ig"`
	Gq                       float64 `json:"Gq"`
	Gc                       float64 `json:"Gc"`
	Gg                       float64 `json:"Gg"`
	Bq                       float64 `json:"Bq"`
	Bc                       float64 `json:"Bc"`
	Bg                       float64 `json:"Bg"`
	UnitWeight               float64 `json:"unitWeight"`
	Cohesion                 float64 `json:"cohesion"`
	FrictionAngle            float64 `json:"frictionAngle"`
	UltimateBearingCapacity  float64 `json:"ultimateBearingCapacity"`
	AllowableBearingCapacity float64 `json:"allowableBearingCapacity"`
	IsSafe                   bool    `json:"isSafe"`
}
