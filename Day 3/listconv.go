package main

import "fmt"

type diagnosticList []bits

type bitCriteria int

const (
	LEAST_COMMON_VALUE bitCriteria = 0
	MOST_COMMON_VALUE  bitCriteria = 1
)

func (d diagnosticList) column(index int, criteria bitCriteria) diagnosticList {
	totalOnes := 0
	totalZeroes := 0

	var criteriaValue int // will filter the diagnosticList based on this value

	if len(d) == 1 {
		return d
	}

	for _, diagnosticValue := range d {
		if diagnosticValue[index] == 1 {
			totalOnes += 1
		} else {
			totalZeroes += 1
		}
	}

	if criteria == MOST_COMMON_VALUE {
		if totalOnes >= totalZeroes {
			criteriaValue = 1
		} else {
			criteriaValue = 0
		}
	} else {
		if totalOnes >= totalZeroes {
			criteriaValue = 0 // we want the least common value
		} else {
			criteriaValue = 1
		}
	}

	var result diagnosticList

	for _, diagnosticValue := range d {
		if diagnosticValue[index] == criteriaValue {
			result = append(result, diagnosticValue)
		}
	}

	return result
}

func (d diagnosticList) getOxygenGeneratorRating() bits {
	filteredList := d

	for column := 0; column < BIT_WIDTH && len(filteredList) > 1; column++ {
		filteredList = filteredList.column(column, MOST_COMMON_VALUE)
		fmt.Println("OxygenRating filtered list length", len(filteredList))
	}

	return filteredList[0]
}

func (d diagnosticList) getCarbonDioxideScrubberRating() bits {
	filteredList := d

	for column := 0; column < BIT_WIDTH && len(filteredList) > 1; column++ {
		filteredList = filteredList.column(column, LEAST_COMMON_VALUE)
	}

	return filteredList[0]
}
