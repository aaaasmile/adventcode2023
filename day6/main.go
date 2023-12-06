package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
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
	flag.IntVar(&part, "part", 1, "part 1 or 2")
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
		ans = part1(input, test)
	case 2:
		ans = part2(input)
	}
	fmt.Println("Output:", ans)
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("total call duration: %v\n", elapsed)
}

type Record struct {
	_duration    int
	_recordDist  int
	_minLoopTime int
	_maxLoopTime int
	_id          int
}

func (rc *Record) CalcDiffWayToBeatRecord() int {
	speed_unit := 1
	wins := 0
	win_hist := []int{}
	dist_hist := []int{}
	for i := rc._minLoopTime; i <= rc._maxLoopTime; i++ {
		hold_time := i
		speed := hold_time * speed_unit
		dist := speed * (rc._duration - hold_time)
		if dist > rc._recordDist {
			//fmt.Printf("[Race %d]: hold time %d beat the record with %d \n", rc._id, hold_time, dist)
			wins++
			win_hist = append(win_hist, 1)
		} else {
			win_hist = append(win_hist, 0)
		}
		dist_hist = append(dist_hist, dist)
	}
	fmt.Printf("[Race %d]: different ways to win (duration %d, record %d, min_hold %d, max_hold %d): %d\n",
		rc._id, rc._duration, rc._recordDist, rc._minLoopTime, rc._maxLoopTime, wins)
	fmt.Println("hist win: ", win_hist, len(win_hist))
	fmt.Println("hist dis: ", dist_hist, len(dist_hist))

	return wins
}

func part1(input string, isTest bool) int {
	races := []Record{}
	times := []int{}
	dist := []int{}
	for ix, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		fmt.Println(line)
		if ix == 0 {
			tt := strings.Split(line, ":")
			times = spaceStrToNumArray(tt[1])
		} else if ix == 1 {
			dd := strings.Split(line, ":")
			dist = spaceStrToNumArray(dd[1])
		}
	}
	for ix := range times {
		races = append(races, Record{_duration: times[ix], _recordDist: dist[ix], _id: ix + 1, _maxLoopTime: times[ix]})
	}
	fmt.Println("records: ", races)
	ways := []int{}
	mult := 1
	if isTest {
		races[1]._minLoopTime = 4
		races[1]._maxLoopTime = 11
		races[2]._minLoopTime = 11
		races[2]._maxLoopTime = 19
	}
	for _, rr := range races {
		ww := rr.CalcDiffWayToBeatRecord()
		ways = append(ways, ww)
		mult *= ww
	}
	log.Println("Number of ways you can beat the record", mult)
	return mult
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
