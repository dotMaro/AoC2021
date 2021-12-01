package main

import (
	"fmt"
	"strconv"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	splitInput := utils.SplitInput("day01/input.txt")

	increased := increaseCount(splitInput)
	fmt.Printf("Task 1. There were %d measurements that increased\n", increased)
	threeSumIncreaseCount := threeSumIncreaseCount(splitInput)
	fmt.Printf("Task 2. There were %d three-sum measurements that increased\n", threeSumIncreaseCount)
}

func increaseCount(s []string) int {
	increased := 0
	lastVal, _ := strconv.Atoi(s[0])
	for _, line := range s[1:] {
		curVal, _ := strconv.Atoi(line)
		if curVal > lastVal {
			increased++
		}
		lastVal = curVal
	}
	return increased
}

func threeSumIncreaseCount(s []string) int {
	increased := 0
	val1, _ := strconv.Atoi(s[0])
	val2, _ := strconv.Atoi(s[1])
	lastVal := 99999999
	for _, line := range s[2:] {
		val3, _ := strconv.Atoi(line)
		sum := val1 + val2 + val3
		if sum > lastVal {
			increased++
		}
		val1 = val2
		val2 = val3
		lastVal = sum
	}
	return increased
}
