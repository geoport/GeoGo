package meyerhof

import (
	"testing"

	pkg "github.com/geoport/GeoGo/internal"
)

func TestCalcBearingCapacityFactors(t *testing.T) {
	expectedNc := 30.14
	expectedNq := 18.4
	expectedNg := 20.09
	Nc, Nq, Ng := calcBearingCapacityFactors(30)

	if !pkg.AssertFloat(Nc, expectedNc, 0.01) {
		t.Errorf("Got %v, want %v for Nc", Nc, expectedNc)
	}
	if !pkg.AssertFloat(Nq, expectedNq, 0.01) {
		t.Errorf("Got %v, want %v for Nq", Nq, expectedNq)
	}
	if !pkg.AssertFloat(Ng, expectedNg, 0.01) {
		t.Errorf("Got %v, want %v for Ng", Ng, expectedNg)
	}
}

func TestCalcShapeFactors(t *testing.T) {
	expectedSc := 1.29
	expectedSq := 1.29
	expectedSg := 0.8

	Sc, Sq, Sg := calcShapeFactors(10, 20, 30)

	if !pkg.AssertFloat(Sc, expectedSc, 0.01) {
		t.Errorf("Got %v, want %v for Sc", Sc, expectedSc)
	}
	if !pkg.AssertFloat(Sq, expectedSq, 0.01) {
		t.Errorf("Got %v, want %v for Sq", Sq, expectedSq)
	}
	if !pkg.AssertFloat(Sg, expectedSg, 0.01) {
		t.Errorf("Got %v, want %v for Sg", Sg, expectedSg)
	}
}

func TestCalcLoadInclinationFactors(t *testing.T) {
	expectedIc := 0.977
	expectedIq := 0.978
	expectedIg := 0.965

	Ic, Iq, Ig := calcLoadInclinationFactors(30, 10, 20, 150, 50)

	if !pkg.AssertFloat(Ic, expectedIc, 0.01) {
		t.Errorf("Got %v, want %v for Ic", Ic, expectedIc)
	}
	if !pkg.AssertFloat(Iq, expectedIq, 0.01) {
		t.Errorf("Got %v, want %v for Iq", Iq, expectedIq)
	}
	if !pkg.AssertFloat(Ig, expectedIg, 0.01) {
		t.Errorf("Got %v, want %v for Ig", Ig, expectedIg)
	}
}

func TestCalcDepthFactors(t *testing.T) {
	expectedDc1 := 1.2
	expectedDc2 := 1.39
	expectedDq1 := 1.14
	expectedDq2 := 1.28

	Dc1, Dq1, _ := calcDepthFactors(5, 10, 30)
	Dc2, Dq2, _ := calcDepthFactors(15, 10, 30)

	if !pkg.AssertFloat(Dc1, expectedDc1, 0.01) {
		t.Errorf("Got %v, want %v for Dc1", Dc1, expectedDc1)
	}
	if !pkg.AssertFloat(Dq1, expectedDq1, 0.01) {
		t.Errorf("Got %v, want %v for Dq1", Dq1, expectedDq1)
	}

	if !pkg.AssertFloat(Dc2, expectedDc2, 0.01) {
		t.Errorf("Got %v, want %v for Dc1", Dc2, expectedDc1)
	}
	if !pkg.AssertFloat(Dq2, expectedDq2, 0.01) {
		t.Errorf("Got %v, want %v for Dq1", Dq2, expectedDq1)
	}

}
