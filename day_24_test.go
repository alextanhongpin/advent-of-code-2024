package main

import (
	"fmt"
	"go-aoc-2025/utils"
	"slices"
	"strings"
)

func ExampleDayN() {
	fmt.Println("part 1:", part1(inputs[0]))
	fmt.Println("part 1:", part1(inputs[1]))
	fmt.Println()
	fmt.Println("part 2:", part2(inputs[1]))

	// Output:
	// part 1: 2024
	// part 1: 43942008931358
	//
	// NOT SOLVED
	// part 2: dvb,fhg,fsq,tnc,vcf,z10,z17,z39
}

type Gate struct {
	a, b string
	out  string
	op   string
}

func part1(input string) int {
	a, b, ok := strings.Cut(input, "\n\n")
	if !ok {
		panic("invalid")
	}

	vals := make(map[string]int)
	for _, line := range strings.Split(a, "\n") {
		a, b, ok := strings.Cut(line, ": ")
		if !ok {
			panic("invalid line")
		}
		if b == "1" {
			vals[a] = 1
		} else {
			vals[a] = 0
		}
	}

	var gates []Gate
	for _, row := range strings.Split(b, "\n") {
		parts := strings.Fields(row)
		g := Gate{
			a:   parts[0],
			b:   parts[2],
			out: parts[4],
			op:  parts[1],
		}
		gates = append(gates, g)
	}

	for len(gates) > 0 {
		waiting := slices.Clone(gates)
		gates = nil
		for _, g := range waiting {
			a, ok := vals[g.a]
			if !ok {
				gates = append(gates, g)
				continue
			}

			b, ok := vals[g.b]
			if !ok {
				gates = append(gates, g)
				continue
			}

			switch g.op {
			case "AND":
				vals[g.out] = a & b
			case "OR":
				vals[g.out] = a | b
			case "XOR":
				vals[g.out] = a ^ b
			}
		}
	}

	return parseBinary("z", vals)
}

func extract(g string, vals map[string]int) map[string]int {
	m := make(map[string]int)
	for k := range vals {
		if strings.HasPrefix(k, g) {
			m[k] = vals[k]
		}
	}

	return m
}

func parseBinary(g string, vals map[string]int) int {
	var res int

	for i := range 100 {
		key := fmt.Sprintf("%s%02d", g, i)
		v, ok := vals[key]
		if !ok {
			break
		}
		res += (1 << i) * v
	}

	return res
}

func toZ(n int) map[string]int {
	z := strings.Split(fmt.Sprintf("%b", n), "")
	slices.Reverse(z)
	zm := make(map[string]int)
	for i := range len(z) {
		zm[fmt.Sprintf("z%02d", i)] = utils.ToInt(z[i])
	}
	return zm
}

func part2(input string) int {
	a, b, ok := strings.Cut(input, "\n\n")
	if !ok {
		panic("invalid")
	}

	vals := make(map[string]int)
	for _, line := range strings.Split(a, "\n") {
		a, b, ok := strings.Cut(line, ": ")
		if !ok {
			panic("invalid line")
		}
		if b == "1" {
			vals[a] = 1
		} else {
			vals[a] = 0
		}
	}

	x := parseBinary("x", vals)
	y := parseBinary("y", vals)

	var gates []Gate
	for _, row := range strings.Split(b, "\n") {
		parts := strings.Fields(row)
		g := Gate{
			a:   parts[0],
			b:   parts[2],
			out: parts[4],
			op:  parts[1],
		}
		gates = append(gates, g)
	}

	z := x + y
	m := mismatch(gates, vals, z)

	wrong := make(map[string]result)
	for _, r := range m {
		wrong[r.z] = r
	}

	isXOrY := func(g string) bool {
		return strings.HasPrefix(g, "x") || strings.HasPrefix(g, "y")
	}

	var walk func(s string, depth int) string
	walk = func(z string, depth int) string {
		if depth != -1 && depth == 0 {
			return z
		}

		for i, g := range gates {
			if g.out != z {
				continue
			}

			// Find the faulty gate.
			// If the OP is XOR, and the expected value is 1, then the side with 0
			// 1 AND 1 = 1
			// 1 AND 0 = 0
			// 0 AND 1 = 0
			// 1 AND 1 = 0
			var lhs, rhs string
			if isXOrY(g.a) {
				lhs = g.a
			} else {
				lhs = walk(g.a, depth-1)
			}
			if isXOrY(g.b) {
				rhs = g.b
			} else {
				rhs = walk(g.b, depth-1)
			}

			_ = i
			return fmt.Sprintf("(%s %s %s)", lhs, g.op, rhs)
		}

		return ""
	}

	fmt.Println(len(wrong))
	for _, g := range gates {
		r, ok := wrong[g.out]
		if strings.HasPrefix(g.out, "z") && ok {
			fmt.Println(g.out, r.got, r.want)
			fmt.Println(walk(g.out, 5))
			fmt.Println()
		}
	}

	return 0
}

