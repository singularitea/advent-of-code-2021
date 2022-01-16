package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	a := strings.Split(string(dat), "\n")

	// Part One
	forward := 0
	depth := 0

	for i := range a {
		cmd := strings.Split(a[i], " ")
		num, err := strconv.Atoi(cmd[1])
		check(err)

		switch cmd[0] {
		case "up":
			depth = depth - num
		case "down":
			depth = depth + num
		case "forward":
			forward = forward + num
		}
	}

	fmt.Println("Part one:", forward*depth)

	// Part Two
	forwardTwo := 0
	depthTwo := 0
	aim := 0

	for i := range a {
		cmd := strings.Split(a[i], " ")
		num, err := strconv.Atoi(cmd[1])
		check(err)

		switch cmd[0] {
		case "up":
			aim = aim - num
		case "down":
			aim = aim + num
		case "forward":
			forwardTwo = forwardTwo + num
			depthTwo = depthTwo + (num * aim)
		}
	}

	fmt.Println("Part two:", forwardTwo*depthTwo)

}
