package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// M is an alias for map[string]int
type M map[string]int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkBoardHasWon(board board) int {
	for i := range board.board {
		countHorizontal := 0
		totalHorizontal := 0
		horizontal := false
		for l := range board.board[i] {
			if board.board[i][l].seen == true {
				countHorizontal++
				totalHorizontal += board.board[i][l].number
			}
			if countHorizontal == 5 {
				horizontal = true
			}
		}

		countVertical := 0
		totalVertical := 0
		vertical := false
		for l := range board.board[i] {
			for j := range board.board {
				if board.board[i][l].seen == true {
					countVertical++
					totalVertical += board.board[i][l].number
				}
				if countVertical == 5 {
					vertical = true
				}
			}
		}
		if vertical == true || horizontal == true {
			return totalHorizontal + totalVertical
		} else {
			return 0
		}
	}
	return 0
}

type number struct {
	number int
	seen   bool
}

type board struct {
	board [][]number
}

func main() {

	dat, err := os.ReadFile("numbers.txt")
	check(err)
	numbers := strings.Split(string(dat), ",")

	fmt.Println(numbers)

	// dat, err = os.ReadFile("boards.txt")
	// check(err)
	// boardsRaw := strings.Split(string(dat), "\n")

	// boards := make([][][]string, len(boardsRaw))

	dat, err = os.ReadFile("boards.txt")
	check(err)
	test := strings.Split(string(dat), "\n\n")

	// boards := make([][]string, len(test))
	var boards []board

	// newString2 := make([][]string, len(test))
	for i := range test {
		newString := test[i]
		newString = strings.Replace(newString, "  ", " ", -1)
		// newString = strings.Replace(newString, "\r ", "\r", -1)
		newString = strings.Replace(newString, "\n ", "\n", -1)
		newString = strings.Replace(newString, " \n", "\n", -1)
		newString = strings.Replace(newString, "\n", " ", -1)
		if string(newString[0]) == " " {
			newString = strings.Replace(newString, " ", "", 1)
		}

		tempBoard := strings.Split(newString, " ")
		// for l := range tempBoard {
		// 	fmt.Println(string(tempBoard[l]))
		// }
		// fmt.Println("hi")
		// os.Exit(1)
		count := 0
		var row []number
		var board board
		for j := range tempBoard {
			// fmt.Println(tempBoard[j])
			tempNum, err := strconv.Atoi(tempBoard[j])
			check(err)
			n := number{number: tempNum, seen: false}
			if count < 5 {
				row = append(row, n)
				count++

			} else {
				count = 0
				board.board = append(board.board, row)
				row = nil
			}
			// fmt.Println(board.board)

			// newString = strings.Replace(newString, " ", "", -1)

			// append(newString, newString2)
			// }
		}
		// fmt.Println("board", board)
		boards = append(boards, board)
	}

	for i := range numbers {
		for j := range boards {
			for k := range boards[j].board {
				for l := range boards[j].board[k] {
					num, err := strconv.Atoi(numbers[i])
					check(err)
					if boards[j].board[k][l].number == num {
						boards[j].board[k][l].seen = true
						checkBoardHasWon(boards[j])
					}
				}
				// fmt.Println(i)
				// fmt.Println(boards[j].board[k])
			}
		}
	}

	// response := fmt.Sprintf("board %d", j)
	// fmt.Println(response)
	// fmt.Println(boards[j].board)
	fmt.Println(boards)

}
