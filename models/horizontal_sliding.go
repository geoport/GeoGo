package models

type HorizontalSliding struct {
	Rth     float64 `json:"Rth"`
	Ptv     float64 `json:"Ptv"`
	Ac      float64 `json:"Ac"`
	RpkX    float64 `json:"RpkX"`
	RpkY    float64 `json:"RpkY"`
	RptX    float64 `json:"RptX"`
	RptY    float64 `json:"RptY"`
	SumX    float64 `json:"sumX"`
	SumY    float64 `json:"sumY"`
	IsSafeX bool    `json:"isSafeX"`
	IsSafeY bool    `json:"isSafeY"`
	VthX    float64 `json:"VthX"`
	VthY    float64 `json:"VthY"`
}
