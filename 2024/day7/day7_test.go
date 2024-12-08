package day7

import (
	"fmt"
	"os"
	"testing"
)

const example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestGenPermutationOps(t *testing.T) {
	ops := GenPermutationOps(3)

	if len(ops) != 8 {
		t.Fatalf("expected = 8, got = %v\nops = %#v\n", len(ops), ops)
	}
}

func TestTotalValidEquationsExample(t *testing.T) {
	equations := Parse(example)

	total := TotalValidEquations(equations)

	if total != 3749 {
		t.Fatalf("expected = 3749, got = %v\n", total)
	}
}

func TestTotalValidEquations(t *testing.T) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Fatalf("Could not read puzzle input: err = %s", err)
	}
	equations := Parse(string(input))

	total := TotalValidEquations(equations)

	fmt.Printf("# Puzzle\nPart 1 Total = %v\n", total)
}
