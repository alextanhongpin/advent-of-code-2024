package main

import (
	"fmt"
	"go-aoc-2025/utils"
)

func ExampleDayN() {
	fmt.Println("part 1:", part1(inputs[0]))
	fmt.Println("part 1:", part1(inputs[1]))
	fmt.Println()
	fmt.Println("part 2:", part2(inputs[0]))
	fmt.Println("part 2:", part2(inputs[1]))

	// Output:
	// part 1: 18
	// part 1: 2358
	//
	// part 2: 9
	// part 2: 1737
}

type Point struct {
	X, Y int
}

// Find all "X", then check if the surrounding characters form "MAS".
func part1(input string) int {
	grid := make(map[Point]rune)
	var xs []Point
	for x, row := range utils.Scan(input) {
		for y, r := range []rune(row) {
			grid[Point{x, y}] = r
			if r == 'X' {
				xs = append(xs, Point{x, y})
			}
		}
	}

	points := [][]Point{
		{{1, 0}, {2, 0}, {3, 0}},       // Right
		{{-1, 0}, {-2, 0}, {-3, 0}},    // Left
		{{0, -1}, {0, -2}, {0, -3}},    // Up
		{{0, 1}, {0, 2}, {0, 3}},       // Down
		{{1, -1}, {2, -2}, {3, -3}},    // Diagonal up-right
		{{-1, -1}, {-2, -2}, {-3, -3}}, // Diagonal up-left
		{{1, 1}, {2, 2}, {3, 3}},       // Diagonal down-right
		{{-1, 1}, {-2, 2}, {-3, 3}},    // Diagonal down-left
	}

	var total int
	for _, p := range xs {
		for _, pt := range points {
			var str string
			for _, d := range pt {
				str += string(grid[Point{p.X + d.X, p.Y + d.Y}])
			}
			if str == "MAS" {
				total++
			}
		}
	}

	return total
}

func part2(input string) int {
	grid := make(map[Point]rune)
	var xs []Point
	for x, row := range utils.Scan(input) {
		for y, r := range []rune(row) {
			grid[Point{x, y}] = r
			if r == 'A' {
				xs = append(xs, Point{x, y})
			}
		}
	}

	points := [][]Point{
		{{-1, 1}, {1, -1}}, // Diagonal up-right
		{{-1, -1}, {1, 1}}, // Diagonal down-right
	}

	var total int
	for _, p := range xs {
		x, y := p.X, p.Y
		ok := true
		for _, pt := range points {
			var str string
			for _, d := range pt {
				str += string(grid[Point{x + d.X, y + d.Y}])
			}
			if !(str == "SM" || str == "MS") {
				ok = false
			}
		}
		if ok {
			total++
		}
	}

	return total
}

