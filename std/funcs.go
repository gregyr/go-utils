package std

// allows inline if statements
func If[T any](cond bool, returnValue T, elseValue T) T {
	if cond {
		return returnValue
	} else {
		return elseValue
	}
}

func IfFunc[T any](cond bool, thenFn func() T, elseFn func() T) T {
	if cond {
		return thenFn()
	}
	return elseFn()
}

func UnwrapOrDefault[T any](value T, err error) T {
	if err != nil {
		return *new(T)
	}
	return value
}

func ClampMin(v, min int) int {
	if v < min {
		return min
	}
	return v
}

func ClampMax(v, max int) int {
	if v > max {
		return max
	}
	return v
}

func Clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
