package main

import "fmt"

type BingoBoard struct {
	board      [5][5]int
	marked     [5][5]int // if a board entry is marked, the corresponding entry here will be set to 1
	rowPointer int
}

func (b *BingoBoard) addRow(numbers [5]int) {
	for column := 0; column < 5; column++ {
		b.board[b.rowPointer][column] = numbers[column]
	}
	b.rowPointer += 1
}

func (b *BingoBoard) empty() {
	for row := 0; row < 5; row++ {
		for column := 0; column < 5; column++ {
			b.board[row][column] = 0
			b.marked[row][column] = 0
		}
	}
	b.rowPointer = 0
}

func (b *BingoBoard) markNumber(number int) {
	for row := 0; row < 5; row++ {
		for column := 0; column < 5; column++ {
			if b.board[row][column] == number {
				b.marked[row][column] = 1
			}
		}
	}
}

func (b BingoBoard) hasWon() bool {
	won := false
	for row := 0; row < 5; row++ {
		total := 0
		for column := 0; column < 5; column++ {
			total += b.marked[row][column]
		}
		won = total == 5
		if won {
			return won
		}
	}

	for column := 0; column < 5; column++ {
		total := 0
		for row := 0; row < 5; row++ {
			total += b.marked[row][column]
		}
		won = total == 5
		if won {
			return won
		}
	}
	return won

}

func (b BingoBoard) printBoard() {
	for row := 0; row < 5; row++ {
		for column := 0; column < 5; column++ {
			if b.marked[row][column] == 0 {
				fmt.Printf("%2d ", b.board[row][column])
			} else {
				fmt.Print("XX ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b BingoBoard) getSumOfUnmarked() int {
	total := 0
	for row := 0; row < 5; row++ {
		for column := 0; column < 5; column++ {
			if b.marked[row][column] == 0 {
				total += b.board[row][column]
			}
		}
	}
	return total
}
