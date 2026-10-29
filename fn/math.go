package fn

import "math"

// clamps a value to a min bound
func ClampMin(v, min int) int {
	if v < min {
		return min
	}
	return v
}

// clamps a value to a max bound
func ClampMax(v, max int) int {
	if v > max {
		return max
	}
	return v
}

// clamps a value between a min and a max bound
func Clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// returns the Ceil of x as an int, if x is NaN or +-Inf it returns 0
func CeilToInt(x float64) int {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return 0
	}
	return int(math.Ceil(x))
}

// returns the Floor of x as an int, if x is NaN or +-Inf it returns 0
func FloorToInt(x float64) int {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return 0
	}
	return int(math.Floor(x))
}
