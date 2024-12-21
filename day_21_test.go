package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

// Just using manhattan distance to find the shortest path between two points doesn't work.
// This of the scenario
func ExampleDayN() {
	fmt.Println("part 1:", part1(inputs[0]))
	fmt.Println("part 1:", part1(inputs[1]))
	fmt.Println()
	fmt.Println("part 2:", part2(inputs[0]))
	fmt.Println("part 2:", part2(inputs[1]))

	// Output:
	// part 1: 126384
	// part 1: 188398
	//
	// part 2: 0
	// part 2: 0
}

func part1(input string) int {
	num := newGrid(numeric)
	dir := newGrid(directional)

	solveNum := func(code string) []string {
		var result []string
		codes := []rune("A" + code)
		for i, r := range codes[1:] {
			from, to := num.Pos(codes[i]), num.Pos(r)
			comb := num.Shortest(from, to)
			if len(result) == 0 {
				for _, c := range comb {
					result = append(result, string(c.hist)+"A")
				}
			} else {
				var tmp []string
				for _, r := range result {
					for _, c := range comb {
						tmp = append(tmp, r+string(c.hist)+"A")
					}
				}
				result = tmp
			}
		}
		return result
	}

	solveDir := func(code string) []string {
		var result []string
		codes := []rune("A" + code)
		for i, r := range codes[1:] {
			from, to := dir.Pos(codes[i]), dir.Pos(r)
			comb := dir.Shortest(from, to)
			if len(result) == 0 {
				for _, c := range comb {
					result = append(result, string(c.hist)+"A")
				}
			} else {
				var tmp []string
				for _, r := range result {
					for _, c := range comb {
						tmp = append(tmp, r+string(c.hist)+"A")
					}
				}
				result = tmp
			}
		}
		return result
	}

	var total int
	for _, code := range strings.Split(input, "\n") {
		n, err := strconv.Atoi(code[:len(code)-1])
		if err != nil {
			panic(err)
		}

		nums := solveNum(code)
		var result []string
		for _, n := range nums {
			result = append(result, solveDir(n)...)
		}
		var shortest int = math.MaxInt
		for _, r := range result {
			for _, d := range solveDir(r) {
				shortest = min(shortest, len(d))
			}
		}

		total += shortest * n
	}
	return total
}

func part2(input string) int {
	return 0
}

func newGrid(input string) *Grid {
	data := make(map[complex128]rune)
	for y, row := range strings.Split(input, "\n") {
		for x, col := range []rune(row) {
			p := complex(float64(x), float64(y))
			if col == '.' {
				continue
			}
			data[p] = col
		}
	}

	return &Grid{
		data:  data,
		cache: make(map[CacheKey][]Path),
	}
}

type CacheKey struct {
	start, end complex128
}

type Grid struct {
	data  map[complex128]rune
	cache map[CacheKey][]Path
}

var dirByRune = map[rune]complex128{
	'^': -1i,
	'v': 1i,
	'<': -1,
	'>': 1,
}

type Path struct {
	pos  complex128
	hist []rune
}

func (g *Grid) Pos(r rune) complex128 {
	for p, v := range g.data {
		if v == r {
			return p
		}
	}
	panic("not found: " + string(r))
}

func (g *Grid) Shortest(start, end complex128) []Path {
	key := CacheKey{start, end}
	if c, ok := g.cache[key]; ok {
		return c
	}

	var res []Path
	visited := make(map[complex128]int)
	q := []Path{{pos: start}}
	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		if _, ok := g.data[s.pos]; !ok {
			continue
		}
		if n, ok := visited[s.pos]; !ok {
			visited[s.pos] = len(s.hist)
		} else if ok && n < len(s.hist) {
			continue
		}

		if s.pos == end {
			res = append(res, s)
			continue
		}
		for d, r := range dirByRune {
			q = append(q, Path{pos: s.pos + r, hist: append(slices.Clone(s.hist), d)})
		}
	}

	g.cache[key] = res

	return res
}

var inputs = []string{
	`029A
980A
179A
456A
379A`,
	`935A
319A
480A
789A
176A`,
}

var numeric = `789
456
123
.0A`

var directional = `.^A
<v>`
