package day6

import (
	"fmt"
	"os"
	"testing"
)

const example = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestCountPositionsExample(t *testing.T) {
	count, possibleObstacles := Run(example)

	if count != 41 {
		t.Fatalf("expected = 41, got = %v\n", count)
	}

	if possibleObstacles != 6 {
		t.Fatalf("possibleObstacles: expected = 6, got = %v\n", possibleObstacles)
	}
}

func TestCountPositions(t *testing.T) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Fatalf("Could not read puzzle input: err = %s", err)
	}
	count, possibleObstacles := Run(string(input))

	fmt.Printf("# Puzzle\ncount = %v\n\nobstacles = %v\n", count, possibleObstacles)
}
