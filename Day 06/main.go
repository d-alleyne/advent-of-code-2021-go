package main

import (
	"fmt"
	"lanternfish/lanternfish"
)

func main() {
	schoolState := make(lanternfish.SchoolState, 1)
	schoolState[0] = make(lanternfish.School, 0)

	schoolState.SetInitialState("input.txt")

	stateAfter80Days := schoolState.GetStateAfterDay(80)
	fmt.Println("There are", len(stateAfter80Days), "after 80 days")
}
