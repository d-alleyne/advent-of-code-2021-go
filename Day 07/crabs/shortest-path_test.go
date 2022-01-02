package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLeastFuel(t *testing.T) {
	var crabPositions = make(CrabPositions, 0)
	crabPositions.SetInitialPositions("sample.txt")
	var numberOfLoadedValues = len(crabPositions)

	if numberOfLoadedValues != 10 {
		assert.Equal(t, 10, numberOfLoadedValues)
	}

	var _, fuel = crabPositions.FindLeastPositionAndFuel()

	assert.Equal(t, 168, fuel)

}

func TestCrabSubmarineFuelUsage(t *testing.T) {
	var horizontalDistance = 11

	var fuel = CrabSubmarineFuelUsage(horizontalDistance)

	assert.Equal(t, 66, fuel)

	horizontalDistance = 4

	fuel = CrabSubmarineFuelUsage(horizontalDistance)

	assert.Equal(t, 10, fuel)

}
