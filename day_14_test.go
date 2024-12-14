package main

import (
	"fmt"
	"go-aoc-2025/utils"
	"regexp"
	"strings"
)

func ExampleDayN() {
	fmt.Println("part 1:", part1(11, 7, 100, inputs[0]))
	fmt.Println("part 1:", part1(101, 103, 100, inputs[1]))

	// For part 2, we just check if none of the robots are on the same position.
	if part2(101, 103, 8258, inputs[1]) {
		fmt.Println("part 2:", 8258)
	}

	// Output:
	// part 1: 12
	// part 1: 228410028
	// .............................................#.........................................................
	// ......#..............................................................................#.................
	// .......................................................................................................
	// ..................................................#....................................#...............
	// .................................#............................................................#........
	// ........................................................................................#..............
	// ...............#......................................#............#...................................
	// .............................#..............#..........................................................
	// ......#................................................................................................
	// ..................................................................#....................................
	// ..........................#.................................#..........................................
	// ................................................#......................................................
	// .........................#.............................................................................
	// .....................................................##..........................................#.#...
	// .........................................................#.................#...........................
	// .........................#....#..................#.....................................................
	// ............................#...................................................#......................
	// .......................................................................................................
	// ...............................#..........................................#....#.......................
	// .....#.............................................................#...................................
	// .......................................................................................................
	// .......................................................................................................
	// .............................#..............................................#............#.............
	// ........................#.................................................#............................
	// ..................#.............................#..............#.......................................
	// .............#..........................#..............................................................
	// ...........................................................#..................................#.....#..
	// ......#......#......................................................................................#..
	// ...........................................#.........#.................................................
	// ................................................#...........#..........................................
	// #...................#.....................................................#...#........................
	// ...............................................................................#.......................
	// ............#.............................................#............................................
	// ....................#.......................................#..........................................
	// ..............#..............................#.....#...................................................
	// .........................#.............................................................................
	// ..................................#....................................................................
	// ......................................................................#................#...............
	// .............................#################################.........................................
	// .............................#...............................#...........#.............................
	// ...........#.................#...............................#.........................................
	// .............#...............#...............................#.........................................
	// .............................#...............................#.........................................
	// .............................#.......................#.......#.........................................
	// .............................#......................##.......#..#......................................
	// .............................#..................#..###.......#..................................#......
	// .............................#.................##.####.......#.......................#...............#.
	// .............................#.............#..########.......#.........................................
	// .............................#............##.#########.......#.........................................
	// .#...........................#........#..#############.......#.........................................
	// .............................#.......##.##############.......#.........................................
	// .............................#......##################.......#.....................................#...
	// .............................#.....######################....#.........................................
	// .............#...............#....#######################....#.........................................
	// ......................#......#.....######################....#...#.....................................
	// .............................#......##################.......#.......................#............#....
	// .............................#.......##.##############.......#..............................#..........
	// .............................#........#..#############.......#.........................................
	// .............................#............##.#########.......#.........................................
	// .............................#.............#..########.......#.....#...................................
	// .............................#.................##.####.......#............................#............
	// .............................#..................#..###.......#..........#..............................
	// .................#...........#......................##.......#....#....................................
	// ............#................#.......................#.......#................................#........
	// .#...........................#...............................#..........................#..............
	// .............................#...............................#............#............................
	// #.....................#......#...............................#.........................................
	// .............................#...............................#.........................................
	// .............................#################################.........................................
	// ...............................................................................#............#..........
	// .........................#.....#...............#.................................#.....................
	// ..........................#.......#...............................................................#....
	// ..#.....................................................................................#...#..........
	// ........#..............................................................................#...............
	// .....................................#.................................................................
	// .................................................................................#.....................
	// .......#...............................................................#................#..............
	// ...............................................................................................#.......
	// ..................................#................................#...................................
	// .......................................................................................................
	// .......#...............................................................................#...............
	// ..............................#........................................................................
	// .......................................................................................................
	// ..........................................................#............................................
	// .....................................#..#..................................................#...........
	// ......................................................................#.......................#........
	// ...#...................................................................................................
	// .......................................................................................................
	// .......................................................................................................
	// ...#...............................#...................................................................
	// ......................................#.............................................................#..
	// .......................................................................................................
	// .......................................................................................................
	// .......................................................#..#............................................
	// .........................................................#.....................#.......................
	// .....#....................................#............................................................
	// ........................................#.............................................#................
	// .............................................................................#.#.......................
	// ...........#..................................................#.............#..........................
	// .......................................................................................................
	// .......................................................................................................
	// part 2: 8258
}

