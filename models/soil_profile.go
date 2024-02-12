package models

import (
	"reflect"

	np "github.com/geoport/numpy4go/vectors"
)

type SoilLayer struct {
	SoilClass                        string  `json:"soilClass"`
	Depth                            float64 `json:"depth"`  // meter
	Center                           float64 `json:"center"` // meter
	IsCohesive                       bool    `json:"isCohesive"`
	DampingRatio                     float64 `json:"dampingRatio"`           // %
	Thickness                        float64 `json:"thickness"`              // meter
	DryUnitWeight                    float64 `json:"dryUnitWeight"`          // t/m^3
	SaturatedUnitWeight              float64 `json:"saturatedUnitWeight"`    // t/m^3
	FineContent                      float64 `json:"fineContent"`            // %
	LiquidLimit                      float64 `json:"liquidLimit"`            // %
	PlasticLimit                     float64 `json:"plasticLimit"`           // %
	PlasticityIndex                  float64 `json:"plasticityIndex"`        // %
	UndrainedShearStrength           float64 `json:"undrainedShearStrength"` // t/m^2
	Cohesion                         float64 `json:"cohesion"`               // t/m^2
	FrictionAngle                    float64 `json:"frictionAngle"`          // degrees
	EffectiveFrictionAngle           float64 `json:"effectiveFrictionAngle"` // degrees
	WaterContent                     float64 `json:"waterContent"`           // %
	PoissonsRatio                    float64 `json:"poissonsRatio"`
	ElasticModulus                   float64 `json:"elasticModulus"` // t/m^2
	VoidRatio                        float64 `json:"voidRatio"`
	RecompressionIndex               float64 `json:"recompressionIndex"`
	CompressionIndex                 float64 `json:"compressionIndex"`
	PreconsolidationPressure         float64 `json:"preconsolidationPressure"` // t/m^2
	VolumeCompressibilityCoefficient float64 `json:"volumeCompressibilityCoefficient"`
	ShearWaveVelocity                float64 `json:"shearWaveVelocity"` // m/s
	Spt_N                            int64   `json:"spt_N"`
	ConeResistance                   float64 `json:"coneResistance"`
	RQD                              float64 `json:"RQD"`
	IS50                             float64 `json:"IS50"`
	Kp                               float64 `json:"kp"`
}

type SoilProfile struct {
	Layers []SoilLayer `json:"layers"`
	Gwt    float64     `json:"gwt"` // meter
}

func NewSoilProfile(layers []SoilLayer, gwt float64) SoilProfile {
	soilProfile := SoilProfile{
		Layers: layers,
		Gwt:    gwt,
	}
	soilProfile.CalcLayerDepths()

	return soilProfile
}

// GetLayerFields returns the fields of a soil layer
func (sp *SoilProfile) GetLayerFields() []string {
	var fields []string
	nonLayerFields := []string{
		"Spt_N", "ConeResistance", "PorePressure", "state", "sizeCache", "unknownFields",
	}
	val := reflect.ValueOf(&sp.Layers[0]).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i).Name
		if !np.Contains(nonLayerFields, field) {
			fields = append(fields, field)
		}
	}
	return fields
}

// CalcLayerDepths calculates center and bottom depth of each layer and inserts them to the object.
func (sp *SoilProfile) CalcLayerDepths() {
	if len(sp.Layers) == 0 {
		return
	}

	bottom := sp.Layers[0].Thickness
	center := bottom / 2

	for i, layer := range sp.Layers {
		if i > 0 {
			thickness := layer.Thickness
			center = bottom + thickness/2
			bottom += thickness
		}
		sp.Layers[i].Center = center
		sp.Layers[i].Depth = bottom
	}
}

func (sp *SoilProfile) GetLayerDepths() []float64 {
	depths := make([]float64, len(sp.Layers))

	for i, layer := range sp.Layers {
		depths[i] = layer.Depth
	}
	return depths
}

// SliceProfile
func (sp *SoilProfile) SliceProfile(minDepth, maxDepth float64) SoilProfile {
	var slicedProfile SoilProfile
	var currentDepth float64

	for _, layer := range sp.Layers {
		currentDepth += layer.Thickness
		// Check if the current layer is within the min and max depth range
		if currentDepth > minDepth && (currentDepth-layer.Thickness) < maxDepth {
			newLayer := layer // Create a copy of the current layer

			// Adjust the top of the first included layer
			if (currentDepth - layer.Thickness) < minDepth {
				newLayer.Thickness -= minDepth - (currentDepth - layer.Thickness)
			}

			// Adjust the bottom of the last included layer
			if currentDepth > maxDepth {
				newLayer.Thickness -= currentDepth - maxDepth
			}

			// Update the center depth of the newLayer based on its new position and thickness

			slicedProfile.Layers = append(slicedProfile.Layers, newLayer)
		}
	}

	slicedProfile.Gwt = sp.Gwt

	slicedProfile.CalcLayerDepths()

	return slicedProfile
}

// GetLayerIndex returns the index of the layer that contains the given depth
func (sp *SoilProfile) GetLayerIndex(depth float64) int {
	layerDepths := sp.GetLayerDepths()
	if layerDepths[len(layerDepths)-1] < depth {
		return len(layerDepths) - 1
	}
	if len(layerDepths) == 1 || depth <= layerDepths[0] {
		return 0
	} else if depth >= layerDepths[len(layerDepths)-1] {
		return len(layerDepths) - 1
	} else {
		diff := np.SumWith(layerDepths, -depth)
		for _, i := range np.Arange(1, float64(len(layerDepths)), 1) {
			prevDiff := diff[int(i-1)]
			currDiff := diff[int(i)]
			if currDiff == 0 {
				return int(i)
			}
			if prevDiff < 0 && currDiff > 0 {
				return int(i)
			}
		}
	}
	return 0
}

// CalcNormalStress returns the normal stress at the given depth
func (sp *SoilProfile) CalcNormalStress(depth float64) float64 {
	stresses := []float64{0}
	layerIndex := sp.GetLayerIndex(depth)

	var H1, H0, H float64
	for i := range np.Arange(0, float64(layerIndex+1), 1) {
		layer := sp.Layers[i]
		if i == layerIndex {
			H1 = depth
		} else {
			H1 = layer.Depth
		}

		if i == 0 {
			H0 = 0
		} else {
			H0 = sp.Layers[i-1].Depth
		}

		if sp.Gwt >= H1 {
			H = H1
		} else if H0 >= sp.Gwt {
			H = H0
		} else {
			H = sp.Gwt
		}

		gammaDry := layer.DryUnitWeight
		gammaSat := layer.SaturatedUnitWeight

		stress := (H-H0)*gammaDry + gammaSat*(H1-H)
		stresses = append(stresses, stress+stresses[i])
	}
	return stresses[len(stresses)-1]
}

func (sp *SoilProfile) Copy() SoilProfile {
	newProfile := SoilProfile{
		Layers: make([]SoilLayer, len(sp.Layers)),
		Gwt:    sp.Gwt,
	}

	copy(newProfile.Layers, sp.Layers)

	return newProfile
}
