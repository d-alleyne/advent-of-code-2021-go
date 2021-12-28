package lanternfish

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetInitialState(t *testing.T) {
	var lanternFish = SchoolState{}

	lanternFish.SetInitialState("sample.txt")

	totalFish := lanternFish.GetCountAfterDay(0)

	if totalFish != 5 {
		t.Errorf("After getInitialState, initial size expected to be 5, actual is %d\n", totalFish)
	}

}

func TestGetCountAfterDay(t *testing.T) {
	var expected uint64
	var lanternFish = SchoolState{}
	lanternFish.SetInitialState("sample.txt")

	var stateAfter1Days = lanternFish.GetCountAfterDay(1)

	expected = 5 // [2, 3, 2, 0, 1]

	assert.EqualValues(t, expected, stateAfter1Days, "Expected values after 1 day are not the same")

	var stateAfter2Days = lanternFish.GetCountAfterDay(2)

	expected = 6 // [1, 2, 1, 6, 0, 8]

	assert.EqualValues(t, expected, stateAfter2Days, "Expected values after 2 days are not the same")

	var stateAfter9Days = lanternFish.GetCountAfterDay(9)

	expected = 11 //School{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 8}

	assert.EqualValues(t, expected, stateAfter9Days, "Expected values after 9 days are not the same")

}
