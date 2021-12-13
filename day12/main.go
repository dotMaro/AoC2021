package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	input := utils.InputString("day12/input.txt")
	grid := parseGrid(input)
	fmt.Printf("Task 1. There are %d points after doing the first fold\n", grid.firstFold().pointsCount())
	fmt.Printf("Task 2. The resulting points look like:\n%s\n", grid.foldAll().GridString(39, 6))
}

func parseGrid(s string) grid {
	lines := utils.SplitLine(s)
	points := make([]point, 0, len(lines))
	var folds []fold
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "fold along ") {
			value, _ := strconv.Atoi(line[13:])
			fold := fold{
				axis:  toAxis(rune(line[11])),
				value: value,
			}
			folds = append(folds, fold)
			continue
		}
		coords := strings.SplitN(line, ",", 2)
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		points = append(points, point{
			x: x,
			y: y,
		})
	}

	return grid{
		points: points,
		folds:  folds,
	}
}

type grid struct {
	points []point
	folds  []fold
}

type point struct {
	x, y int
}

func (p point) shouldFold(f fold) (onFold, shouldFold bool) {
	switch f.axis {
	case xAxis:
		return p.x == f.value, p.x >= f.value
	case yAxis:
		return p.y == f.value, p.y >= f.value
	}
	return false, false
}

func (p point) fold(f fold) point {
	switch f.axis {
	case xAxis:
		diff := p.x - f.value
		p.x -= diff * 2
	case yAxis:
		diff := p.y - f.value
		p.y -= diff * 2
	}
	return p
}

type fold struct {
	axis  axis
	value int
}

type axis rune

const (
	xAxis axis = 'x'
	yAxis axis = 'y'
)

func toAxis(r rune) axis {
	var axis axis
	switch r {
	case 'x':
		axis = xAxis
	case 'y':
		axis = yAxis
	}
	return axis
}

func (g grid) foldAll() grid {
	for _, fold := range g.folds {
		g = g.fold(fold)
	}
	return g
}

func (g grid) firstFold() grid {
	return g.fold(g.folds[0])
}

func (g grid) fold(fold fold) grid {
	var newPoints []point
	addedPoints := make(map[point]struct{})
	addIfNoOverlap := func(p point) {
		if _, overlap := addedPoints[p]; !overlap {
			addedPoints[p] = struct{}{}
			newPoints = append(newPoints, p)
		}
	}
	for _, p := range g.points {
		onFold, shouldFold := p.shouldFold(fold)
		if !onFold {
			if shouldFold {
				newPoint := p.fold(fold)
				addIfNoOverlap(newPoint)
			} else {
				addIfNoOverlap(p)
			}
		}
	}
	return grid{
		points: newPoints,
		folds:  g.folds,
	}
}

func (g grid) pointsCount() int {
	return len(g.points)
}

func (g grid) GridString(x, y int) string {
	var b strings.Builder
	b.WriteRune('\n')
	for curY := 0; curY < y; curY++ {
		for curX := 0; curX < x; curX++ {
			hasPoint := false
			for _, p := range g.points {
				if p.x == curX && p.y == curY {
					hasPoint = true
					break
				}
			}
			if hasPoint {
				b.WriteRune('#')
			} else {
				b.WriteRune('.')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}
