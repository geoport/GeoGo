package soilcoefficient

func CalcSoilCoefficientBySettlement(settlement float64, foundationLoad float64) float64 {
	if settlement <= 0 {
		return 999999
	}
	return 100 * foundationLoad / settlement // t/m3
}

func CalcSoilCoefficientByBearingCapacity(bearingCapacity float64) float64 {
	return 400 * bearingCapacity // t/m3
}
