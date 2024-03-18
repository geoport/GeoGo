package models

import (
	"reflect"
	"testing"

	pkg "github.com/geoport/GeoGo/internal"

	np "github.com/geoport/numpy4go/vectors"
)

var soilProfile = SoilProfile{
	Layers: []SoilLayer{
		{SoilClass: "SC", DryUnitWeight: 1.8, SaturatedUnitWeight: 2, Thickness: 1},
		{SoilClass: "SP", DryUnitWeight: 1.7, SaturatedUnitWeight: 2.1, Thickness: 1.4},
		{SoilClass: "SM", DryUnitWeight: 1.9, SaturatedUnitWeight: 2.2, Thickness: 3.4},
	},
	Gwt: 1,
}

func TestSoilProile_GetLayerFields(t *testing.T) {
	expected := []string{
		"SoilClass",
		"Depth",
		"Center",
		"IsCohesive",
		"DampingRatio",
		"Thickness",
		"DryUnitWeight",
		"SaturatedUnitWeight",
		"FineContent",
		"LiquidLimit",
		"PlasticLimit",
		"PlasticityIndex",
		"UndrainedShearStrength",
		"Cohesion",
		"FrictionAngle",
		"EffectiveFrictionAngle",
		"WaterContent",
		"PoissonsRatio",
		"ElasticModulus",
		"VoidRatio",
		"RecompressionIndex",
		"CompressionIndex",
		"PreconsolidationPressure",
		"VolumeCompressibilityCoefficient",
		"ShearWaveVelocity",
		"RQD",
		"IS50",
		"Kp",
	}

	output := soilProfile.GetLayerFields()
	if reflect.DeepEqual(output, expected) == false {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}

func TestSoilProfile_CalcLayerDepths(t *testing.T) {
	soilProfile.CalcLayerDepths()
	l := soilProfile.Layers[2]

	if !pkg.AssertFloat(l.Center, 4.1, 0.1) {
		t.Errorf("Expected %v, got %v", 4.1, l.Center)
	}

	if !pkg.AssertFloat(l.Depth, 5.8, 0.1) {
		t.Errorf("Expected %v, got %v", 5.8, l.Depth)
	}
}

func TestSoilProfile_GetLayerDepths(t *testing.T) {
	soilProfile.CalcLayerDepths()
	depths := soilProfile.GetLayerDepths()

	expected := []float64{1, 2.4, 5.8}

	if !pkg.AssertFloatArray(depths, expected, 0.1) {
		t.Errorf("Expected %v, got %v", expected, depths)
	}

}

func TestSoilProfile_GetLayerIndex(t *testing.T) {
	soilProfile.CalcLayerDepths()

	testInputs := []float64{-1, 2.4, 7, 3}
	expectedOutputs := []int{0, 1, 2, 2}

	for i, inp := range testInputs {
		output := soilProfile.GetLayerIndex(inp)
		if output != expectedOutputs[i] {
			t.Errorf("Expected %v, got %v", expectedOutputs[i], output)
		}
	}
}

func TestSoilProfile_SliceProfile(t *testing.T) {
	soilProfile.CalcLayerDepths()

	newProfile := soilProfile.SliceProfile(2, 4)

	if len(newProfile.Layers) != 2 {
		t.Errorf("Expected %v layers, got %v layers", 2, len(newProfile.Layers))
	}

	if !pkg.AssertFloat(newProfile.Layers[0].Thickness, 0.4, 0.1) {
		t.Errorf("First layer thickness expected to be %v , but got %v ", 0.4, newProfile.Layers[0].Depth)
	}

	if !pkg.AssertFloat(newProfile.Layers[0].DryUnitWeight, 1.7, 0.1) {
		t.Errorf("First layer dry unit weight expected to be %v , but got %v ", 1.7, newProfile.Layers[0].DryUnitWeight)
	}
}

func TestSoilProfile_CalcNormalStress(t *testing.T) {
	soilProfile.CalcLayerDepths()
	SP1 := soilProfile.Copy()
	SP2 := soilProfile.Copy()
	SP3 := soilProfile.Copy()
	SP2.Gwt = 0.5
	SP3.Gwt = 10

	checkPoints := []float64{0, 1, 1.5, 3, 8}
	expectedOutputs1 := []float64{0, 1.8, 2.85, 6.06, 17.06}
	expectedOutputs2 := []float64{0, 1.9, 2.95, 6.16, 17.16}
	expectedOutputs3 := []float64{0, 1.8, 2.65, 5.32, 14.82}

	output1 := np.Apply(checkPoints, SP1.CalcNormalStress)
	output2 := np.Apply(checkPoints, SP2.CalcNormalStress)
	output3 := np.Apply(checkPoints, SP3.CalcNormalStress)

	if reflect.DeepEqual(np.Round(output1, 2), expectedOutputs1) == false {
		t.Errorf("Expected %v, got %v", expectedOutputs1, output1)
	}
	if reflect.DeepEqual(np.Round(output2, 2), expectedOutputs2) == false {
		t.Errorf("Expected %v, got %v", expectedOutputs2, output2)
	}
	if reflect.DeepEqual(np.Round(output3, 2), expectedOutputs3) == false {
		t.Errorf("Expected %v, got %v", expectedOutputs3, output3)
	}
}

func TestSoilProfile_CalcEffectiveStress(t *testing.T) {
	soilProfile.CalcLayerDepths()
	checkPoints := []float64{0.5, 1.5}

	expectedOutputs := []float64{0.9, 2.36}

	output := np.Apply(checkPoints, soilProfile.CalcEffectiveStress)

	if reflect.DeepEqual(np.Round(output, 2), expectedOutputs) == false {
		t.Errorf("Expected %v, got %v", expectedOutputs, output)
	}
}
