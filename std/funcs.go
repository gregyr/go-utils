package std

import "math"

// allows inline if statements
func If[T any](cond bool, returnValue T, elseValue T) T {
	if cond {
		return returnValue
	} else {
		return elseValue
	}
}

// executes a function if a condition is true, if the condition is false the else function will be executed
func IfFunc[T any](cond bool, thenFn func() T, elseFn func() T) T {
	if cond {
		return thenFn()
	}
	return elseFn()
}

// when provided with a value and an error, this functions returns either the value if the error is nil, otherwise the default value for the type
func UnwrapOrDefault[T any](value T, err error) T {
	if err != nil {
		return *new(T)
	}
	return value
}

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

// provides a slice over a range
func Range(n int) []int {
	r := make([]int, n)
	for i := range n {
		r[i] = i
	}
	return r
}

// checks whether all the values of a slice equal one given value
func AllEqual[T comparable](s []T, t T) bool {
	for _, e := range s {
		if t != e {
			return false
		}
	}
	return true
}

func CeilToInt(x float64) int {
	return int(math.Ceil(x))
}

func FloorToInt(x float64) int {
	return int(math.Floor(x))
}
