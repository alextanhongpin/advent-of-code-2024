package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ExampleDayN() {
	fmt.Println("part 1:", part1(input1))
	fmt.Println("part 1:", part1(input2))
	fmt.Println()
	fmt.Println("part 2:", part2(input1))
	fmt.Println("part 2:", part2(input2))

	// Output:
	// part 1: 0
	// part 1: 0
	//
	// part 2: 0
	// part 2: 0
}

func part1(input string) int {
	var total int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		text := scanner.Text()
		_ = text
	}

	return total
}

func part2(input string) int {
	var total int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		text := scanner.Text()
		_ = text
	}

	return total
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

var input1 = ``

var input2 = ``
