package effective_depth

import (
	"testing"

	dt "github.com/geoport/GeoGo/data"

	"github.com/geoport/GeoGo/internal"
)

func TestCalcEffectiveDepth(t *testing.T) {
	soilProfile := dt.SoilProfile.Copy()
	soilProfile.CalcLayerDepths()

	result := CalcEffectiveDepth(soilProfile, dt.FoundationData, 50)
	expected := 34.41

	if !internal.AssertFloat(result, expected, 0.01) {
		t.Errorf("Found %v: Expected %v", result, expected)
	}
}
