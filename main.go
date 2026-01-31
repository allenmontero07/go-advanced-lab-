package main

import (
	"errors"
	"math"
)

func factorial(n int) (int, error) {

	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}

	result := 1

	for i := 1; i <= n; i++ {
		result = result * i
	}

	return result, nil
}

func IsPrime(n int) (bool, error) {

	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}

	limit := int(math.Sqrt(float64(n)))

	for i := 2; i <= limit; i++ {

		if n%i == 0 {

			return false, nil

		}
	}

	return true, nil

}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}

	return result, nil
}

func MakeCounter(start int) func() int {
	counter := start
	return func() int {
		counter++
		return counter
	}

}

func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	total := initial

	add = func(x int) {
		total += x
	}

	subtract = func(x int) {
		total -= x
	}

	get = func() int {
		return total
	}

	return add, subtract, get
}
