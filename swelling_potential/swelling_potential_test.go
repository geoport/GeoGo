package swelling_potential

import (
	"GoPrime/src/TestData"
	"reflect"
	"testing"

	np "github.com/geoport/numpy4go/vectors"
)

func TestCalcSwellingPotential(t *testing.T) {
	expectedPressures := []float64{8.89, 0, 0}
	expectedCheckSafeties := []bool{false, true, true}
	output := CalcSwellingPotential(TestData.TestSoilProfile, TestData.TestFoundationData, TestData.TestLoadingData)
	outputPressures := np.Round(output.SwellingPressures, 2)
	outputCheckSafeties := output.IsSafe

	if reflect.DeepEqual(outputPressures, expectedPressures) == false {
		t.Errorf("Expected pressures: %v, got: %v", expectedPressures, outputPressures)
	}
	if reflect.DeepEqual(outputCheckSafeties, expectedCheckSafeties) == false {
		t.Errorf("Expected check safeties: %v, got: %v", expectedCheckSafeties, outputCheckSafeties)
	}
}
