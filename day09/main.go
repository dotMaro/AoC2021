package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	splitInput := utils.SplitInput("day09/input.txt")
	heightMap := parseHeightMap(splitInput)
	lowPoints := heightMap.lowPoints()
	fmt.Printf("Task 1. The risk of the low points is %d\n", heightMap.risk(lowPoints))
	fmt.Printf("Task 2. The product of the three largest basins is %d\n", heightMap.largestBasins(lowPoints))
}

type heightMap [][]int

type coord struct {
	y, x int
}

func parseHeightMap(lines []string) heightMap {
	heightMap := make([][]int, len(lines))
	for row, line := range lines {
		heightMap[row] = make([]int, len(line))
		for col, r := range line {
			heightMap[row][col], _ = strconv.Atoi(string(r))
		}
	}
	return heightMap
}

func (m heightMap) lowPoints() []coord {
	var lowPoints []coord
	for y, row := range m {
		for x := range row {
			if m.lowerThanNeighbors(y, x) {
				lowPoints = append(lowPoints, coord{y: y, x: x})
			}
		}
	}
	return lowPoints
}

func (m heightMap) risk(points []coord) int {
	total := 0
	for _, p := range points {
		total += m[p.y][p.x] + 1
	}
	return total
}

func (m heightMap) lowerThanNeighbors(y, x int) bool {
	cur := m[y][x]

	if y > 0 {
		if m[y-1][x] <= cur {
			return false
		}
	}
	if y < len(m)-1 {
		if m[y+1][x] <= cur {
			return false
		}
	}

	if x > 0 {
		if m[y][x-1] <= cur {
			return false
		}
	}
	if x < len(m[y])-1 {
		if m[y][x+1] <= cur {
			return false
		}
	}

	return true
}

func (m heightMap) largestBasins(points []coord) int {
	basins := make([]int, len(points))
	for i, p := range points {
		basins[i] = m.basin(p.y, p.x)
	}
	sort.Ints(basins)
	len := len(basins)
	return basins[len-1] * basins[len-2] * basins[len-3]
}

func (m heightMap) basin(y, x int) int {
	return m.basinTraverse(y, x, make(map[coord]struct{}))
}

func (m heightMap) basinTraverse(y, x int, alreadyVisited map[coord]struct{}) int {
	if y < 0 || y >= len(m) || x < 0 || x >= len(m[y]) {
		// Coords out of bounds.
		return 0
	}
	if m[y][x] == 9 {
		return 0
	}
	coord := coord{y: y, x: x}
	if _, isAlreadyVisited := alreadyVisited[coord]; isAlreadyVisited {
		return 0
	}
	alreadyVisited[coord] = struct{}{}
	return 1 + m.basinTraverse(y-1, x, alreadyVisited) + m.basinTraverse(y+1, x, alreadyVisited) +
		m.basinTraverse(y, x-1, alreadyVisited) + m.basinTraverse(y, x+1, alreadyVisited)
}
