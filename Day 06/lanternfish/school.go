package lanternfish

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type School []int
type SchoolState []School

func (s SchoolState) Print() {
	for _, row := range s {
		fmt.Println(row)
	}
}

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

			if err == nil {
				(*s)[0] = append((*s)[0], value)
			} else {
				log.Fatal(err)
			}
		}
	}
}

func (s *SchoolState) GetStateAfterDay(day int) School {
	if len(*s) > day {
		return (*s)[day]
	}

	for start := len(*s) - 1; start < day; start++ {
		newSchool := make(School, len((*s)[start]))
		copy(newSchool, (*s)[start])

		for index, timerValue := range (*s)[start] {
			timerValue--
			newSchool[index] = timerValue
			if timerValue == -1 {
				newSchool[index] = 6
				newSchool = append(newSchool, 8)
			}
		}
		*s = append(*s, newSchool)
	}

	return (*s)[day]
}
