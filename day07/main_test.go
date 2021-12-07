package main

import (
	"testing"

	"github.com/dotMaro/AoC2021/utils"
)

const input = `16,1,2,0,4,2,7,1,2,14`

func Test_findCheapestPosition(t *testing.T) {
	positions := parse(input)
	res := findCheapestPosition(positions, linearCost)
	if res != 37 {
		t.Errorf("Should return 2, not %d", res)
	}

	res = findCheapestPosition(positions, exponentialCost)
	if res != 168 {
		t.Errorf("Should return 168, not %d", res)
	}
}

func Benchmark_findCheapestPosition_exponential(b *testing.B) {
	input := utils.InputString("input.txt")
	positions := parse(input)
	for i := 0; i < b.N; i++ {
		findCheapestPosition(positions, exponentialCost)
	}
}
