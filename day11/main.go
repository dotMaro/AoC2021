package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	m := parseOctopusMap(utils.InputString("day11/input.txt"))
	fmt.Printf("Part 1. After 100 steps there have been %d flashes\n", m.stepN(100))
	fmt.Printf("Part 2. After %d steps all octopuses flash at the same time\n", m.findSynchronization())
}

type octopusMap [][]int

type coord struct {
	y, x int
}

func parseOctopusMap(s string) octopusMap {
	lines := utils.SplitLine(s)
	m := make([][]int, len(lines))
	for row, line := range lines {
		m[row] = make([]int, len(line))
		for col, level := range line {
			m[row][col], _ = strconv.Atoi(string(level))
		}
	}
	return m
}

func (m octopusMap) findSynchronization() int {
	step := 1
	for {
		flashCount := m.step()
		if flashCount == len(m)*len(m[0]) {
			return step
		}
		step++
	}
}

func (m octopusMap) stepN(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += m.step()
	}
	return total
}

func (m octopusMap) step() int {
	flashes := make(map[coord]struct{})
	m.increaseEnergyLevel()
	hadNewFlash := true
	for hadNewFlash {
		hadNewFlash = m.findFlashes(flashes)
	}
	m.resetFlashes(flashes)
	return len(flashes)
}

func (m octopusMap) increaseEnergyLevel() {
	for y, row := range m {
		for x := range row {
			m[y][x]++
		}
	}
}

func (m octopusMap) resetFlashes(flashes map[coord]struct{}) {
	for coord := range flashes {
		m[coord.y][coord.x] = 0
	}
}

func (m octopusMap) findFlashes(flashes map[coord]struct{}) bool {
	hadNewFlash := false
	for y, row := range m {
		for x, level := range row {
			_, alreadyFlashed := flashes[coord{y: y, x: x}]
			if level > 9 && !alreadyFlashed {
				hadNewFlash = true
				flashes[coord{y: y, x: x}] = struct{}{}
				m.flash(y, x)
			}
		}
	}
	return hadNewFlash
}

func (m octopusMap) flash(y, x int) {
	// Make sure curY and curX stay within bounds.
	minY, maxY := max(0, y-1), min(len(m)-1, y+1)
	minX, maxX := max(0, x-1), min(len(m[0])-1, x+1)
	for curY := minY; curY <= maxY; curY++ {
		for curX := minX; curX <= maxX; curX++ {
			m[curY][curX]++
		}
	}
}

func (m octopusMap) String() string {
	var b strings.Builder
	for _, row := range m {
		for _, level := range row {
			b.WriteString(fmt.Sprint(level))
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
