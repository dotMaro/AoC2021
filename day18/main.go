package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	lines := utils.SplitInput("day18/input.txt")
	fmt.Printf("Part 1. The magnitude of the sum of all numbers is %d\n", sumOfAllNumbers(lines).magnitude())
	fmt.Printf("Part 2. The largest magnitude is %d\n", findLargestMagnitude(lines))
}

func findLargestMagnitude(lines []string) int {
	largestMagnitude := 0
	for _, line1 := range lines {
		for _, line2 := range lines {
			if line1 == line2 {
				continue
			}
			pair1 := parseNextPair(line1)
			pair2 := parseNextPair(line2)
			sum := pair{
				left:  pair1,
				right: pair2,
			}
			sum = sum.reduce()
			magnitude := sum.magnitude()
			if magnitude > largestMagnitude {
				largestMagnitude = magnitude
			}
		}
	}
	return largestMagnitude
}

func sumOfAllNumbers(lines []string) pair {
	res := parseNextPair(lines[0])
	for _, line := range lines[1:] {
		curPair := parseNextPair(line)
		res = pair{
			left:  res,
			right: curPair,
		}
		res = res.reduce()
	}
	return res
}

func parseNextPair(s string) pair {
	if s[0] != '[' {
		panic("Pair should start with [")
	}
	pair := pair{}
	if s[1] == '[' {
		pair.left = parseNextPair(s[1:])
	} else {
		nbr, _ := strconv.Atoi(string(s[1]))
		pair.left = regular(nbr)
	}
	// Find comma position.
	paranthesesDepth := 0
	var afterComma int
findCommaLoop:
	for i, r := range s[1:] {
		switch r {
		case '[':
			paranthesesDepth++
		case ']':
			paranthesesDepth--
		case ',':
			if paranthesesDepth == 0 {
				afterComma = i + 2
				break findCommaLoop
			}
		}
	}
	if s[afterComma] == '[' {
		pair.right = parseNextPair(s[afterComma:])
	} else {
		nbr, _ := strconv.Atoi(string(s[afterComma]))
		pair.right = regular(nbr)
	}

	return pair
}

type pair struct {
	left, right element
}

type element interface {
	isRegular() bool
	value() int
	explosionFromLeft(val int) element
	explosionFromRight(val int) element
	magnitude() int
}

type regular int

func (r regular) value() int {
	return int(r)
}

func (r regular) isRegular() bool {
	return true
}

func (r regular) explosionFromLeft(val int) element {
	return regular(int(r) + val)
}

func (r regular) explosionFromRight(val int) element {
	return regular(int(r) + val)
}

func (r regular) magnitude() int {
	return int(r)
}

func (p pair) isRegular() bool {
	return false
}

func (p pair) value() int {
	return p.left.value() + p.right.value()
}

func (p pair) magnitude() int {
	return 3*p.left.magnitude() + 2*p.right.magnitude()
}

func (p pair) reduce() pair {
	for i := 0; i < 200; i++ {
		// Find pairs to explode.
		exploded := true
		for exploded {
			p, _, _, exploded = findNextExplosion(p, 0)
		}

		// Find regular numbers to split.
		p, _ = findNextSplit(p, 0)
	}
	return p
}

func (p pair) explosionFromLeft(val int) element {
	p.left = p.left.explosionFromLeft(val)
	return p
}

func (p pair) explosionFromRight(val int) element {
	p.right = p.right.explosionFromRight(val)
	return p
}

func findNextExplosion(p pair, depth int) (pair, int, int, bool) {
	explodeToLeft := 0
	explodeToRight := 0
	hasExploded := false
	if !p.left.isRegular() {
		// Left is pair.
		leftPair := p.left.(pair)
		if depth >= 3 {
			// Explode left pair.
			hasExploded = true
			p.left = regular(0)
			p.right = p.right.explosionFromLeft(leftPair.right.value())
			explodeToLeft = leftPair.left.value()
		} else {
			newLeft, toLeft, toRight, exploded := findNextExplosion(leftPair, depth+1)
			hasExploded = exploded
			explodeToLeft = toLeft
			p.right = p.right.explosionFromLeft(toRight)
			p.left = newLeft
		}
	}
	if !hasExploded && !p.right.isRegular() {
		// Right is pair.
		rightPair := p.right.(pair)
		if depth >= 3 {
			// Explode right pair.
			hasExploded = true
			p.right = regular(0)
			p.left = p.left.explosionFromRight(rightPair.left.value())
			explodeToRight = rightPair.right.value()
		} else {
			newRight, toLeft, toRight, exploded := findNextExplosion(rightPair, depth+1)
			hasExploded = exploded
			explodeToRight = toRight
			p.left = p.left.explosionFromRight(toLeft)
			p.right = newRight
		}
	}

	return p, explodeToLeft, explodeToRight, hasExploded
}

func findNextSplit(p pair, depth int) (pair, bool) {
	hasSplit := false
	if p.left.isRegular() {
		if p.left.value() >= 10 {
			// Split left.
			hasSplit = true
			p.left = pair{
				left:  regular(p.left.value() / 2),
				right: regular(math.Ceil(float64(p.left.value()) / 2.0)),
			}
		}
	} else {
		p.left, hasSplit = findNextSplit(p.left.(pair), depth+1)
	}
	if !hasSplit {
		if p.right.isRegular() {
			if p.right.value() >= 10 {
				// Split right.
				hasSplit = true
				p.right = pair{
					left:  regular(p.right.value() / 2),
					right: regular(math.Ceil(float64(p.right.value()) / 2.0)),
				}
			}
		} else {
			p.right, hasSplit = findNextSplit(p.right.(pair), depth+1)
		}
	}

	return p, hasSplit
}
