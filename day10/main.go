package main

import (
	"fmt"
	"sort"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	lines := utils.SplitInput("day10/input.txt")
	fmt.Printf("Part 1. The syntax error score is %d\n", firstErrorSum(lines))
	fmt.Printf("Part 2. The middle score is %d\n", corruptedLinesMiddleScore(lines))
}

func firstErrorSum(lines []string) int {
	sum := 0
	for _, line := range lines {
		error, _ := parseUntilFirstError(line)
		sum += error
	}
	return sum
}

func corruptedLinesMiddleScore(lines []string) int {
	var scores []int
	for _, line := range lines {
		error, closing := parseUntilFirstError(line)
		if error == 0 {
			score := 0
			// Reverse closing.
			for i, j := 0, len(closing)-1; i < j; i, j = i+1, j-1 {
				closing[i], closing[j] = closing[j], closing[i]
			}
			for _, c := range closing {
				score *= 5
				switch c {
				case ')':
					score++
				case ']':
					score += 2
				case '}':
					score += 3
				case '>':
					score += 4
				}
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[(len(scores)-1)/2]
}

// parseUntilFirstError and return the error score along with the remaining expected closing brackets.
// If the score is returned as 0 then no corruption was encountered but there will still always be a
// non-empty expected closing brackets slice returned.
func parseUntilFirstError(s string) (int, []rune) {
	matchingClosing := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	errorScore := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	var expClosing []rune
	for _, r := range s {
		if closing, hasMatchingClosing := matchingClosing[r]; hasMatchingClosing {
			// r is an opening bracket.
			expClosing = append(expClosing, closing)
		} else if expClosing[len(expClosing)-1] != r {
			// r is an unexpected closing bracket.
			return errorScore[r], expClosing
		} else {
			// r is an expected closing bracket, pop the stack.
			expClosing = expClosing[:len(expClosing)-1]
		}
	}
	return 0, expClosing
}
