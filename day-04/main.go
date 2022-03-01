package main

import (
	"fmt"
	"os"
	"regexp"
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

type number struct {
	number int
	seen   bool
}

type board struct {
	board [][]number
}

func splitText(word string) []string {
	array := regexp.MustCompile("[\\ \\,\\\n\\s]+").Split(word, -1)
	return array
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
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
		var board board
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
		// os.Exit(1)
		count := 0
		var row []number
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

			boards = append(boards, board)
		}
	}
	fmt.Println(boards)

}