var inputs = []string{
	`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
	`XASXMAXXMSXXSMMSXMMSMXSMXMSSMSSSMMSMAMXMXSMMMMAXAMXSASXSSMMSSMXAMXMSAMXMMXAXXXSAMXXXXXMMXSXMXXSMASAMXMXAXXMASAMXXXMAMMMSXSXMXMMMSAASXSMSSMMS
XASAMSSMAMMMAMAAAMASMAMAAXMASAAASAAMAMAMAAAAASMSSSSMAXAAAXASASMMMMMMXMASXMAMXXXXSMSMSXSAAXAMXMASMXMMSMSAXSXMMSXSMMAAMXAMAAMMAAAAXMXSAAAAXAAM
MXSXMAAMXAAAXMMSSMAXMASXSSSSMMSMMSXSASMSMSSMMMAAAXXXXMMMMMXXAMMAAAASMMAAAXMASMXXXAAASAMMXSASASMMMAMXSASASASMAXXMAMSSSMSSMMAXSMMSXMMMMMMMXAMX
SAMMMSSMMXXSXXXAAMSMSASXAAXMXXAAXXXMAXAAAXXMASMMMMMMXMXXXXSMXMSSMSASASMMXMMXMAAMASMSMSMAASXSAAAASAXAMAMXAAXMASMSSMXAMMMAMXSAMAMAASXMMSSMSAMX
MASAAAAMSMMMMMMSMMMAMAXMMMMMXMSSMMXMSMMMMMAMXASAXAASXMMASAAXAXAXMXMXMXAASXMAMSXSAMMXMMMMMSAMXMSMSMSMMAMSMXMMASAMXAMMSASAMMXASAMSXMAMXAAXSAMA
SSMXMSMAAMAAXXAMXASMMSMAMXSXSAMXXMAAAAXSXSAMXMASMSXMASXAMSAMXMMMSASMXMXMMAMSMXAMASXMXAASXMMMMXMMXAAXXAMXXAAMAMAMXXXASXSASXMMSAMMASXMMMMMMMMS
XXXXMAMMMSSSMXSXSASXAXSXMASAMXSMXMASMSMSXMASXXMMXMXSAMMXMAXAMAXXSXSAASMASXMXAMXMMMAASMXXAMXMMAXMXMSMSAMSSMSMMSSSSMMMSXMXMASXSMMXAXAXSSSMXAAX
MMXSSXSMXAMAMAAAMSXMXMMXMASXMMSAMXXAAMXSASAMMXXSXMXMASXSSSMXSSMXXMXMMSAXMASMMMSMASAMXMASXSAMSSMSMAAAMAMAXAMAMAMXAXAXSAMXMXAXMMSMMSXMAAAMSAMX
AXMXMASMMASAMXSXMMXSAMAAMXMMMAXMMSMMXMASAMXSXMXMAAMAMSAMAAAMSAMMSXMXSXMXAXMASAMXAXMASXMAMMAMAAASMMMSSSMSSMXSAMXSMSMXXAMXSSMASAMAMMASMMMASMMM
XMAMSXSASASASMXXMAMMAASXSAAAMMXAAXXAMMMMSMSMXMAXMXMAMXAMSMMMSAMASAMXSAMMMMXMMAMMSXXAMAXXAMSMMSMMAXXXMAAAAMAMAXXAMXXSMSAXAAXAMMMAMSXMAXXMSASX
MSSMXASAMXSMMMAXASMSMMAXMAMSASXMAXMMSAXSASAAASMXMASMSSMMMAAXSAMMSAMAMAMMSAMSSSMSAMMSSSMSMXMAXXASXMMMMMMMSMASASXSXMASAXXMSSMMXMXMMMMSAMMXSAMX
MAMMMMMMMMMASAMMMXAAXXMASAAXASMMSXMAXMXSMSMSMSXAMMXXAMXASXMXXAMASAMSSXAAMAMAAMXMASAAAAAAXSSSMMAMASASXXXXMMXSASAXXAAMAMAXXMXSASASAMASAMXAMAMX
MASXAXAMXXSAMMAAXMSMMXMAMXXMMMAAMMMSXMMSXMXAAXXMSAMMSSMMSMMMMAMASAMAXMASXMMMSXAXXMMMSMMMAXAAXXSSMSASMMSAMXXMAMAMXSSMSMSMAMAXXSXSASXMSSMXSAMM
SXMXMSSSMMMASMSMSMAAMASXMMXSASMMXXAMXXASASMMXMAMAAXAMAMAXAASXMMXMMSMXXAXMSXMAXXSSSXAXAAXMMMMMMXAXMMMXAAMAMMMXMAMAMXAMAXAAMASAMASAMXMAMMXSASA
MMMXMXXAXAMAMMAAAMSMMAXAASASAXXXSMSSSMAMXMAMXXASXSMMSMMXSMMXAXSXMAXAMMMXXMAMSAMXAAMSSSMMXAAAAXSAMXXAASMXAASMASXMSSMAMAMMXSAAAMAMAMSMAXSAXAMM
SASASMXMMSXSXSMSMMXSMXXXMMASAXSXMAXAAMMMSSMMMXASXMXMAAXMXMMMXMSAMXSXSASMASAMMSSMMMMAAAXMSSSSMXXAMXMAXXXMSMSAAMAMAMAMMMMXXMMSSMMXAMXXAXMASXSM
MASMMAMSAMMMMAMMXSAMXSMSAMXMMMAAMAMXMAMAMAAAAMSMXSASMXMAAAXSAMXAMAAMSAMSAMXSAMMASAMXSSMMMAMMXXMSMMXMSSMMXAMMXSAMMSAAAAMXAAMAXASXSMSMMSMAMAAX
MMMMMSMMAMXAAMMAAMMXAMMAAAXAXSSXMAMXMAMASMMMXMAAXSASAMSSMSASASXMMMSXMAMXXMMMMXMAMAXAMXAXXSAAASAAAXXXAAMAMAMAAXAMXSASMXSASXMSSMMAMASAMSMMMSXM
SAAMAMAMAMSSSXMMXSSMSSSSSMSSXAAMSXSXSASASAMAMSMSMMMMXAAXAXMSAMASXSMMXSMSMSXSSSMMXMMXSSSMAXMXSXMMAMXMSXMASMMMXSAMAXAXMASMMMMASAMXMAMAMXAXAMXA
SSSMASXMMXMAMXMMAMXAAAAMMAAXAASMSXAAXAMASMMAXAMAAMSAMMMMSMMMXSAMXMAAMXAMAAXMAAXAAMMXMAMMAMMXXAXMAXSAAXSXSXXAASAMAMSMMMXXMXMMSMMSMXSXMSXMXMMS
XAMMAMAMXMMSMMSMAXMMMSMMSMMXMAMAXXMSMSMMMXXMMSSSSMMXMAAAMAXXMMMSXMMMSMAMSMMMSMMSXXASMAMMASMMMAMSSMASMXSAXXMMMXXXMXMXXAMMSASMMMAXMXMXAMXSAMXA
MMXMXSAMMXAXAXMSMSXXAAAAAMSMMXMAMXXAAMASMSXSAAXXAAMAMSSSSSMSXAAXMSMXAMXMAMAMXAMASXXMMMSMMMAMMSMMAMMAXAMAMSXSSMSSMMMASAXASMXAAMXSXAMXMAMXASMS
AXMSMSXXAMXSMMMMXAMMMSMMMSAASAMMMSSMSSXMAAAMXSMXSSMAMMAMMAASXMMSAAXSXSSSSSSMMAMASXAXXASXSSMXAMAXAMXAMMMSMMASXAAXAASAMXMMXXSSMMSMMXSAMSMSMMMA
SMAAAMASXMMSMMXMMMMXAAXASMMMMASXAAAXAXAMSMMMAAAXXAMXXMAMSMMMAXSMMMMMSAAAAAXAXSMMSMMMSAMAXXXMSSMSSSMMSAAXAMSMMMMXSMSASXXSXMMXAXMAAXXASMAAXASM
XMSMSMAMAAAMMSASAAAMSSSXSAMXSAMMMSSMMSMMAMAXSXSXSAMAXSAMXMMSSMXAAXASMMMMMSMMXMAXMASASAMSMMSAMXXMXAAASMSSMMASAAXAXASAMAMMASMXSMSMMSSMMMSMSMSX
MXXMAXASMMSSXMASMMMAMXMAMXMAMAMXAAXMMAAXMSAXMSMASMMSASAXSMAAXAMXMSXXSMXSAAMMAMAMXMMXSAMXMAXMASMSMMMMSMXAMAMXSMSXMAMXMXASMMSAAAMXAXXASAMXAASA
MMSXMMMXMAAXMMXMMMSSMMMAMSMSMSMMSSSSSSMMXSXSMAMAMXAXASMMSMMSSSMSMAXMXSAMXSASASXSSXMASMMSXMMSMMAMXSAXXMXMAXXMAXSAMAMMAMXXSAMXMMMMMSSMMAXSMMMS
MASMSASXMMMSSMAAAAMAAXSSMSAAAAXXAAMAAAMSXMASMXMASMMMMMAAMAMMXXMASMAMAMMSMXAMAMAAXXSASAXMASMASMMMASXXAXMSMSSSMAXMXSMSAXMAMXMSMSXXMAMASXMASAAX
MXXAXAXMXAAMASXXMSSSMMMMAXSMSMSAMXMMMMMSAMAMAMAMXXAASAMXXAMMSSSMSXAMAMXAAMXMSMMMMXMAMXMSXASAMXSMAMMSSMXAAMAAXAXSAMXSASMSMAXAAMSAXASMMASAMMSM
MAMSMSMSXMXSAMXASMMMAMAMMMMXXXAMXAMAMAXSAMASAMSSSSSSSMAMSMSAAAAAMXMSASXMSMSAXXXXMAMXXSMXSXMMSMSMSXMAMASMSMSMMSAMXSXMAMAMSSXMXMXXMAXXXAMASAXM
MASAAMAAXSAMXSSXMASMMMAMXAMAXMMMSMXAMSXSAMAXAMXAAAXMXMAXAMXMMMMMMSXXMAXMMMSASXSMSMMMMSAMAXMAAAXAAMMMXMXXAMXMAXXMASAMMMXMAXXMSSMMASXMMSSXMASX
MXSMMMMMXMASAXAMSAMAASMMSASXSAXAAAMSXMASMMSMSSMMMMMMXSSSXSMMSSMXXSAMXMSMAAMAMAXAAMSSMXAMMSASMSMSMAASMSMMSSSMMSAMXXAMXSSMAMXSASAMXXAXAAMAMAXA
XMXMAXXAXSAMMXSAMAMXMMAAXAAASMMSXSXXAMAMXXMAAAMSAMASMAXAMMMXAAMSAXSAMXSMMSSSMMMSMAAAMSMMXMAXAXAXMSMMAAAXMAXAXSXXAMSMMAAMMMXXAXMMAMSMMMSAMASM
MAASXXMMXMAXSXMXMASMMMMMMXMAMMAXXMAMSMXMAAMMMSXMASAXMMMMMAAMSSMSAXXXMXMAXAAMXAAMXMMSMAAASMMMSMAMAXAMSSSMMMMSMMMSMMXAMSXMASMMSMAMXSXXXXSXXAXX
AMXMSXAMMMSMXMSMSASAXAXXMMMMSMMSSMAMMMMSMSXAAMMSMMXSAAASXXXXAAXMMMSMMASAMMSMMMMXAMAMXMXMMAAAMMXMAMXMXAMXMXMAAAAXXASXMAMSXSAAAMAMXAXXSAMXMXSM
XXAAMSSMAAMXAAAAMXSMMSAMXAAXAAASASXMAAMAAAXMMSAXAAAXXMSSMAXMSSMSAMXASMSMSMMMSMMMSMAXAXSASXMSSXSMMSXSXMSAMSSSSMMXMXSAXXAMAMMSSSMMSMSMSXSASAAX
MSMSMAMXMSSSXSMSSMMXMMXMSSSSMSMSAMXSSXSMSMMXAMXSMMSXXXAXMSMMXMASMSSXMAXMAMAAAAMAXMMMAXMAMAAAMXMAXXMMAXSASMMMXMXSAAXMMMSMAMXAXAMMAAAASAAASXSM
XAAXMMSXMAXXAXMAMAMMMXMXAXMXAMXMAMXAMXSXXXSAMXMXMAMMAMMMXAMMAMAMXXXASXMSMSMXSXMXSAXMSMMASMMMSMSSMMMSXMSSMASMMAAMMSMMSAAMSSMASMMSMSMSMSMAMAXX
MMMMAAMMXSMMSMMASAMAAASMMSSMXMMMAMXAXXSASAMXMASASASAAMAAXMMSASMXSAMMMAAXMAMXMXSXMAMSAMSAXAXMAXAMSAMMAMXXSAMAMMXSAXXAMXXSAMAAMXAXAMMXXXXAMMMS
MASMMMMXAXMAMASASXSMSXSAMAMXMMMSSMSMMXMAMMMSSMSAMXSXMMXSSXASXSMMMMSSSMMXXAMASAMXMAMMAXMXSSMMSMMSMSXMAMMMMXSXMAMMMXMSXSXMMSSMMXSSXMASMSMMMXAX
MAXAMXSMAXMAMMAMXASAMXSXMAMSXSAAAMAAXAMXMXAAAMMXMMMAMSAMXMMMAXAXSAMMAXAMSSXSAMXASXSMXMXAXAAAAAXXAAMASXMAMXSXMASAMSAXASXAXAXAXXMXMMXMASASMMMS
MSSMMAMMMMSSSMASMMMAMXXASMMSAMMXSASMMMSSSMMSSMAXMASAMMXMXMMSMSMMMASXSMXXAMXXAMXXSASAASMSSSMMSMMMMMSAMASASMXASXXAXSAMAMSSMMSXMSSSMMAMSMAMAAXX
XAAXMASASMMAAXSMMSMSMMSMMSAMXMAMAAAAAASMXMMMAMASMMMMSAMXMAAAXAAXSAMAXAXMMMAMAMMSMAMSMMAXAXXMAMMAMXMASXSAAASAMMMSMMAMMMMXXXMASAAAXMASMMASXMMM
ASMMSXSXMAMSMMMAMXSAAMASMMMSXXAMMMMSMMSAASASXMASMASAXMASMMSXSXSMMXMSMSMMXMMSMMMXMAMXXMSMXMASMMXSMMSAMMMMMMMAMAXXASXMSAMXMMMXMMSMMMXXAXAXAMXM
MXXAXXMASMMXSASAMXSSXSASXAXMASXSXMMXXSMMMSASAMAMXAMXSSXXAAXASAMXSAMAXAASMXXAMAMASMSMSXXAMAMAAMAXAMSAMXAMXMSSMMSSMMMMSASMSASXSAMAMSMMSMSSMMAM
MSSMMXSASAASMMSSMAMAXMAXMMSXXMXAAMAAMXXAXMAMXMSSMXSAMXMSMMMAMAMXSASMSSSMAMSASASMXMAASAMSASXSAMAMAMMASMASMXAXAAXASAAXSAMXMMAMMXXAMASAXAMXXSAS
XAAAAAMXSMMSSXXMMXSXMASMSXMASMMSSMMSSMSSSSXSAXXMXAMXXAMASAMXMAMMSAMXMMAMSXXASAMXSMMMMAMXSMAXAMMSMMSAMMXMAMXSMSSMMSAMMMSXSAASXMSSMMMXSAMAASAS
MSSMMSSMXXAMAXMXXAXAMXXAXASMAAXAMMXAAAMAXAASMMSMXMMSSMSASMMMSMSAMMSXMSAMXMMXMXMXAAXAMXMXXXAMASAAAMMASXMMSAMXXAAAXMXSXAXMASXSXAMAAXXAXAMXMMAM
AXAXMXAXSMSSMSMMMMSMMMXXSAMSMMMSSSSSMMMMMMMMXAAXXAAAAAMXMAMAAMMXMAMAXXMXAMAXMASXSMMSSMMXSMSSXMXSMMMAMAAAMXXXASXMMMAMMSMMAMMMASMAMMMMSSMMXSAS
SMSMSXMMMAMXXSAMAAAXAMSAMXAXXMAMXMAAMXAXXSAMMSSMSMMSMMMXSAMMXSAMXAXMMSMSAMXAXXMAMAAMAAAXXAXMASAMMSMMSMMSSSMMMMSAAMAMAMAMASASAMXXMAXAAAAAXSAM
XAXASASAMSMMMSAMSSXSAMASXMAXXMAXSMSMMSMSXSASAMAMXXXAXAXXMAXMAMASXMSAAAXMMMXSMXMAMMMSSSMXMSMMMMXSAAAXXXXAAAXSAAMMSMMXXXAMAXAMXXMXSXXMSXMMXMAM
MAMAMSMXSAAAASAMAMAMMSMXXXMASMXSAAAXXMASXXXMMSXMSMMXSSSSSSSMMSXMAAAMMMSSSSMAAASASAXXAAXAXXMASMMMXSSMMMMMSMMSMXSXMMSAXXSMXMXXXAXAASXXAMMMASXM
MXMAMASXSXXMXSAMAMXMXAMMASMMSASAMSMSSSMXASXSAMXMASAMXAAAAAAMAMXSMMMSXXAAAASMSMSSSXSMMXSXMAXASXAXXMAMAXAXAAXXMASAMAMMSAXAMSAMMAMSAMXMMASMXMMM
MSSXSASXMAXSXSXMAMXXXAXSASAAMMMMMAXSAAMAMAAMAMXXAMMAMMMMMMMMAXXSAMAMXSXMSAMXXMXMMMXAMAMAMSMMSMSMAMSXMXSXXAMXXAMAMXAAMXMAAMASAXMMMSXASAMXAAAX
MXXXMAMAMAXSAMXSXXMXSSMMASMMXAMXSXMMSMMAXMSXMMSMSXSSXXAAXSASXSASXMXSAMXXMASMSMAXAAXMMAXAAAAXXAASMMAASAMASMXMMXSAMSMXSXMAMSAMMSSXMAXMMMSSSMMA
XMXMASMMMXMMAMXXAASXXXAMAMAMSMSMMAAMMXXXMXXAMXAAAASAMSMSSMAMAAMMXMMMXSMXSAMAXMASXSSSSSSMXSSMMSMAMMXAMAMMASAXXMXMAMXASASAXMASAXAMAMMXAMXAAASX
SAXASXAXMAASXMMMXMMAAXXMMXXMAXAAMXSMMSSSMASMMSMSMMMAMAMMMMAMAMXMAMSSXAAAMMMSMMAMAAAAAAAAAXAMAAMMXSAMSMMSAMMMSAMXSAMXMAMMXMMMXMXXAMXXMMSMSMMA
XXAXAMXMSSXMASAMASXMMMSAMSSMSSSXSAMAMXASMAMMASAXMASXMASMASASAXXSAAAMSMXMXSXMXMSSMMMMMMMMMXAASXSAAMAMXMAMAMAAMAMAMAAMMSMMASAAMSASXXSAXAXXAAXS
MMSSSMSXAXAMASASMMAMMXMAMAAAAAMXMXSAMMMMMSSMAMMMMXAXSAXMAMXSXMXAMMXXXMAXXAASXXXAMXMASXMASMMMMASAMSSMMMAXSAMXSMMSXSMSAXAMAAMMSAASAASXMAXXXMXM
SXAAAAAMSSMMAXMMMSAMMXSMMSSMMMMXSXSXMAXMXMAMMSMXMAMXMAMMMMXSAMXXXMXMASXSMMMMMSSMMXMASASMMASMMAMAMAMAAXASAXXXMASMAMSMASXMXSXXXMMMMMMXMASMAMAS
AMMSMXMAMAXMXMSAMMXMMAMMMMAMXXSMMASMSSSMASXMAMXAXSXMSXSAXSAMAMASMMMXAMXMXAMAAAAXAAMMSXMXSMMMMSXXMAMSMSSMMMSMSXMMAMAMMXMXAMXMMXXSXAXMAAAMAMAS
MXXXMAXMMAMXSXMASMSMMASAXSAMXAMAMAMMAAAXAMMMASMMXXAMAASAMMASAMAXAAAMSSMMXSSMMSMMSASASXMXSAMAAMXSMMXAMXAAXAXAXMAMSSMMMAMMSSMAAXXMAMXSMXXMXMXS
XXSASXSSMSSXMXSAMXAASASXXAMXXMSSMMSMMMMMMXASAMMMAMMMMXMMMSAMAMSSMMMSAAAASMXXMAXAAMMMMAMAXAXMMSXMASXMSSSMMXMAMAXMAMAAMXSAMAASMMSMXAAXMAXMAAAX
MMSASAAAAAXMSAMASXXMMAXMXMSASXAXAAMASXXMASMMMSSMMSSMMSAAAMASXMAAMMSMMXMMSXAXSSMMMXAAMMMAXSMSASAMAMAAXMAMSAMXSMXMAXMMSAMXSXMAAASXAMASXAAAXMMS
MAMXMMSMMMSAMXSAMAMSSSXSAAAMAMXSAMXASXXMASAAAAMAAAAAAXXMXSAMMSXMMXAAXSAMXMAMMXMAXMSSSSMMMMAMASAMASMMMXSASXSAXMASMXSAMASXMMMMMMMAMSAAMMSXSAXX
MXSXSXAXAMAXXXMAXAAXAAAAMSMMMSXMAAMXSXXMASXMMSSMMSMMMSMMMMASASAXXSXSMAMSSMMXSAMMXMAAMMASAMXMMMXMMXXMASMMMXMMMSMSAXXAMSXMAAXAXXSMMMXSXAAAMAMX
AMAXMXMSSSMSMMMSXSSMMMMMXAASXMASMMSXMMMSAMAAXXXXXMAXMAMAAXMMASAMXMMXMXMAXMAMXASAXXMXMMASMSMSMMXSMSSSSXXAMMXSAMMMXMSAMXASMMSXMXSXSXMAAMMSMMXX
MMMXSAXMAAAAAAMMAMXMXXSXSSSMASMMMAMAAAMMMSXMAMMAMSMMSSSSSSMMAMMMAAAAMAMXSMMSSMMMSSXAXMAXXMAAXXAASAAXMAXAMSAMXSXMXAMXMSAMXASXSMMMMXXMXXAMASMA
XAAAXMMMMMSMSSSMSMASMASXXXAXAMXAMAXMXMMAMAAMXAAMXSAAAXMXMAMMAXAXXMMAMSXXSAAXAXAMXMASXSMMMMSMSMXSMMSMMMMMAMXSAMAMMMSAMXSMMXXMAAAAMXSXMAXSSMAX
SMMSMXAAXMAXXXXAMXAMMAMXASXMSMSSSSSXSXSMSSXMASXSXSMMSXSMSAMSMMSAMXXSAMXASAMXMSXSAMXMAXAASXMAXXMMXSAMXASAAXASASASAAMASAMXMMMAXSMSAAMASMXMXASX
XAXMASXSMAMMSMMXMMSMMASXMMSMMAAAAAMAXMMMAXXMXXASAMXXMAMAMSMXMAXXMAMXASAMXAXXXAMMAMAMXMSASAMMMSSSSSXMMASMSMMMMAAMMSMASXSAAAAMXAMXMXXMAMAMXSMX
MAMMMMAAMXAXAAXAXAMASMAXMASMMMMSMMMMSASMMSSMAMSMSMAXMAMXAMMAMAMXXSAMXMXMXMMAMMMSAMXXXMMMSMMMXMAMAMMSMXMAMAXAAMXXXXMAMAMXSXSXMMMMMSMSASAXSAMX
MMSXMMSMMSASMSSSMASXMAAXMASMAMXXAXAAMXMAXAAXXMMAMSMSSXMMSXSAMASMMMXMXMASXXSMMSASASMSMXAAXXASXMAMAXAAMAMXXMSSXXAASMMAAAXXXAMMAMXAXAAMAMAXSASM
XMAAXAMAMXMSXAAXMXMAAMMSMSMMAMSSMMMSSMSXMXSMXMMMMAXAMXXAAAXMSASAAAAMAMXSMAMMAMASAMXAMSAMASXSASXSMSSMASXMXMAXAMXMMXSXSASAMAMXXASMSMSMSMSMSAMA
XAMMMAXAMAAMMMMMSXSMMMAXXAMMSMMAAMAMAXMAMAXAMXAAMXMSSMMMXSXXMMSMSMMXASAMMMMMASMMMMSSXXAXXMASAMMXMAXXXMASXMAXXSAMSAMXMASMSAMASXXMAAAXMXXAMXMX
MMSSXMSMMSMMASXMXASMSMXSMXXAXMXSMMAXSMSAMMXXAMXXSSMAAXMAAXMSMASAMASMMMASXMAXMXAXXXAXMXMXXXXMASAAMSSSMSXMMMMXASAXXAXXMXMXSXSMMMASMSMXSAMXMMAM
XAAMAXAMAAASASAXMXMASAXAAMMSSMAMASXMXAMASAAXSSXXAMMSMMMMXSAMMASASASASXMMASMSSSMMMMXMXXMASXMSXAAXMXAAASMAMASMAMMMMAMXSXMXMASAASAMXAAAAASAMXAM
MMMSXMAMSMMMASAMSMMXMXMMSMAAAMASAMXASMMAMMSMMAMMAMAMAAXAAMAMXAXXMASMMAAMMMAAAAXAXMASMMMAMXAAXMASXMMMMMAXMAMMMMSMASXMAXMASXMMMMASMMXMSAAXXMMS
XMXMAMXMXXMMXMAMAASAMXSXMMSSSMAMAXMMMSMXMMAAMAMSAMASAMXMSXMMMSMMMAMMSSMMSMMMMMXSASMMAAMMSMMMSXXMXMXMMMSXMAXSAAAAMXAXMMXASASXXSAMXMMXXXMSMSMS
MXAMSMAXMAMXSSXMSSMAMAAAMAAMXMSSMMSMAXSXXXSMMSXSXSXAXMAXXAASAMAAMMAAXXMAMMXMAXXMAXXSSMSMAAMSXMAMMXMSAAAXSAAMMMXXXXMMXSMMMAAAMMXMAXMAXMXAAAAX
SSMSMMSASAMXMAMMXMMSSMMSMSMMSXMAMAMMMMMMSMAAXXAMXMAXAASASXMMSSSMSSMXSMMASXSMMMSMAMMXMASMMSSMAMAMAAASMSSMMMSXSASMSMXMASAMMSMMMSMXMAMXXAXMSMSM
XASAMXMAAASXMSSSSMMXAAXAAXXMXASAMXMAAAMAAMSMMMAMMMMSSSMAAXMXXMAAAMMMMXMSAMXAAAAMMXMAMMMAMMMMSMAMSSXXMAAMXAAAMAMAAXXSASAMAAAXSAAASXSAASXMMXAA
SMMMSAMMMMAAXMAXAASXSMMMMMAMMAMXSXSSSSSSSXXXASAMAXXAXXMXMSMSSMMMMXASAAXAAASXMSMMXAMMSXMAMMAAMXAMXXMAMSMMMSSXMAMMMMXMMMAMSMSXSMXMSASAMMAASMSS
AXAASASAAMSMMMMSSMMAXAAXASAMXSMASAMXAMAAMASMMSASXSMMXSAAXAMAMXSASAMXASMSAMXXXXXSSMMMAMSMSMMMSMSMSASXMXAMAMXAMMSAAAMSXSAMMXXAMXSAMMMASXMMMAAM
MMMMXXMXXMASAXXAMAMSMSASMXSMMXAXMASMMMMMMAMXMXXMMXASASMSXSMSXMMASAXXXMAMXSMXSAXMASAMAXAMMMMSXAAAXAMMAMAMSSSXMAXXMXAAXMASXMMAMSAXMASAMXASMMMX
XMSMSSSMMSASMSMMSSMAAMMMXMAXAMMMSSXMXAMXASXMMSXMXSXMASAMAXAMXSMMMMMSSSMSAMXMMMMSAMXMMXMXAAMXMMMSMMAXMSSMXAXMXMSAMMMSAAAMAXMAMXMXSXXAMXMXAXAS
XSAAXAAAAMMMMXAXAASMSMSAASXMMMAMXMASXXSAAMAAASAMAMXMXMASXMAMAXAAAXXAAAAMMMAMAMXMASAMSASXSSSMAMMXAMAXMAAXMMMSAMAMXXMAMMAMAMSMMXXMMMSMMMMSSMAM
MSMSMSMMMSSMSMSMSXMXXAMSMSXAASMSMSAMXAMXXMXMAXAMSSSMXMXMASAMXSXMSSMMXMMMMMMMAMASXMAXAAMAMAMMAMXMAMAXMMSMASAMASMSAMSAMSXMXXAXMSSMAAAMMXAAXMAM
AMAMAXXXAAAMAAAAXXMSMMMMMSMSMSAAXMASMSMAMSMMXSAMXMAMMSMMMMMSMMMAMMMSAAAXSASXXMAMMXSXMSMAMAMMXSSXSMSXXXMMMMASAMXMAMSAMAASXSXMAAAMXSSSXMMSXSSS
MMAMXMMMMMSMMMMMMMMXAMXAASMMAMMMMMSMXMAMMAXAMXXMASAMMASAMMAMAAAXXSAMMMMXMASAMMSMMAXAMMMXSXSMSMMAMAMXMASAMSXMXSAMAMSAMSMXAMAMMSSMMMMAXMAMASMM
XSMSSSMMSAMXXXSSMMSSSMSMMSAMSMXSAMXMXMAMSSMMMAMSAXAMSASMXSASMMSMXMXXASMMSAMAMAXAMASXMAMXMMAXMAMAMAMXAMMXMXAAMSASAMSXMMAMSSMMAAMMAAXMAMMSXMAS
AXAMXXAAMAMSSXMAAAAXAAAXAMAMMMMMAMAMMMMMXXAAMAMMXXAMMAMMASAXXMAMXSSSMSAAMXMXMAXAMXSAMASAMSMMMAMXSXMMSSMASXMMMSAMMXMAMMAMAAXMMSSSSSSXSAASASAM
MMMMMSMMXAMXXASXMMXMMMMMXSMMMAASASASMSASASMMSASXSAMSMSMMXSAMSSMSAMXAMSMMSMMSMSSMSAMXSMSXMAMASASXAAAAAAMMMMSAMMAMXAXMASASMMMMSXAAAXAAMMXXAMMS
MAAAAAMMMSSMSXMXXXSMSAXXAAAASMMSASAMASASXAXASASAMSMAAAAXAMAMXAAMMSSSMSAAXAXSAXAAAMMMMAMMSXSAAMAMMMMMMSAMSASASXMAMMMSXSASMMMXAXMMSSMXMSSMAXMX
SSMSSSXSAAAAXMSAMXAXAASMMSSMSMMMMMMMMMAMMMSXSAMXMAMMMMXMMSAMXMMMAMAAASMMSSMMAMMMMXAMSMSASAMXXSAMXAMMMMASMASXMASMXSASMMAMAASMMXSAAAXSXAXMMMMM
MMAAXAAMMSMMMXSXAXMASMMMXAXXXXMAMAAMMMMMAXMAMMMXSSMSASAAMMASXMXMMSMMMMMAMAXMAXSSXMXXAAMXMXXSMMAMSASAAMSMMMMMMMMAAMAMXMSMSMSAXAMMXMMMMXXAMAAA
AMXMSMMMXAXAXAXXXAXAMAAXMASXMMSMSSSMXAMMSAMSMMMAMAASASMSMMMMAAXSMMXXAXMASMSSSSXXAMAMMSMSAMXMASAMXAXMMSXMAMXSASMMMMAMAAXAXMSAMXSAASMMAMSXSASX
XXAXAMXXSMMMMMSMSMMMSMMMMMMXMAAMMMAASAMAXXXAAAMXSMMMAMMMMASMMMMSAMASMXSASAAAXAMXAMAAMMMMAAAXMMAMMSMSMXMSMSAXAXSAMXXASAMXMASASAMMXMAMMXAAXAAM
MSMMMSXMXXMSAMAMAXMASXMAXAXXMSSSMMMMMAMMSMSSSMXASMSMXMAMSAXAAMAXAMMSMAMAMAMXMAASXSSSMAASXSMSXSSMAAAXMAMAAMMSMSAMXASXMXSASAMXMASMSSMMSMMMMXMA
MAXXAMMMAAXASXSSSMMASASMSMXSAAAAXSMXSXMXMAXMAMMASAAMXSAMMSSSMMSSXMXXMAMSMSMAAXAAAXMAXMMSXAMAMAXMSMMMMMMMMMXAXMAMXMMXMMXMASXMMXXAAMAXSAMXSAXM
SASMAMSMMSMXMMMAXXMAXMMAAASMMMSMMSMAXSMAMSMXAXMAMXAMAMXXAAMASAMAMSXSMMSAAAMMSSMMXMSAMMMMXXMASMMMMAAAXAXMASXMMSASAMXAXAMSAMSAXAMMMSSMSAMAXAMX
MASXMAAXAAMASMMAMXMSSMMMMMMAAMAXAMMMXMMASAMMAMMXSSSMXSAMMMSAMXMAMMAAASXMSMSMXMASMAMASMAMXXXAMMAAMMSMSXSAAXAAAXXXXMXAXSMSAAASMXSMAAAMSXMXSMSS
XMXAMSSSSMSASAMMXAAAAASMSMSMMSAMXSASASMMSAMSSMSSMAMMAXMXSXMXSAMMXMMMMMAMAAAMMXMSAMSAMMAMMSMSSSSXSAAXXAAMASMMMSSXSMSSXXASMXSXMAXMMSMMMMSXAAAX
MMMXMXAAMXMXXMSASMMMSMMAMAAXXMASXSASASAXMMMMAAXAMAMMMSMXSAMAXMASMMSAXXSMMSXMXSXXMAMXSSMSAAAXAAAMMMMSMXMMASAXMXMAMAAXAMAXAXXAMSMMXMXSAAXSMMMS
AXMASMMXMMSSMXXMAMXAAAMSMMMXXXAMXSXMAMMSAXSSMMMASXSXMSAAMAMXSXAAAASXSMXXXMAMSMSAMXAAXAMMXSMMMMXMASXSXMXSASMMSAMAMMMSXMSMMMXAMAAXAMASMSXMMAMA
XSMMXXMMMXAASAMXMMSSSMMAMXSAMSSXMSAMXMXMXMXXAMSXSAAAMXMASXMMMXAMMMMASMXMXMAMAAAMXSXMMXMAXMAASASMAXXXAXXMASMASAMXMAAMAMMAMASXMSMMXSAXMMAMXMXA
MSASMMXAXMSMMXMAXXAAMASXSAMAMAMMMSAMXSASAMXSXMMAMMMSXMXMAMAASASMXMMXMMAXASASMSMMAAAXXSMSSSMMSASMMXSSMMMMAMMMSAMSSMMSAMMXMMMAAAMAMMMXMSAMMSMX
ASAMXAXSSXMASMSSSMMSMAMAMMMAMASXAXXMASASASAXAMMXMSAMASXMAMXMAAXAAXSMMSSMXSASAMAMXSAMAAAMXXXAMAMAAAXXSASMAMXAMAMAAAASMSMSMSSMSMMMSAMAAXAMXAAM
XMAMMSMAAMSAMAAAAXXXMAMMMXSXSASMMSXSAMXMAMXMMMAAAMAXAMXMASMXMAMSXMSAMAAAMMXMMSAMAXAMXMXMASMMMSSMMMSASXSSXSAXSMMSSMMMXAMXAMAXMAMMSASXSMMSSMSX
MSAMAAMMMMMSSSMSMMMMSXSAMAAAMASAAAMMXSSSSXSXMAXMSSSMSSXMMSAASXMMMXSSMMMMSAAXAMMSMXXMAMXMMSAMXXAXAXMXMAXAMMMMXMMAMMXSSMSMMMMXMAMXMXMAAAXAAMAX
XSAMSXSAXSAAXAAXXSAAXASMSSSXMASMMSMXMAXAAAMASMSMAAMAMXMXMSMMXAAMXAMXSXAAXXMSMSMXMAMXASXMASMMMMMMMMMAMSMMMAMSAXMXXSAMMAXMXSXSXXXAXMMSSMMSSMAS
ASAMXXXASMMXSXMMASAMSAMXAMAXMASMAAASMMMMMMMAMAAMMMMAMMMSMMXXSSMAMMSAMSMSMMMXXAMAMAMSAMAMASXMAAXAXXXAXASMMAMMXXMXMMASMMXMASXSAXSMSAMAMAAXXMAS
MMXXXMSMMAAMMXXMAMAXMMMMSSMSMMSMMXAXAAAXAAMASXSMXSSSMSAXAXSAXASAAXMAXXMAMAMXSAXAXAMXAMXMASAMSXSAASMMSAMSMSMMSSXMASAMAXAMAXAMXMAAXAMXMSMMAMAS
XAXSMAAMMXMAAASMXSAMXSAMXAXSAMXMAXSSMSAMMXMMXAAAAAAXAMMSXMMMSAMMSAXMMXXAXMSAMMSSSSMSSMXSASMXAAMXMXAAMMMSMXAMXAASXMMSAAXMMSXSASMMMSXMAMASXMAS
MMSMASMXXAXMXMXAAXASASASXSMSAMAMMSMAXMASMSSSMSMMMSMMSMXXMAAAAXXAAAASASXSSMMASAAAAAAXMAXMASXMMMMAXMMMMXXMASXMMSMMAMMXMASMXAMMASMMMAASASAMXMAX
MXAXMAMXSMSAMXMMMSMMASAMXXAXAMXAMMXMMMMAAAAXMAMXXXXAAMXSAXASMSASMXMMASMMAMSAMXSMXMMMASXMAMMXAASXSXAXMXXMXMXSXAXSXMAAXMAMMXSMAMAAMMAMAMXMXMMM
SSSSXSXMAAAMAAMAXXAMAMMMMMAMSMSSSMASASASMMMMSMSSXAMSSSMMSXAAAXMAMAXMAMASAMXXMAXXXSSSMAMSXSAMSXMAMXMMSMASAXXXXMMMMSSSSMAMXSXMASMMSXAMXMAMAAAA
XAXMAMSAMXMSSXSAXXMMXSMMASAMAAAXAXMSASAXMSSMSXAMAMXMAMAAXMMMXMXXXMMMMAMXXMASMASXAMAAXMMMAMXXMAMXMMSAAAMSASXSMSAAAMAAXXASXMAMAAAXMMMMAMAMXXAS
MAMAASAMXSXAMAMXSMMAXAASASAXXMMXSMMMMMXMASAASMAXAMXXAMSMMAXAMSMSASMSMSXSAMXXMASXMMXMMXAMAXMXMAAXMAMMXSAMAMXSASMMSSMSMSMSAMAMSSMASXMSSSSSMSAA
MXASXSASAXMASMAMAAMMSSMMXSMMSASXAMSSSMSAXSMMMSASMXSMXXAASAMXSAAMAMAAAAXAMXSXMXSASXSSXSASAMXMSASMMSSMSXMMAMXMAMAXAAXMASAMXSMMAAAAAAAAXAAAAXAM
MXMMASAMXMAXAAAAXXMAAAMMMXMASAMSMMAAAAMSXMXXAMXXAASAMSSMMSMXMMSMSMXMMMSXMASMMAMAMAAXMMAMXAMXMSAAXMAMSAXSXSASASAMSSMSSXAXAAXMSMMMSMMMSMMMMXAX
XMMMMMMMAASAMSSSSMMMSXXXXXMXMAMMXMMSMMMMMMAMXSSMMMMAMMAMAAXXMXAXMASXXAAAMAMAMXMSMMMMXAAMXSAMXMXSMSSMSXMXMSMSAMMXXMAXMMMMMSSXXXMAXXXXAXMMSSMM
MAXSSXMSXSMMXMAAAXXXMMSMSMXXSAMXXMXMAXXAAMXSAAMMXXMSMMAMMSMMSMSMSMAMMMSXMSMSAMXXAXMSSMXSAMXMXSAMXMXAMMXAMMAMMMSXXMSMSMMAXAXAMXMMSMMXMMMAAAAX
ASMXSAAXMXMASMMSMMXMAXAAMSAXMXMASXMXMASMMSAMMSMMXSXXASMMXAXMAAAAAAAXAXMMMXAAMMMSMMSAMAMMXMAXAMMXXMMSMMSSSMAMAXXAXXAAXMASXXMSMXSAAAASMSMMSSSM
MXMMSMMSXSMAMSAXAMXMSSMMMXXSMAMAMAMSSMSAAMXMMMAMASMAMMAASAXMAXMSMSMSMXMAAMXMMAXAAXMAMSMAAMMMMSSMSMAAAXSAMXMMAMXMMSMSMMAXAMXAXAMSSSMMAAMAMAMA
ASAMXMASMMMSSMSSSSSMAAMSMMMAMAMAMAMMAXXMXMMMAXXMAMMAMAXMMASXSMAXAAAAASXSMSMSSXSMSXSXMAMSMSASAAAXXMSSSMMAMXXMMSAXMXXAAMXSAMXMMMMMXAMMXMMMMMMA
XSAMAMAXAXMMAAASAAAMSSMAAAAMSAXSXSMXMSMSAMSSMXAMXSMXMXSMXXMAMXAMSMSMSMAAXMAMXAXXMASMSXMAXMAMMSMSAMAAMASXMASXMSASXMMSSMMAMMAXAAAMSMMSMMMXASXS
ASAMXMAXSMMSMMMMMSMMMAXSSMMXMAXXAMSMXAMXASXAAASXAMASXMMAXMAMXAXAXMXXMMSMMMMMXXSSMMMXMXAMMMSMMAXSAMMSMXAMMAXMAMAMXAAAAAXXMSXMXXXXAMXAAXASXSAM
XSAMXMMMMAXSXSSXXAMAMXAMXAASMMSMAMASMASXMMXMMMXAMMAMAASMXXMMXSSSXMMXMAMXMSAMXXMAMSXMAMXAMMXASMXSAMXXMMMAMASXXMMMSMMSSSMMASXSASXMSAMMSMMXMMAM
MSAMMAMASXMXAAMXMSMMSMMSMSMSAAXMXMAMXAMXMAMXAXAMXMSSSMMMMSMMAMAMAAAAMXSAASASMXSAMMAMAASXSMSMMXMSXMXXAAXXSASMASMAMAMAMXAXAXAMASMAAMAMMXMAAMXM
ASMMAXSASAMMSMMSMASAAXXAAAASMMMSMMMSAMXAMASMMMMMXSAMXMAXAAAMXSMSXMSXXAMXMMXAAXAASAMXMXMAAASMSMMMAMASMMSXMAMMAMMAMMMSMMSMMSMMAMMSMXSXMASXMXAX
XXMASXMXSAMMAXAAXMASMSAMXMXMASXMAXAXAMSSXXXAMASAMXMXASMMSSSMAMXMMMAAMSAMXMSSSMMXMAMXSAMXMAMAXXASAMXXAAMAMAMMXSXMXSAMAAXMXSXAMXAXXAMMSMMAMSXM
AAXXAMMASXSSMMSSXSAAXMAXXXAAXMASAMSSSMAMMSSXMASAASXMXSAAAMAMSMAMAMMXMAAAXAAAXMXSSMMASASAMSMXMSXSMXSXMMMXMAMMMSAAXMASMMXSAMASXMASMXXMAMMAXMAM
SXMASAMASAMAAAAAAMSMSAMXSSMXSXXMAMMAMMXMAMAXMMSMMSAXMSMMMSXMAAMMSSXASXMMMMXSMMAAAAMMSAMXSXSAASAMMASMAAXXMXSAAXMMMMMMAMSMMSAMAMSXMAXSMXSMXMAA
XMAAAXMASMSSMMMMMMMXSAMXAXMASMSSSMMAMSMMSSMXXMXMXSAMAMXMASASMMXAAMXXMAAXXMAMAMXSSMMMMMMXMASMAMAMMASASMSMSASMSSXMASASAMAAAMASXMXAMSXMXXAMXSXM
MAMXXXSXSXMMXAMASXMASAMXMXSAXAMAAXMAMMMAMAXSMMMMMMSAMXAXMXXAXAMSSSSSSSMMAMAMSMMXMASAAAMAMXMMASAMMMSAXAXXMAMAAXMXMSASXSMMMSAMXSMMMXAMASASMMSM
ASMMSASAMMMMSXSASAXXXMAXXMMAMSMSMSSSMSMSXMMMAAXMAAAMMXMSMSMMMMMAAMAXAXMSASXSAMMASXMMSMSMSMXSASASXXMAMXMMMSMMMSSMMMXMAXXAXMASAXXXAMAMXSAMMAAX
MXAASAMAMAAXMXMASMMMSAMXMSSSMAMMAMAMASAMASMSSMSSMXMMSAAXAASAAMMMSMMMASXMASXAAXMASAXAMXAAMAXMMSAMXAMMMASAAXASAMXASXSMSMSMSSMMMMXMSSSMXMAMMSSX
XSMMMASMMSSXSAMXMAAMASXXMAAASXMMSMAMMMASAMXAASAXXXSASMMMSMSMSSXXAMXMXMASXMASMMMASXMASXMMMXSAAMAMSMAAAASMXSAMSSSMMASAAAAMMAXMAMXAXAAMAMAMXAXX
AXXXSXMAXXMXSSSMSSMSXMSSMMSMMMXXXAMXSMXMXMMSSMSMSAMXSAMXXXXMMMMAMXXXSSXMXSAMXXMXMXAXXMASXMSMMSMMSXSAMMSXMMMSXMAXMAMXMSMXSAMSXSXSMSMMXSASMASX`,
}
