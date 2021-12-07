package main

import "testing"

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
