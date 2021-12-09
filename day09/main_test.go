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

func Test_heightMap_risk(t *testing.T) {
	heightMap := parseHeightMap(utils.SplitLine(input))
	points := heightMap.lowPoints()
	res := heightMap.risk(points)
	if res != 15 {
		t.Errorf("Should return 15, but returned %d", res)
	}
}

func Test_heightMap_largestBaisins(t *testing.T) {
	heightMap := parseHeightMap(utils.SplitLine(input))
	coords := heightMap.lowPoints()
	res := heightMap.largestBaisins(coords)
	if res != 1134 {
		t.Errorf("Should return 1134, but returned %d", res)
	}
}
