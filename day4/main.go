// skeleton template copied and modified from https://github.com/alexchao26/advent-of-code-go/blob/main/scripts/skeleton/tmpls/main.go

// started        ;
// finished part1 , 'go run' time s, run time after 'go build' s
// finished part2 , 'go run' time s, run time after 'go build' s

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
	flag.IntVar(&part, "part", 2, "part 1 or 2")
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

type NumColl struct {
	_serie []int
	_id    int
}

func (nc *NumColl) FindInto(other NumColl) NumColl {
	res := NumColl{_id: nc._id}
	for _, vo := range other._serie {
		idx := slices.IndexFunc(nc._serie, func(c int) bool { return c == vo })
		if idx >= 0 {
			res._serie = append(res._serie, nc._serie[idx])
		}
	}
	return res
}

func (nc *NumColl) Points() int {
	points := 0
	for ix := range nc._serie {
		if ix == 0 {
			points = 1
		} else {
			points = points * 2
		}
	}
	return points
}

type Scratches struct {
	_freq_orig []int
	_freq_copy []int
}

func (sc *Scratches) IncCount(ixCurr int, nc NumColl) {
	sc._freq_orig[ixCurr] += 1
	sc.setCopies(ixCurr+1, nc)
	// process previous copies
	for ip := 0; ip < sc._freq_copy[ixCurr]; ip++ {
		sc.setCopies(ixCurr+1, nc)
	}
}

func (sc *Scratches) setCopies(ixNext int, nc NumColl) {
	if len(nc._serie) == 0 {
		return
	}
	max_len := ixNext + len(nc._serie)
	if max_len > len(sc._freq_orig) {
		max_len = len(sc._freq_orig)
	}
	for ic := ixNext; ic < max_len; ic++ {
		sc._freq_copy[ic] += 1
	}
}

func (sc *Scratches) Total() int {
	fmt.Println("orig: ", sc._freq_orig)
	fmt.Println("copy: ", sc._freq_copy)
	count := 0
	for ix, vv := range sc._freq_orig {
		count += vv
		count += sc._freq_copy[ix]
	}
	return count
}

// something "34 53  9 36 52 30 70 60 65 96" to [34, 53, 9, 36, 52, 30, 70, 60, 65, 96]
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

func part1(input string) int {
	row := 1
	sum_points := 0
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		i := strings.IndexByte(line, ':')
		ss := line[i+2:]
		wt := strings.Split(ss, " | ")
		win_num := spaceStrToNumArray(wt[0])
		my_num := spaceStrToNumArray(wt[1])
		fmt.Println("scratch: ", row, win_num, my_num)
		ww := NumColl{
			_serie: win_num,
			_id:    row,
		}
		mm := NumColl{
			_serie: my_num,
			_id:    row,
		}
		winner := mm.FindInto(ww)
		points := winner.Points()
		fmt.Println("winner", winner, points)
		sum_points += points
		row++
	}
	log.Println("Score is ", sum_points)
	return sum_points
}

func part2(input string) int {
	row_ix := 0
	win_arr := []NumColl{}
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		i := strings.IndexByte(line, ':')
		ss := line[i+2:]
		wt := strings.Split(ss, " | ")
		win_num := spaceStrToNumArray(wt[0])
		my_num := spaceStrToNumArray(wt[1])
		//fmt.Println("scratch: ", row, win_num, my_num)
		ww := NumColl{
			_serie: win_num,
			_id:    row_ix,
		}
		mm := NumColl{
			_serie: my_num,
			_id:    row_ix,
		}
		winner_item := mm.FindInto(ww)
		win_arr = append(win_arr, winner_item)
		fmt.Println("winner size", winner_item, len(winner_item._serie), row_ix)
		row_ix++
	}
	scr := Scratches{
		_freq_orig: make([]int, row_ix),
		_freq_copy: make([]int, row_ix),
	}
	for ix, vv := range win_arr {
		scr.IncCount(ix, vv)
	}
	tot := scr.Total()
	log.Println("Total scratch: ", tot)
	return tot
}
