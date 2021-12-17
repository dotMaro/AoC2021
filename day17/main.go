package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = `target area: x=85..145, y=-163..-108`

func main() {
	targetArea := parseTargetArea(input)
	fmt.Printf("Task 1. The highest possible altitude is %d\n", targetArea.findHighestY())
	fmt.Printf("Task 2. The amount of valid velocity variations is %d\n", targetArea.validVelocityCount())
}

func (a targetArea) findHighestY() int {
	highestY := 0
	for x := 0; x <= 2000; x++ {
		for y := 0; y <= 2000; y++ {
			pos := vec{0, 0}
			vel := vec{x: x, y: y}
			localHighestY := 0
			for pos.x <= a.maxX && pos.y >= a.minY {
				pos, vel = step(pos, vel)
				if pos.y > localHighestY {
					localHighestY = pos.y
				}
				if localHighestY > highestY && pos.x >= a.minX && pos.x <= a.maxX && pos.y >= a.minY && pos.y <= a.maxY {
					highestY = localHighestY
				}
			}
		}
	}
	return highestY
}

func (a targetArea) validVelocityCount() int {
	validCount := 0
	for x := 1; x <= a.maxX; x++ {
		for y := a.minY; y <= 200; y++ {
			pos := vec{0, 0}
			vel := vec{x: x, y: y}
			for pos.x <= a.maxX && pos.y >= a.minY {
				pos, vel = step(pos, vel)
				if pos.x >= a.minX && pos.x <= a.maxX && pos.y >= a.minY && pos.y <= a.maxY {
					validCount++
					break
				}
			}
		}
	}
	return validCount
}

func parseTargetArea(s string) targetArea {
	// Beautiful, I know.
	split := strings.SplitN(s[15:], "..", 3)
	minX, _ := strconv.Atoi(split[0])
	maxY, _ := strconv.Atoi(split[2])
	split2 := strings.SplitN(split[1], ", y=", 2)
	minY, _ := strconv.Atoi(split2[1])
	maxX, _ := strconv.Atoi(split2[0])

	return targetArea{
		minX: minX,
		maxX: maxX,
		minY: minY,
		maxY: maxY,
	}
}

type targetArea struct {
	minX, maxX int
	minY, maxY int
}

type vec struct {
	x, y int
}

func step(pos vec, vel vec) (vec, vec) {
	// The probe's x position increases by its x velocity.
	pos.x += vel.x
	// The probe's y position increases by its y velocity.
	pos.y += vel.y

	// Due to drag, the probe's x velocity changes by 1 toward the value 0; that is, it decreases by 1 if it is greater than 0, increases by 1 if it is less than 0, or does not change if it is already 0.
	if vel.x > 0 {
		vel.x--
	} else if vel.x < 0 {
		vel.x++
	}

	// Due to gravity, the probe's y velocity decreases by 1.
	vel.y--
	return pos, vel
}
