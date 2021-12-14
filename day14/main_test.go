package main

import "testing"

const input = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

func Test_polymerFormula_mostCommonElementSubtractedWithLeastCommon(t *testing.T) {
	formula := parsePolymerFormula(input)
	afterTen := formula.stepN(10)
	res := afterTen.mostCommonElementSubtractedWithLeastCommon()
	if res != 1588 {
		t.Errorf("Should return 1588, but returned %d", res)
	}

	afterForty := formula.stepN(40)
	res = afterForty.mostCommonElementSubtractedWithLeastCommon()
	if res != 2188189693529 {
		t.Errorf("Should return 2188189693529, but returned %d", res)
	}
}
