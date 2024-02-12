package models

type Foundation struct {
	FoundationDepth            float64 `json:"foundationDepth"`
	FoundationBaseAngle        float64 `json:"foundationBaseAngle"`
	FoundationWidth            float64 `json:"foundationWidth"`
	FoundationLength           float64 `json:"foundationLength"`
	SurfaceFrictionCoefficient float64 `json:"surfaceFrictionCoefficient"`
	SlopeAngle                 float64 `json:"slopeAngle"`
}
