package Pressuremeter

import (
	"testing"

	dt "github.com/geoport/GeoGo/data"

	np "github.com/geoport/numpy4go/vectors"
)

func TestCalcEffectivePressure(t *testing.T) {
	expectedEffectivePressure := 47.85
	expectedNetEffectivePressure := 42.90
	outputEffectivePressure, outputNetEffectivePressure := calcEffectivePressure(dt.Pressuremeter)
	outputEffectivePressure = np.Round(outputEffectivePressure, 2).(float64)
	outputNetEffectivePressure = np.Round(outputNetEffectivePressure, 2).(float64)
	if outputEffectivePressure != expectedEffectivePressure {
		t.Errorf("Expected : %f, Actual : %f", expectedEffectivePressure, outputEffectivePressure)
	}
	if outputNetEffectivePressure != expectedNetEffectivePressure {
		t.Errorf("Expected : %f, Actual : %f", expectedNetEffectivePressure, outputNetEffectivePressure)
	}
}

func TestGetKp(t *testing.T) {
	testSoilData := dt.SoilProfile.Copy()
	TestFoundationData := dt.FoundationData
	testSoilData.CalcLayerDepths()

	expected := 0.84
	output := np.Round(getKp(TestFoundationData, testSoilData, 47.848), 2)
	if output != expected {
		t.Errorf("Expected Kp: %f, Actual Kp: %f", expected, output)
	}
}

func TestCalcBearingCapacity(t *testing.T) {
	testSoilData := dt.SoilProfile.Copy()
	testFoundationData := dt.FoundationData
	testSoilData.CalcLayerDepths()
	testPMData := dt.Pressuremeter

	expectedQuNet := 36.21
	output := CalcBearingCapacity(testSoilData, testFoundationData, 20, testPMData)
	outputQuNet := np.Round(output.AllowableBearingCapacity, 2)
	if outputQuNet != expectedQuNet {
		t.Errorf("Expected QuNet: %f, Actual QuNet: %f", expectedQuNet, outputQuNet)
	}
}
