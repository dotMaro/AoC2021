package main

import (
	"testing"
)

const input = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func Test_point_fold(t *testing.T) {
	grid := parseGrid(input)
	t.Log(grid.GridString(11, 15))
	foldedGrid := grid.firstFold()
	res := foldedGrid.pointsCount()
	if res != 17 {
		t.Log(foldedGrid.GridString(11, 7))
		t.Errorf("Should return 17, not %d", res)
	}
}
