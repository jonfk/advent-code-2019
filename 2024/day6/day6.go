package day6

import (
	"log"
	"strings"
)

type cell int
type direction int

const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

const (
	EMPTY cell = iota
	OBSTACLE
)

type Input struct {
	direction direction
	start     []int
	matrix    [][]cell
}

func Parse(input string) Input {
	var direction direction
	var start []int
	var matrix [][]cell

	for y, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		var row []cell
		for x, ch := range line {
			switch ch {
			case '.':
				row = append(row, EMPTY)
			case '#':
				row = append(row, OBSTACLE)
			case '^':
				row = append(row, EMPTY)
				direction = UP
				start = []int{y, x}
			case '>':
				row = append(row, EMPTY)
				direction = RIGHT
				start = []int{y, x}
			case 'v':
				row = append(row, EMPTY)
				direction = DOWN
				start = []int{y, x}
			case '<':
				row = append(row, EMPTY)
				direction = LEFT
				start = []int{y, x}
			default:
				log.Fatal("Unrecognized char in input: " + string(ch))
			}
		}
		matrix = append(matrix, row)
	}

	var length int
	for _, row := range matrix {
		if length == 0 {
			length = len(row)
		} else if length != len(row) {
			log.Fatalf("Non uniform row lengths")
		}
	}
	return Input{direction, start, matrix}
}

func Run(input string) (int, int) {
	parsedInput := Parse(input)
	positionCount, visitedCells, _ := CountGuardPatrolPositions(parsedInput)

	loopObstaclesCount := CountLoopingObstacles(parsedInput, visitedCells)

	return positionCount, loopObstaclesCount
}

func CountGuardPatrolPositions(input Input) (int, [][]bool, bool) {
	var count int
	var position []int = make([]int, len(input.start))
	copy(position, input.start)

	currentDirection := input.direction
	var visitedCells [][]bool
	var directions [][][]direction
	var isOffMatrix bool = false

	for _, row := range input.matrix {
		visitedCells = append(visitedCells, make([]bool, len(row)))
		directions = append(directions, make([][]direction, len(row)))
	}

	for {
		if !visitedCells[position[0]][position[1]] {
			count += 1
			visitedCells[position[0]][position[1]] = true
		}
		if !Contains(directions[position[0]][position[1]], currentDirection) {
			directions[position[0]][position[1]] = append(directions[position[0]][position[1]], currentDirection)
		} else {
			return count, visitedCells, true
		}
		position, currentDirection, isOffMatrix = NextPosition(input.matrix, position, currentDirection)
		if isOffMatrix {
			break
		}

	}

	return count, visitedCells, false
}

func CountLoopingObstacles(input Input, visitedCells [][]bool) int {
	var count int
	for y, visitedCellsRow := range visitedCells {
		for x, visitedCell := range visitedCellsRow {
			if visitedCell {
				newMatrix := Clone(input.matrix)
				newMatrix[y][x] = OBSTACLE
				_, _, isLoop := CountGuardPatrolPositions(Input{
					matrix:    newMatrix,
					start:     input.start,
					direction: input.direction,
				})
				if isLoop {
					count += 1
				}
			}
		}
	}
	return count
}

// Calculates the next position given the arguments and whether the next position would be off
// the matrix
func NextPosition(matrix [][]cell, current []int, direction direction) ([]int, direction, bool) {
	maxX, maxY := len(matrix[0]), len(matrix)
	var nextPosition []int
	switch direction {
	case UP:
		nextPosition = []int{current[0] - 1, current[1]}
	case DOWN:
		nextPosition = []int{current[0] + 1, current[1]}
	case RIGHT:
		nextPosition = []int{current[0], current[1] + 1}
	case LEFT:
		nextPosition = []int{current[0], current[1] - 1}
	default:
		log.Fatalf("unrecognized direction: %v\n", direction)
	}
	if nextPosition[0] < 0 || nextPosition[0] >= maxY || nextPosition[1] < 0 || nextPosition[1] >= maxX {
		return nextPosition, direction, true
	}

	if matrix[nextPosition[0]][nextPosition[1]] == OBSTACLE {
		return NextPosition(matrix, current, RotateNextDirection(direction))
	}

	return nextPosition, direction, false
}

func RotateNextDirection(direction direction) direction {
	if direction == LEFT {
		return UP
	} else {
		return direction + 1
	}
}

func Contains(s []direction, e direction) bool {
	for _, x := range s {
		if x == e {
			return true
		}
	}
	return false
}

func ContainsPos(s [][]int, p []int) int {
	for i, p1 := range s {
		if p1[0] == p[0] && p1[1] == p[1] {
			return i
		}
	}
	return -1
}

func Clone(matrix [][]cell) [][]cell {
	clone := make([][]cell, len(matrix))

	for i := range matrix {
		clone[i] = make([]cell, len(matrix[i]))
		copy(clone[i], matrix[i])
	}
	return clone
}
