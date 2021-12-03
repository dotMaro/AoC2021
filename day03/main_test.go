package main

import (
	"testing"

	"github.com/dotMaro/AoC2021/utils"
)

const input = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func Test_powerConsumption(t *testing.T) {
	splitInput := utils.SplitLine(input)
	res := powerConsumption(splitInput)
	if res != 198 {
		t.Errorf("Should return 198, not %d", res)
	}
}

func Test_lifeSupportRating(t *testing.T) {
	splitInput := utils.SplitLine(input)
	res := lifeSupportRating(splitInput)
	if res != 230 {
		t.Errorf("Should return 230, not %d", res)
	}
}
