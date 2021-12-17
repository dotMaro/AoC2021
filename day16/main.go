package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const test = "8A004A801A8002F478"
	binary := convertHexToBinary(test)
	fmt.Println(binary)
	fmt.Println(versionSum(binary))
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

func versionSum(s string) int {
	versionSum := 0

	offsetSum := 0
	lastOffset := 0
	for offsetSum < len(s)-3 {
		_, offset, sum := parseNextPackage(s[lastOffset:])
		// fmt.Printf("%d - %d\n", offset, len(s))
		offsetSum += offset
		versionSum += sum
		// offset2 := offset
		// if offset2%4 != 0 {
		// 	offset2 += 4 - offset2%4
		// }
		lastOffset = offset
	}
	return versionSum
}

func parseNextPackage(s string) (bitsPackage, int, int) {
	version, _ := strconv.ParseUint(s[0:3], 2, 8)
	typeID, _ := strconv.ParseUint(s[3:6], 2, 8)
	var endOffset int

	versionSum := int(version)

	switch typeID {
	case 4: // Literal value.
		reachedEnd := false
		for start := 6; !reachedEnd; start += 5 {
			end := start + 5
			group := s[start:end]
			if group[0] == '0' {
				reachedEnd = true
			}
			endOffset = end
			// For now, ignore the value.
		}
	default: // Operator.
		subpackageOffset := 0
		switch s[6] {
		case '0': // Length type.
			length, _ := strconv.ParseUint(s[7:22], 2, 16)
			// subpackageOffsets := 0
			for subpackageOffset < int(length) {
				_, offset, sum := parseNextPackage(s[22+subpackageOffset:])
				versionSum += sum
				subpackageOffset += offset
			}
			endOffset = 22 + subpackageOffset
		case '1': // Packet type.
			numberOfSubpackets, _ := strconv.ParseUint(s[7:18], 2, 16)
			count := 0
			for count < int(numberOfSubpackets) {
				_, offset, sum := parseNextPackage(s[18+subpackageOffset:])
				versionSum += sum
				subpackageOffset += offset
				count++
			}
			endOffset = 18 + subpackageOffset
		}
		// if endOffset%4 != 0 {
		// 	endOffset += 4 - endOffset%4
		// }
	}

	// if endOffset%4 != 0 {
	// 	endOffset += 4 - endOffset%4
	// }
	return bitsPackage{
		version: int(version),
		typeID:  int(typeID),
	}, endOffset, versionSum
}

type bitsPackage struct {
	version int
	typeID  int
}
