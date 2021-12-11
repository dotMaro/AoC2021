package main

import "testing"

const input = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

func Test_octopusMap_stepN(t *testing.T) {
	m := parseOctopusMap(input)
	res := m.stepN(100)
	if res != 1656 {
		t.Errorf("Should return 1656, but returned %d", res)
	}
}
