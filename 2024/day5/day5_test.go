package day5

import (
	"fmt"
	"os"
	"testing"
)

const exampleInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestParse(t *testing.T) {
	input, err := Parse(exampleInput)
	if err != nil {
		t.Fatalf("Err parsing: %s\n", err)
	}

	if len(input.Rules) != 21 {
		t.Fatalf("expected = 21 rules, got = %v rules\n", len(input.Rules))
	}

	if len(input.Updates) != 6 {
		t.Fatalf("expected = 6 updates, got = %v updates\n", len(input.Updates))
	}
}

func TestRunExample(t *testing.T) {
	sum, err := Run(exampleInput)
	if err != nil {
		t.Fatalf("Err parsing: %s\n", err)
	}

	if sum != 143 {
		t.Fatalf("Expected = 143, got = %v\n", sum)
	}
}

func TestRun(t *testing.T) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Fatalf("Could not read puzzle input: err = %s", err)
	}
	sum, err := Run(string(input))
	if err != nil {
		t.Fatalf("Err parsing: %s\n", err)
	}

	fmt.Printf("# Puzzle\nSum = %v\n", sum)
}

func TestRun2Example(t *testing.T) {
	sum, err := Run2(exampleInput)
	if err != nil {
		t.Fatalf("Err parsing: %s\n", err)
	}

	if sum != 123 {
		t.Fatalf("Expected = 123, got = %v\n", sum)
	}
}

func TestRun2(t *testing.T) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Fatalf("Could not read puzzle input: err = %s", err)
	}
	sum, err := Run2(string(input))
	if err != nil {
		t.Fatalf("Err parsing: %s\n", err)
	}

	fmt.Printf("# Puzzle Part 2\nSum = %v\n", sum)
}
