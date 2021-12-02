package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	splitInput := utils.SplitInput("day02/input.txt")
	fmt.Printf("Day 1. Horizontal multiplied with depth is %d\n", traverse(splitInput))
	fmt.Printf("Day 2. Horizontal multiplied with depth is %d\n", traverseWithAim(splitInput))
}

func traverse(lines []string) int {
	horizontal := 0
	depth := 0
	for _, line := range lines {
		words := strings.Split(line, " ")
		direction := words[0]
		length, _ := strconv.Atoi(words[1])
		switch direction {
		case "forward":
			horizontal += length
		case "up":
			depth -= length
		case "down":
			depth += length
		}
	}

	return horizontal * depth
}

func traverseWithAim(lines []string) int {
	horizontal := 0
	depth := 0
	aim := 0
	for _, line := range lines {
		words := strings.Split(line, " ")
		direction := words[0]
		length, _ := strconv.Atoi(words[1])
		switch direction {
		case "forward":
			horizontal += length
			depth += aim * length
		case "up":
			aim -= length
		case "down":
			aim += length
		}
	}

	return horizontal * depth
}