var re = regexp.MustCompile(`(-?\d+)`)

func digits(input string) []int {
	matches := re.FindAllStringSubmatch(input, -1)
	res := make([]int, len(matches))
	for i, m := range matches {
		res[i] = utils.ToInt(m[1])
	}
	return res
}

type Robot struct {
	pos complex128
	vel complex128
}

func part1(x, y int, seconds int, input string) int {
	var robots []Robot
	rows := strings.Split(strings.TrimSpace(input), "\n")
	for _, row := range rows {
		if row == "" {
			continue
		}
		nums := digits(row)
		robot := Robot{
			pos: complex(float64(nums[0]), float64(nums[1])),
			vel: complex(float64(nums[2]), float64(nums[3])),
		}
		robots = append(robots, robot)
	}

	for range seconds {
		for i := range robots {
			robots[i].pos += robots[i].vel
			x := (int(real(robots[i].pos)) + x) % x
			y := (int(imag(robots[i].pos)) + y) % y
			robots[i].pos = complex(float64(x), float64(y))
		}
	}
	grid := make(map[complex128]int)
	for _, r := range robots {
		grid[r.pos]++
	}

	midX := x / 2
	midY := y / 2

	var a, b, c, d int
	for i := range midX {
		for j := range midY {
			if n, ok := grid[complex(float64(i), float64(j))]; ok {
				a += n
			}
		}
	}

	for i := midX + 1; i < x; i++ {
		for j := range midY {
			if n, ok := grid[complex(float64(i), float64(j))]; ok {
				b += n
			}
		}
	}

	for i := range midX {
		for j := midY + 1; j < y; j++ {
			if n, ok := grid[complex(float64(i), float64(j))]; ok {
				c += n
			}
		}
	}

	for i := midX + 1; i < x; i++ {
		for j := midY + 1; j < y; j++ {
			if n, ok := grid[complex(float64(i), float64(j))]; ok {
				d += n
			}
		}
	}

	return a * b * c * d
}

