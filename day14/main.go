package main

import (
	"fmt"
	"strings"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	input := utils.InputString("day14/input.txt")
	fmt.Println(parsePolymerFormula(input).stepN(40).mostCommonElementSubtractedWithLeastCommon())
}

type polymerFormula struct {
	pairs            map[string]uint
	rightmostElement byte
	rules            map[string]byte
}

func parsePolymerFormula(s string) polymerFormula {
	lines := utils.SplitLine(s)
	template := lines[0]
	pairs := make(map[string]uint)
	for i := 1; i < len(template); i++ {
		pairs[template[i-1:i+1]] = pairs[template[i-1:i+1]] + 1
	}
	rules := make(map[string]byte, len(lines)-2)
	for _, line := range lines[2:] {
		split := strings.SplitN(line, " -> ", 2)
		rules[split[0]] = split[1][0]
	}
	return polymerFormula{
		pairs:            pairs,
		rightmostElement: template[len(template)-1],
		rules:            rules,
	}
}

func (f polymerFormula) stepN(n int) polymerFormula {
	for i := 0; i < n; i++ {
		f = f.step()
	}
	return f
}

func (f polymerFormula) step() polymerFormula {
	newPairs := make(map[string]uint)

	for pair, count := range f.pairs {
		insert := f.rules[pair]
		newPairs[string(pair[0])+string(insert)] = newPairs[string(pair[0])+string(insert)] + count
		newPairs[string(insert)+string(pair[1])] = newPairs[string(insert)+string(pair[1])] + count
	}

	return polymerFormula{
		pairs:            newPairs,
		rightmostElement: f.rightmostElement,
		rules:            f.rules,
	}
}

func (f polymerFormula) mostCommonElementSubtractedWithLeastCommon() uint {
	elementCount := make(map[byte]uint)
	for pair, count := range f.pairs {
		elementCount[pair[0]] = elementCount[pair[0]] + count
	}
	elementCount[f.rightmostElement] = elementCount[f.rightmostElement] + 1

	var mostCommon, leastCommon uint = 0, 999999999999999
	for _, count := range elementCount {
		if count > mostCommon {
			mostCommon = count
		}
		if count < leastCommon {
			leastCommon = count
		}
	}
	return mostCommon - leastCommon
}
