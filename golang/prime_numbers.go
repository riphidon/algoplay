package main

import (
	"math"
)

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func NthPrime(position int) int {
	num := 2
	count := 2

	for count <= position {
		num++
		if ok := IsPrime(num); ok {
			count++
		}
	}

	return num
}
