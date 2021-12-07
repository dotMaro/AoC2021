package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	input := utils.InputString("day07/input.txt")
	positions := parse(input)
	fmt.Printf("Task 1. The cheapest linear fuel consumption is %d\n", findCheapestPosition(positions, linearCost))
	fmt.Printf("Task 2. The cheapest exponential fuel consumption is %d\n", findCheapestPosition(positions, exponentialCost))
}

func parse(s string) []int {
	split := strings.Split(s, ",")
	positions := make([]int, len(split))
	for i, s := range split {
		positions[i], _ = strconv.Atoi(s)
	}
	return positions
}

type costFunc func(positions []int, to int) uint

func findCheapestPosition(positions []int, cost costFunc) uint {
	const maxSearchPos = 2000
	var lowestCost uint = math.MaxUint64
	for i := 0; i < maxSearchPos; i++ {
		cost := cost(positions, i)
		if cost < lowestCost {
			lowestCost = cost
		}
	}
	return lowestCost
}

func linearCost(positions []int, to int) uint {
	var totalCost uint = 0
	for _, pos := range positions {
		cost := pos - to
		if cost < 0 {
			cost *= -1
		}
		totalCost += uint(cost)
	}
	return totalCost
}

func exponentialCost(positions []int, to int) uint {
	var totalCost uint = 0
	for _, pos := range positions {
		steps := pos - to
		if steps < 0 {
			steps *= -1
		}
		cost := steps * (steps + 1) / 2
		totalCost += uint(cost)
	}
	return totalCost
}
