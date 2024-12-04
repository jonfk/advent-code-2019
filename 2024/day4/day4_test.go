package day4

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const matrix = `abc
def
ghi`

// ..X...
// .SAMX.
// .A..A.
// XMAS.S
// .X....
const smallExample = `..X...
.SAMX.
.A..A.
XMAS.S
.X....`

const example = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func expectedLines(input string) int {
	lines := strings.Split(input, "\n")
	numLines := len(lines)
	numRows := len(lines[0])

	return ((numLines + numRows - 1) * 2) + numLines + numRows

}

func TestCollectAllLines(t *testing.T) {

	lines := collectAllLines(matrix)
	expected := expectedLines(matrix)

	if len(lines) != expected {
		t.Fatalf("Expected = %d, got = %d\nlines = %#v\n", expected, len(lines), lines)
	}

	lines = collectAllLines(smallExample)
	expected = expectedLines(smallExample)

	if len(lines) != expected {
		t.Fatalf("Expected = %d, got = %d\nlines = %#v\n", expected, len(lines), lines)
	}
}

func TestCountXmas(t *testing.T) {
	count := CountAllXmas(smallExample)

	// if count != 4 {
	// 	t.Fatalf("expected = 4, got = %d\n", count)
	// }

	count = CountAllXmas(example)

	if count != 18 {
		t.Fatalf("expected = 18, got = %d\n", count)
	}
}

func TestCountXmasInput(t *testing.T) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Fatalf("Could not read puzzle input: err = %s", err)
	}

	count := CountAllXmas(string(input))

	fmt.Printf("# Puzzle\nCount: %d\n", count)
}

func TestCountMasInX(t *testing.T) {
	count := CountMasInX(example)
	if count != 9 {
		t.Fatalf("expected = 9, got = %d\n", count)
	}
}

func TestCountMasInXInput(t *testing.T) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Fatalf("Could not read puzzle input: err = %s", err)
	}

	count := CountMasInX(string(input))

	fmt.Printf("# Puzzle Part 2\nCount: %d\n", count)
}
