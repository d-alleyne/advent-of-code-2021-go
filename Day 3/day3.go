package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var diagnosticEntries diagnosticList
	var sumOfEachBit intbits
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		diagnosticValue := scanner.Text()
		diagnosticEntries = append(diagnosticEntries, diagnosticValue)
		sumOfEachBit = addBits(sumOfEachBit, strings.Split(diagnosticValue, ""))
	}

	totalDiagnosticEntries := len(diagnosticEntries)

	var gammaRate = sumOfEachBit.convertToBinary(totalDiagnosticEntries / 2)
	epsilonRate := gammaRate.invert()

	gamma := gammaRate.convertToDecimal()
	epsilon := epsilonRate.convertToDecimal()
	power := gamma * epsilon
	fmt.Println(gammaRate, gamma, epsilonRate, power)

	oxygenGeneratorRating := diagnosticEntries.getOxygenGeneratorRating()
	carbonDioxideScrubberRating := 0

}

func addBits(sumOfEachBit intbits, strbits []string) intbits {
	for i, bit := range strbits {
		value, err := strconv.Atoi(bit)
		if err != nil {
			log.Fatal(err)
		}
		sumOfEachBit[i] += value
	}
	return sumOfEachBit
}
