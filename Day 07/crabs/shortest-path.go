package day07

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CrabPositions []int
type PositionFuel map[int]int

func (c *CrabPositions) SetInitialPositions(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	if scanner.Scan() {
		var line = scanner.Text()
		var values = strings.Split(line, ",")

		for _, value := range values {
			var position, err = strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			*c = append(*c, position)
		}
	}

}

func (c CrabPositions) FindLeastPositionAndFuel() (position int, fuel int) {
	sort.Ints(c)
	var min, max = c[0], c[0]

	for _, value := range c {
		min = int(math.Min(float64(min), float64(value)))
		max = int(math.Max(float64(max), float64(value)))
	}

	fuel = math.MaxInt

	for rangeOfValues := min; rangeOfValues <= max; rangeOfValues++ {
		total := 0
		for _, value := range c {
			var distance = value - rangeOfValues
			if distance < 0 {
				distance *= -1
			}
			total += CrabSubmarineFuelUsage(distance)
		}
		temp := int(math.Min(float64(fuel), float64(total)))
		if temp != fuel {
			fuel = temp
			position = rangeOfValues
		}
	}
	return position, fuel
}

func CrabSubmarineFuelUsage(distance int) int {
	var total = 0

	for start := 1; start <= distance; start++ {
		total = total + start
	}

	return total
}
