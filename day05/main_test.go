package main

import (
	"testing"

	"github.com/dotMaro/AoC2021/utils"
)

const input = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func Test_twoOrMoreCount(t *testing.T) {
	vents := parseVents(utils.SplitLine(input), false)
	t.Log(vents)
	res := vents.twoOrMoreCount()
	if res != 5 {
		t.Errorf("Should return 5, but returned %d", res)
	}
	vents = parseVents(utils.SplitLine(input), true)
	t.Log(vents)
	res = vents.twoOrMoreCount()
	if res != 12 {
		t.Errorf("Should return 12, but returned %d", res)
	}
}
