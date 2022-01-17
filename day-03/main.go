package main

import (
	"os"
	"strconv"
	"strings"
)

// M is an alias for map[string]interface{}
type M map[string]int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.ReadFile("input.txt")
	check(err)
	a := strings.Split(string(dat), "\n")

	// using custom type for map
	count := make([]M, len(a[0]))

	// create the map for storing the count
	for i := range count {
		m1 := M{"0": 0, "1": 0}
		count[i] = m1
	}

	for i := range a {
		for j := range a[i] {
			switch a[i][j : j+1] {
			case "0":
				count[j]["0"] = count[j]["0"] + 1
			case "1":
				count[j]["1"] = count[j]["1"] + 1
			}
		}
	}

	gammaStr := ""
	epsilonStr := ""

	for i := range count {
		if count[i]["1"] > count[i]["0"] {
			gammaStr = gammaStr + "1"
			epsilonStr = epsilonStr + "0"
		} else {
			gammaStr = gammaStr + "0"
			epsilonStr = epsilonStr + "1"
		}
	}

	println("gamma:", gammaStr)
	println("epsilon:", epsilonStr)

	gamma, err := strconv.ParseInt(gammaStr, 2, 0)
	check(err)

	epsilon, err := strconv.ParseInt(epsilonStr, 2, 0)
	check(err)

	println("consumption:", gamma*epsilon)

}
