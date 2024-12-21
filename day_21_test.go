package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

// Just using manhattan distance to find the shortest path between two points doesn't work.
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
	// part 2: 154115708116294
	// part 2: 230049027535970
}

func part1(input string) int {
	return solve(input, 2)
}

func part2(input string) int {
	return solve(input, 25)
}

func solve(input string, depth int) int {
	num := newGrid(numeric)
	dir := newGrid(directional)

	var total int
	for _, code := range strings.Split(input, "\n") {
		n, err := strconv.Atoi(code[:len(code)-1])
		if err != nil {
			panic(err)
		}
		total += n * dir.Min(num.Solve(code), depth)
	}

	return total
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
		cache: make(map[string]int),
	}
}

type Grid struct {
	data  map[complex128]rune
	cache map[string]int
}

var dirByRune = map[rune]complex128{
	'^': -1i,
	'v': 1i,
	'<': -1,
	'>': 1,
}

func (g *Grid) Pos(r rune) complex128 {
	for p, v := range g.data {
		if v == r {
			return p
		}
	}
	panic("not found: " + string(r))
}

func (g *Grid) Min(codes []string, depth int) int {
	length := math.MaxInt
	for _, code := range codes {
		length = min(length, g.expand(code, depth))
	}
	return length
}

func (g *Grid) expand(code string, depth int) int {
	if depth == 0 {
		return len(code)
	}

	// Always start with the "A" key.
	code = "A" + code
	key := fmt.Sprintf("%d:%s", depth, code)
	if n, ok := g.cache[key]; ok {
		return n
	}

	var length int
	codes := []rune(code)
	for i, r := range codes[1:] {
		prev := g.Pos(codes[i])
		curr := g.Pos(r)

		// There can be multiple paths to the next key, so we need to find the shortest one.
		comb := g.Shortest(prev, curr)
		short := math.MaxInt
		for _, c := range comb {
			short = min(short, g.expand(c, depth-1))
		}
		length += short
	}

	g.cache[key] = length

	return length
}

func (g *Grid) Solve(code string) []string {
	aCode := "A" + code
	var result []string
	codes := []rune(aCode)
	for i, r := range codes[1:] {
		from, to := g.Pos(codes[i]), g.Pos(r)
		comb := g.Shortest(from, to)
		if len(result) == 0 {
			for _, c := range comb {
				result = append(result, c)
			}
		} else {
			var tmp []string
			for _, r := range result {
				for _, c := range comb {
					tmp = append(tmp, r+c)
				}
			}
			result = tmp
		}
	}

	return result
}

func (g *Grid) Shortest(start, end complex128) []string {
	type Path struct {
		pos  complex128
		hist []rune
	}
	var res []string
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
			res = append(res, string(s.hist)+"A")
			continue
		}
		for d, r := range dirByRune {
			q = append(q, Path{pos: s.pos + r, hist: append(slices.Clone(s.hist), d)})
		}
	}

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
