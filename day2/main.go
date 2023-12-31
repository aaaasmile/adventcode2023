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

func part1(input string) int {
	roundNum := 1
	roundSum := 0

	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		i := strings.IndexByte(line, ':')
		redCount := 0
		greenCount := 0
		blueCount := 0
		game := strings.Split(line[i+2:], "; ")
		gameValid := true
		
		for _, round := range game {
			singleRound := strings.Split(round, ", ")
			for _, pull := range singleRound {
				singlePull := strings.Split(pull, " ")
				count := singlePull[0]
				color := singlePull[1]

				switch color {
				case "red":
					redCount = stringToInt(count)
				case "green":
					greenCount = stringToInt(count)
				case "blue":
					blueCount = stringToInt(count)
				}

				if redCount > 12 || greenCount > 13 || blueCount > 14 {
					gameValid = false
					break
				}
			}
		}
		
		if gameValid {
			roundSum += roundNum
		}
		roundNum++
	}

	return roundSum
}

func part2(input string) int {
	ans := 0

	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		i := strings.IndexByte(line, ':')
		biggestRedCount := 0
		biggestGreenCount := 0
		biggestBlueCount := 0
		game := strings.Split(line[i+2:], "; ")
		
		for _, round := range game {
			singleRound := strings.Split(round, ", ")
			for _, pull := range singleRound {
				singlePull := strings.Split(pull, " ")
				count := singlePull[0]
				color := singlePull[1]

				switch color {
				case "red":
					redCount := stringToInt(count)
					if redCount > biggestRedCount {
						biggestRedCount = redCount
					}
				case "green":
					greenCount := stringToInt(count)
					if greenCount > biggestGreenCount {
						biggestGreenCount = greenCount
					}
				case "blue":
					blueCount := stringToInt(count)
					if blueCount > biggestBlueCount {
						biggestBlueCount = blueCount
					}
				}

			}
		}
		ans += biggestRedCount * biggestGreenCount * biggestBlueCount 
	}

	return ans
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
