package day14

import (
	"fmt"
	"os"
	"testing"
)

const example = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func TestParse(t *testing.T) {
	robots := Parse(example)

	if len(robots) != 12 {
		t.Fatalf("incorrect len(robots), expected = 12, got = %v", len(robots))
	}

	if robots[0].Pos.X != 0 || robots[0].Pos.Y != 4 || robots[0].Vel.X != 3 || robots[0].Vel.Y != -3 {
		t.Fatalf("incorrectly parsed 1st robot")
	}
}

func TestSafetyFactorExample(t *testing.T) {
	robots := Parse(example)

	Tick(100, robots, 11, 7)
	safetyFactor := CalculateSafetyFactor(robots, 11, 7)

	if safetyFactor != 12 {
		t.Fatalf("expected = 12, got = %v", safetyFactor)
	}
}

func TestSafetyFactor(t *testing.T) {

	inputTxt, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("err reading file: %v", err)
	}
	robots := Parse(string(inputTxt))
	xLen := 101
	yLen := 103

	Tick(100000000, robots, xLen, yLen)
	safetyFactor := CalculateSafetyFactor(robots, xLen, yLen)

	fmt.Printf("# Puzzle\nsafety factor: %v\n", safetyFactor)
}
