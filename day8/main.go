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

func (t *Turn) String() string {
	if *t == Left {
		return "L"
	} else if *t == Right {
		return "R"
	}
	panic("not recognized")
}

type LRItem struct {
	_leftKey  string
	_rightKey string
}

type TurnArr []Turn

type GuideMap struct {
	_instr     TurnArr
	_guideBook map[string]LRItem
}

func (tt *TurnArr) String() string {
	str := ""
	for _, r := range *tt {
		if str == "" {
			str += r.String()
		} else {
			str = fmt.Sprintf("%s %s", str, r.String())
		}

	}
	return str
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
			continue
		}
		if line == "" {
			continue
		}
		arr := strings.Split(line, "=")
		kk := strings.TrimSpace(arr[0])
		arg := strings.TrimSpace(arr[1])
		arg = strings.Replace(arg, "(", "", -1)
		arg = strings.Replace(arg, ")", "", -1)
		keysarr := strings.Split(arg, ",")
		lri := LRItem{
			_leftKey:  strings.TrimSpace(keysarr[0]),
			_rightKey: strings.TrimSpace(keysarr[1]),
		}
		gm._guideBook[kk] = lri
	}
	fmt.Println(gm._instr.String())
	fmt.Println(gm._guideBook)
	return 0
}

func part2(input string) int {
	return 0
}
