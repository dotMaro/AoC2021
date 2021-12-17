package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	input := utils.InputString("day16/input.txt")
	binary := convertHexToBinary(input)
	pkg, _, versionSum := parseNextPackage(binary)
	fmt.Printf("Task 1. The sum of all versions is %d\n", versionSum)
	fmt.Printf("Task 2. The evaluated expression is %d\n", pkg.resolve())
}

func convertHexToBinary(s string) string {
	mapping := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
	var binary strings.Builder
	for _, r := range s {
		binary.WriteString(mapping[r])
	}
	return binary.String()
}

func parseNextPackage(s string) (bitsPackage, int, int) {
	version, _ := strconv.ParseUint(s[0:3], 2, 8)
	typeID, _ := strconv.ParseUint(s[3:6], 2, 8)
	var endOffset int
	value := -1
	var subpackages []bitsPackage

	versionSum := int(version)

	switch typeID {
	case 4: // Literal value.
		reachedEnd := false
		var valueStr strings.Builder
		for start := 6; !reachedEnd; start += 5 {
			end := start + 5
			group := s[start:end]
			if group[0] == '0' {
				reachedEnd = true
			}
			endOffset = end
			valueStr.WriteString(group[1:5])
		}
		val, _ := strconv.ParseUint(valueStr.String(), 2, 64)
		value = int(val)
	default: // Operator.
		subpackageOffset := 0
		switch s[6] {
		case '0': // Length type.
			length, _ := strconv.ParseUint(s[7:22], 2, 16)
			for subpackageOffset < int(length) {
				pkg, offset, sum := parseNextPackage(s[22+subpackageOffset:])
				subpackages = append(subpackages, pkg)
				versionSum += sum
				subpackageOffset += offset
			}
			endOffset = 22 + subpackageOffset
		case '1': // Packet type.
			numberOfSubpackets, _ := strconv.ParseUint(s[7:18], 2, 16)
			count := 0
			for count < int(numberOfSubpackets) {
				pkg, offset, sum := parseNextPackage(s[18+subpackageOffset:])
				subpackages = append(subpackages, pkg)
				versionSum += sum
				subpackageOffset += offset
				count++
			}
			endOffset = 18 + subpackageOffset
		}
	}

	return bitsPackage{
		version:     int(version),
		typeID:      int(typeID),
		value:       value,
		subpackages: subpackages,
	}, endOffset, versionSum
}

type bitsPackage struct {
	version     int
	typeID      int
	value       int
	subpackages []bitsPackage
}

func (p bitsPackage) resolve() int {
	switch p.typeID {
	case 0: // Sum.
		sum := 0
		for _, s := range p.subpackages {
			sum += s.resolve()
		}
		return sum
	case 1: // Product.
		product := 1
		for _, s := range p.subpackages {
			product *= s.resolve()
		}
		return product
	case 2: // Minimum.
		minimum := math.MaxInt32
		for _, s := range p.subpackages {
			val := s.resolve()
			if val < minimum {
				minimum = val
			}
		}
		return minimum
	case 3: // Maximum.
		maximum := 0
		for _, s := range p.subpackages {
			val := s.resolve()
			if val > maximum {
				maximum = val
			}
		}
		return maximum
	case 4: // Value (not an operator).
		return p.value
	case 5: // Greater than.
		if p.subpackages[0].resolve() > p.subpackages[1].resolve() {
			return 1
		}
		return 0
	case 6: // Less than.
		if p.subpackages[0].resolve() < p.subpackages[1].resolve() {
			return 1
		}
		return 0
	case 7: // Equal to.
		if p.subpackages[0].resolve() == p.subpackages[1].resolve() {
			return 1
		}
		return 0
	default:
		panic(fmt.Sprintf("Bad operation typeID %d", p.typeID))
	}
}
