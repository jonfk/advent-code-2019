package day4

import (
	"strings"
)

func CountAllXmas(input string) int {
	lines := collectAllLines(input)
	return countAllXmas(lines)
}

func collectAllLines(input string) []string {
	var lines []string
	var rows []string

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}

	collectHorizontalAndDiagonalLines := func(lines []string) {
		var rowLength int
		rowLength = len(lines[0])

		for i := range lines {
			row := new(strings.Builder)
			for j := 0; j < rowLength; j++ {
				row.WriteByte(lines[i][j])
			}
			rows = append(rows, row.String())
		}

		// x...
		// ax..
		// bax.
		// cbax
		// dcba
		var j int
		for l := 0; l < len(lines); l++ {
			row := new(strings.Builder)

			for j = 0; j < rowLength; j++ {
				if l+j >= len(lines) {
					continue
				}
				row.WriteByte(lines[l+j][j])
			}

			rows = append(rows, row.String())
		}

		// // .abc
		// // ..ab
		// // ...a
		// // ....
		// // ....
		// // 0,0 1,1 2,2
		// // 0,1 1,2 2,3
		// // 0,2 1,3 2,4
		for j := rowLength - 1; j > 0; j-- {
			row := new(strings.Builder)
			for i := 0; i < len(lines); i++ {
				if i+j >= rowLength {
					break
				}
				row.WriteByte(lines[i][i+j])
			}
			if row.Len() > 0 {
				rows = append(rows, row.String())
			}
		}

	}

	rotateLines := func(lines []string) []string {
		var rotatedLines []string

		for j := len(lines[0]) - 1; j >= 0; j-- {
			line := new(strings.Builder)
			for i := 0; i < len(lines); i++ {
				line.WriteByte(lines[i][j])
			}
			rotatedLines = append(rotatedLines, line.String())

		}
		return rotatedLines
	}
	collectHorizontalAndDiagonalLines(lines)
	collectHorizontalAndDiagonalLines(rotateLines(lines))

	return rows
}

func countAllXmas(lines []string) int {
	var count int

	for _, line := range lines {
		// count += countSubstringOverlapping(line, "XMAS")
		// count += countSubstringOverlapping(line, "SAMX")
		count += strings.Count(line, "XMAS")
		count += strings.Count(line, "SAMX")
	}
	return count
}

func countSubstringOverlapping(input, substr string) int {
	var count int
	for i := range input {
		if strings.HasPrefix(input[i:], substr) {
			count += 1
		}
	}
	return count
}

// Part 2

func CountMasInX(input string) int {
	lines := strings.Split(input, "\n")
	var matrix [][]rune

	for _, line := range lines {
		var row []rune
		if len(line) <= 1 {
			continue
		}
		for _, ch := range line {
			row = append(row, ch)
		}
		matrix = append(matrix, row)
	}

	var count int
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 'A' {
				if ((i-1 >= 0 && j+1 < len(matrix[i]) && matrix[i-1][j+1] == 'S' && i+1 < len(matrix) && j-1 >= 0 && matrix[i+1][j-1] == 'M') || (i-1 >= 0 && j+1 < len(matrix[i]) && matrix[i-1][j+1] == 'M' && i+1 < len(matrix) && j-1 >= 0 && matrix[i+1][j-1] == 'S')) && ((i-1 >= 0 && j-1 >= 0 && matrix[i-1][j-1] == 'S' && i+1 < len(matrix) && j+1 < len(matrix[i]) && matrix[i+1][j+1] == 'M') || (i-1 >= 0 && j-1 >= 0 && matrix[i-1][j-1] == 'M' && i+1 < len(matrix) && j+1 < len(matrix[i]) && matrix[i+1][j+1] == 'S')) {
					count += 1
				}
			}
		}
	}
	return count
}
