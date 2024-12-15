package main

import (
	"fmt"
	"log"
	"strings"
)

func ExampleDayN() {
	// The moves are split into multiple lines for readability... This caught me off guard.
	fmt.Println("part 1:", part1(inputs[0]))
	fmt.Println("part 1:", part1(inputs[1]))
	fmt.Println("part 1:", part1(inputs[2]))
	fmt.Println()
	fmt.Println("part 2:", part2(inputs[0]))
	fmt.Println("part 2:", part2(inputs[1]))
	fmt.Println("part 2:", part2(inputs[2]))

	// Output:
	// part 1: 2028
	// part 1: 10092
	// part 1: 1552463
	//
	// part 2: 1751
	// part 2: 9021
	// part 2: 1554058
}

var movesByDir = map[rune]complex128{
	'^': -1i,
	'v': 1i,
	'<': -1,
	'>': 1,
}

func robot(grid map[complex128]string) complex128 {
	for p, v := range grid {
		if v == "@" {
			return p
		}
	}
	panic("no robot")
}

func display(grid map[complex128]string) {
	for y := range 20 {
		for x := range 20 {
			ch := complex(float64(x), float64(y))
			fmt.Print(grid[ch])
		}
		fmt.Println()
	}
}

func part1(input string) int {
	parts := strings.Split(input, "\n\n")
	grid := make(map[complex128]string)
	moves := strings.TrimSpace(strings.ReplaceAll(parts[1], "\n", ""))

	for y, row := range strings.Split(parts[0], "\n") {
		for x, col := range strings.Split(row, "") {
			p := complex(float64(x), float64(y))
			grid[p] = col
		}
	}

	for _, dir := range moves {
		move := movesByDir[dir]

		curr := robot(grid)
		next := curr + move
		switch grid[next] {
		case "#":
			// Wall, continue.
		case ".":
			// Next is a space, moves robot.
			grid[curr] = "."
			grid[next] = "@"
		case "O":
			// Next is a box, check for subsequent boxes.
			for grid[next] == "O" {
				next += move
			}
			if grid[next] == "#" {
				continue
			}
			grid[next] = "O"
			next = curr + move
			grid[curr] = "."
			grid[next] = "@"
		default:
			log.Fatal("invalid grid")
		}
	}

	var total int
	for p, b := range grid {
		if b != "O" {
			continue
		}
		x := real(p)
		y := imag(p)
		total += int(x) + int(y)*100
	}

	return total
}

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	grid := make(map[complex128]string)
	moves := strings.TrimSpace(strings.ReplaceAll(parts[1], "\n", ""))

	for y, row := range strings.Split(parts[0], "\n") {
		for x, col := range strings.Split(row, "") {
			mx := float64(x) * 2
			my := float64(y) * 1
			p := complex(mx, my)
			switch col {
			case "#":
				grid[p] = col
				grid[p+1] = col
			case ".":
				grid[p] = col
				grid[p+1] = col
			case "O":
				grid[p] = "["
				grid[p+1] = "]"
			case "@":
				grid[p] = col
				grid[p+1] = "."
			}
		}
	}

	isBox := func(tile string) bool {
		return tile == "[" || tile == "]"
	}

	for _, dir := range moves {
		move := movesByDir[dir]

		curr := robot(grid)
		next := curr + move

		switch tile := grid[next]; tile {
		case "#":
			// Wall, continue.
		case ".":
			// Next is a space, moves robot.
			grid[curr] = "."
			grid[next] = "@"
		default:
			switch {
			case dir == '>':
				steps := 0
				ahead := next
				for grid[ahead] == "[" {
					ahead += 2 * move
					steps++
				}
				if grid[ahead] == "#" {
					continue
				}

				grid[curr] = "."
				grid[next] = "@"
				ahead = next + move
				for range steps {
					grid[ahead] = "["
					grid[ahead+1] = "]"
					ahead += 2 * move
				}
			case dir == '<':
				steps := 0
				ahead := next
				for grid[ahead] == "]" {
					ahead += 2 * move
					steps++
				}
				if grid[ahead] == "#" {
					continue
				}

				grid[curr] = "."
				grid[next] = "@"
				ahead = next + move
				for range steps {
					grid[ahead] = "]"
					grid[ahead-1] = "["
					ahead += 2 * move
				}
			case (dir == '^' || dir == 'v'):
				// Collect all the boxes position, then add up or down.
				boxesBySteps := make(map[int][]complex128)
				if tile == "[" {
					boxesBySteps[0] = []complex128{next}
				} else {
					boxesBySteps[0] = []complex128{next - 1}
				}
				var steps int
				var valid = true
				for {
					boxes := boxesBySteps[steps]
					if len(boxes) == 0 {
						break
					}
					for _, box := range boxes {
						left := grid[box+move]
						right := grid[box+move+1]
						isObstacle := left == "#" || right == "#"
						if isObstacle {
							valid = false
							break
						}
						if !(isBox(left) || isBox(right)) {
							continue
						}
						//  []
						// []
						//
						// []
						// []
						//
						// []
						//  []

						if left == "[" {
							boxesBySteps[steps+1] = append(boxesBySteps[steps+1], box+move)
						} else if left == "]" {
							boxesBySteps[steps+1] = append(boxesBySteps[steps+1], box+move-1)
						}
						if right == "[" {
							boxesBySteps[steps+1] = append(boxesBySteps[steps+1], box+move+1)
						}
					}
					if !valid {
						break
					}
					steps++
				}
				if !valid {
					continue
				}

				for i := steps; i > -1; i-- {
					for _, box := range boxesBySteps[i] {
						grid[box] = "."
						grid[box+1] = "."
						grid[box+move] = "["
						grid[box+move+1] = "]"
					}
				}

				grid[curr] = "."
				grid[next] = "@"
			default:
				log.Fatal("what")
			}
		}
	}

	var total int
	for p, b := range grid {
		if b != "[" {
			continue
		}
		x := real(p)
		y := imag(p)
		total += int(x) + int(y)*100
	}

	return total
}

