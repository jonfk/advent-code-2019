package day16

import (
	"fmt"
	"os"
	"testing"
)

const example = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

const example2 = `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`

func TestExample(t *testing.T) {
	puzzle := Parse(example)

	cost := puzzle.FindMinPathCostBFS()
	if cost != 7036 {
		t.Fatalf("expected = 7036, got = %v", cost)
	}

	_, spots := puzzle.FindAllMinCostPaths()
	if spots != 45 {
		t.Fatalf("expected = 45, got = %v", spots)
	}
}

func TestExample2(t *testing.T) {
	puzzle := Parse(example2)

	cost := puzzle.FindMinPathCostBFS()
	if cost != 11048 {
		t.Fatalf("expected = 11048, got = %v", cost)
	}
	_, spots := puzzle.FindAllMinCostPaths()
	if spots != 64 {
		t.Fatalf("expected = 64, got = %v", spots)
	}
}

func TestRealInput(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}
	puzzle := Parse(string(input))

	cost := puzzle.FindMinPathCostAStar()
	fmt.Printf("# Puzzle\ncost = %v\n", cost)
	_, spots := puzzle.FindAllMinCostPaths()
	fmt.Printf("spots = %v\n", spots)
}
