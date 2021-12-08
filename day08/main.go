package main

import (
	"fmt"
	"strings"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	input := utils.SplitInput("day08/input.txt")
	fmt.Printf("Task 1. There are %d ones, fours, sevens and eights\n", countUniqueNumbers(input))
	fmt.Printf("Task 2. The sum of all outputs is %d\n", sumOfOutputs(input))
}

func countUniqueNumbers(s []string) int {
	count := 0
	for _, line := range s {
		split := strings.Split(line, " | ")
		splitOutput := strings.Split(split[1], " ")
		for _, output := range splitOutput {
			len := len(output)
			if len == 2 || len == 3 || len == 4 || len == 7 {
				count++
			}
		}
	}
	return count
}

func sumOfOutputs(s []string) int {
	sum := 0
	for _, line := range s {
		sum += decodeOutput(line)
	}
	return sum
}

func decodeOutput(line string) int {
	split := strings.Split(line, " | ")
	signals := strings.Split(split[0], " ")
	var one, four, seven, eight string
	for _, signal := range signals {
		switch len(signal) {
		case 2:
			one = signal
		case 3:
			seven = signal
		case 4:
			four = signal
		case 7:
			eight = signal
		}
	}

	three := findLenAndHasAllSegments(5, one, signals, nil)
	nine := findLenAndHasAllSegments(6, three, signals, nil)
	zero := findLenAndHasAllSegments(6, one, signals, &nine)
	six := findLen(6, signals, []string{zero, nine})
	f := findCommonSegment(six, one)
	var five string
	for _, s := range signals {
		if len(s) == 5 && s != three && hasSegment(s, f) {
			five = s
			break
		}
	}
	two := findLen(5, signals, []string{five, three})

	res := 0
	for _, output := range strings.Split(split[1], " ") {
		match := findMatch(output, zero, one, two, three, four, five, six, seven, eight, nine)
		res = res*10 + match
	}

	return res
}

func findMatch(output string, zero, one, two, three, four, five, six, seven, eight, nine string) int {
	switch {
	case areTheSame(output, zero):
		return 0
	case areTheSame(output, one):
		return 1
	case areTheSame(output, two):
		return 2
	case areTheSame(output, three):
		return 3
	case areTheSame(output, four):
		return 4
	case areTheSame(output, five):
		return 5
	case areTheSame(output, six):
		return 6
	case areTheSame(output, seven):
		return 7
	case areTheSame(output, eight):
		return 8
	case areTheSame(output, nine):
		return 9
	}
	return -1
}

func areTheSame(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, r1 := range a {
		hasR1 := false
		for _, r2 := range b {
			if r1 == r2 {
				hasR1 = true
				break
			}
		}
		if !hasR1 {
			return false
		}
	}
	return true
}

func hasSegment(s string, segment rune) bool {
	for _, r := range s {
		if r == segment {
			return true
		}
	}
	return false
}

func findCommonSegment(a, b string) rune {
	for _, r1 := range a {
		for _, r2 := range b {
			if r2 == r1 {
				return r2
			}
		}
	}
	return ' '
}

func findLen(length int, signals []string, exclude []string) string {
signalLoop:
	for _, s := range signals {
		if len(s) != length {
			continue
		}
		for _, e := range exclude {
			if s == e {
				continue signalLoop
			}
		}
		return s
	}
	return ""
}

func findLenAndHasAllSegments(length int, segs string, signals []string, exclude *string) string {
	for _, s := range signals {
		if len(s) != length {
			continue
		}
		if exclude != nil && s == *exclude {
			continue
		}
		hasAllSegments := true
		for _, r := range segs {
			hasSegment := hasSegment(s, r)
			if !hasSegment {
				hasAllSegments = false
				continue
			}
		}
		if hasAllSegments {
			return s
		}
	}
	return ""
}
