package main

import (
	"day07/crabs"
	"fmt"
)

func main() {
	var crabPositions = make(day07.CrabPositions, 0)
	crabPositions.SetInitialPositions("input.txt")

	var position, fuel = crabPositions.FindLeastPositionAndFuel()
	fmt.Println("Fuel needed to align to", position, "is", fuel)

}