var inputs = []string{
	`########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`,
	`##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`,
	`##################################################
#..O...OOO..O..O..O.O.OOOO.#OO...O.O.......OO....#
#O...OO#..#.#O...OOOO...O#O..O#.O..OO...O.O......#
#..OO......OO....#O...OO...O...#O........O......O#
##..#..OOO..O.....OO....##......OO..O.....O.O..O.#
#..O#.O.O..O..O.#OOOO...OO......#OO..OO..#.OOO...#
#O.....O.OOO....O....O#.....O.#....#........O.O..#
#...O....O.OO..O#..#O.#..O...........O#.....#O..O#
#....O....OO.O.##.......#....O.O.O.O.O..#O...O...#
#O......#..OO.....O...OO.OOO.#.O#.#..OO.....#.O.O#
#O....O.#O..O..O...O#.OO..O.........OOO..O...OOO.#
#OO..O....#OOO...#.O..#.....O#O#..#.......O..OO.##
#..O.O.......#O...#.#O..O#.O.O.OO...O......O.....#
#....O...O##.....O..OO#..#O..O.....O...O.#.#O....#
#...#.OO#.O.........O.....O..O.O#O#.....OO.O....##
#..#...OO.O#OOO..O.#..........O.......O..##O...O##
#.O..O.OO#.....OO.O.O....O.#.....O#OO..OOO.##....#
#O...OO#O......O...#.O....O...O.#......OO..O.O...#
#.#..O.......O..OO....OOOO...O#...O....#.O.O.....#
#...OO.OO..OO....OO...........#.O#...O#.O.O#.OO.##
#O....O..#O.....O#.......#.O....O.....O....O.O..O#
##..O..O.O..#..O#......O......O..O#...#.O..O...#O#
#OO..OO.O..OO....OO.O#.OO.O.O#OO.#.OOOO...#.OO#..#
#.#O.O..OO.O....O...O...O.......#O..#O#.........##
#O.#O..#.O....#.#O.O.O..@..O.......O#.O...#..O...#
#O...O..#.#..O.OO....O.....O.O#....O..O.O#.OO..O.#
#....O.....O#..O....#......O.O.#..OO#OOO.O#.....O#
#O.#.O..#O#.....OO.O..O......#O...O...OO.....O...#
##O.O.O....OO.OOOOO..O..#.....O#OO.......OO....#.#
#.OO..O#.OO.......OOO#.O..O.OO.O..OO...O...O..O#.#
#O.#O..#O..O...O...O..O..O..O.....OO..O..O.....OO#
#O......O...O.OO..O.#..O..OO..OO...OOO.O.......O.#
#.OO....#....OO........O..O...OOO...O.O...O#..O..#
#..O....#..OO.O.OO..O#OO...#..O...#....O.....O..O#
#O...O#O.O..OO.#O.OOO.#.O..O#.O.......OO.......O##
#..##...O..O...#O.OOO.OO.#....O..O.O.O...##O...O.#
###......#.OO#.....O.O......O#.O...OO......O..O#.#
#.O.......O#OO#.O.O..#.....OOO....#.....OOO.OO..O#
#.O.#....O....#...O.....O....#....O............O.#
#OO.O..OO..OO...O.O.....O.##O..............O....O#
#....O..OO.O..O.......OOO..O..#..#OOO.O.O....OO..#
#O..O#..O#.#...O......O..O..O.O.#....O....OO#..OO#
#...OO........#.........O........O..O.....O......#
#O..OOO..O...O.O.OO#O#..O##..#OOO.O.OO..O..O...O.#
#.....#..#.O.....O....O#OO....O..O......#........#
#.......O.O....O.OO.#..O#....#.O..O..OO..OO#O...##
#..OO..O..O..O..#OO.OOO.....#.O.O..O.OOO..#O.....#
#O.OOO.#..OO...O......O.#.O...#...O..O......OO...#
#..OOOO.O#..........#.OOO......#....OO..OO..O.O..#
##################################################

v<v^^><^<vv>^<v^>^<^><>^>>^vvvv>^<^><^vv^v><<>><v>^>v<<^>^^^^^>v<v<^v^v<v>>vv<^>v^^v><<<^<><^>^<^<^^^<^^^v><<>v<><<<v<v<<v>>>^v<<<>^<>^v<<^v<^>v<<>><<v^^^^^^v>>vv^>>^^>^v^v^<<>^v<v<^>><^<v^^^v^<^><>v^<^>vv>v><^<><>^^>>^^<>>^v>v<<^v^^>v^<<^v^^><^^>v>v<><>>vv<^^v>><v^vvv^v^^<>>v<^v^v<>v<vvv>^<v<^<v<>^^<<<<v<<<v^^><^<^^<vv^v^v>><^<>^^><<><>><><v<>^v^<^>v<<>><<<v>^>vv><<<<v^^>^vv>v<>>^><v^^^>vv<v^>v>>v^^^^>>^<v>v<<^><^<v>v^>>>^v>^<v>>^<vv<>>v>v>v^<v^v<^vv<^><<^>vv><^<>>^>>v^>^v<<><<v<>^<<^^>v^^<v<^^^^^^>^<^v>v><vv^^^>vv<^vv<>vv><^v>^v<>v^>^<>>>>><^>^<vv<<><<v^>>v><^<^<<<>><>vv^>>vv>^^<vv<><v>vvv^>^v><v^^>v><^>vv<v^v>v><v>>v<vv<v<^vv^vv>>^>>^v>^>^v<vv>^v^<><v<>v<<<<^v>>>v>>v<^v>^<^>^<^<vv<<^v^>^^>><v<<^<v>v<>^>v>^^vv<<v>^>^^<^>v^>>v>>><<vvvv^vv^><^^v^^v<<v>><^^>v>^v<v>v><>vv^><>><v><^v<v><vv>^<<>>><<>^^<^v<<^^>v^<^^><^v^^>>v>^^>>^<v><^>><<^^v<><<<v<>vv^^<vv><^>>>^<vv^><<<<<^<vv<><v^>>^>>>v>>>^^<<>^v>^^<^<<><vv^>v^>v<>>><<vv^<vvv<<<<^<^>^^^^vvv<^>v^<^>>v^<^^^<<vv>^<<^^v>^^<<^><vv><^v<>>v><^^
>>vv<vvv<^<>><<^vv<^><vv^>><>>^^<^>vvvv^^v>>^><^vv><v^^>^v<^v^>>v^^^<<<<<v^^<>><><>^>v>><^>v>v^<<<^>^<^<v><><v>^v^v<>vv<^<<>vv><^^<><<^^^^<vv>^^><^>v^<v^v^^v<<>><>v^<>v>>^^v^v^^vvvvv<>^>v^>>>>>v<<v<>>>^><^<<^^^v<>v<^<>>vv<>^<v<^>vv^^<<<<<>v<^v>><<^v^><<v<^>^v^<>><^<><<<v<^^><v<^^<>v^<>^><><v<^<v>>v<v^<<^<v<^v<>><<^><^<><<vv>^<^vv><<^><^>^^vvv<^v^^<>v><<v<v<v^>v^>^>^vv^v>v<<>><<<v^>>v<^>^^^<v><v<<<<v<>>v^<>>v<vv>v^v>^>^v>>^^vv<v^^>><<v><<><>vv>vv^<>v^v^^v><>>vvv^v<vv<>^>v<^<v><>^v>><<^^^<^^><v^^<vv^vv^<^><vv<^><^^>>><^>^>vvv><<<><vv>v<<^>><><<<v<<^v>><<<>^^v<<>>^>><^^v>>vv<^^^v<vv^^<^<<>^v>>^v<v><<^v^^v>><>vvv>>v^<v<<<<v^>^v^<^<^v^vv^<^v^^^>^v><v^v^v<>^<^<>^^^>^^<>v<<>^<><>>>>><><<>vvv>v>>v><>>>>v^^v^<^><><<>^v^^v><<^^v><<v<v>>>^vvv^>vvvvv<v^^>>>v^>><^>v<^>^^v^<<>v<<<>vvv^<v<v>^v^vv^^<>^<vv<>^><<><v<<<v^><<v^<<vv><v>><^<<^<vv^<<<^^^<^v^vv><<<<><>v^vv<<<^vv^>^<<<v<^^<^v<>><><^^^v>>^vv^<>vv^^<^vvv^<>>><<^<>v<^>^^<^>v<^<vv>>>^^<^<v<^^>>^^>>v>>v<v<>>v<<>>>>^v^vvv>><^>>^v><^^>^<>^>^^vv>^<<>>
<<v<v>v^v>v>^^<>v<><^>^v<^>v^><>^>^vvv>^<^v>vv^>vv^^^<>>>v^<>^^<v<^>v>^<>^<>^><vv><<<v<^^^v>vvvv^^v<v<><>v>vv<<v^^vv>^v^v<v><<^^^^<v<^v<vvv>v<vvv>v>v><<vvvv^<<>v><^>v>^^^><v>>^v^v>v^^>v>^^v>v<<^<v>><<v^>><<<>>vv><<^><<v>vv>>^<<>^^<^v><v<v<>v>^v>>^<>>vv<vv<<^<^v^<v<>>^v^^>^>v>><v^^v<><<>v>><><^vv^^^v>v>v<>v<v<<v<<<^<<<^>v^<<vvv<v<v><^><v>^>v><<^>>v>>vvvv<v>v>^^><^><^vv^<><vv<<<>v<v<^<^>v^<^>>vv><^>v<>v<v><>v<^^<^^<<vv<>v>^v>^^>^<<^vv^>v<>vvv<<<<>>>>v>>v<>>><vvv>v<vv<>v<<>vv>>^vv^^<v>>v>v^<^><^^><^<<>^>v<vvvv<^^v^vv><^<^>vv^<>vv<v^^v><<>v<^^v>vv>^<vv>>^<>>v<v^^<><>^<v<vv><>v<<v>v><>v^<>>v^v<>>>^>>>v>^<^v<^<^^vvvvv>^<v<<v>>>><<^v^^^<v^<^<<v<>v<<<^<^^v<<v>>vv<>>><><><<<>^<vv>>v><<><<>^^v><v>>vvv<^>v>^v<>>>^>>>^^^>v>><^^>v^vv^^v>>>^<>>vvv>v<v<vv^vv<^v^>>>^v>v>^vvv>^v>><<<v><^>^<^><vv<<>v>>>>^v<vvv>>^<>v>^^^><>>v^>>v>^^>v>^^>^^<v<>>><vvvvv<vvv^<<<v<<<>^^v<^><>^>v<^^>>v>>>>>^^<<^<><>^^>>>v><^>>^>^^^^><v^^><<<^<^vv>v<>^<<v^<^v^<<<<>><>>><^<><><v^v^<vv<>v<^>^>>>^<v><^v^^<<><>>v<<^v<<^v^^<<vvv^^
>^<<^>><>>^v>^<<^^vv<<>^^v^<v<^^<>><>^<<^^v^<v>><><v<v^<>v<^>^v>^v^v^><<>>^v>><>>^<<>vvv^<v>^v>v<vv<><<>^^^>>><<>^><vvvv^<<<v><<^v^^^>>vv^<^<v>v<<<><v>><>v^^^<^^^>^>>>v><v<v<<><><>vv^v^<^<<<^>vv^vv>vv<>>>^vvv<><>^><>^^vvvv<^v<<<v>><>^<<v<^<v<^><>^v^^>^^<<<>v<><vvvv>v>v<^vv>vv^^<^^>vv>>^<<<v^<<<<^>^><^>>^>>^^<vv^^v<^>><^^^v>^><^>^^<v^^^>>^^v>>>^<>>^v><vv^>v<>^<<^<^^v>><>>^^^>>v>><<^vvv<<^^v^>^vvvvv^v^vv><^^<v<vv>v>><^^v<>^v>>>^v<><v<<>^<v^>>vv>v^>><^<^<>v^<v^vv^^^^^>^><^>^<><^<v<><v>>>>v^>v<vv^<^vv^^v><><><vv^vvv^>v<v>^^<<^><<>vvv>^vvvv^<>^>><^>v<v<^v<v^^><>v^v^<v><<^<v><v^vvv>^vv>vvvv^^^vvv>>v>>^v>v^<<v^^<<<^>v<v>v>v<vv^><<v>^vv>>v<<^<>^>^v<<<^<>v<^<<^^>>><vv<^<^^v><vv>><vvv>v<^v>><^^v<v^<^<^v^v>^><vv>^^vvv<<v><<>v<v<vv>^<^><><<>v^^v^^^^^>vv^^v<^><^vv<<<v<>><^<^>vvv<<v<^^^^<^v^<><vv^>v<vv<>^^v^>^<v^^>v^>^^><>^><>^<v>^v><<<>^^<v^>^^<v^^>^><>^^<^<vv>vv<<<v>>^>v>vv^v><>vv^vv>^^^><<vv>^v>v^v^vv<<^^<^<^^^v^<v><<<v<<^v<vv<<>v^^<^<<v>v<^^<<vv^><<<v<v<vv<vv^^^>^v^>v^<<^^><^>v<<^<^^>><<>><<><><
>v<^^>>v<><vvv^<><><v>vv^<>^^^<v<>>><>><<<>v^^^<<^v>>><v>^>v^<>^vvvvv^>v^<<<<>^>v><^>^>>><v>><v>^^vv^<v^^^<<<>>>>^><>>>>vvv>^v<<<>v^<^v^<^<>^v><>^>v^^<><^v<^<<<>^v>><vv><><<>^<vv^>^>>v^vvv^<<<vv<vvv<<v>vvv<v^<>v>>^^<<vv>v<v<^^>^^>^>>>>vv<v>v^>>^<>v<>v^>v^>v<^<<^<<<^<>><v<^vv><<v><>vv<<<<v<>^^>v<^>v<>^><>^<vvv>><<><>>^vv<>>^<<<v^v^^v<<v<^<v<^^>v><<^<v^v^^v<^^><>>v<>><^^>>>v^^<>v><<v<>>v>>vv>^v>vv><>vv^^<<<v<vv>>><^<>^^<^>>>^<v<<v>><^^<<^<><v<v<v<><v>vv<<<^><<v^^>v<v><v^^vv^>v>^<^v>v>>><<^v><<>^v^^^>v>vv<v>^vv>vv<<v^<><<^>v>^v^<>v^^^<>^<<^<v^<>v<v<<v<v<<>^^vv>^>^>^v>>^>^>vv^vv<^v<<v>v^<><><^^v^>v^<^>^v^<v><<>v>><^<<^<^>^<<v<><<^^<v<v>v^<<v<<>^vv>>v^>><<^^v>^v^v<<^<^v>>^<^>v>^^v^v><^<^>^^v<<^v^<<^>v^>^<<v>><>v<>>>>><><><v^^v<v<<><^v<^<>>v<>v>^<v<^vv^>>><^v<v<>vv^v^<^<><<^>>^^<<<<>>v<v^^v>><^^^v>^v>v>><^<>>^v>>vv>^^>v><<v^<>^^<^^>^<v<<>><v^<>><v>v<^v<^v<<vv>><v<v<v^vv<<>^^>^<v>><v^^v^><><<^<^><^vv<<vvvv^>>^>vv^<^>><^<v>^><><^>v^>>^><^<^<v>^>^>>^v>v<>>vvv><v^<>^vv<<v<^^^^^^vv<<<^^v<^<^vv^v<
>><v<^<v^><>v^<>v^<<v^>>v<^v><v^<v<^><<>^>v<>>><><^^^vvv^^^^v>>^^v>><v<v<^<>vv^><><^^v^^vv>>^v<><<v^v>v^<^^<^^^><vv<v>^vv<v<^^<^vvv><<<^>><^^<^>^>>>^v<^^><>^>v><><v^<>v><v<v<<^<<vv>>^v<^v>^>>>v>>^>^<<><<^>>>^<>v^^v><^^<^^v<<><v>>^<>^><<<>>>^>><<vv>^<^<>><v>^<>v^>>>v<vv<^<^vv^vv>>>^^>>^v^v^>><><<<>vvvv<v>>><^<><><>^>v^^<v^<<>>>>v><v>vv^<>v>^^v^><^vv>><v>>>v^^vv><<^vv<^><><vv<><<v<v>^vvv^<>><^>><<<><<^>>v>><vv>v^v<^<>^<^v^>>^><><v>^vv>^<v<vvv^v^><>>>^v<<>><v<>v>^v>>v><<<^v<v><<>>v<^>^><v>v^^v^<>^><^<>^v^^^v^><<^vv<vv>^<v<v><<<vv><>>>v<<v^<>>>vv><v<><vvv>><<^>>>v^v^<^>^^vv>>v<v><^<>^^<<<^>v><>v>v<v^<>>^>v><>^>><<>^vv^v>^^v>v^^v>v><>^>^<>vv^<v^<>><^v<v><<^><^<<<v^^^<v^<>v<^>^v>^<^<>^>v>v>^>vv^^v<<><v>>><^<vv^<>v<<>>v<><v<vv<<^<v><>v^^^<v^v^<<^>^<<<<<<^>>>>>^^><<^>>^v>^><>v^>^v^v>>vvvv>vvv><v<<<^vv^>^^^><<<v<<>vv^>^<<v^<v>^v<<<^>v<<v<v<<^^^<<<^>><v<vv><vv>><<<<v^<^v<^^<<v<<<^<v>v^^^<<<<>vv>^v>^<v<^^v^>^v^><>>v^^>v^^^<^^<v>>^>>^<^><vv^v^<><><>^v>>>^<^v>>^vvv><><v<<^v<^><<v><vvv<>><<vv>><^>^v
>>^>^^^<v<>^v^^<<^^<v<<v>>^<vv<^<v^<^^v<^>^<vv>v<<v><v>vvv^>^v>>>>v<v^>v^<<>^<<>^v^^^><<vv>^^^<^v>^<>v<^><vv^>v><^^>>>^>v<<>^><>^^^>>vvv>^v^v>>>^<v>><^<<v^v>^v^>v>^v^<>^>>^>>>>>^<v^><<<>v<vv<<<^v>vv^^>^>>v><^<v<v<vvv<<<v>^<^><><vv><^^^^>v<v>>>v<<><^^>v^v<<v^^^v>><<>v^v<^^vv^>^v>^<^>^>v<<>v<v<^^<<<<^^>>v<>>v<><^^^v^<<><<^v^<v<<^>^>>^v^<^<>>^<<^>v<>>>^<vv^^vv<>><^^<v<v>^><vvv^^<>><>vvvv^>^v^><<<vv^^v^v<v<v<<^<>vv>^>>><>>vv^>^>>^^<v^<v^^<v<>^><<^>v<>^v><<v<v><^<^^<v>v>vv>^^^^vvv>^v>>^^^^<v^>>^^<^>vv^>^<^^^<^^<^>><>^<>^v<<v>vvv^v^<v^<<^v^<v>>v^><>><v<<^>^^^>vv^<<^<^<v^^<<<^<<><>^><>v^^><>^^^><<<vvvvv^^^<v<^^<>vv>>>^vv><v<>>^v>>><<>v><<<<<>><^vv><>>v<>^<v><<vvv<>><<><v^^<<<^<>^>^<^<<^^>^><>v<>vv<v<<><>^^v^>v<^^<v^^v><^^vv<>><>v>>v<>>v^^>v><><>v^v<^>><<>>^<^<v<>^><^<><>>><<<<><>^v<>^^<><<v>^>v<^<^<v>>^v^vvv<^vv^^>^^v<>>^^<^>^v>^><v<v<>v^<<<vvv^<>>vv^>>^<^v<<>v>v>^^<^<^<v>^<v><>vv>>^<^<^^^>>^^v^^^^<^vv<^^>^<><>v<>vv^>^v<><>>>><v^<v^><<^^^v>^>^>>^>>v>^vv>v^>vv^v<v^v<^<^>^vvvv^<><v>><<v>><^^v>>
^vv^<vv><^v<>^v><^v>><^<><^v^>v>^<^<>vv>>^<^^^v<<^^>v<vvvv^<<v>^<v^<v<v<^>>><v^>v><<vv^^<v>^<<<^>^v>v>>v>><>v^>v><>^>^>v<<<>^v><><><<>^>^v><>v<^<v<^v<>^<^v>v>^vv<v<>><<v<v>><vv>^<^<^^^<<>^<<<>v<<^^v>vv<^v^v>>^<v<^^vv^vv>>^^^v><>>^v<<^^>^>^>>^<<<v^^><<<^^v^vv<<>v^>>v>^vvvv<<><><>vvv^v><<^v^v^>^><vvv^<<^^<><^<<>^v^<v^v<<<>^^^^v^v<^^><<>^^>v><<^vv>>vv<v<v<<>v^^^<<v^^>^v>v>vvvv^v<v>v^^<>>^>^^>vv>^>>v>>><v^^>v<^<<><v<v^<^<v><><><^>>^^^>v<^v^><^v^>^><><>^v>^v^<><^>>vvv>v^v>><^^^vv<><>>>v>^<v<<^^v>>v<^vv<vv^^<v>v<^<^>v>v><v^><^><<^v>v<<<>^^>>>^vv<^v^<<><<^<^<v^^<v><>^>vvv<v<vv>^<^^^^><<>>v^>^<vv<v><<^^v<^>v^^^v^^<>v<<<v<><v<<>^v<>><<v^<^<^^^<<<<<>v>v^<^^<^><>^vv<>><^^v<^^^<^<<>v><<v^<<>v<<<^^><^<^vvv><vv^v>^><v^^v<v<v<<<vv>><^^>^><^>v<v<>vvv<<v>><v^v<vv<^vv<^v^>v<^^vv^^<>^v>^^><^^<vvv<^v><^v<<vv^>^v<v^<v^>v><^v^<>>vv^^>^<>v^v>v><<^<>>>v<><<<<>v<<<<<v<>^v>v^<^>><^^<v><vvvv^>^^>vv><><<^<^>^<>><^<v^^^<>>^<^^<vvvvvvv>^^<v<v<^>>v>vvvv>^<vv>^<<v^v>v^>>v^>^^v<v^>^>^^vvvv>v<>><^v>v<>vv><^v<^>>^^<vv^<
v<<v^v^>^<^<><>>^<v<^v^>^>^><v>v><^>^^>v><^<><<<<>^^v^^^v>^v^<^^vv^vvv>v^vv^>^<><vvv<<<vv<>^<v^v>^^^>><<><^>^>^^v>v^<>vv>v^>^v<<v<v^>>v^<>vv<<<<<^<>^vv>vvv><><>v^>^v<^vv><^vvv^^>>^^>v^^vvv^<^>>v<^>v<<v<vvv>>>>>v<><>v>>v^<^^>vv><<>^vv<^<>>>^>>^<<>><^>v<><>^>>>vv^<>><<>v>v^>>^vvv^v^<>><^^^v><<vv^^><v^v^><v<<^<<v<>^vv^<^><^v>v<><^vv>^^^<vv><vv^^><<v^vv>vv^<>^<v<v<vv>^<v^<<>vv<^v>^vv<>^vvv>>><<v<v>^v>>><<^>^<>^v>v<vv^^v<^^<>><^v^v<v<<^>>v<<<^v^>^>^><^vvv>>><^>>v^^^<<>><vv<><v<>>>v>^>vvv<<^>v^<^><<v<^v^>>>>v<^<vv>v<<^>^><><<vv>>>v^<>>v^^v<>v<v^v<v^<v<^^vv^>>v^>^<v<vv><v<<vv^>v^<<^^>v^v<^vv<v<vv>>^>vv^vv<>v^<<><<>vv<><^v<^v<^^>^<<^<>><v>v>v><<>^v^<<>><vvv>^v^v>^>>v^<^<vvvv^vv^v>v>>^><<<^^^vv>^<><^v>^v^v^<^>v<>^v^><v>><>><<<<<><>^^<>^<>^<v<v^<<<^v^>>>>>^^<v^^>><>^^^v>^<>^>^>>^<v<>^v<>^>v>^^>^^^^^vv<^<<^<<<^<^>>>^><^^v>^^^>>^>>v^v^v^^<vvv<>>>>v>vv>v>>v>^>^<>>^<><>>^>^vv>>>v^<>><<^^><v>v^>>><>v<v^<v<vv<^>^<><<^><v<v<vv^>>^><v^^><<<>><>>>v^>^<<<v<<v<vv>v>>>>>>^v^vvv^v^<vv>v>v>^v^v<>^^<v<v^vv>><<
v>^^^v<v<>>^^<>^>^^<<^^vvv><<>vv^^^<vv^^><>vvv<^v><<>>^v^^>^v<^^v<>>^>v><<^>^^<<<v<<>^<v>>>v<>v<^v^^vvv^<^><v<>v<v<<v<vv^<<<^<>v<>>^>v^<v>v>^<>>v^^<vv><v>^^^<v>^^v^^>^^><<v^><v>^^^v^^>v^v>^<><^>^<<>>v<^<^>^<><<^v><<^><^v<v>^<v<^^^^^^><v^v<>v<>><^^^^<^><v<<^>>v<^>v<<<^<>vv<>v<<v<v^v^v^<<><v<<<vvv>>^<^>^>v^vv<^v<v<><^<v<>v<<^>v^v>v^^<v>>^<v^^<vv<<>><^><>^>>v><^v^<>v<^v<^>^>v><vvv>>>><>^>>vv^><>><<^>>vv^>><>v>>>^^<^^>vvvv>^^<<><vv>>^>^^><>>>v^>v><vv>v^<vv<<><>><v<vv<v<v>^^<><v^>v^^^v<<<^<<v^^<v><^^<vv>^vv>><^^<><<v^<>><<><>^^^vvv^v<^^^^<v^>^>^^<v>v>v<v>v>>v<<^>>^^^v<vv^v<^v^<^v^>^<<><v<>>>v^><^><>>v>^>^^vv^^<^<^v<^>^<<>>>^^>v>>^^^^^>v<<^^^<<^^^v>v^^>v>>v<^<^<v>>>>>>^>^^^v^v^^<v<<<<><<^^<^v>><^v<>^>^^v>>v^^<>vvvv>v>^<^vv><v>v>><v<v^<^^^><<>^v^<v><>^v<vv><>^><>^^v<<v>v><>^<v<^v^v<<<><vv><<<>^<<v>><>><v^>>vvv^<vv>>v^<^^<v<><vv>>v^>>^>>>v>^>>>>^<v^<<vv>^v^^>><vv>>v>>>v>v<>^vv^v><<<>>vvv<><^^<^^^^>>v<vv^><v^v^^^vv><>v<^^^^vv<^^<v>^^^><^>^^><<^><<^><v<v>^vvv<><<v^<<^^<>>^v<v<v^<>>^<^v>vv^^<>^>v
<v<v<v^>^<>^v<><>v><<>v^<<><>^^<<v>><<^v<v^v<>>v<^<<v<^v<v^v>><><><<<v>^v^>^<v>^<<^v<vv^^^<>>v^^<vv>>^^^><<vvv>^^<<v>>><v^v>>>^v^<vv<^>^<^<<v<>v>v<<vv^>v<><>><<>v<><^^<^v^^^v><^^^<v>^>^v^^v>^<^><v^<>^>v<<v>><^vv><^<<<<vv^<^^^v>^v<^<<><^^<>^v>v<>>^^v>^>^^v<<^>vv<vv^<>vv^<<^<^>vvvvv^v<<v^^<^v<v>^<v^v>^>v^^>^>v^><v>^vv<vv^v^v>>v^><<<>^<<<^<<><>^v<v><vv^v^v^v^v<<v^>^<><^^><^><<>>><>vv>>>>^<<^>v<^v<<<v<>>^^>^>>vv<v<v<<v<v^^><>><>^>^<^<<v<v^vv<^<>vv^>^<>^vv^^^<<><v^>v<v^v^>>v<^>v<<<>^^>vv^<v^>^v<>^^>v^v<^><v^vv^>^v>vv^>^v^^>v^<^<^<v<<v^<<^^^v>><>>^^^^<^^>^>><v<<<><<^^<^<vv<^<><<v<<<><<v^^><<>vvv>^><<<vv>^>v>v>vvv^^v>>^v><>>v<><<>>>v^><^^<v<vv^<<<v^^>><^>v><<><<v<^v<<>^^><vv<^<vv>><>>><<><v^<^^<v<^vv>v>^>^>>^^<v<^v<vv>^<^>>^><>v>v<<<^><vv>^^<^^<>vvvv>>>>>^^<<<^>>vv^<vv<v^<<^>^^v^^>vv<<>vvv>v^>^v^<^^^>v>^<v<<^vv><v><v>><v>>>^v>^v^v>^<^>v><vv^^v<>>><<><>^>vv^vv<vvvv>vv<<vvvvvv>^>v><>v^vv<^v<v^^<v<v><^>>v^>>>vv^<v>v>v>>vv<v<><>>v^<v<<>>>>v^>^v>>>v^v<<<v>v><vv<>v<^>^<^>>v<v<>^v>v>^v<<<^^<<^^<<v^^
<>v>v^^>v<<>>>^v^v<>^>^>v^v><v>^>vv>vv<<v<v<^^v><>^^^<<^v<^^><>v>><v><^^vv><^v>>vv>vv^^<><vv^><>vv<v>>vv>vv<^^>v>v^v<>^^<<^<>>>v>^^<^>v>v^^v>vv<>^^v<<<>^<v<v^^<<<>^>^<>^>v<><v>>>>vv><v><<v<<><>><<<^>v^^^vv<v<<<<v>v>^<><>v>>^<<<>>^<vv^>v^v<v^>>>v>^>v>>^v<<><<vv<v^>v^^vv^>>v>v>v^v<v>v^>v>v><^<<vv><>>^<>^v<v>vvv>v>^v^>>^><vv^^^><>^<vv<^<>^v>><v>^v><><^vv<><^><<<v<^<v>v^<^v<>^vvv<><^v<v^><>v><>><<>^^v><^^<^>^^<>vv^^v<^^vv^<v<<^v<vv>><>v^<^v^<<v><v>><^v^^><^><v><>^^^<<^v<v>>>v<<v<<<vvv<>>^^>>^vv>><^^<<^^vvv^<v>>>vv^v><<>>^>^vv<>vv<>vv<>^^<><>>v<^<^^^>v<^^<>vv<^><v^><^<^^v^^>v<>>v>^v>>>vvvv>^v<<>v>^^<>vv><^v<^>^v<>v>^<^v>>^vv^v^>v>><v^><<>^^<>>^v^^^^>^^><v>^v>v^^vv>^v<><^v<<^v^^>>v>>vv<>^^>>>^^^^v>>^^^<v>^>v><vv><v<<v>^><^vv<>>vvv^^><^^<<v>^>^>v^^v<^>v><v^<>v>^vv<v>^v>v^^<>vv>><>v>v^>^^>v<>^>vv>>><^^v<vv>^^>vv><v^<<^<v^<^<v<vvvv^^^<^v><^<>>^^>^<v<v>^>>^^<<><^^v<v^v>vvv<vv>^v<v<<^<>>^vv>^>v^<<^<>><vv<>^^^><v<v><^>>><>^^><<^^^<vv<<<<v>>v>^><<><<<^vvv^^v<vv>^v>>^<<v>>^v^^>^><v>^<v<v^<>v<<<<>><^
^v<^vv>>>>vv>v>><^<<vv><<^<>>>^vv^v^v>vv^<<<v>>>v><v><>^><<<^>v^>v<v<v<<v>^>vv^>v^^>>^^><>>><<^<>^^<>><><><><><v<v^<v>v^^^vvv>>^vv>^>>>vvvvv><<^v<<>v^v<>^><>>>vv<^>vv^<^v<><v>^^>v<v<<>vv^^<>vvv^>>^>><v>v<^>>>>>>><<v>^<vv<>^<>v^>^><>v<vv><<>><<vv^>>^<^v<>>^v>^>>^vv>>v<^>^^<v^^vvv<<<vv><><v^><><>v>>v><<<<>^>><^<>v>^^<vv<^<v<<>v>>^^<<<<>>v>^>v>v<^<^>><vv>v>v><>^<>^^>>^<v<<^>^>^><>vv^vv>^^<>>^<v<^^<<<v><^^<^>vvv>v^^^>vvv><^>v<^<v^<>>>^v^^<<v^<<^v^><<><<vv>^><<v>>^v><v^v<^>v<v^^^<>v<v>^><<v<vv><>>v<<<<><<v><<vv<v>>^^>>^vvvv^<^<<v^<v<v><<^>v>v>^^><^v^^<>v^^>>^vv><^^v<<<<vv><v>^^><v^>>^>^v<<^><<>>^^<^v<v><vv^<<><<v>vvv>^<<vvvv^>v<><v^v^vv^v<v>v<>>^v^<v<<>^v^^><><<^^^^^^<vv^<<<vv<<v><vv<^^<^^vv<<<>><^^>v^v>^><><<^vvv>^^>>v<^<<<^>v<v<>v>vvvv^>>>v>^vv>^<<<v<<^>^<^><vv><<<^>>^^>>^vv>>>v>^vvvv^><v<^vv<^^<v<<><^vv>><>>^^^vv^^v^<>v<>><<<v^^>v>^<v>^v^<<<><>>>v^>^^^vv^v<<v^>^^><>v>v<^>><v<>v>^>v>vv^>^^^^^vv<v>>>^^^<^<<>>^^v<^^><vv^v><v<v^>>v>v<vv^^>><v^><vvvv^<vv<v<>>v><vvv>v<^^>^>^v><^>^v<v><<v<<>><^
>v^vv^v>^^<vv>^^v^<<v>^>><>^vv^>^<vv^><<<<^v^<<^><v<^vv^^v^<^vv><>vv<v>>vvvv^vv<^>vvvv^^^><>v><>^v><<<^<<^v>>vvv><^<<vv^<^^>v>><>v^^v>^><^>><>>>>><>v^<^v^>v^><v^>v<>vvv^<<v>>vvv>vv>v><<<>^^^<^>^^vv<>^<^v>^^v^v<><v<vv>>^<v>v^<^>^<><v<<<<>^vv>>v^>^v^<><v^<<^v^<>^v<^<v^^>v<<>><^^v>^>v>><v^<>>v><^v<^^vv^<^>>>^>v>>^>vv^><v<^v>^v<>><>^><<vv^<>vv>^<><^^v><><v>^>^<<v>>^^^>^v^v<^>>^vv^><<><<>>><^>>^>>>><^^^v^><<v^>^v>vv^<>><^^>v><<><v^<v^<v>>>v^>>>>><<vv>v<>v^^>v^v<^vv^^><vv^><v>^v^^><v<<>^>>><^^<^v^>v>><<^vv>vv>>^<vv<vv>^><<<vv<v<<^<^>^<vv^v^vv^vv<><^<<^^^v<^<v<>^^^v^^>v>^^>vvvv><^<^^^<v>vv<vv<^><^<>^v^<vv^<<v>>v>^<^^v^<^>v>><^^^<<>^<v>^><<>v><^^v><vv^>^^<vv>^<<v>v<^^>v<v^>>^^<^>v<<^v>^>^v>><v>^<>>>v^^v<v>>^vv>^><v>>^vv^>^>^>v^v><v>>v^v<^v^<^<v^><<<v<v^^v>>><><<^><<<^>>>vvv<<^<^>vv^^^^<<^v^><^>^v^>v^^>vv<<^<^^^v<v>^>>v^>^^><>>v><>>^<v<>^>>vvv<^<<<^vv>vv<v^<^v>><><<vvv^><<<^^<>^^<<^<^<^^^<>^<>^^v><^^vv<^<v<^^><<v^>>v^^v>^vv>^v<v<<><<^vvvvv>vv^v^>vv<^v<^<^<>^>^v^><v^<^<v<<^vv<>^>v^v>^v<^v^<>^<v>
>><vv<v>^<v<^v><^<>v>^><^v><vv>>^<v><v><^^<^>v>vv<<<v^^>^^^>>v<v><^^>v^>><v<<^vvvv<^>>vvv^>>>^^><>v<<<>^v>><^<v>vvv>vvv<vvvv>vvvv^vv^^>v<^>>>v<<vv<^>v^<<v^vv><<<^<<^^<>^><>><>^vv<^v<vv<<><v>^v><^<<v<<v^<>v<^v^>v^<>^vvv<<v>^v^>>v^^>v>^<v>>^><v>^v<vv><>v^>>>v^vvvvv<>v<><>v<^>>v<<v^<><v^>^<v^<^v^^v^^^>v<v<>vv^^<>^<^v^>>>^vv^v><^>^v<^<>v>v>^<v<>^^v^>^>><<vv^^^>><><>^^<>^<<>v>v<v<^<^vvv><<><><^>v>><><^^>>v<v<<v><>>>>v><<><<^<v<v^^>>^><^^>^<<>>>v<vv^^>v^v^>>>vv>^^^^<v^^>^<v>^><v^<vv^>>v^v>^vv<vvv>>>><^^v><^v^<>^<<>^>^^>vv^^v<>^^<^^^>><v<<v<v>^v^v><>^<^>>v<vv<v>v>>^^^^v^v<<<>v<^<><>>v>>>>>v>>>>^<^vvvv>>v<>>><v^^^>^^v<v^v^vv<><>^<v<^<<<<^^<<v^<>v>v<>^<^v^v<v^<^vvv><^>v<v<><<^<v^>vv^<<v^<^<<^<<<v><v<>^^v<>><v>^^^^v^>vv>v^^^^<^^<<v<<>><vv<>>><<^v<>>v<^><^<>v<vvv>v>^<>vvv^vv^v<<<><v<<<>v>v><><v<^v<>^^<>>>><>>^<<<^<^<v<>v>>>><v>>^>>>^<v>^v^<>>^<v^^^><^vv>>v><<^<v^v^v>v>^v^^v^>><v>^<>>^>v^^<^vv>v^v><v<<v>^<<>v^>vv><<^<v<><v>>^>>v<^<^^^^^<>>>>>vv<v>>^>^<>^<v>v><^v<v^^^^^<^><<<v^^><<><^<<v>>^^v<^<vv^
>><<^<<v<>^><^v>>>vv^<<><v>^<<<^>^><><v<^^^<^<>>>><vv<v^<v>^^vv>v^<<><vv^<^<vv>^>vv<^v^v>>v^<>v<<<<>>^<v>^<^<>^^<<>>>>^>><^>^<^^^>^^<vv>v^<><<^v^v^^^><^v>>>v>^^>^v>^>v>v<<<vv>><>v<^^v<>v>^<<<v^<>>>^^<<^^>>>^<^<^v^>><<v<<>^<^<v><<^<<v<<vvv<<>v<v<^<^>v^>vv^>vv<v^<<<v<^v^<v^v>v<v><v<^>>^^v><><^<v^>^^<<^^>^^^<>v>^>^<^<v>v>v<^^>vv<<><><>>>^^^>>^^^<><v><vv^<<>vv^v<v^>^v<^<vv^^>vvv><>^v^>^>>v<v^v<>>^^^<v^>^^<^<^v<<v^vv><>vv^<>>vvvvv<^^v^>^<vv^v>>><vvv>>>>v<^v>vv<v^v^<^><^<>v<>^v^^^^^<v>v^<>>>^v>^vv<<v<<v<v<vv<v>v^>v>^v>>v<>><>>v>v<<>>^v^vv<<vv^<<<>v^^>^<<v>^^^><v^^<^<vv<vvvv>v^^<<><<<>^<^>v>v^vv^><>>v>^>>>>>^^v><^<>v^<^v><^vv^<v<v<><<v><^<>vv<^^vv>>>>^vv<>v^v>^^<<<<v>v>>v><^>v^>>v^><^v^vv^vv<>^vv>v>>v>vv>^v<v^v^<^><<v^>vv^vv>>>v>v<v^<<>^vv<^^<^vv<>v^>^>>>vv><v^<v>>^v>>v<^v>>>>^<^v<<<>v^v^>^v<<vv^^^vvv>v<v><>^>^vvv^>^^>><^><><v><^^><<><<^<v>^vv>>^vv>>>>^^v<^<^^><<>vv^<^v<v<>^^<><^<<v^v^<<<>^vv^^^vv>>^^>^<<v>v>v^^v^v<>^v^^><><>v^^vv^v^v>>><<<^v>v^v><^^<>vvv>^v<^^v^^vv>^<^<v^^>^vvv^<>>^<>><<v><v
>^^vvv^><<^>^^v<<^v<^v^>vvv><vvv>><v<<<v^v^^<v>><<v><<v>^>vvv>>>v<<^>^<>^^<^<^^^v>^<>>^vvv^^<>vvv^^^<<>>>v<^v<<>^^^vv<<><<>v<>v>^<>>v^>^^v^>v<<^v<v<>^v^vvv>vvv<^<^><v>v>><<^<>>^v>^vv>vv^>^v^><>v^<^>vvv^v>>>^<>v<>^^^^<vv>>^>>>><><<>>v<<<vv^^>>v>^>vv^>^<<<>>^v^><^<vvvv^v^^<v^v^^<v^><>v>v<>v>>^^vvv>^v^<v<v^^<<^<vv^v<>v<v<<vvvv>^><v>>>>^>^^^<<<<<>^<><v<^<<^^vv^>v><v><^^>v^>^>^^^<<<><<><>^>><<^<<^v>^><>^^^^^<^^<^^^vvv^v^^<v<vv^<^<^v^v><<<>vv^v<><<v<v<^^>>^>>^<^v^v<>v^>^v<><<v^>v<<v<v^v>>^<<<^><>^^^<<><^<^^^>v<><v>v<>^>><vvv><v><^<^v^^<>^^^>>^^><vvv^<^<<><^^^v<v<>^^v^<v>><v^<<vv^>^>v^^<>^v<v>^vv<><^<^v<v<>>v<v><><^v^<vvv^>>v^vv<<v>>>^v>^<vv><v<>>^<v>>><vv>><v^^>>>>^^<>^v>>v<v^^^v^v<<^^>>^^^<>>>v>v><v><vv>><>>><>>v<>^^>>>^v<^^^^v^v><v><<^><>v^^<><<>v>>vvvvv^v>>>v^<v<<<^<<^^>v^^v^^^<vv<>^^<>>vv>^^v><v><v>v<<>^vv^<v>^v<v<v>>v^<<^vv^>>>vv>>^<<vv<<><^><^^vv^>^<<v^v^^<>>^>^v<<v^v>>v><<<<><v^>>>^<>>>^>^>>>^>><vv>>^<v>>v<<<<>^^<>^>v>v>^^v>v>v>>^^<<<<><v^^^>^v<v>><^vv^>><v<v<v>vv^<v><<vv<v><^^<><<^v<
v^<>v<<^><<<<>^v>^^<<><^>v<^>v<v>^v<<v>^^^vv<<<^v^^<^<v^v^v<><><<<^<v<^^<vv^^vvv<^v<>><>>^v>^v^^<>vvv^>^<>>v<^<<vv<^<v><v^>>^>v>vvvv<>>>vv><<^^>^<^^>v>><^^^^<v><>v<^^<<><<vv>v<vv<>vvv>^<<>>vv^>v>>vv^>>v<<^<<^^>><v^^<>>v<v<>v<<^<<^<vvv^^v><<>v^v<<v<<>vv^vvv>>^>^<<^><><<^<^v^>><vvv^<>vvv^^>^><^<<v^^>>^^<<>^^vv><>^vv><>v>>><v>>^<v<v<vv<^vv^^<v>v>>v<>>>>^>v><^^v<>>^>>v<^^^^>><^v^^v^v><<<^^>v<v>^>v^<><<vv<v<^^v>^>^><><>^v<^^^v>>^<^<v><^vvv>v^<^^^^^<v><><v^^<>><v^>^<v<^^>vvv<v<<<>>v<^>>>v<<>^vv>v^^vv<^<^v<>^v<<>v^v<^vv<>v>><<v>>^<>^<><v<<v^>v<v>><^>>v<>>^v<<^v<v<<^><^><<>>>>><v<^^><>><v>>v^^v<^<>^>>v><v<vv^vv><>v><v>>^v><>v<>^>v^<v^v<>>^>^vv><>^vvv^v^v>v^^v^<>^<>^^>v>>v^v^^><>^>^^<^>^>^v><v^<^>>vv>^<<<<^>^<^^>>>^>>^v<<>^<^<v^^vv^^^vv<<v>>v<vv<v<^<^v<<^<^^^<v<^v^<v><v^v>v<v<^>^^vv>^^<><<v>>^v<v<v><>>^^v<vv<>vvv<v^>v>vv<v><vvvv<vvvvv<v>v>^vv^>>v>^^^v>>><>^vv^^<v^v>>v<v^v<>v>^^v^^<>v^>>>><v><v^^<v^^>^^^>^>>^>^v^><vv>^v^^<^>>^vv^^v>v^><>vv<vv^>v><<<<^<<>>vvv^vvv<<<>><v>^<>>vv><>>>^v<>v<>^v>v<>>>
^>^v<><>v>>^^^v<<<^^<v^<v>v^<<<<v<>v^^^vvv>v>^<><><^<<^>>v>><<>>^^<<<>^<^<<>^<^vv>><^vv>>vv>v^><<><^^><>>>^^v<^<v><vv<<^^^^^>^v>>^>>>>^>v^>v^vv^>>>v^>v<>^><v^^^^>^<>^<vv<><>>v<<><^^<><vvv<<>^^<v>^<<^^^vvv><<^v<<<>^^^^v>v^<<<^<^<<v>^^><vv^><<>^^>><<>vvv^<><>>><v>v<v<<>>v>vv<><^<>>><<^<v>^^v^v<<><v<<>v^^<v>^^<^>v^v^^^^>^^>^>v^>^v><^vv>^><>>>^<^<v<>^^^vv^>^<<^>v>>v^^>^v^>^<vv^>^><vv<<>v^<vvvvv><^>>v<<>>>^>vv<>vv^<vv<v^<v^vv<^<^^vv<v>^^v<^<v>vvv^^^<^<v>^v<<>v><v<<<><>><><<<v^><>>v<vv^<>^<<>^><v^<<<>v>v<>v<<^<^<<^v<^>^>>>^>^^<v<v^<<^<^>vvv^<v<>^<><>v><<<<>v^^v^v^vv^v<^^^>^v>^^v>vv<>^<v<<><><^><<<v<>v^^<v>><^vv^^>vv<<<^^vvv<><>>>^v^><<v^vv><>^<v><<>>vv>^<v<>^<v^^^^v>^>>vv<vv>vv>>v><<^v<>><<><v<<<<^<^<^^>v<>>v^^<<<>v>^v><<>>>v<^<^>v>^v<<v>>^>^>vv<>>^<^vv^v^>^><^>^^<^<v^v><^v<<v^^<^^>>>^<<v<<vv^>v<v^>^vv^<vv><<<>vvvvv><<^>vvv<>>>v<<v<^>>^>^>v^vv^^^><v^<v^v<<vv<^<><><>v><v<<^vv<v<^<><<<><>>vv^>^v>v<><<>>v^^v^>^<v<<<v^v>vv<^<<vv^v^<v^^^^<vv^<v>^^<^>>^<<v>^>>v<<v^^<<v<vv<<v^><^^^^^<^v<^v<<v^^v>^v
>vv>vv>^^<v<vv^>><^>v<<>v<^^<^<^>v^<vv><>^><v^^^>><>^<^v>v^>vv>>^>><v^^^<^<^<^v<v><^^<>vv^<v<<<^<>v<v>v<<^v^^^>v><><^<^<<v<^v<vv>^vvvv<v<>^>>v<<^>>vv^vv>vv^<^<<<>^vv^<<>v^<^<^>>v>^<v>>>^><><^<v^>v^^>^>>v<<>^>><><<<<vvv^^vv<^<^>vv>><v<<>^vv<><<vv^<^<>^<<^v^v><<<>v>>^vv>>^><><>^^><^v<<>vv^>^v<^>^^^>><^v>v^v<>^^^^v><v<^^^^<v>>><^><<<v<vv^<vvvv^v>>v>v>^v^>v^^>^>>><>^<<^<^v^v<vv<<v>^>><^><<v<<<><^<<v<^<^<><<<v>v^>><vvv>v>v^<>v<vv<>v<<^vvv<^<<><>vvv<<^^v<>^<<^>><v^^v>>^>>v<>>^v^v>vvv^<<v<<<<>vv><><<v<>vv>^<v>^<vvvv^>^vv<><v<<>v^<><<^v<<vv^v>v^^^v^<<>v><>><<^>^^<v<<^vvv^^>^v<>>><<>><<>>>>>>>v>^v^>>^>>>>^vv<<v^^<v<<v^^>v<>^<^^>>v<<^v^v<<><v^^<v>v^>^<^vv>v><v<^<>v<<^v>v<>>^v>>v>^<vv>^^v>>^v>v>>>^<<vv<>>^<>^>^v^v<><^><>vv^<v^><>^^vvvv<<<><v^<<vvv^<>^<<^<^v>>><^>v^>><><<>^^^>vvvvv>v<<<vv<<>^vv^<<>^^<>>^>v^<v^^v^v<v^>><<v^^^v^v^^^v^vvvv<^^v^^<>vv>^><>>v<^^<v<^>>^^<<v>^>^^><v^<^^>^><^^>^vv>>>vv>vvv>>v<v>><^<<<vvv^><<>^>^><>^v<<vvv<^^v^<<<v>v>v><<^v><v>>^^<>>>vv>^^<v<^>v^^>>>vvv<>v<<<vv>>^v><^<<^<^v`,
	`#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`,
}
