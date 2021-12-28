package main

import (
	"fmt"
	"lanternfish/lanternfish"
)

func main() {
	schoolState := lanternfish.SchoolState{}

	schoolState.SetInitialState("input.txt")

	countAfter80Days := schoolState.GetCountAfterDay(80)
	fmt.Println("There are", countAfter80Days, "fish after 80 days")

	countAfter256Days := schoolState.GetCountAfterDay(256)

	fmt.Println("There are", countAfter256Days, "fish after 256 days")

}
