package main

import (
	"fmt"
	"math"
)

type bits [12]int
type diagnosticList []bits

func (b bits) toString() string {
	result := ""
	for _, value := range b {
		result += fmt.Sprintf("%d", value)
	}
	return result
}

func (b bits) convertToBinary(midpoint int) bits {
	var result bits

	for i, bit := range b {
		if bit > midpoint {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}

func (b bits) invert() bits {
	var result bits

	for i, bit := range b {
		if bit == 0 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	return result
}

func (b bits) convertToDecimal() int {
	var result int

	for i, bitValue := range b {
		power := 11 - i
		result += bitValue * int(math.Pow(2, float64(power)))
	}
	return result
}