func mismatch(input []Gate, vals map[string]int, z int) []result {
	gates := slices.Clone(input)
	for len(gates) > 0 {
		waiting := slices.Clone(gates)
		gates = nil
		for _, g := range waiting {
			a, ok := vals[g.a]
			if !ok {
				gates = append(gates, g)
				continue
			}

			b, ok := vals[g.b]
			if !ok {
				gates = append(gates, g)
				continue
			}

			switch g.op {
			case "AND":
				vals[g.out] = a & b
			case "OR":
				vals[g.out] = a | b
			case "XOR":
				vals[g.out] = a ^ b
			}
		}
	}

	var swap []result
	want := toZ(z)
	got := extract("z", vals)
	for k := range want {
		if got[k] != want[k] {
			swap = append(swap, result{
				want: want[k],
				got:  got[k],
				z:    k,
			})
		}
	}

	return swap
}

type result struct {
	want int
	got  int
	z    string
}

var inputs = []string{
	`x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj`,
	`x00: 1
x01: 1
x02: 0
x03: 0
x04: 0
x05: 0
x06: 0
x07: 0
x08: 0
x09: 0
x10: 1
x11: 1
x12: 0
x13: 1
x14: 0
x15: 0
x16: 0
x17: 0
x18: 0
x19: 0
x20: 0
x21: 1
x22: 1
x23: 1
x24: 0
x25: 0
x26: 0
x27: 1
x28: 0
x29: 1
x30: 1
x31: 0
x32: 1
x33: 1
x34: 1
x35: 1
x36: 0
x37: 1
x38: 0
x39: 1
x40: 0
x41: 1
x42: 1
x43: 0
x44: 1
y00: 1
y01: 1
y02: 0
y03: 1
y04: 1
y05: 0
y06: 0
y07: 0
y08: 0
y09: 0
y10: 1
y11: 1
y12: 1
y13: 0
y14: 0
y15: 0
y16: 1
y17: 1
y18: 0
y19: 0
y20: 1
y21: 1
y22: 1
y23: 1
y24: 0
y25: 1
y26: 0
y27: 0
y28: 0
y29: 1
y30: 0
y31: 1
y32: 1
y33: 1
y34: 1
y35: 1
y36: 0
y37: 0
y38: 1
y39: 1
y40: 1
y41: 0
y42: 0
y43: 0
y44: 1

hvk XOR hpr -> z41
y27 AND x27 -> qqr
sfm XOR wsn -> z06
fvb OR hqb -> z45
x15 XOR y15 -> mng
pvs AND wgc -> mjr
dpc OR pwj -> jjf
jsn AND dvb -> ftc
x13 AND y13 -> sdf
dvb XOR jsn -> z35
rvd OR wrj -> z39
hmg XOR jjd -> z43
krc AND bcj -> rhp
cnr XOR hct -> z04
tsd OR wfw -> dqp
y35 XOR x35 -> fsq
x07 XOR y07 -> nvh
y23 AND x23 -> fjp
qjn OR fhg -> jfb
x43 XOR y43 -> jjd
wps XOR rnc -> z42
qkw XOR dqp -> z29
x02 AND y02 -> dmq
qjg AND jjf -> z17
x20 XOR y20 -> shw
kwg OR fwc -> skm
gmw XOR cqn -> z08
x19 XOR y19 -> pjs
x08 AND y08 -> qwf
wsn AND sfm -> kjs
x00 XOR y00 -> z00
brb OR wnc -> sfm
x27 XOR y27 -> qbg
y19 AND x19 -> ccq
pmn AND pfq -> rnj
x17 XOR y17 -> qjg
gvj AND qbw -> brq
y28 XOR x28 -> cws
htw AND mpj -> sdk
vks XOR vwg -> z23
y24 XOR x24 -> swd
qqr OR dhp -> kqp
bgh XOR ctc -> z13
x09 AND y09 -> kwg
ctc AND bgh -> jmr
crj AND ghf -> mns
dcj XOR rfv -> z16
hfm OR wwm -> kcd
hvk AND hpr -> ngm
jjb OR bfn -> hkj
qnv AND vbh -> tgn
psp OR kss -> mdv
y25 AND x25 -> dwh
y35 AND x35 -> dvb
y20 AND x20 -> mkd
qbw XOR gvj -> z12
y39 AND x39 -> rvd
ncq XOR bwc -> z36
bmd AND fgb -> jgv
x03 AND y03 -> psr
btr OR nnb -> vks
x37 XOR y37 -> rhj
y13 XOR x13 -> ctc
x06 XOR y06 -> wsn
rhp OR whq -> pfq
jfb XOR vgw -> z18
kqp AND cws -> wfw
x42 XOR y42 -> rnc
y05 AND x05 -> wnc
ggt OR nbp -> ffv
ddd XOR tvq -> z21
x44 XOR y44 -> qgg
y31 AND x31 -> vqt
hkj XOR rsn -> z38
y38 AND x38 -> gqm
x30 XOR y30 -> hqf
dcj AND rfv -> dpc
x06 AND y06 -> kbv
dqp AND qkw -> dpr
x11 XOR y11 -> bmd
vgw AND jfb -> jsh
x17 AND y17 -> qjn
y31 XOR x31 -> crj
gqm OR kfp -> mnd
y40 AND x40 -> jmc
x10 XOR y10 -> kck
vqt OR mns -> bcj
dmq OR sdk -> gwf
ccq OR jfs -> trb
kmh AND mnd -> wrj
kmh XOR mnd -> tnc
x10 AND y10 -> z10
y16 AND x16 -> pwj
y00 AND x00 -> kvj
y26 AND x26 -> jsr
ctm XOR hqf -> z30
pjs AND nwg -> jfs
x11 AND y11 -> gvd
tgn OR bbf -> jsn
ctm AND hqf -> dmw
krc XOR bcj -> z32
ftc OR fsq -> bwc
fgb XOR bmd -> z11
nbw AND gwf -> btm
hmg AND jjd -> wfc
y28 AND x28 -> tsd
rhj XOR mdv -> z37
y36 AND x36 -> kss
mpf OR jsh -> nwg
x36 XOR y36 -> ncq
rsn AND hkj -> kfp
ffw XOR snd -> z26
y44 AND x44 -> hqb
y23 XOR x23 -> vwg
sst OR vcf -> fgb
x04 XOR y04 -> cnr
jqv XOR nvh -> z07
nbg OR pnm -> rfv
y16 XOR x16 -> dcj
jmr OR sdf -> wgc
hsp OR rft -> hmg
nvh AND jqv -> tfs
x14 AND y14 -> dtj
kbv OR kjs -> jqv
y32 XOR x32 -> krc
bwc AND ncq -> psp
skm AND kck -> sst
y42 AND x42 -> rft
psr OR btm -> hct
tcj AND swd -> wwm
x34 XOR y34 -> vbh
pmh XOR ffv -> z05
shw XOR trb -> z20
rtf AND tnc -> gtv
cws XOR kqp -> z28
x22 XOR y22 -> tfj
y41 AND x41 -> mdh
tvq AND ddd -> fcj
mkd OR nrr -> tvq
x26 XOR y26 -> snd
mng AND rhw -> nbg
x25 XOR y25 -> vnt
knf XOR qgg -> z44
ngm OR mdh -> wps
y09 XOR x09 -> gff
x22 AND y22 -> nnb
ffw AND snd -> msw
y04 AND x04 -> ggt
y41 XOR x41 -> hpr
pgc XOR tfj -> z22
y29 AND x29 -> jbf
y30 AND x30 -> ttk
ttk OR dmw -> ghf
nbw XOR gwf -> z03
kvj AND bhq -> wkc
knf AND qgg -> fvb
x12 XOR y12 -> qbw
ghf XOR crj -> z31
x21 AND y21 -> drw
vnt AND kcd -> vdd
nfp OR wfc -> knf
x18 AND y18 -> mpf
x12 AND y12 -> kdb
x34 AND y34 -> bbf
fcj OR drw -> pgc
y37 AND x37 -> jjb
dpr OR jbf -> ctm
kdb OR brq -> bgh
x40 XOR y40 -> rtf
mpj XOR htw -> z02
trb AND shw -> nrr
y21 XOR x21 -> ddd
y08 XOR x08 -> gmw
jgv OR gvd -> gvj
kcd XOR vnt -> z25
jgj AND gff -> fwc
qnv XOR vbh -> z34
bhq XOR kvj -> z01
y39 XOR x39 -> kmh
jgj XOR gff -> z09
y01 AND x01 -> rtg
pmn XOR pfq -> z33
msw OR jsr -> wgr
pgc AND tfj -> btr
wgr AND qbg -> dhp
y33 AND x33 -> dfv
y02 XOR x02 -> mpj
x14 XOR y14 -> pvs
rtf XOR tnc -> z40
jjf XOR qjg -> fhg
vks AND vwg -> jfd
dtj OR mjr -> rhw
x38 XOR y38 -> rsn
tcj XOR swd -> z24
vdd OR dwh -> ffw
ffv AND pmh -> brb
cnr AND hct -> nbp
y03 XOR x03 -> nbw
gtv OR jmc -> hvk
y15 AND x15 -> pnm
y24 AND x24 -> hfm
rtg OR wkc -> htw
y05 XOR x05 -> pmh
gmw AND cqn -> jfn
rnc AND wps -> hsp
jfd OR fjp -> tcj
x29 XOR y29 -> qkw
y01 XOR x01 -> bhq
y07 AND x07 -> ghp
x33 XOR y33 -> pmn
pjs XOR nwg -> z19
x18 XOR y18 -> vgw
dfv OR rnj -> qnv
qbg XOR wgr -> z27
kck XOR skm -> vcf
rhj AND mdv -> bfn
qwf OR jfn -> jgj
mng XOR rhw -> z15
y32 AND x32 -> whq
x43 AND y43 -> nfp
pvs XOR wgc -> z14
ghp OR tfs -> cqn`,

	`x00: 1
x01: 1
x02: 1
y00: 0
y01: 1
y02: 0

x00 AND y00 -> z00
x01 XOR y01 -> z01
x02 OR y02 -> z02`,

	`x00: 0
x01: 1
x02: 0
x03: 1
x04: 0
x05: 1
y00: 0
y01: 0
y02: 1
y03: 1
y04: 0
y05: 1

x00 AND y00 -> z00
x01 AND y01 -> z01
x02 AND y02 -> z02
x03 AND y03 -> z03
x04 AND y04 -> z04
x05 AND y05 -> z05`,
}
