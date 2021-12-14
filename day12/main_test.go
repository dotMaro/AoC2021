package main

import "testing"

const input1 = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

func Test_findPathCount(t *testing.T) {
	caves := parse(input1)
	res := findPathCount(caves)
	if res != 10 {
		t.Errorf("Should return 10, but returned %d", res)
	}
}

func Test_findPathCountVisitOneSmallCaveTwice(t *testing.T) {
	caves := parse(input1)
	res := findPathCountVisitOneSmallCaveTwice(caves)
	if res != 36 {
		t.Errorf("Should return 36, but returned %d", res)
	}
}
