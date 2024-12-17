package main

import (
	"fmt"
	"go-aoc-2025/utils"
	"math"
	"reflect"
	"strings"
)

func ExampleDayN() {
	fmt.Println("part 1:", part1(inputs[0]))
	fmt.Println("part 1:", part1(inputs[1]))
	fmt.Println()
	fmt.Println("part 2:", part2(inputs[0]))
	fmt.Println("part 2:", part2(inputs[1]))

	// Output:
	// part 1: 4,6,3,5,6,3,5,2,1,0
	// part 1: 1,5,0,1,7,4,1,0,3
	//
	// part 2: 47910079998866
	// part 2: 47910079998866

}

func equation(a, b, c int) []int {
	var output []int
	for {
		b = a % 8      // 2,4
		b = b ^ 6      // 1,6
		if 1<<b != 0 { // 7,5
			c = a / (1 << b)
		}
		b = b ^ c                    // 4,4
		b = b ^ 7                    // 1,7
		a = a / 8                    // 0,3
		output = append(output, b%8) // 5,5
		if a == 0 {                  // 3,0
			break
		}
	}

	return output
}

func parse(input string) (program []string, registers []int) {
	parts := strings.Split(input, "\n\n")
	registers = make([]int, 3)
	for i, line := range strings.Split(parts[0], "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		_, d, ok := strings.Cut(line, ": ")
		if !ok {
			panic("invalid input:" + line)
		}
		registers[i] = utils.ToInt(d)
	}

	program = strings.Split(strings.TrimPrefix(parts[1], "Program: "), ",")
	return
}

func solve(program []string, registers []int) []int {
	combo := func(o int) int {
		switch o {
		case 0, 1, 2, 3:
			return o
		case 4, 5, 6:
			return registers[o-4]
		default:
			panic("invalid operand")
		}
	}
	var output []int
	var i int
	for i+1 < len(program) {
		inst := utils.ToInt(program[i])
		operand := utils.ToInt(program[i+1])

		switch inst {
		case 0, 6, 7: // adv, division
			r := 0
			if inst == 6 {
				r = 1
			} else if inst == 7 {
				r = 2
			}

			num := registers[0]
			den := int(math.Pow(2, float64(combo(operand))))
			if den == 0 {
				registers[r] = 0
			} else {
				registers[r] = num / den
			}
		case 1:
			registers[1] = registers[1] ^ operand
		case 2:
			registers[1] = combo(operand) % 8
		case 3:
			if registers[0] != 0 {
				i = operand
				continue
			}
		case 4:
			registers[1] = registers[1] ^ registers[2]
		case 5:
			output = append(output, combo(operand)%8)
		}
		i += 2
	}

	return output
}

func part1(input string) string {
	program, registers := parse(input)
	output := solve(program, registers)
	var out []string
	for _, o := range output {
		out = append(out, fmt.Sprint(o))
	}
	return strings.Join(out, ",")
}

func part2(input string) int {
	return findValueOfA([]int{2, 4, 1, 6, 7, 5, 4, 4, 1, 7, 0, 3, 5, 5, 3, 0})
}

func findValueOfA(inst []int) int {
	var a, b, c int
	var output []int
	a = 1
	continueLoop := true

	for continueLoop {
		output = equation(a, b, c)

		if reflect.DeepEqual(output, inst) {
			return a
		}

		if len(inst) > len(output) {
			a *= 2
			continue
		}

		if len(inst) == len(output) {
			for j := len(inst) - 1; j >= 0; j-- {
				if inst[j] != output[j] {
					// Key Insight: every nth digit increments at every 8^n th step.
					// https://www.reddit.com/r/adventofcode/comments/1hg38ah/comment/m2gkd6m/
					a += int(math.Pow(8, float64(j)))
					break
				}
			}
		}

		if len(inst) < len(output) {
			a /= 2
		}
	}
	return a
}

var inputs = []string{
	`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`,
	`Register A: 37293246
Register B: 0
Register C: 0

Program: 2,4,1,6,7,5,4,4,1,7,0,3,5,5,3,0`,
}
