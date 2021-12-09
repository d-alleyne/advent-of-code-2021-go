package main

import (
	"fmt"
	"math"
)

type intbits [12]int
type diagnosticList []string

func (b intbits) convertToBinary(midpoint int) intbits {
	fmt.Println("Midpoint is ", midpoint)
	var result intbits

	for i, bit := range b {
		if bit > midpoint {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}

func (b intbits) invert() intbits {
	var result intbits

	for i, bit := range b {
		if bit == 0 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	return result
}

func (b intbits) convertToDecimal() int {
	var result int

	for i, bitValue := range b {
		power := 11 - i
		result += bitValue * int(math.Pow(2, float64(power)))
	}
	return result
}

func (d diagnosticList) column(index int) intbits {
	var result intbits

	return result
}

func (d diagnosticList) getOxygenGeneratorRating() int {
	var rating int

	for column, bitValue := range d {

	}

	return rating
}
