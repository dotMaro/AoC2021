package main

import (
	"testing"

	"github.com/dotMaro/AoC2021/utils"
)

const input = `2199943210
3987894921
9856789892
8767896789
9899965678`

func Test_heightMap_lowPoints(t *testing.T) {
	heightMap := parseHeightMap(utils.SplitLine(input))
	res := heightMap.lowPoints()
	if res != 15 {
		t.Errorf("Should return 15, but returned %d", res)
	}
	// t.Fail()
}
