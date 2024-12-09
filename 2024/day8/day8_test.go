package day8

import (
	"fmt"
	"os"
	"testing"
)

const example = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestCountAntiNodesExample(t *testing.T) {
	input := Parse(example)

	count := CountAntiNodes(input)

	if count != 14 {
		t.Fatalf("expected = 14, got = %v\n", count)
	}
}

func TestCountAntiNodes(t *testing.T) {
	inputTxt, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Fatalf("Could not read puzzle input: err = %s", err)
	}

	input := Parse(string(inputTxt))
	count := CountAntiNodes(input)

	fmt.Printf("# Puzzle\nCount = %v\n", count)
}
