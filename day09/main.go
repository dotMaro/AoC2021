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
	fmt.Println(heightMap.largestBaisins(lowPoints))
	// fmt.Println(heightMap.lowPoints().risk())
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
		for x, height := range row {
			if m.lowerThanNeighbors(y, x) {
				fmt.Printf("Point %d %d is low, height %d\n", x, y, height)
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

func (m heightMap) largestBaisins(points []coord) int {
	baisins := make([]int, len(points))
	for i, p := range points {
		baisins[i] = m.baisin(p.y, p.x)
	}
	sort.Ints(baisins)
	len := len(baisins)
	return baisins[len-1] * baisins[len-2] * baisins[len-3]
}

func (m heightMap) baisin(y, x int) int {
	return m.baisinTraverse(y, x, make(map[coord]struct{}))
}

func (m heightMap) baisinTraverse(y, x int, alreadyVisited map[coord]struct{}) int {
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
	return 1 + m.baisinTraverse(y-1, x, alreadyVisited) + m.baisinTraverse(y+1, x, alreadyVisited) +
		m.baisinTraverse(y, x-1, alreadyVisited) + m.baisinTraverse(y, x+1, alreadyVisited)
}
