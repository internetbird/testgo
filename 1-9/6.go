package main

import "math"

func Euler6() {
	//Project Euler Problem #6
	sumOfSquares := 0
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
		sumOfSquares += int(math.Pow(float64(i), 2))

	}
	squareOfSums := int(math.Pow(float64(sum), 2))

	PrintAnswer(6, squareOfSums-sumOfSquares)

}
