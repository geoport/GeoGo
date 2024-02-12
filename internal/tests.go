package internal

import "math"

// AssertFloat returns true if the tolerance value is equal or higher than the difference of given values
func AssertFloat(var1, var2, tolerance float64) bool {
	diff := math.Abs(var1 - var2)

	return diff <= tolerance
}

// AssertFloatArray returns true if the tolerance value is equal or higher than the difference of each array elements
func AssertFloatArray(arr1, arr2 []float64, tolerance float64) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if !AssertFloat(arr1[i], arr2[i], tolerance) {
			return false
		}
	}

	return true
}
