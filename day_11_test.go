package main

import (
	"fmt"
	"go-aoc-2025/utils"
	"maps"
	"strings"
)

func ExampleDayN() {
	fmt.Println("part 1:", part1(inputs[0]))
	fmt.Println("part 1:", part1(inputs[1]))
	fmt.Println()
	fmt.Println("part 2:", part2(inputs[0]))
	fmt.Println("part 2:", part2(inputs[1]))

	// Output:
	// part 1: 55312
	// part 1: 175006
	//
	// part 2: 65601038650482
	// part 2: 207961583799296
}

func blink(input string, n int) int {
	stones := make(map[string]int)
	for _, s := range strings.Fields(input) {
		stones[s]++
	}

	for range n {
		oldStones := maps.Clone(stones)
		clear(stones)
		for s, i := range oldStones {
			switch {
			case s == "0":
				stones["1"] += i
			case len(s)%2 == 0:
				m := len(s) / 2
				lhs, rhs := s[:m], fmt.Sprint(utils.ToInt(s[m:]))
				stones[lhs] += i
				stones[rhs] += i
			default:
				stones[fmt.Sprint(utils.ToInt(s)*2024)] += i
			}
		}
	}

	var total int
	for _, n := range stones {
		total += n
	}
	return total
}

func part1(input string) int {
	return blink(input, 25)
}

func part2(input string) int {
	return blink(input, 75)
}

var inputs = []string{
	`125 17`,
	`64554 35 906 6 6960985 5755 975820 0`,
}
