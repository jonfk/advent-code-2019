package day13

import (
	"fmt"
	"os"
	"testing"
)

const example = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

func TestFindMinTokensExample(t *testing.T) {

	machines := Parse(example)

	if len(machines) != 4 {
		t.Fatalf("Incorrect number of machines parsed. expected = 4, got = %v, machines = %#v\n", len(machines), machines)
	}

	tokens := FindMinTokenForAllPossiblePrizes(machines)

	if tokens != 480 {
		t.Fatalf("expected = 480 tokens, got = %v tokens", tokens)
	}
}

func TestSolve(t *testing.T) {
	countA, countB, possible := SolvePrizeSystem(8400, 5400, 94, 22, 34, 67)
	if countA != 80 || countB != 40 || !possible {
		t.Fatalf("Not solved")
	}
}

func TestFindMinTokens(t *testing.T) {
	inputTxt, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	machines := Parse(string(inputTxt))
	tokens := FindMinTokenForAllPossiblePrizesCorrected(machines)

	fmt.Printf("# Puzzle\ntokens = %v\n", tokens)
}
