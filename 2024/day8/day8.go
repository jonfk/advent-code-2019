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
	var an1, an2 Coord

	xDiff := a.x - b.x
	yDiff := a.y - b.y

	an1.x = a.x + xDiff
	an1.y = a.y + yDiff

	if IsValid(an1, xLen, yLen) {
		res = append(res, an1)
	}

	an2.x = b.x - xDiff
	an2.y = b.y - yDiff

	if IsValid(an2, xLen, yLen) {
		res = append(res, an2)
	}

	return res
}

func CountAntiNodesWithResonance(input Input) int {
	var antiNodes map[Coord]bool = make(map[Coord]bool)

	for _, coords := range input.antennas {
		combinations := GenCombinations(coords)

		for _, combo := range combinations {
			if len(combo) != 2 {
				log.Fatalf("Incorrect combo generated")
			}
			for _, c := range GenAntiNodesWithResonance(combo[0], combo[1], input.xLen, input.yLen) {
				antiNodes[c] = true
			}
		}
	}

	return len(antiNodes)
}

func GenAntiNodesWithResonance(a, b Coord, xLen, yLen int) []Coord {
	var res []Coord

	xDiff := a.x - b.x
	yDiff := a.y - b.y

	res = append(res, a)

	i := 1
	for {
		var an Coord
		an.x = a.x + (xDiff * i)
		an.y = a.y + (yDiff * i)

		if IsValid(an, xLen, yLen) {
			res = append(res, an)
		} else {
			break
		}
		i++
	}

	i = 1
	for {
		var an Coord
		an.x = a.x - (xDiff * i)
		an.y = a.y - (yDiff * i)
		if IsValid(an, xLen, yLen) {
			res = append(res, an)
		} else {
			break
		}
		i++
	}

	return res
}

func getLinePoints(point1, point2 Coord, xLen, yLen int) []Coord {
	points := make([]Coord, 0)

	// Handle vertical line
	if point1.x == point2.x {
		for y := 0; y <= yLen; y++ {
			points = append(points, Coord{point1.x, y})
		}
		return points
	}

	// Calculate slope and y-intercept
	slope := float64(point2.y-point1.y) / float64(point2.x-point1.x)
	b := float64(point1.y) - slope*float64(point1.x)

	// Calculate points for the entire width of the space
	for x := 0; x <= xLen; x++ {
		y := slope*float64(x) + b
		roundedY := round(y)

		// Only add points that are within the vertical bounds
		if roundedY >= 0 && roundedY < yLen {
			points = append(points, Coord{x, roundedY})
		}
	}

	return points
}

// round implements math.Round functionality manually
// Adds 0.5 and truncates to achieve rounding to nearest integer
func round(x float64) int {
	if x < 0 {
		return int(x - 0.5)
	}
	return int(x + 0.5)
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
