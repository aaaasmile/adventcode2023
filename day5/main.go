package main

import (
	"cmp"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
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
	flag.IntVar(&part, "part", 2, "part 1 or 2")
	flag.BoolVar(&test, "test", false, "run with test.txt inputs?")
	flag.Parse()
	fmt.Println("Running part", part, ", test inputs = ", test)
	start := time.Now()
	if test {
		input = testInput
	}

	var ans int
	switch part {
	case 1:
		ans = part12(input, false)
	case 2:
		ans = part12(input, true)
	}
	fmt.Println("Output:", ans)
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("total call duration: %v\n", elapsed)
}

type SectionLine struct {
	_dest_start   int
	_source_start int
	_range_len    int
}

type Section struct {
	_lines []SectionLine
	_corr  map[int]SectionLine
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

	sc._lines = append(sc._lines, sl)
	slices.SortFunc(sc._lines, func(a, b SectionLine) int {
		return cmp.Compare(a._source_start, b._source_start)
	})
}

func (sc *Section) addLineRev(ll []int) {
	if len(ll) != 3 {
		panic("format error")
	}
	sl := SectionLine{
		_dest_start:   ll[0],
		_source_start: ll[1],
		_range_len:    ll[2],
	}

	sc._lines = append(sc._lines, sl)
	slices.SortFunc(sc._lines, func(a, b SectionLine) int {
		return cmp.Compare(b._source_start, a._source_start)
	})
}

func (sc *Section) CalcSource(dst int) int {
	for _, sl := range sc._lines {
		if dst == sl._dest_start {
			return sl._source_start
		}
	}
	for _, sl := range sc._lines {
		offset := dst - sl._dest_start
		if dst < sl._dest_start {
			return dst
		}
		if dst > sl._dest_start+sl._range_len {
			continue
		}
		return sl._source_start + offset
	}
	return dst
}

func (sc *Section) CalcDestination(ss int) int {
	for _, sl := range sc._lines {
		if ss == sl._source_start {
			return sl._dest_start
		}
	}
	for _, sl := range sc._lines {
		offset := ss - sl._source_start
		if ss < sl._source_start {
			return ss
		}
		if ss > sl._source_start+sl._range_len {
			continue
		}
		return sl._dest_start + offset
	}
	return ss
}

func NewSection() *Section {
	sc := Section{
		_lines: []SectionLine{},
	}
	return &sc
}

type SeedWithRange struct {
	_startIx int
	_length  int
	_endIx   int
}

type Almanc struct {
	_seeds          []int
	_seedsWithRange []SeedWithRange
	_detail         map[string]*Section
}

func (al *Almanc) PrintSeedToSoil() {
	sct := al._detail["seed-to-soil"]
	for _, v := range al._seeds {
		d := sct.CalcDestination(v)
		log.Printf("Seed number %d corresponds to soil number %d\n", v, d)
	}
}

func (al *Almanc) PrintfindSeedLocation(ss int) int {
	seq := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
	curr_src := ss
	dest := 0
	for _, seqkey := range seq {
		sct, ok := al._detail[seqkey]
		if !ok {
			panic("map not found")
		}
		dest = sct.CalcDestination(curr_src)
		fmt.Printf("[%s]%d  to %d", seqkey, curr_src, dest)
		curr_src = dest
	}
	return dest
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
		dest = sct.CalcDestination(curr_src)
		//fmt.Printf("[%s]%d  to %d", seqkey, curr_src, dest)
		curr_src = dest
	}
	return dest
}

func (al *Almanc) findSeedRangeLocation(ss_s, ss_e, b_loc_s, b_loc_e int) int {
	loc_s := b_loc_s
	loc_e := b_loc_e
	if loc_s == -1 {
		loc_s = al.findSeedLocation(ss_s)
	}
	if loc_e == -1 {
		loc_e = al.findSeedLocation(ss_e)
	}

	if loc_s == loc_e {
		return loc_e
	}
	if ss_e-ss_s == 1 {
		if loc_e < loc_s {
			return loc_e
		} else {
			return loc_s
		}
	}
	half := ss_s + (int)((ss_e-ss_s)/2)
	if loc_e < loc_s {
		return al.findSeedRangeLocation(half, ss_e, -1, loc_e)
	} else {
		return al.findSeedRangeLocation(ss_s, half, loc_s, -1)
	}
}

