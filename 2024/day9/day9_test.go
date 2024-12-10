package day9

import (
	"fmt"
	"os"
	"testing"
)

const example = `2333133121414131402`

func TestChecksumExample(t *testing.T) {
	input := Parse(example)
	compactedBlocks := Compact(input.blocks)
	checksum := CalculateChecksumCompacted(compactedBlocks)

	if checksum != 1928 {
		t.Fatalf("expected = 1928, got = %v\n", checksum)
	}
}

func TestCompactedByFileChecksumExample(t *testing.T) {

	input := Parse2(example)
	CompactByFile(input)
	blocks := ContiguousBlocks(input).toBlocks()

	checksum := CalculateChecksumCompacted(blocks)
	if checksum != 2858 {
		t.Fatalf("expected = 2858, got = %v\n", checksum)
	}

}

func TestChecksum(t *testing.T) {
	inputTxt, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Fatalf("err reading file: %v", err)
	}
	input := Parse(string(inputTxt))
	compactedBlocks := Compact(input.blocks)
	checksum := CalculateChecksumCompacted(compactedBlocks)

	fmt.Printf("# Puzzle\nchecksum = %v\n", checksum)
}

func TestChecksum2(t *testing.T) {
	inputTxt, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Fatalf("err reading file: %v", err)
	}
	input := Parse2(string(inputTxt))
	CompactByFile(input)
	compactedBlocks := ContiguousBlocks(input).toBlocks()
	checksum := CalculateChecksumCompacted(compactedBlocks)

	fmt.Printf("# Puzzle Part 2\nchecksum = %v\n", checksum)
}

// Evil input that should show worst case perf
func TestChecksumRedditEvil(t *testing.T) {
	inputTxt, err := os.ReadFile("./input2.txt")
	if err != nil {
		t.Fatalf("err reading file: %v", err)
	}

	input := Parse(string(inputTxt))
	compactedBlocks := Compact(input.blocks)
	checksum := CalculateChecksumCompacted(compactedBlocks)

	inputByFile := Parse2(string(inputTxt))
	CompactByFile(inputByFile)
	compactedBlocksByFile := ContiguousBlocks(inputByFile).toBlocks()
	checksum2 := CalculateChecksumCompacted(compactedBlocksByFile)

	fmt.Printf("# Reddit Evil Input\nchecksum = %v, checksum2 = %v\n", checksum, checksum2)
}
