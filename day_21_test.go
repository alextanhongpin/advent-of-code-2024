package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

var dirByRune = map[rune]complex128{
	'^': -1i,
	'v': 1i,
	'<': -1,
	'>': 1,
}

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
	num := newKeypad(numeric)
	dir := newKeypad(directional)

	var total int
	for _, code := range strings.Split(input, "\n") {
		a, err := strconv.Atoi(code[:len(code)-1])
		if err != nil {
			panic(err)
		}
		b := dir.Min(num.Solve(code), depth)
		total += a * b
	}

	return total
}

func newKeypad(input string) *Keypad {
	keys := make(map[complex128]rune)
	for y, row := range strings.Split(input, "\n") {
		for x, col := range []rune(row) {
			p := complex(float64(x), float64(y))
			if col == '.' {
				continue
			}

			keys[p] = col
		}
	}

	return &Keypad{
		keys:  keys,
		cache: make(map[string]int),
	}
}

type Keypad struct {
	keys  map[complex128]rune
	cache map[string]int
}

func (k *Keypad) GetPos(r rune) complex128 {
	for p, v := range k.keys {
		if v == r {
			return p
		}
	}

	panic("key not found: " + string(r))
}

func (k *Keypad) Min(codes []string, depth int) int {
	dist := math.MaxInt
	for _, code := range codes {
		dist = min(dist, k.expand(code, depth))
	}

	return dist
}

func (k *Keypad) expand(path string, depth int) int {
	if depth == 0 {
		return len(path)
	}

	// Always start with the "A" key.
	path = "A" + path
	key := fmt.Sprintf("%d:%s", depth, path)
	if n, ok := k.cache[key]; ok {
		return n
	}

	var dist int
	chars := []rune(path)
	for i, r := range chars[1:] {
		prev := k.GetPos(chars[i])
		curr := k.GetPos(r)

		steps := math.MaxInt
		paths := k.ShortestPaths(prev, curr)
		for _, p := range paths {
			steps = min(steps, k.expand(p, depth-1))
		}
		dist += steps
	}

	k.cache[key] = dist

	return dist
}

func (k *Keypad) Solve(path string) []string {
	chars := []rune("A" + path)
	paths := make([][]string, len(chars)-1)
	for i, r := range chars[1:] {
		prev := k.GetPos(chars[i])
		next := k.GetPos(r)
		paths[i] = k.ShortestPaths(prev, next)
	}

	return cartesianProduct(paths...)
}

func (k *Keypad) ShortestPaths(start, end complex128) []string {
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

		if _, ok := k.keys[s.pos]; !ok {
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
			q = append(q, Path{
				pos: s.pos + r,
				// NOTE: We need to clone the slice here, otherwise the slice will be
				// modified in the next iteration.
				hist: append(slices.Clone(s.hist), d),
			})
		}
	}

	return res
}

func cartesianProduct(vs ...[]string) []string {
	switch len(vs) {
	case 0:
		return nil
	case 1:
		return vs[0]
	default:
		var result []string
		for _, v := range vs[0] {
			for _, u := range cartesianProduct(vs[1:]...) {
				result = append(result, v+u)
			}
		}
		return result
	}
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
