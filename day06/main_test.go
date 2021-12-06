package main

import "testing"

const input = `3,4,3,1,2`

func Test_fishTimers_tick(t *testing.T) {
	timers := parseFishTimers(input)
	timers = timers.tickN(18)
	res := len(timers)
	if res != 26 {
		t.Errorf("After 18 days there should be 26 fish, not %d", res)
	}
}

func Test_optimizedfishTimers_tick(t *testing.T) {
	timers := parseOptimizedFishTimers(input)
	timers = timers.tickN(18)
	res := timers.total()
	if res != 26 {
		t.Errorf("After 18 days there should be 26 fish, not %d", res)
	}
}
