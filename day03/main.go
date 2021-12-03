package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	splitInput := utils.SplitInput("day03/input.txt")
	fmt.Printf("Task 1. The power consumption is %d\n", powerConsumption(splitInput))
	fmt.Printf("Task 2. The life support rating is %d\n", lifeSupportRating(splitInput))
}

func powerConsumption(lines []string) uint {
	lineLen := len(lines[0])
	mostCommonBits := make([]int, lineLen)
	for _, line := range lines {
		for i, c := range line {
			if c == '1' {
				mostCommonBits[i]++
			}

		}
	}
	var gammaRateBuilder strings.Builder
	var epsilonRateBuilder strings.Builder
	for _, count := range mostCommonBits {
		if count > len(lines)/2 {
			gammaRateBuilder.WriteRune('1')
			epsilonRateBuilder.WriteRune('0')
		} else {
			gammaRateBuilder.WriteRune('0')
			epsilonRateBuilder.WriteRune('1')
		}
	}
	gammaRateString := gammaRateBuilder.String()
	gammaRate, err := strconv.ParseUint(gammaRateString, 2, lineLen)
	if err != nil {
		panic(err)
	}
	epsilonRate, err := strconv.ParseUint(epsilonRateBuilder.String(), 2, lineLen)
	if err != nil {
		panic(err)
	}
	return uint(gammaRate) * uint(epsilonRate)
}

func lifeSupportRating(lines []string) uint64 {
	return oxygenGeneratorRating(lines) * co2ScrubberRating(lines)
}

func oxygenGeneratorRating(lines []string) uint64 {
	filterByMostCommon := func(bit byte, mostCommonBit byte) bool {
		return bit == mostCommonBit
	}
	return filterLines(lines, filterByMostCommon)
}

func co2ScrubberRating(lines []string) uint64 {
	filterByLeastCommon := func(bit byte, mostCommonBit byte) bool {
		return bit != mostCommonBit
	}
	return filterLines(lines, filterByLeastCommon)
}

type filterFunc func(bit byte, mostCommonBit byte) bool

func filterLines(lines []string, filter filterFunc) uint64 {
	lineLen := len(lines[0])

	// Copy lines into filteredLines.
	filteredLines := make([]string, len(lines))
	for i, line := range lines {
		filteredLines[i] = line
	}

	// Go through each bit.
	for i := 0; i < lineLen; i++ {
		mostCommonBit := mostCommonBit(filteredLines, i)
		newFilteredLines := make([]string, 0, len(filteredLines))
		// Check what lines pass the filter for the current bit.
		for _, line := range filteredLines {
			if filter(line[i], mostCommonBit) {
				newFilteredLines = append(newFilteredLines, line)
			}
		}
		filteredLines = newFilteredLines
		if len(filteredLines) == 1 {
			nbr, _ := strconv.ParseUint(filteredLines[0], 2, lineLen)
			return nbr
		}
	}

	return 0
}

func mostCommonBit(lines []string, i int) byte {
	oneCount := 0
	for _, line := range lines {
		if line[i] == '1' {
			oneCount++
		}
	}
	zeroCount := len(lines) - oneCount
	if oneCount >= zeroCount {
		return '1'
	}
	return '0'
}
