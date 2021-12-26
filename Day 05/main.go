package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var oceanFloor = readVentLocations(scanner)
	// oceanFloor.print()
	fmt.Println(oceanFloor.getNumberOfOverlappingLines(2))
}

func readVentLocations(scanner *bufio.Scanner) OceanFloor {
	var oceanFloor OceanFloor

	for scanner.Scan() {
		line := scanner.Text()
		var x1 int
		var x2 int
		var y1 int
		var y2 int
		_, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		if err != nil {
			log.Fatal(err)
		}
		oceanFloor.drawLine(x1, y1, x2, y2)
	}
	return oceanFloor
}
