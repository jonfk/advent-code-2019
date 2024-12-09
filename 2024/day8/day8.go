package day8

import (
	"log"
	"strings"
)

// ..........
// ...#......
// ..........
// ....a.....
// ..........
// .....a....
// ..........
// ......#...
// ..........
// ..........

// ......,....,
// ...,....0...
// ....,0....,.
// ..,....0....
// ....0....,..
// .,....A.....
// ...,........
// #......,....
// ........A...
// .........A..
// ..........,.
// ..........,.

type Coord struct {
	x, y int
}

type Input struct {
	xLen     int
	yLen     int
	antennas map[rune][]Coord
}

func Parse(input string) Input {
	lines := strings.Split(input, "\n")
	var yLen int
	var xLen int
	var antennas map[rune][]Coord = make(map[rune][]Coord)

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		if xLen == 0 {
			xLen = len(line)
		} else if xLen != len(line) {
			log.Fatalf("Uneven row lengths xLen = %v, len(line) = %v", xLen, len(line))
		}
		for x, ch := range line {
			if ch != '.' {
				antennas[ch] = append(antennas[ch], Coord{y: y, x: x})
			}
		}
		yLen += 1
	}

	return Input{xLen: xLen, yLen: yLen, antennas: antennas}
}

func CountAntiNodes(input Input) int {
	var antiNodes map[Coord]bool = make(map[Coord]bool)

	for _, coords := range input.antennas {
		combinations := GenCombinations(coords)

		for _, combo := range combinations {
			if len(combo) != 2 {
				log.Fatalf("Incorrect combo generated")
			}
			for _, c := range GenAntiNodes(combo[0], combo[1], input.xLen, input.yLen) {
				antiNodes[c] = true
			}
		}
	}

	return len(antiNodes)
}

func GenAntiNodes(a, b Coord, xLen, yLen int) []Coord {
	var res []Coord

	var an1 Coord
	an1XDiff := a.x - b.x
	an1.x = a.x + an1XDiff
	an1YDiff := a.y - b.y
	an1.y = a.y + an1YDiff

	if IsValid(an1, xLen, yLen) {
		res = append(res, an1)
	}

	var an2 Coord
	an2XDiff := b.x - a.x
	an2.x = b.x + an2XDiff
	an2YDiff := b.y - a.y
	an2.y = b.y + an2YDiff

	if IsValid(an2, xLen, yLen) {
		res = append(res, an2)
	}

	return res
}

func IsValid(c Coord, xLen, yLen int) bool {
	return c.x >= 0 && c.x < xLen && c.y >= 0 && c.y < yLen
}

func GenCombinations(coords []Coord) [][]Coord {
	var temp []Coord = make([]Coord, len(coords))
	copy(temp, coords)

	var res [][]Coord

	for len(temp) != 1 {
		coord1 := temp[0]
		for i := 1; i < len(temp); i++ {
			res = append(res, []Coord{coord1, temp[i]})
		}
		temp = temp[1:]
	}

	return res
}
