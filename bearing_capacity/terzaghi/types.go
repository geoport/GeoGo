package terzaghi

type Result struct {
	BearingCapacityFactors  BearingCapacityFactors `json:"bearingCapacityFactors"`
	SoilParams              BCSoilParams           `json:"soilParams"`
	UltimateBearingCapacity float64                `json:"ultimateBearingCapacity"`
}

type BCSoilParams struct {
	UnitWeight    float64 `json:"unitWeight"`
	Cohesion      float64 `json:"cohesion"`
	FrictionAngle float64 `json:"frictionAngle"`
}

type BearingCapacityFactors struct {
	Nq float64 `json:"Nq"`
	Nc float64 `json:"Nc"`
	Ng float64 `json:"Ng"`
}
