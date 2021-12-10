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
		error, _ := firstError(line)
		sum += error
	}
	return sum
}

func corruptedLinesMiddleScore(lines []string) int {
	var scores []int
	for _, line := range lines {
		error, closing := firstError(line)
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

func firstError(s string) (int, []rune) {
	paranthesesCount := 0
	bracketCount := 0
	curledCount := 0
	angleCount := 0
	var expClosing []rune
	for _, r := range s {
		switch r {
		case '(':
			paranthesesCount++
			expClosing = append(expClosing, ')')
		case ')':
			if expClosing[len(expClosing)-1] != r {
				return 3, expClosing
			}
			expClosing = expClosing[:len(expClosing)-1]
			paranthesesCount--
			if paranthesesCount < 0 {
				return 3, expClosing
			}
		case '[':
			bracketCount++
			expClosing = append(expClosing, ']')
		case ']':
			if expClosing[len(expClosing)-1] != r {
				return 57, expClosing
			}
			expClosing = expClosing[:len(expClosing)-1]
			bracketCount--
			if bracketCount < 0 {
				return 57, expClosing
			}
		case '{':
			curledCount++
			expClosing = append(expClosing, '}')
		case '}':
			if expClosing[len(expClosing)-1] != r {
				return 1197, expClosing
			}
			expClosing = expClosing[:len(expClosing)-1]
			curledCount--
			if curledCount < 0 {
				return 1197, expClosing
			}
		case '<':
			angleCount++
			expClosing = append(expClosing, '>')
		case '>':
			if expClosing[len(expClosing)-1] != r {
				return 25137, expClosing
			}
			expClosing = expClosing[:len(expClosing)-1]
			angleCount--
			if angleCount < 0 {
				return 25137, expClosing
			}
		}
	}
	return 0, expClosing
}
