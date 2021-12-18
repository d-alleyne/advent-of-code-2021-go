package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var diagnosticEntries diagnosticList
	var sumOfEachBit bits
	for scanner.Scan() {
		diagnosticValue := scanner.Text()
		var binaryEntry bits
		for i, bit := range diagnosticValue {
			bitValue, err := strconv.Atoi(string(bit))
			if err != nil {
				log.Fatal(err)
			}
			binaryEntry[i] = bitValue
		}
		diagnosticEntries = append(diagnosticEntries, binaryEntry)
		sumOfEachBit = addBits(sumOfEachBit, binaryEntry)
	}

	totalDiagnosticEntries := len(diagnosticEntries)

	gammaRate := sumOfEachBit.convertToBinary(totalDiagnosticEntries / 2)
	epsilonRate := gammaRate.invert()

	gamma := gammaRate.convertToDecimal()
	epsilon := epsilonRate.convertToDecimal()
	power := gamma * epsilon
	fmt.Println("The power consumption of the submarine is", power)

}

func addBits(sumOfEachBit bits, binaryEntry bits) bits {
	for i, value := range binaryEntry {
		sumOfEachBit[i] += value
	}
	return sumOfEachBit
}
