package day3

import (
	"regexp"
	"strconv"
	"strings"
)

type patternType int

const (
	mulPattern patternType = iota
	dontPattern
	doPattern
)

type match struct {
	Pattern patternType
	Text    string
	Index   int
}

func ParseWithRegex(input string) []Mul {
	patterns := map[patternType]*regexp.Regexp{
		mulPattern:  regexp.MustCompile(`mul\((\d+),(\d+)\)`),
		dontPattern: regexp.MustCompile(`don't\(\)`),
		doPattern:   regexp.MustCompile(`do\(\)`),
	}

	var allMatches []match

	// Find matches for each pattern
	for patternName, regex := range patterns {
		// Find all matches with their indices
		matches := regex.FindAllStringIndex(input, -1)

		// Store each match with its pattern type and position
		for _, matchIndex := range matches {
			allMatches = append(allMatches, match{
				Pattern: patternName,
				Text:    input[matchIndex[0]:matchIndex[1]],
				Index:   matchIndex[0],
			})
		}
	}

	sort(allMatches)

	var program []Mul
	var isDontActive bool

	for _, match := range allMatches {
		if match.Pattern == doPattern {
			isDontActive = false
		} else if match.Pattern == dontPattern {
			isDontActive = true
		} else if !isDontActive && match.Pattern == mulPattern {
			mul := extractMulWithStrings(match)
			program = append(program, mul)
		}
	}

	return program
}

func extractMulWithStrings(match match) Mul {
	input := match.Text

	// Remove first 4 characters ("mul(") and last character (")")
	content := input[4 : len(input)-1]
	// Find comma position
	commaPos := strings.IndexByte(content, ',')
	// Extract and convert numbers
	num1, _ := strconv.Atoi(content[:commaPos])
	num2, _ := strconv.Atoi(content[commaPos+1:])
	return Mul{val1: num1, val2: num2}
}

func sort(l []match) {
	quicksort(l, 0, len(l)-1)
}

func quicksort(l []match, lo, hi int) {
	if lo >= hi || lo < 0 {
		return
	}

	pivotIdx := partition(l, lo, hi)

	quicksort(l, lo, pivotIdx-1)
	quicksort(l, pivotIdx+1, hi)
}

func partition(l []match, lo, hi int) int {
	pivot := l[hi]
	i := lo

	for j := lo; j < hi; j++ {
		if l[j].Index <= pivot.Index {
			l[j], l[i] = l[i], l[j]
			i++
		}
	}

	l[i], l[hi] = l[hi], l[i]
	return i

}
