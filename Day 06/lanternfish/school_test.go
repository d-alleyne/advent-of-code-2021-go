package lanternfish

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInitialState(t *testing.T) {
	var lanternFish = make(SchoolState, 1)
	lanternFish[0] = make(School, 0)

	lanternFish.SetInitialState("sample.txt")

	if len(lanternFish[0]) != 5 {
		t.Errorf("After getInitialState, initial size expected to be 5, actual is %d\n", len(lanternFish))
	}

}

func TestGetStateAfterDay(t *testing.T) {
	var lanternFish = make(SchoolState, 1)
	lanternFish[0] = make(School, 0)
	lanternFish.SetInitialState("sample.txt")

	var state = lanternFish.GetStateAfterDay(0)

	if len(state) != 5 {
		t.Errorf("After getStateAfterDay(0), initial size expected to be 5, actual is %d\n", len(state))
	}

	var stateAfter1Days = lanternFish.GetStateAfterDay(1)

	var expected = School{2, 3, 2, 0, 1}

	assert.EqualValues(t, expected, stateAfter1Days, "Expected values after 1 day are not the same")

	var stateAfter2Days = lanternFish.GetStateAfterDay(2)

	expected = School{1, 2, 1, 6, 0, 8}

	assert.EqualValues(t, expected, stateAfter2Days, "Expected values after 2 days are not the same")

	var stateAfter9Days = lanternFish.GetStateAfterDay(9)

	expected = School{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 8}

	assert.EqualValues(t, expected, stateAfter9Days, "Expected values after 9 days are not the same")

}
