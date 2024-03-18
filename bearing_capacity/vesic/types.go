package vesic

type Result struct {
	BearingCapacityFactors   BearingCapacityFactors `json:"bearingCapacityFactors"`
	ShapeFactors             ShapeFactors           `json:"shapeFactors"`
	DepthFactors             DepthFactors           `json:"depthFactors"`
	LoadInclinationFactors   LoadInclinationFactors `json:"loadInclinationFactors"`
	GroundFactors            GroundFactors          `json:"groundFactors"`
	BaseFactors              BaseFactors            `json:"baseFactors"`
	SoilParams               BCSoilParams           `json:"soilParams"`
	UltimateBearingCapacity  float64                `json:"ultimateBearingCapacity"`
	AllowableBearingCapacity float64                `json:"allowableBearingCapacity"`
	IsSafe                   bool                   `json:"isSafe"`
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

type ShapeFactors struct {
	Sq float64 `json:"Sq"`
	Sc float64 `json:"Sc"`
	Sg float64 `json:"Sg"`
}

type DepthFactors struct {
	Dq float64 `json:"Dq"`
	Dc float64 `json:"Dc"`
	Dg float64 `json:"Dg"`
}

type LoadInclinationFactors struct {
	Iq float64 `json:"Iq"`
	Ic float64 `json:"Ic"`
	Ig float64 `json:"Ig"`
}

type GroundFactors struct {
	Gq float64 `json:"Gq"`
	Gc float64 `json:"Gc"`
	Gg float64 `json:"Gg"`
}

type BaseFactors struct {
	Bq float64 `json:"Bq"`
	Bc float64 `json:"Bc"`
	Bg float64 `json:"Bg"`
}
