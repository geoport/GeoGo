package HorizontalSliding

import (
	"testing"

	dt "github.com/geoport/GeoGo/data"

	np "github.com/geoport/numpy4go/vectors"
)

func TestCalcHorizontalSliding(t *testing.T) {
	testSoilData := dt.SoilProfile
	testSoilData.CalcLayerDepths()

	testfoundationData := dt.FoundationData

	result := CalcHorizontalSliding(testSoilData, testfoundationData, 50, 150, 150)
	Rth := np.Round(result.Rth, 2)
	rpkX := np.Round(result.RpkX, 2)
	rpkY := np.Round(result.RpkY, 2)
	rptX := np.Round(result.RptX, 2)
	rptY := np.Round(result.RptY, 2)
	sumX := np.Round(result.SumX, 2)
	sumY := np.Round(result.SumY, 2)

	if Rth != 5454.55 {
		t.Errorf("Rth is expected %f, got %f", 5454.55, Rth)
	}
	if rpkX != 76.21 {
		t.Errorf("rpkX is expected %f, got %f", 76.21, rpkX)
	}
	if rpkY != 152.43 {
		t.Errorf("rpkY is expected %f, got %f", 152.43, rpkY)
	}
	if rptX != 54.44 {
		t.Errorf("rptX is expected %f, got %f", 54.44, rptX)
	}
	if rptY != 108.88 {
		t.Errorf("rptY is expected %f, got %f", 108.88, rptY)
	}
	if sumX != 5470.88 {
		t.Errorf("sumX is expected %f, got %f", 5470.88, sumX)
	}
	if sumY != 5487.21 {
		t.Errorf("sumY is expected %f, got %f", 5487.21, sumY)
	}
}
