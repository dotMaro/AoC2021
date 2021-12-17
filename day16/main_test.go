package main

import "testing"

func Test_parseNextPackage_versionSum(t *testing.T) {
	testCases := []struct {
		input string
		exp   int
	}{
		{input: "8A004A801A8002F478", exp: 16},
		{input: "620080001611562C8802118E34", exp: 12},
		{input: "C0015000016115A2E0802F182340", exp: 23},
		{input: "A0016C880162017C3686B18A3D4780", exp: 31},
	}

	for _, tc := range testCases {
		binary := convertHexToBinary(tc.input)
		_, _, res := parseNextPackage(binary)
		if res != tc.exp {
			t.Errorf("Should return %d, not %d, for input %s", tc.exp, res, tc.input)
		}
	}
}

func Test_bitPackage_resolve(t *testing.T) {
	testCases := []struct {
		input string
		exp   int
	}{
		{input: "C200B40A82", exp: 3},
		{input: "04005AC33890", exp: 54},
		{input: "880086C3E88112", exp: 7},
		{input: "CE00C43D881120", exp: 9},
		{input: "D8005AC2A8F0", exp: 1},
		{input: "F600BC2D8F", exp: 0},
		{input: "9C005AC2F8F0", exp: 0},
		{input: "9C0141080250320F1802104A08", exp: 1},
	}

	for _, tc := range testCases {
		binary := convertHexToBinary(tc.input)
		pkg, _, _ := parseNextPackage(binary)
		res := pkg.resolve()
		if res != tc.exp {
			t.Errorf("Should return %d, not %d, for input %s", tc.exp, res, tc.input)
		}
	}
}
