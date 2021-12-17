package main

import "testing"

func Test_versionSum(t *testing.T) {
	testCases := []struct {
		input string
		exp   int
	}{
		{input: "8A004A801A8002F478", exp: 16},
		{input: "620080001611562C8802118E34", exp: 12},
		{input: "C0015000016115A2E0802F182340", exp: 23},
	}

	for _, tc := range testCases {
		binary := convertHexToBinary(tc.input)
		t.Log("\n", binary)
		res := versionSum(binary)
		if res != tc.exp {
			t.Errorf("Should return %d, not %d", tc.exp, res)
		}
	}
}
