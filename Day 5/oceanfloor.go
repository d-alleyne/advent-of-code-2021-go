package main

import (
	"fmt"
)

const MAX_X = 1000
const MAX_Y = 1000

type OceanFloor [MAX_Y][MAX_X]int

func (o *OceanFloor) print() {
	for y := 0; y < MAX_Y; y++ {
		for x := 0; x < MAX_X; x++ {
			fmt.Printf("%2d", o[y][x])
		}
		fmt.Println()
	}
}

func (o *OceanFloor) drawLine(x1 int, y1 int, x2 int, y2 int) {

	// validate inputs
	if !isValidX(x1) || !isValidX(x2) || !isValidY(y1) || !isValidY(y2) {
		return
	}

	var lowerBound int
	var upperBound int

	if isHorizontalLine(y1, y2) {
		if x1 > x2 {
			lowerBound = x2
			upperBound = x1
		} else {
			lowerBound = x1
			upperBound = x2
		}
		for x := lowerBound; x <= upperBound; x++ {
			o[y1][x]++
		}
	} else if isVerticalLine(x1, x2) { //vertical line
		if y1 > y2 {
			lowerBound = y2
			upperBound = y1
		} else {
			lowerBound = y1
			upperBound = y2
		}
		for y := lowerBound; y <= upperBound; y++ {
			o[y][x1]++
		}
	}

}

func isValidX(x int) bool {
	return x >= 0 && x < MAX_X
}

func isValidY(y int) bool {
	return y >= 0 && y < MAX_Y
}

func isHorizontalLine(y1 int, y2 int) bool {
	return y1 == y2
}

func isVerticalLine(x1 int, x2 int) bool {
	return x1 == x2
}

func (o *OceanFloor) getNumberOfOverlappingLines(minimum int) int {
	var total int
	for y := 0; y < MAX_Y; y++ {
		for x := 0; x < MAX_X; x++ {
			if o[x][y] >= minimum {
				total++
			}
		}
	}
	return total
}
