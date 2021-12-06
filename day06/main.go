package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	input := utils.InputString("day06/input.txt")
	timers1 := parseFishTimers(input)
	fmt.Printf("Task 1. The amount of fish after 80 days is %d\n", len(timers1.tickN(80)))
	timers2 := parseOptimizedFishTimers(input)
	fmt.Printf("Task 2. The amount of fish after 256 days is %d\n", timers2.tickN(256).total())
}

type fishTimers []int

func parseFishTimers(s string) fishTimers {
	split := strings.Split(s, ",")
	timers := make([]int, len(split))
	for i, time := range split {
		timers[i], _ = strconv.Atoi(time)
	}
	return timers
}

func (t fishTimers) tickN(n int) fishTimers {
	for i := 0; i < n; i++ {
		t = t.tick()
	}
	return t
}

func (t fishTimers) tick() fishTimers {
	newCount := 0
	for i, timer := range t {
		if timer > 0 {
			t[i] = timer - 1
		} else {
			newCount++
			t[i] = 6
		}
	}
	for i := 0; i < newCount; i++ {
		t = append(t, 8)
	}
	return t
}

type optimizedFishTimers [9]int

func parseOptimizedFishTimers(s string) optimizedFishTimers {
	split := strings.Split(s, ",")
	var timers [9]int
	for _, time := range split {
		timeNbr, _ := strconv.Atoi(time)
		timers[timeNbr]++
	}
	return timers
}

func (t optimizedFishTimers) tickN(n int) optimizedFishTimers {
	for i := 0; i < n; i++ {
		t = t.tick()
	}
	return t
}

func (t optimizedFishTimers) tick() optimizedFishTimers {
	var newTimers [9]int
	for i := 1; i < len(t); i++ {
		newTimers[i-1] += t[i]
	}
	newTimers[6] += t[0]
	newTimers[8] = t[0]
	return newTimers
}

func (t optimizedFishTimers) total() int {
	total := 0
	for _, count := range t {
		total += count
	}
	return total
}
