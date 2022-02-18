package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// M is an alias for map[string]int
type M map[string]int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type board struct {
	number int
	seen   bool
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

	dat, err = os.ReadFile("test.txt")
	check(err)
	test := strings.Split(string(dat), "\n")

	boards := make([][]strings, len())


	for i := range test {
		newString := ""
		newString = strings.Replace(test[i], " ", "", -1)
		newString = strings.Replace(newString, " ", "", -1)
		newString = strings.Replace(newString, "\n", "", -1)
		for j := range 
		test[i] == make(newString

	}
	for i := range test {
		fmt.Println(test[i])
	}

	// for i := range test {
	// 	test2 := strings.Split(string(dat), " ")
	// 	print(i)
	// 	for j := range test2 {
	// 		fmt.Println(test2[j])
	// 	}
	// }

	// i :=0
	// for i < 5{
	// 	append(boards[i], board)
	// }

	// newBoards := make([][]string, len(boards))
	// fmt.Println(newBoards)

	// for i := range boards {
	// 	nb := strings.Split(string(boards[i]), " ")
	// 	append(newBoards[i], nb)
	// }

	// // using custom type for map
	// count := make([]M, len(a[0]))

	// // create the map for storing the count
	// for i := range count {
	// 	m1 := M{"0": 0, "1": 0}
	// 	count[i] = m1
	// }

	// for i := range a {
	// 	for j := range a[i] {
	// 		switch a[i][j : j+1] {
	// 		case "0":
	// 			count[j]["0"] = count[j]["0"] + 1
	// 		case "1":
	// 			count[j]["1"] = count[j]["1"] + 1
	// 		}
	// 	}
	// }

	// gammaStr := ""
	// epsilonStr := ""

	// for i := range count {
	// 	if count[i]["1"] > count[i]["0"] {
	// 		gammaStr = gammaStr + "1"
	// 		epsilonStr = epsilonStr + "0"
	// 	} else {
	// 		gammaStr = gammaStr + "0"
	// 		epsilonStr = epsilonStr + "1"
	// 	}
	// }

	// println("gamma:", gammaStr)
	// println("epsilon:", epsilonStr)

	// gamma, err := strconv.ParseInt(gammaStr, 2, 0)
	// check(err)

	// epsilon, err := strconv.ParseInt(epsilonStr, 2, 0)
	// check(err)

	// println("consumption:", gamma*epsilon)

	// // part two

	// fmt.Printf("%v", count)

	// ox := make([]string, len(a))
	// copy(ox, a)

	// c02 := make([]string, len(a))
	// copy(c02, a)

	// for j := range a[0] {
	// 	for i := range ox {
	// 		fmt.Println(len(ox))
	// 		if len(ox) == 1 {
	// 			break
	// 		} else if count[j]["0"] > count[j]["1"] {
	// 			if a[i][j] != 0 {
	// 				ox = remove(ox, i)
	// 			}
	// 		} else {
	// 			if a[i][j] != 1 {
	// 				ox = remove(ox, i)
	// 			}
	// 		}

	// 	}
	// }
	// fmt.Println(ox)

}