func (al *Almanc) CheckSeedLocation() int {
	locations := []int{}
	for i := 0; i < 100; i++ {
		loc := al.findSeedLocation(i)
		locations = append(locations, loc)
	}
	fmt.Println("Locations", locations)
	mm := slices.Min(locations)
	return mm
}

func (al *Almanc) findReverseLocationToSeed(dst int) int {
	seq := []string{"humidity-to-location", "temperature-to-humidity", "light-to-temperature", "water-to-light", "fertilizer-to-water", "soil-to-fertilizer", "seed-to-soil"}
	curr_dst := dst
	src := 0
	for _, seqkey := range seq {
		sct, ok := al._detail[seqkey]
		if !ok {
			panic("map not found")
		}
		src = sct.CalcSource(curr_dst)
		fmt.Printf("[%s]%d  to %d", seqkey, curr_dst, src)
		curr_dst = src
	}

	return src
}

func (al *Almanc) IsSeedInRange(seed int) bool {
	ix := slices.IndexFunc(al._seedsWithRange, func(c SeedWithRange) bool {
		return c._startIx <= seed && seed <= c._endIx
	})
	return ix >= 0
}

func (al *Almanc) FindMinLocationWithReverse() int {
	locations := []int{}
	loc := 0
	//for _, ss := range al._seeds {
	sct := al._detail["humidity-to-location"]
	for ix := range sct._lines {
		sctline := sct._lines[ix]
		loc = sctline._dest_start
		old_size := len(locations)
		locations = al.reverselocation(loc, locations)
		if len(locations) > old_size {
			break
		}
		loc = sctline._dest_start + sctline._range_len
		locations = al.reverselocation(loc, locations)
	}
	for i := 0; i < 2; i++ {
		old_size := len(locations)
		locations = al.reverselocation(i, locations)
		if len(locations) > old_size {
			fmt.Println("*** Wov found it", i)
			break
		}
	}

	fmt.Println("Locations", locations)
	mm := slices.Min(locations)
	return mm
}

func (al *Almanc) reverselocation(loc int, locations []int) []int {
	seed := al.findReverseLocationToSeed(loc)

	if al.IsSeedInRange(seed) {
		fmt.Println("Seed is a min", seed)
		locations = append(locations, loc)
	}
	return locations
}

func (al *Almanc) FindSeedRangeMinLocation() int {
	locations := []int{}
	rr := -1
	ss_s := -1
	for _, vs := range al._seeds {
		if ss_s == -1 {
			ss_s = vs
		} else {
			rr = vs
			loc := al.findSeedRangeLocation(ss_s, ss_s+rr, -1, -1)
			locations = append(locations, loc)
			rr = -1
			ss_s = -1
		}
	}
	fmt.Println("Locations", locations)
	mm := slices.Min(locations)
	return mm
}

func (al *Almanc) FindSimpleMinLocation() int {
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

func part12(input string, rangeSearch bool) int {
	alm := Almanc{
		_seeds:          []int{},
		_seedsWithRange: []SeedWithRange{},
		_detail:         make(map[string]*Section),
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
			if rangeSearch {
				dd := spaceStrToNumArray(tp[1])
				ssl := -1
				ssr := -1
				for _, ddv := range dd {
					if ssl == -1 {
						ssl = ddv
					} else {
						ssr = ddv
						alm._seedsWithRange = append(alm._seedsWithRange, SeedWithRange{_startIx: ssl, _length: ssr, _endIx: ssl + ssr})
						ssl = -1
						ssr = -1
					}
				}
			} else {
				alm._seeds = spaceStrToNumArray(tp[1])
			}
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
			if rangeSearch {
				section.addLineRev(dd)
			} else {
				section.addLine(dd)
			}

		}
	}
	if key != "" {
		alm._detail[key] = section
	}
	fmt.Println("seeds: ", alm._seeds)
	for k, v := range alm._detail {
		fmt.Println("section: ", k, v)
	}

	mm := 0
	if rangeSearch {
		// part 2
		//mm = alm.PrintfindSeedLocation(82)
		//fmt.Println("Min location for seed 82: ", mm)
		mm = alm.FindMinLocationWithReverse()
		//alm.CheckSeedLocation()
	} else {
		// Part 1
		alm.PrintSeedToSoil()
		alm.findSeedLocation(79)
		mm = alm.FindSimpleMinLocation()
	}

	log.Println("min location is ", mm)

	return mm
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
