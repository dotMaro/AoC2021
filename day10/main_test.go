package main

import (
	"testing"

	"github.com/dotMaro/AoC2021/utils"
)

const input = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func Test_firstErrorSum(t *testing.T) {
	lines := utils.SplitLine(input)
	res := firstErrorSum(lines)
	if res != 26397 {
		t.Errorf("Should return 26397, but returned %d", res)
	}
}

func Test_corruptedLines(t *testing.T) {
	lines := utils.SplitLine(input)
	res := corruptedLinesMiddleScore(lines)
	if res != 288957 {
		t.Errorf("Should return 288957, but returned %d", res)
	}
}
