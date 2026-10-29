package fn

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

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

func AtoiOrZero(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

// Logs an error on functions that only return an error
func LogIfError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ExitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
