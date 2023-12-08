package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
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
		ans = part1(input)
	case 2:
		ans = part2(input)
	}
	fmt.Println("Output:", ans)
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("total call duration: %v\n", elapsed)
}

type Turn int

const (
	Left = iota
	Right
)

type LRItem struct {
	_left  Turn
	_right Turn
}

type GuideMap struct {
	_instr     []Turn
	_guideBook map[string]LRItem
}

func part1(input string) int {
	gm := GuideMap{
		_instr:     []Turn{},
		_guideBook: map[string]LRItem{},
	}
	for ix, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		//fmt.Println(line)
		if ix == 0 {
			for _, ll := range line {
				if ll == 'L' {
					gm._instr = append(gm._instr, Left)
				} else if ll == 'R' {
					gm._instr = append(gm._instr, Right)
				}
			}
		}
	}
	fmt.Println(gm._instr)
	return 0
}

func part2(input string) int {
	return 0
}
