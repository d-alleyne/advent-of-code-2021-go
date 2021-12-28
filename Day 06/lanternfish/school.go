package lanternfish

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type SchoolState [9]uint

/*
 * Initial state: 3,4,3,1,2
 * After  1 day:  2,3,2,0,1
 * After  2 days: 1,2,1,6,0,8
 * After  3 days: 0,1,0,5,6,7,8
 * Timer Value         | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
 * Number of Fish @T0    0   1   1   2   1   0   0   0   0
 * Number of Fish @T1    1   1   2   1   0   0   0   0   0
 * Number of Fish @T2    1   2   1   0   0   0   1   0   1
 * Number of Fish @T3    2   1   0   0   0   1   1   1   1
 */
func (s *SchoolState) SetInitialState(pathToFile string) {
	file, err := os.Open(pathToFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		contents := scanner.Text()
		lanternFish := strings.Split(contents, ",")

		for _, internalTimer := range lanternFish {
			value, err := strconv.Atoi(internalTimer)

			if err != nil {
				log.Fatal(err)
			} else {
				s[value]++
			}
		}
	}
}

/*
 * After  1 day:  2,3,2,0,1
 * After  2 days: 1,2,1,6,0,8
 * After  3 days: 0,1,0,5,6,7,8
 * Timer Value         | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
 * Number of Fish @D1    1   1   2   1   0   0   0   0   0
 * Number of Fish @D2    1   2   1   0   0   0   1   0   1
 * Number of Fish @D3    2   1   0   0   0   1   1   1   1
 */

func (s SchoolState) addDay() SchoolState {
	nextState := SchoolState{}

	for index := 7; index >= 0; index-- {

		if index > 0 {
			nextState[index] = s[index+1]
		} else {
			nextState[6] += s[0]
			nextState[8] += s[0]
			nextState[0] = s[1]
		}
	}
	return nextState
}

func (s SchoolState) GetCountAfterDay(day int) uint64 {
	state := s
	for start := 0; start < day; start++ {
		state = state.addDay()
	}

	var total uint64

	for _, timerCount := range state {
		total += uint64(timerCount)
	}
	return total
}
