package main

import (
	"fmt"
	"strconv"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	splitInput := utils.SplitInput("day09/input.txt")
	heightMap := parseHeightMap(splitInput)
	fmt.Println(heightMap.lowPoints())
}

type heightMap [][]int

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

func (m heightMap) lowPoints() int {
	// var lowPoints []int
	total := 0
	for y, row := range m {
		for x, height := range row {
			if m.lowerThanNeighbors(y, x) {
				fmt.Printf("Point %d %d is low, height %d\n", x, y, height)
				// lowPoints = append(lowPoints, height+1)
				total += height + 1
			}
		}
	}
	// total := 0
	// for _, risk := range lowPoints {
	// 	total += risk
	// }
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

func (m heightMap) baisin(y, x int) int {
	return 0
}
