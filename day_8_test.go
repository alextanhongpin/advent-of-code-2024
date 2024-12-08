package main

import (
	"fmt"
	"strings"
)

func ExampleDayN() {
	fmt.Println("part 1:", part1(inputs[0]))
	fmt.Println("part 1:", part1(inputs[1]))
	fmt.Println()
	fmt.Println("part 2:", part2(inputs[0]))
	fmt.Println("part 2:", part2(inputs[1]))

	// Output:
	// part 1: 14
	// part 1: 367
	//
	// part 2: 34
	// part 2: 1285
}

func part1(input string) int {
	var nodes []complex128
	grid := make(map[complex128]rune)

	for i, row := range strings.Split(input, "\n") {
		for j, col := range row {
			p := complex(float64(j), float64(i))
			grid[p] = col
			if col != '.' {
				nodes = append(nodes, p)
			}
		}
	}

	antinodes := make(map[complex128]bool)
	for _, node := range nodes {
		for _, other := range nodes {
			if grid[node] == grid[other] && node != other {
				dist := other - node
				next := node + 2*dist
				if _, ok := grid[next]; !ok {
					continue
				}
				antinodes[next] = true
			}
		}
	}

	return len(antinodes)
}

func part2(input string) int {
	var nodes []complex128
	grid := make(map[complex128]rune)

	for i, row := range strings.Split(input, "\n") {
		for j, col := range row {
			p := complex(float64(j), float64(i))
			grid[p] = col
			if col != '.' {
				nodes = append(nodes, p)
			}
		}
	}

	antinodes := make(map[complex128]bool)
	for _, node := range nodes {
		for _, other := range nodes {
			if grid[node] == grid[other] && node != other {
				dist := other - node
				antinodes[other] = true
				antinodes[node] = true
				next := node + 2*dist
				for {
					if _, ok := grid[next]; !ok {
						break
					}
					antinodes[next] = true
					next += dist
				}
			}
		}
	}

	return len(antinodes)
}

var inputs = []string{
	`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`,
	`....1.y.D...Y..........w....m.....................
..R..D..5....Y...1.........w.........G............
........R........D..o.............................
.............H......Y...w.....m...................
.......R..................3.........v.............
..1...D..5.........o................0.Bm..........
......5y.....o.........................3..........
....H...y......Z...............................0..
..............H.x..............m........w..g......
..........................A.......................
.........................................fg.......
...8.............v.....e............3B.....2......
.............5.....r......B.......2...........G..0
......................v....................3g.....
......P..............Y...c...........M.2.G........
..................................................
.....H....Z.............................K.......0.
....8d..Z......................u....X......f.g....
......d..P..r..............B.........E.........9..
.......r...........E..............q...M...........
...k...............v......Eb........q...........f.
.....R................b..............U.q9...2.....
.J......i.............M....q...................K..
..........d........................M.....A........
.......Zj..........h................9S............
.........j..........P..........Q....7.....c.......
.j........................a.......................
....j.6.....h.....F..a......L......c.X............
.................I.......a..b.............A......V
x........................p..........EK............
.......6.....................................X....
..J....................bf.....r.....K.............
.e..k................................7......X.....
...x..kP..................u...........U...........
J.8.....h....d........U....Q........F.c....iC.O...
...J...h.I..e......................i...7..........
..............................L.QU.....A......7...
...............k....t.........a.WO..i.............
.....4..6..............l...............T..........
........z...4.....p..........LS...Q...............
....e..z................t........pS..........C....
..............I........W.............9..........C.
..................l..........F...u...O............
....l............T.t.6...F.........S..s........V..
.......................t4.........................
.........z...........................CV....s......
..z.........IL.......W....p.........V...u.........
.....................l............................
........T.......................s.................
..........T..........4............................`,
}
