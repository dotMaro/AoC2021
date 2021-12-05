package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	splitInput := utils.SplitInput("day05/input.txt")
	vents := parseVents(splitInput, false)
	ventsWithDiagonal := parseVents(splitInput, true)
	fmt.Printf("Task 1. There are %d coordinates with two or more vents\n", vents.twoOrMoreCount())
	fmt.Printf("Task 2. There are %d coordinates with two or more vents, if you include diagonal lines\n", ventsWithDiagonal.twoOrMoreCount())
}

type hydrothermalVents [][]int

func parseVents(lines []string, includeDiagonal bool) hydrothermalVents {
	const size = 999
	vents := make([][]int, size)
	for _, line := range lines {
		coordPairs := strings.SplitN(line, " -> ", 2)
		startCoords := strings.SplitN(coordPairs[0], ",", 2)
		endCoords := strings.SplitN(coordPairs[1], ",", 2)

		x1, _ := strconv.Atoi(startCoords[0])
		y1, _ := strconv.Atoi(startCoords[1])
		x2, _ := strconv.Atoi(endCoords[0])
		y2, _ := strconv.Atoi(endCoords[1])

		if x1 == x2 {
			for y := min(y1, y2); y <= max(y1, y2); y++ {
				if vents[y] == nil {
					vents[y] = make([]int, size)
				}
				vents[y][x1]++
			}
		} else if y1 == y2 {
			if vents[y1] == nil {
				vents[y1] = make([]int, size)
			}
			for x := min(x1, x2); x <= max(x1, x2); x++ {
				vents[y1][x]++
			}
		} else if includeDiagonal {
			// Diagonal line.
			// Start at x1, y1 and go towards x2, y2.
			x1Lower := x1 < x2
			y1Lower := y1 < y2
			goTowardsEnd := func(x, y *int) {
				if x1Lower {
					*x++
				} else {
					*x--
				}
				if y1Lower {
					*y++
				} else {
					*y--
				}
			}
			for x, y := x1, y1; y1Lower && y <= y2 || !y1Lower && y >= y2; goTowardsEnd(&x, &y) {
				if vents[y] == nil {
					vents[y] = make([]int, size)
				}
				vents[y][x]++
			}
		}
	}
	return vents
}

func (v hydrothermalVents) twoOrMoreCount() int {
	count := 0
	for _, row := range v {
		for _, ventCount := range row {
			if ventCount >= 2 {
				count++
			}
		}
	}
	return count
}

func (v hydrothermalVents) String() string {
	var b strings.Builder
	b.WriteRune('\n')
	for _, row := range v {
		for _, ventCount := range row {
			if ventCount > 0 {
				b.WriteString(fmt.Sprint(ventCount))
			} else {
				b.WriteRune('.')
			}
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
