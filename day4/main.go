// skeleton template copied and modified from https://github.com/alexchao26/advent-of-code-go/blob/main/scripts/skeleton/tmpls/main.go

// started        ;
// finished part1 , 'go run' time s, run time after 'go build' s
// finished part2 , 'go run' time s, run time after 'go build' s

package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		i := strings.IndexByte(line, ':')
		ss := line[i+2:]
		wt := strings.Split(ss, " | ")
		win_num := spaceStrToNumArray(wt[0])
		my_num := spaceStrToNumArray(wt[1])
		fmt.Println("scratch: ", row, win_num, my_num)
		// scratch := strings.Split(line[i+2:], " | ")
		// fmt.Println(scratch, len(scratch))
		row++
	}

	return 0
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (parsedInput []int) {
	for _, line := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, stringToInt(line))
	}
	return parsedInput
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}
