package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

//PrintAnswer -  Prints a formatted answer for problemNum and answer
func PrintAnswer(problemNum int, answer int) {

	fmt.Printf("The answer to Project Euler Problem #%d is :%d\n", problemNum, answer)
}

func PrintBigIntAnswer(problemNum int, answer *big.Int) {

	fmt.Printf("The answer to Project Euler Problem #%d is :%d\n", problemNum, answer)
}

func PrintStringAnswer(problemNum int, answer string) {

	fmt.Printf("The answer to Project Euler Problem #%d is :%s\n", problemNum, answer)
}

func PrintMatrix15(matrix [15][15]int) {

	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Print("\n")
	}
}

func PrintIntArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}

//IsPalidrome - Checks if num is a palidrome
func IsPalidrome(num int) bool {

	numStr := strconv.Itoa(num)
	revesredNum := Reverse(numStr)

	return strings.Compare(numStr, revesredNum) == 0

}

//Reverse - revesrses string s
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func SumDigits(digits string) int {

	sum := 0
	for i := 0; i < len(digits); i++ {
		digit, _ := strconv.Atoi(string(digits[i]))
		sum += digit
	}

	return sum
}

func SumIntDigits(num int) int {
	digits := GetDigits(num)

	sum := SumIntArray(digits)
	return sum
}

func ProductIntDigits(num int) int {
	digits := GetDigits(num)

	product := ProductIntArray(digits)
	return product
}

//GetDigits - Retruns an array of digits that compose the given integer
func GetDigits(num int) []int {

	digits := make([]int, 0)
	numStr := strconv.Itoa(num)

	for i := 0; i < len(numStr); i++ {

		digit, _ := strconv.Atoi(string(numStr[i]))
		digits = append(digits, digit)
	}

	return digits
}

//Calculates the sum of an integer array
func SumIntArray(array []int) int {
	sum := 0

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	return sum
}

func ProductIntArray(array []int) int {
	product := 1

	for i := 0; i < len(array); i++ {
		product *= array[i]
	}

	return product
}

//FindIndex returns the index of the integer in the integer array if it is found, otherwise it will return -1
func FindIndex(intArray []int, intToSearchFor int) int {
	for i, item := range intArray {
		if item == intToSearchFor {
			return i
		}
	}

	return -1
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsBig(s []big.Int, e big.Int) bool {
	for _, a := range s {
		if a.Cmp(&e) == 0 {
			return true
		}
	}
	return false
}
