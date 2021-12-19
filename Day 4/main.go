package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var drawNumbers []int
	bingoFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer bingoFile.Close()
	scanner := bufio.NewScanner(bingoFile)

	if scanner.Scan() {
		header := scanner.Text()
		numberStrings := strings.Split(header, ",")
		for _, value := range numberStrings {
			number, err := strconv.Atoi(value)
			if err == nil {
				drawNumbers = append(drawNumbers, number)
			}
		}
	}

	var boards []BingoBoard
	var singleBoard BingoBoard
	for scanner.Scan() {
		line := scanner.Text()
		var temp [5]int
		_, err := fmt.Sscanf(line, "%d %d %d %d %d", &temp[0], &temp[1], &temp[2], &temp[3], &temp[4])
		if err == nil {
			singleBoard.addRow(temp)
			if singleBoard.rowPointer == 5 {
				singleBoard.rowPointer = 0
				boards = append(boards, singleBoard)
				singleBoard.empty()
			}
		}
	}

	boardsThatWon := make([]int, len(boards))

	for _, number := range drawNumbers {
		for i := range boards {
			boards[i].markNumber(number)
			if boardsThatWon[i] == 0 && boards[i].hasWon() {
				boardsThatWon[i] = 1
				total := 0
				for _, won := range boardsThatWon {
					total += won
				}
				if total == len(boards) {
					boards[i].printBoard()
					sumOfUnmarked := boards[i].getSumOfUnmarked()
					score := sumOfUnmarked * number
					fmt.Println("Sum of Unmarked", sumOfUnmarked)
					fmt.Println("Last winning number", number)
					fmt.Println("The score is", score)
					os.Exit(0)
				} else {
					fmt.Println("Board", i+1, "won")
					boards[i].printBoard()
				}
			}
		}
	}

}