func part2(x, y int, seconds int, input string) bool {
	var robots []Robot
	rows := strings.Split(strings.TrimSpace(input), "\n")
	for _, row := range rows {
		if row == "" {
			continue
		}
		nums := digits(row)
		robot := Robot{
			pos: complex(float64(nums[0]), float64(nums[1])),
			vel: complex(float64(nums[2]), float64(nums[3])),
		}
		robots = append(robots, robot)
	}

	for range seconds {
		for i := range robots {
			robots[i].pos += robots[i].vel
			x := (int(real(robots[i].pos)) + x) % x
			y := (int(imag(robots[i].pos)) + y) % y
			robots[i].pos = complex(float64(x), float64(y))
		}
	}
	grid := make(map[complex128]int)
	for _, r := range robots {
		grid[r.pos]++
		if grid[r.pos] > 1 {
			return false
		}
	}

	var sb strings.Builder
	for i := 0; i < x; i++ {
		sb.Reset()
		for j := 0; j < y; j++ {
			if _, ok := grid[complex(float64(i), float64(j))]; ok {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		fmt.Println(sb.String())
	}
	return true
}

var inputs = []string{
	`p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`,
	`p=7,85 v=-65,-36
p=4,17 v=-37,-76
p=60,71 v=-8,-19
p=82,99 v=20,-59
p=56,11 v=50,-21
p=20,25 v=66,12
p=57,97 v=-67,-20
p=89,37 v=10,52
p=32,51 v=46,56
p=72,40 v=26,-80
p=65,23 v=1,91
p=74,9 v=-37,46
p=44,72 v=-89,96
p=49,34 v=67,1
p=16,83 v=11,-47
p=56,90 v=42,78
p=42,37 v=71,-34
p=64,16 v=-33,-50
p=70,21 v=-42,-90
p=31,81 v=20,-94
p=32,8 v=-5,-38
p=100,10 v=44,64
p=24,93 v=-90,63
p=45,20 v=-3,-31
p=13,18 v=-65,30
p=61,26 v=59,-79
p=30,67 v=-25,39
p=57,13 v=13,25
p=56,27 v=-38,-55
p=68,9 v=-92,-32
p=81,34 v=-54,97
p=65,40 v=-4,87
p=7,87 v=-10,-2
p=48,78 v=-8,-59
p=52,62 v=-17,-98
p=75,6 v=-79,-32
p=6,53 v=-82,-44
p=45,28 v=37,90
p=53,10 v=88,-77
p=21,35 v=-62,-36
p=25,61 v=12,44
p=23,96 v=54,1
p=76,43 v=-6,-35
p=20,47 v=-2,22
p=82,15 v=-83,47
p=45,71 v=-19,-11
p=87,35 v=-45,-68
p=81,101 v=14,-26
p=58,42 v=17,58
p=71,78 v=-9,-22
p=30,97 v=20,43
p=49,91 v=-13,-54
p=27,86 v=45,-42
p=26,0 v=45,-32
p=22,92 v=41,-48
p=95,11 v=22,-52
p=79,47 v=81,17
p=74,74 v=-46,5
p=65,100 v=21,-73
p=4,72 v=-6,73
p=24,43 v=-18,-51
p=56,58 v=53,-39
p=62,27 v=-12,70
p=43,77 v=-81,33
p=11,15 v=-6,29
p=60,7 v=59,71
p=86,79 v=94,55
p=57,29 v=-7,-6
p=62,39 v=-33,14
p=28,5 v=-77,26
p=5,49 v=-67,-5
p=10,33 v=-48,1
p=69,4 v=89,82
p=98,95 v=33,23
p=4,27 v=48,-71
p=3,24 v=-78,-38
p=54,2 v=68,20
p=69,14 v=30,59
p=39,29 v=92,-16
p=45,77 v=-55,-53
p=98,70 v=82,27
p=8,76 v=-4,-93
p=26,39 v=36,-6
p=96,88 v=-95,49
p=24,48 v=58,-97
p=71,67 v=-73,-16
p=50,89 v=67,9
p=93,3 v=-49,-89
p=4,24 v=-94,-44
p=28,16 v=45,-50
p=64,36 v=72,81
p=0,40 v=-10,-46
p=89,54 v=10,-30
p=16,99 v=78,-22
p=95,32 v=69,35
p=82,72 v=-42,-85
p=60,96 v=-88,-42
p=49,52 v=-76,-58
p=40,8 v=-30,-90
p=28,32 v=70,-27
p=93,14 v=12,-85
p=80,26 v=-7,-73
p=85,74 v=-49,-70
p=75,66 v=5,73
p=76,30 v=39,-85
p=35,32 v=46,17
p=89,36 v=72,-23
p=89,33 v=-61,52
p=8,4 v=74,71
p=92,102 v=-82,95
p=66,50 v=35,97
p=78,63 v=-61,-60
p=42,98 v=-39,39
p=39,83 v=-64,61
p=51,30 v=30,75
p=44,29 v=25,47
p=87,89 v=24,71
p=89,12 v=-87,-44
p=32,60 v=79,-24
p=96,21 v=31,-67
p=27,24 v=58,24
p=1,84 v=-23,-7
p=18,67 v=-35,79
p=35,51 v=71,97
p=54,51 v=-42,57
p=0,52 v=-87,80
p=20,79 v=-33,-85
p=89,101 v=-28,-95
p=70,22 v=-10,-56
p=57,22 v=29,-22
p=87,94 v=-21,35
p=80,15 v=95,32
p=70,89 v=-18,30
p=15,99 v=95,83
p=60,56 v=55,6
p=92,66 v=-91,-99
p=53,82 v=89,-65
p=20,100 v=24,-60
p=56,58 v=-67,97
p=7,79 v=2,-13
p=29,23 v=-47,30
p=94,26 v=94,87
p=42,22 v=-46,86
p=46,44 v=78,75
p=83,6 v=43,-32
p=84,2 v=-44,-89
p=22,82 v=-1,-39
p=49,80 v=-4,-13
p=77,93 v=-83,43
p=35,82 v=-89,-71
p=24,7 v=24,-32
p=98,64 v=13,-32
p=40,75 v=59,73
p=41,30 v=-5,-56
p=39,102 v=-73,25
p=31,40 v=58,-57
p=95,96 v=-58,18
p=51,64 v=-36,3
p=74,3 v=-63,7
p=16,67 v=66,-24
p=69,54 v=-96,91
p=17,62 v=-72,-64
p=9,57 v=3,-1
p=82,86 v=63,-81
p=37,24 v=21,-50
p=35,76 v=-29,-49
p=41,4 v=69,23
p=47,71 v=-88,32
p=97,51 v=65,-12
p=9,29 v=36,-33
p=19,101 v=66,20
p=18,19 v=-34,48
p=45,99 v=73,15
p=24,25 v=92,35
p=43,83 v=97,66
p=22,16 v=-73,-38
p=31,82 v=-5,60
p=67,3 v=78,-29
p=45,69 v=59,22
p=99,101 v=-82,66
p=96,88 v=-6,15
p=27,62 v=-56,51
p=80,78 v=-62,-99
p=76,88 v=-67,-26
p=74,15 v=64,53
p=77,25 v=37,-83
p=55,47 v=-30,-57
p=0,54 v=-32,-30
p=46,43 v=58,-12
p=66,66 v=-93,-97
p=5,14 v=-23,-10
p=60,3 v=-89,-30
p=69,38 v=9,-28
p=69,12 v=-25,-62
p=69,98 v=-8,83
p=59,32 v=-13,35
p=15,12 v=41,-62
p=32,71 v=-25,7
p=93,100 v=-49,-3
p=63,25 v=-20,-67
p=36,28 v=37,-56
p=44,54 v=46,-98
p=22,99 v=24,-95
p=33,8 v=-73,-78
p=10,68 v=74,10
p=84,92 v=74,70
p=76,9 v=34,-21
p=17,3 v=12,-51
p=76,1 v=77,72
p=89,76 v=-73,-23
p=74,94 v=64,-3
p=18,49 v=-14,68
p=40,90 v=-9,-48
p=60,47 v=47,58
p=9,37 v=-27,-28
p=66,84 v=72,66
p=11,57 v=62,-17
p=95,64 v=15,-69
p=55,8 v=46,94
p=8,92 v=-99,-24
p=63,36 v=-97,-29
p=38,24 v=46,-99
p=2,39 v=-82,63
p=66,45 v=5,29
p=34,9 v=64,-9
p=24,35 v=-56,81
p=24,90 v=-90,38
p=94,22 v=-86,78
p=98,55 v=15,69
p=23,40 v=-94,92
p=51,10 v=-11,-85
p=13,54 v=11,-30
p=6,37 v=-61,-40
p=0,71 v=-74,-71
p=79,32 v=-75,24
p=15,10 v=-35,-83
p=100,96 v=-57,32
p=6,65 v=-45,54
p=20,73 v=-56,10
p=4,102 v=41,63
p=17,19 v=-91,68
p=73,58 v=-79,91
p=96,57 v=-40,74
p=20,22 v=-9,-27
p=56,73 v=32,84
p=22,11 v=-56,-61
p=21,19 v=54,-68
p=81,24 v=-79,93
p=52,25 v=14,71
p=37,15 v=96,-45
p=13,31 v=41,-67
p=58,90 v=72,-94
p=90,4 v=10,-3
p=93,86 v=81,20
p=65,85 v=-8,-19
p=93,23 v=31,-50
p=26,87 v=42,21
p=73,1 v=31,-84
p=94,52 v=-40,-51
p=52,79 v=10,-46
p=52,81 v=63,32
p=85,18 v=-78,-21
p=40,1 v=12,-95
p=82,54 v=-36,-22
p=77,70 v=-71,-21
p=1,32 v=-66,-63
p=45,3 v=-68,-66
p=78,99 v=35,-84
p=13,18 v=-57,87
p=89,69 v=94,90
p=5,35 v=-19,-34
p=98,14 v=53,-84
p=52,78 v=38,10
p=84,47 v=-83,58
p=75,86 v=-79,-42
p=71,0 v=5,-71
p=30,87 v=62,-71
p=43,57 v=4,51
p=19,18 v=83,75
p=26,86 v=-77,-34
p=45,4 v=79,59
p=61,3 v=-7,86
p=96,77 v=-45,-37
p=86,26 v=-65,83
p=26,39 v=-77,35
p=9,10 v=-23,25
p=40,83 v=-8,48
p=44,10 v=-82,-66
p=20,84 v=-51,-93
p=42,67 v=71,78
p=40,21 v=-89,76
p=87,29 v=-63,75
p=90,82 v=32,56
p=30,18 v=79,-78
p=79,72 v=-62,-82
p=29,85 v=-18,26
p=49,48 v=-97,-86
p=29,27 v=-81,98
p=45,68 v=-72,22
p=67,34 v=-38,77
p=12,27 v=91,58
p=67,33 v=76,-79
p=81,38 v=-95,-97
p=70,3 v=-42,-20
p=45,42 v=-34,-17
p=98,70 v=27,73
p=26,33 v=24,12
p=71,21 v=-75,-84
p=16,84 v=7,95
p=86,93 v=-71,-69
p=28,72 v=-42,-90
p=87,90 v=-96,60
p=3,74 v=78,5
p=21,65 v=83,96
p=6,9 v=-49,45
p=43,11 v=-40,-42
p=77,14 v=21,43
p=44,61 v=-18,-4
p=80,19 v=-96,-32
p=13,84 v=20,50
p=23,2 v=70,-66
p=35,68 v=96,-47
p=51,34 v=59,-50
p=2,58 v=73,-12
p=30,26 v=-26,-22
p=79,57 v=70,50
p=33,46 v=-5,-63
p=99,91 v=-28,67
p=26,35 v=71,-78
p=80,93 v=-45,32
p=14,39 v=-91,-77
p=38,93 v=46,-15
p=20,93 v=-98,-14
p=98,12 v=-52,92
p=21,42 v=24,-91
p=14,9 v=70,-67
p=100,14 v=2,76
p=49,76 v=4,27
p=20,29 v=32,70
p=61,24 v=-88,24
p=45,85 v=-46,-65
p=59,28 v=9,30
p=2,62 v=-69,-63
p=77,17 v=-79,25
p=7,37 v=2,-92
p=93,44 v=91,-67
p=96,15 v=62,-76
p=62,29 v=8,-94
p=34,13 v=49,-21
p=93,60 v=-32,-93
p=92,42 v=78,68
p=9,72 v=91,27
p=32,86 v=58,72
p=48,19 v=-97,76
p=4,73 v=7,68
p=84,26 v=98,64
p=72,9 v=-79,-79
p=12,21 v=91,-4
p=37,74 v=16,-7
p=2,37 v=-61,-63
p=69,83 v=-27,-89
p=86,64 v=-32,68
p=9,16 v=-98,81
p=39,8 v=-4,-34
p=41,46 v=43,-8
p=16,9 v=-19,99
p=82,4 v=-45,48
p=81,95 v=18,54
p=11,75 v=86,62
p=60,86 v=25,14
p=63,54 v=89,-30
p=43,68 v=50,-18
p=51,17 v=-42,-16
p=85,70 v=52,16
p=20,91 v=-27,-36
p=78,36 v=-79,12
p=83,80 v=-3,-13
p=31,10 v=75,-78
p=42,2 v=41,-94
p=10,1 v=64,57
p=59,33 v=18,14
p=10,60 v=63,-69
p=92,47 v=65,-17
p=52,55 v=4,28
p=33,59 v=-73,33
p=46,70 v=33,-8
p=48,33 v=77,14
p=80,51 v=73,63
p=5,40 v=-6,-91
p=5,45 v=43,-92
p=66,24 v=-50,-33
p=10,36 v=-97,28
p=0,84 v=14,38
p=14,88 v=-40,-58
p=90,23 v=-43,-73
p=16,75 v=-73,-59
p=36,16 v=39,-43
p=45,51 v=79,-87
p=67,96 v=55,34
p=55,12 v=-42,-95
p=84,13 v=39,36
p=95,52 v=11,92
p=68,82 v=99,-19
p=36,102 v=75,-38
p=39,45 v=-85,-74
p=98,4 v=24,74
p=2,40 v=40,52
p=0,38 v=-78,53
p=4,34 v=27,-50
p=43,89 v=-93,55
p=21,78 v=78,-1
p=64,63 v=-25,21
p=12,63 v=-60,-46
p=83,3 v=-11,32
p=97,71 v=10,16
p=39,83 v=-20,10
p=9,27 v=-74,-67
p=84,25 v=75,-31
p=30,100 v=-64,8
p=37,96 v=21,32
p=1,66 v=78,62
p=93,75 v=-99,-58
p=75,99 v=-66,-84
p=57,49 v=87,28
p=53,28 v=73,-89
p=4,76 v=-44,96
p=38,95 v=-71,-66
p=48,73 v=-76,-59
p=3,40 v=-61,-23
p=98,91 v=-99,32
p=47,81 v=59,-19
p=40,95 v=71,-94
p=93,63 v=81,28
p=35,36 v=37,46
p=98,73 v=63,90
p=28,3 v=94,-15
p=43,96 v=20,38
p=42,70 v=-72,27
p=66,26 v=-33,-22
p=59,4 v=34,-3
p=68,91 v=67,-86
p=87,71 v=5,-41
p=43,45 v=62,79
p=79,86 v=-77,-96
p=39,22 v=54,-79
p=99,67 v=14,34
p=62,29 v=93,-62
p=85,99 v=35,-94
p=15,57 v=53,-46
p=35,43 v=-26,-58
p=23,66 v=66,-1
p=76,47 v=34,34
p=84,54 v=14,54
p=29,92 v=60,97
p=72,12 v=-84,36
p=6,65 v=-48,-81
p=43,24 v=-89,1
p=95,25 v=24,-78
p=34,86 v=7,67
p=90,5 v=-11,99
p=63,4 v=-60,99
p=17,45 v=-65,92
p=59,75 v=-37,-8
p=63,59 v=81,85
p=12,85 v=34,-83
p=45,48 v=84,26
p=29,5 v=12,70
p=57,94 v=-97,93
p=29,50 v=22,-57
p=57,3 v=50,-58
p=90,33 v=90,52
p=78,82 v=36,60
p=4,64 v=44,17
p=43,34 v=63,98
p=87,69 v=-13,-62
p=10,66 v=-15,38
p=21,70 v=7,84
p=65,52 v=-71,-80
p=3,76 v=-2,10
p=42,25 v=-19,59
p=20,100 v=-32,-84
p=44,0 v=-34,26
p=4,26 v=2,-17
p=84,17 v=-32,2
p=32,62 v=-26,73
p=21,66 v=-94,-44
p=51,54 v=-59,85
p=18,91 v=53,78
p=39,55 v=80,62
p=53,48 v=38,-86
p=79,84 v=-66,-2
p=33,24 v=4,-51
p=54,76 v=-17,27
p=15,55 v=-86,-29
p=26,64 v=50,-70
p=87,53 v=17,6
p=90,71 v=14,-70
p=42,74 v=-72,15
p=70,100 v=68,-88
p=25,98 v=-1,49`,
}
