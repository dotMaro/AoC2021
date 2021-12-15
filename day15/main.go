package main

import (
	"math"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	input := utils.InputString("day15/input.txt")
}

const (
	width  = 30
	height = 30
)

type riskMap [][]uint8

type coord struct {
	x, y int
}

func parseRiskMap(s string) riskMap {
	lines := utils.SplitLine(s)
	riskMap := make([][]uint8, len(lines))
	for y, line := range lines {
		riskMap[y] = make([]uint8, len(line))
		for x, r := range line {
			riskMap[y][x] = r
		}
	}
	return riskMap
}

func (r riskMap) neighbors(c coord) []coord {
	var neighbors []coord
	if c.x > 0 {
		neighbors = append(neighbors, coord{x: c.x - 1, y: c.y})
	}
	if c.x < len(r[0])-1 {
		neighbors = append(neighbors, coord{x: c.x + 1, y: c.y})
	}
	if c.y > 0 {
		neighbors = append(neighbors, coord{x: c.x, y: c.y - 1})
	}
	if c.y < len(r)-1 {
		neighbors = append(neighbors, coord{x: c.x, y: c.y + 1})
	}
	return neighbors
}

func (r riskMap) h(c coord) int {
	return len(r) - c.y + len(r[0]) - c.x
}

func (r riskMap) reconstructPathCost(cameFrom map[coord]coord, current coord) {
	totalPath := 0
	for current, prev := range cameFrom {

	}
}

func (r riskMap) aStar(start, goal coord, h int) int {
	openSet := map[coord]struct{}{
		start: struct{}{},
	}

	cameFrom := make(map[coord]coord)

	gScore := map[coord]int{
		start: 0,
	}

	fScore := map[coord]int{
		start: 0, // ??
	}

	for len(openSet) != 0 {
		var lowestScoreNode coord
		lowestScore := math.MaxUint32
		for node := range openSet {
			score, hasScore := fScore[node]
			if hasScore && score < lowestScore {
				lowestScoreNode = node
				lowestScore = score
			}
		}
		current := lowestScoreNode
		if current == goal {
			// return
		}

		delete(openSet, current)
		for _, neighbor := range r.neighbors(current) {
			tentativeGScore := gScore[current] + int(r[neighbor.y][neighbor.x])
			if tentativeGScore < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + r.h(neighbor)
				_, neighborInOpenSet := openSet[neighbor]
				if !neighborInOpenSet {
					openSet[neighbor] = struct{}{}
				}
			}
		}
	}
	return 0
}
