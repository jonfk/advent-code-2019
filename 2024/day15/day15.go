package day15

import (
	"fmt"
	"log"
	"strings"
)

type Cell int

const (
	EMPTY Cell = iota
	BOX
	WALL
	ROBOT
)

type State struct {
	grid       [][]Cell
	robotPos   Vect
	moves      []Vect
	xLen, yLen int
}

type Vect struct {
	x, y int
}

func Parse(input string) *State {
	inputSplit := strings.Split(input, "\n\n")

	if len(inputSplit) != 2 {
		log.Fatalf("incorrect input. Could not split into grid and moves.")
	}

	gridTxt := inputSplit[0]
	movesTxt := inputSplit[1]

	var grid [][]Cell
	var robotPos Vect
	var moves []Vect

	for y, line := range strings.Split(gridTxt, "\n") {
		var row []Cell

		for x, ch := range line {
			switch ch {
			case '#':
				row = append(row, WALL)
			case '.':
				row = append(row, EMPTY)
			case 'O':
				row = append(row, BOX)
			case '@':
				row = append(row, ROBOT)
				robotPos.x = x
				robotPos.y = y
			}
		}
		grid = append(grid, row)
	}

	if grid[robotPos.y][robotPos.x] != ROBOT {
		log.Fatalf("Robot could not be found on grid")
	}

	for _, line := range strings.Split(movesTxt, "\n") {
		for _, ch := range line {
			switch ch {
			case '<':
				moves = append(moves, Vect{x: -1})
			case '>':
				moves = append(moves, Vect{x: 1})
			case '^':
				moves = append(moves, Vect{y: -1})
			case 'v':
				moves = append(moves, Vect{y: 1})
			}
		}
	}

	return &State{grid: grid, robotPos: robotPos, moves: moves, xLen: len(grid[0]), yLen: len(grid)}
}

func (s *State) ConsumeAllMoves() {
	for len(s.moves) > 0 {
		s.ConsumeMove()
		// s.Print()
	}
}

func (s *State) ConsumeMove() bool {
	if len(s.moves) == 0 {
		return false
	}
	move := s.moves[0]
	s.moves = s.moves[1:]

	newRobotPos := s.Move(s.robotPos, move)
	// log.Printf("newrobotpos: %v\n", newRobotPos)
	s.robotPos.x = newRobotPos.x
	s.robotPos.y = newRobotPos.y

	return true
}

func (s *State) Move(pos Vect, move Vect) Vect {
	var stack []Vect
	nextPos := pos
	for s.IsMovable(nextPos) {
		toAdd := nextPos
		stack = append(stack, toAdd)

		nextPos.x += move.x
		nextPos.y += move.y
	}
	// log.Printf("move: %v, stack: %v\n", move, stack)

	var moveTo Vect = pos
	var lastPos Vect
	for len(stack) > 0 {
		lastPos = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !s.CanMove(lastPos, move) {
			break
		}

		moveTo = Vect{x: lastPos.x + move.x, y: lastPos.y + move.y}

		s.grid[moveTo.y][moveTo.x] = s.grid[lastPos.y][lastPos.x]
		s.grid[lastPos.y][lastPos.x] = EMPTY
	}
	return moveTo
}

func (s State) IsValid(pos Vect) bool {
	return pos.x >= 0 && pos.x < s.xLen && pos.y >= 0 && pos.y < s.yLen
}

func (s State) IsMovable(pos Vect) bool {
	return s.IsValid(pos) && s.grid[pos.y][pos.x] != EMPTY && s.grid[pos.y][pos.x] != WALL
}

func (s State) CanMove(pos Vect, move Vect) bool {
	nextPos := pos
	nextPos.x += move.x
	nextPos.y += move.y
	return s.IsValid(nextPos) && s.grid[nextPos.y][nextPos.x] == EMPTY
}

func (s State) SumBoxesGPSCoordinates() int {
	sum := 0
	for y := range s.grid {
		for x := range s.grid[y] {
			if s.grid[y][x] == BOX {
				sum += (100 * y) + x
			}
		}
	}
	return sum
}

func (s State) Print() {
	fmt.Println("Debug:")
	for y := range s.grid {
		for x := range s.grid[y] {
			switch s.grid[y][x] {
			case EMPTY:
				fmt.Printf(".")
			case WALL:
				fmt.Printf("#")
			case ROBOT:
				fmt.Printf("@")
			case BOX:
				fmt.Printf("O")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
