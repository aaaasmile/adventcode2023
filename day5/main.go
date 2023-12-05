package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var testInput string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
	testInput = strings.TrimRight(testInput, "\n")
	if len(testInput) == 0 {
		panic("empty test.txt file")
	}
}

func main() {
	var part int
	var test bool
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.BoolVar(&test, "test", false, "run with test.txt inputs?")
	flag.Parse()
	fmt.Println("Running part", part, ", test inputs = ", test)

	if test {
		input = testInput
	}

	var ans int
	switch part {
	case 1:
		ans = part1(input)
	case 2:
		ans = part2(input)
	}
	fmt.Println("Output:", ans)
}

type SectionLine struct {
	_dest_start   int
	_source_start int
	_range_len    int
}

type Section struct {
	_lines []SectionLine
	_corr  map[int]int
}

func (sc *Section) addLine(ll []int) {
	if len(ll) != 3 {
		panic("format error")
	}
	sl := SectionLine{
		_dest_start:   ll[0],
		_source_start: ll[1],
		_range_len:    ll[2],
	}
	for i := 0; i < sl._range_len; i++ {
		ix_s := sl._source_start + i
		ix_d := sl._dest_start + i
		sc._corr[ix_s] = ix_d
	}
	sc._lines = append(sc._lines, sl)
}

func (sc *Section) DestToSource(ss int) int {
	if vv, ok := sc._corr[ss]; ok {
		return vv
	}
	return ss
}

func NewSection() *Section {
	sc := Section{
		_lines: []SectionLine{},
		_corr:  map[int]int{},
	}
	return &sc
}

type Almanc struct {
	_seeds  []int
	_detail map[string]*Section
}

func (al *Almanc) PrintSeedToSoil() {
	sct := al._detail["seed-to-soil"]
	for _, v := range al._seeds {
		d := sct.DestToSource(v)
		log.Printf("Seed number %d corresponds to soil number %d\n", v, d)
	}
}

func (al *Almanc) findSeedLocation(ss int) int {
	seq := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
	curr_src := ss
	dest := 0
	for _, seqkey := range seq {
		sct, ok := al._detail[seqkey]
		if !ok {
			panic("map not found")
		}
		dest = sct.DestToSource(curr_src)
		//fmt.Printf("[%s]%d  to %d", seqkey, curr_src, dest)
		curr_src = dest
	}
	return dest
}

func (al *Almanc) FindMinLocation() int {
	locations := []int{}
	for _, ss := range al._seeds {
		loc := al.findSeedLocation(ss)
		locations = append(locations, loc)
	}
	fmt.Println("Locations", locations)
	mm := slices.Min(locations)
	return mm
}

type ParState int

const (
	LookSeed ParState = iota
	LookKey
	LookData
	LookEofData
)

func part1(input string) int {
	alm := Almanc{
		_seeds:  []int{},
		_detail: make(map[string]*Section),
	}
	state := LookSeed
	key := ""
	section := NewSection()
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		if line == "" {
			if state == LookData {
				alm._detail[key] = section
			}
			state = LookKey
			key = ""
			continue
		}
		if state == LookSeed {
			tp := strings.Split(line, ":")
			alm._seeds = spaceStrToNumArray(tp[1])
			state = LookKey
		} else if state == LookKey {
			tp := strings.Split(line, ":")
			kk := tp[0]
			kn := strings.Replace(kk, " map :", "", 0)
			kn = strings.Split(kn, " ")[0]
			key = strings.TrimSpace(kn)
			section = NewSection()
			state = LookData
		} else if state == LookData {
			dd := spaceStrToNumArray(line)
			section.addLine(dd)
		}
	}
	if key != "" {
		alm._detail[key] = section
	}
	fmt.Println("seeds: ", alm._seeds)
	for k, v := range alm._detail {
		fmt.Println("section: ", k, v)
	}
	alm.PrintSeedToSoil()
	alm.findSeedLocation(79)
	mm := alm.FindMinLocation()
	log.Println("min location is ", mm)

	return 0
}

func part2(input string) int {
	return 0
}

func spaceStrToNumArray(s string) []int {
	res := []int{}
	sa := strings.Split(s, " ")
	for _, vv := range sa {
		if vv == "" {
			continue
		}
		num, err := strconv.Atoi(vv)
		if err != nil {
			panic(err)
		}
		res = append(res, num)
	}
	return res
}
