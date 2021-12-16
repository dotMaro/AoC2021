package main

import "testing"

const input = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

func Test_riskMap_aStar(t *testing.T) {
	riskMap := parseRiskMap(input)
	res := riskMap.aStar(coord{x: 0, y: 0}, coord{x: len(riskMap[0]), y: len(riskMap)})
	if res != 40 {
		t.Errorf("Should return 40, not %d", res)
	}
}

func Test_riskMap_dijkstras(t *testing.T) {
	riskMap := parseRiskMap(input)
	res := riskMap.dijkstras()
	if res != 40 {
		t.Errorf("Should return 40, not %d", res)
	}
}

func Test_riskMap_dijkstrasPriority(t *testing.T) {
	riskMap := parseRiskMap(input)
	res := riskMap.dijkstrasPriority()
	if res != 40 {
		t.Errorf("Should return 40, not %d", res)
	}
}

// func Test_riskMap_extend(t *testing.T) {
// 	riskMap := parseRiskMap(input)
// 	extended := riskMap.extend()
// 	t.Error(extended)
// }
