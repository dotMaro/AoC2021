package main

import (
	"fmt"
	"strconv"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	splitInput := utils.SplitInput("day01/input.txt")

	increaseCount := increaseCount(splitInput)
	fmt.Printf("Task 1. There were %d measurements that increased\n", increaseCount)
	threeSumIncreaseCount := threeSumIncreaseCount(splitInput)
	fmt.Printf("Task 2. There were %d three-sum measurements that increased\n", threeSumIncreaseCount)
}

func increaseCount(s []string) int {
	count := 0
	lastVal, _ := strconv.Atoi(s[0])
	for _, line := range s[1:] {
		curVal, _ := strconv.Atoi(line)
		if curVal > lastVal {
			count++
		}
		lastVal = curVal
	}
	return count
}

func threeSumIncreaseCount(s []string) int {
	count := 0
	val1, _ := strconv.Atoi(s[0])
	val2, _ := strconv.Atoi(s[1])
	lastSum := 99999999
	for _, line := range s[2:] {
		val3, _ := strconv.Atoi(line)
		sum := val1 + val2 + val3
		if sum > lastSum {
			count++
		}
		val1 = val2
		val2 = val3
		lastSum = sum
	}
	return count
}
