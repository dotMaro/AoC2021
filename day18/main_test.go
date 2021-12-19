package main

import (
	"reflect"
	"testing"

	"github.com/dotMaro/AoC2021/utils"
)

func Test_findNextExplosion(t *testing.T) {
	testCases := []struct {
		input string
		exp   string
	}{
		{
			input: "[[[[[9,8],1],2],3],4]",
			exp:   "[[[[0,9],2],3],4]",
		},
		{
			input: "[7,[6,[5,[4,[3,2]]]]]",
			exp:   "[7,[6,[5,[7,0]]]]",
		},
		{
			input: "[[6,[5,[4,[3,2]]]],1]",
			exp:   "[[6,[5,[7,0]]],3]",
		},
		{
			input: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			exp:   "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			input: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			exp:   "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
	}

	for _, tc := range testCases {
		pair := parseNextPair(tc.input)
		res, _, _, _ := findNextExplosion(pair, 0)
		exp := parseNextPair(tc.exp)
		if !reflect.DeepEqual(res, exp) {
			t.Errorf("Should return %s for input %s, but returned %v", tc.exp, tc.input, res)
		}
	}
}

func Test_magnitude(t *testing.T) {
	const input = `[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]`
	pair := parseNextPair(input)
	res := pair.magnitude()
	if res != 4140 {
		t.Errorf("Should return 4140, not %d", res)
	}
}

const example = `[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]`

func Test_sumOfAllNumbers(t *testing.T) {
	res := sumOfAllNumbers(utils.SplitLine(example))
	t.Error(res)
}

func Test_sumOfAllNumbers2(t *testing.T) {
	const input = `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]`
	res := sumOfAllNumbers(utils.SplitLine(input))
	t.Error(res)
}

func Test_reduce(t *testing.T) {
	// const input = `[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]`
	const input = `[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]`
	pair := parseNextPair(input)
	res := pair.reduce()
	t.Error(res)
}

// func Test_split(t *testing.T) {
// const input = `[[[[0,7],4],[15,[0,13]]],[1,1]]`

// }
