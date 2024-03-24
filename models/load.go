package models

type Load struct {
	FoundationPressure float64 `json:"foundationPressure"`
	VerticalLoad       float64 `json:"verticalLoad"`
	HorizontalLoadX    float64 `json:"horizontalLoadX"`
	HorizontalLoadY    float64 `json:"horizontalLoadY"`
	MomentLoadX        float64 `json:"momentLoadX"`
	MomentLoadY        float64 `json:"momentLoadY"`
}
