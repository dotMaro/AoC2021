package main

import (
	"fmt"
	"strings"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	input := utils.InputString("day12/input.txt")
	caves := parse(input)
	fmt.Printf("Part 1. There are %d unique paths\n", findPathCount(caves))
	fmt.Printf("Part 2. There are %d unique paths when you can visit one small cave twice\n", findPathCountVisitOneSmallCaveTwice(caves))
}

type cave struct {
	name        string
	large       bool
	connections []string
}

func (c cave) String() string {
	var b strings.Builder
	for i, name := range c.connections {
		b.WriteString(name)
		if i < len(c.connections)-1 {
			b.WriteRune(',')
		}
	}
	return fmt.Sprintf("%s->%v", c.name, b.String())
}

func parse(s string) []cave {
	var caves []cave
	for _, line := range utils.SplitLine(s) {
		names := strings.SplitN(line, "-", 2)
		var start, end *cave
		var startIndex, endIndex int
		for i, c := range caves {
			switch c.name {
			case names[0]:
				start = &c
				startIndex = i
			case names[1]:
				end = &c
				endIndex = i
			}
		}
		if start == nil {
			start = &cave{
				name:        names[0],
				large:       names[0] == strings.ToUpper(names[0]),
				connections: []string{names[1]},
			}
			caves = append(caves, *start)
		} else {
			caves[startIndex].connections = append(caves[startIndex].connections, names[1])
		}
		if end == nil {
			end = &cave{
				name:        names[1],
				large:       names[1] == strings.ToUpper(names[1]),
				connections: []string{names[0]},
			}
			caves = append(caves, *end)
		} else {
			caves[endIndex].connections = append(caves[endIndex].connections, names[0])
		}
	}
	return caves
}

func findPathCount(caves []cave) int {
	var start, end cave
	for _, c := range caves {
		switch c.name {
		case "start":
			start = c
		case "end":
			end = c
		}
	}
	return traverse(caves, make(map[string]struct{}), start, end)
}

func traverse(caves []cave, alreadyTraversedMap map[string]struct{}, current, end cave) int {
	if !current.large {
		alreadyTraversedMap[current.name] = struct{}{}
	}
	if current.name == end.name {
		return 1
	}
	pathCount := 0
	for _, name := range current.connections {
		var c cave
		for _, c = range caves {
			if c.name == name {
				break
			}
		}
		_, alreadyTraversed := alreadyTraversedMap[c.name]
		if c.large || !alreadyTraversed {
			newMap := make(map[string]struct{})
			for k, v := range alreadyTraversedMap {
				newMap[k] = v
			}
			pathCount += traverse(caves, newMap, c, end)
		}
	}
	return pathCount
}

func findPathCountVisitOneSmallCaveTwice(caves []cave) int {
	var start, end cave
	for _, c := range caves {
		switch c.name {
		case "start":
			start = c
		case "end":
			end = c
		}
	}
	traverseMap := make(map[string]uint8)
	traverseMap["start"] = 2
	return traverseVisitOneSmallCaveTwice(caves, traverseMap, start, end, false)
}

func traverseVisitOneSmallCaveTwice(caves []cave, alreadyTraversedMap map[string]uint8, current, end cave, hasTraversedTwice bool) int {
	if !current.large {
		alreadyTraversedMap[current.name] = alreadyTraversedMap[current.name] + 1
		if alreadyTraversedMap[current.name] == 2 {
			hasTraversedTwice = true
		}
	}
	if current.name == end.name {
		return 1
	}
	pathCount := 0
	for _, name := range current.connections {
		var c cave
		for _, c = range caves {
			if c.name == name {
				break
			}
		}
		traverseCount := alreadyTraversedMap[c.name]
		if c.large || !hasTraversedTwice && traverseCount < 2 || hasTraversedTwice && traverseCount < 1 {
			copiedMap := make(map[string]uint8)
			for k, v := range alreadyTraversedMap {
				copiedMap[k] = v
			}
			pathCount += traverseVisitOneSmallCaveTwice(caves, copiedMap, c, end, hasTraversedTwice)
		}
	}
	return pathCount
}
