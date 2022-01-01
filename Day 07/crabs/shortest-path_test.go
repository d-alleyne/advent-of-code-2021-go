package day07

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLeastFuel(t *testing.T) {
	var crabPositions = make(CrabPositions, 0)
	crabPositions.SetInitialPositions("sample.txt")
	var numberOfLoadedValues = len(crabPositions)

	if numberOfLoadedValues != 10 {
		assert.Equal(t, 10, numberOfLoadedValues)
	}

	var _, fuel = crabPositions.FindLeastPositionAndFuel()

	assert.Equal(t, 37, fuel)

}
