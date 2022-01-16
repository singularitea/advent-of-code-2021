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

func sliceAtoi(ss []string) ([]int, error) {
	si := make([]int, 0, len(ss))
	for i := range ss {
		j, err := strconv.Atoi(ss[i])
		if err != nil {
			return si, err
		}
		si = append(si, j)
	}
	return si, nil
}

func main() {

	dat, err := os.ReadFile("input.txt")
	check(err)

	a := strings.Split(string(dat), "\n")
	si, err := sliceAtoi(a)
	check(err)

	// Part One
	total := 0
	for i := range si {
		if i > 0 {

			if si[i] > si[i-1] {
				total = total + 1
			}
		}
	}

	fmt.Println("Previous comparison total:", total)

	// Part Two
	slidingTotal := 0
	for i := range si {
		if i > 3 {
			if si[i]+si[i-1]+si[i-2] > si[i-1]+si[i-2]+si[i-3] {
				slidingTotal = slidingTotal + 1
			}
		}
	}

	fmt.Println("Sliding window comparison total:", slidingTotal)
}
