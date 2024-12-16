package day15

import (
	"fmt"
)

type Cell2 int

const (
	EMPTY2 Cell2 = iota
	BOXSTART
	BOXEND
	WALL2
	ROBOT2
)

type State2 struct {
	grid       [][]Cell2
	robotPos   Vect
	moves      []Vect
	xLen, yLen int
}

// Convert converts a State to a State2, doubling the width of the map
func (s *State) Convert() *State2 {
	// Create a new grid with double width
	newGrid := make([][]Cell2, s.yLen)
	for y := range newGrid {
		newGrid[y] = make([]Cell2, s.xLen*2)
	}

	// Fill the new grid based on conversion rules
	var newRobotPos Vect
	for y := 0; y < s.yLen; y++ {
		for x := 0; x < s.xLen; x++ {
			switch s.grid[y][x] {
			case WALL:
				newGrid[y][x*2] = WALL2
				newGrid[y][x*2+1] = WALL2
			case BOX:
				newGrid[y][x*2] = BOXSTART
				newGrid[y][x*2+1] = BOXEND
			case ROBOT:
				newGrid[y][x*2] = ROBOT2
				newGrid[y][x*2+1] = EMPTY2
				newRobotPos = Vect{x: x * 2, y: y}
			case EMPTY:
				newGrid[y][x*2] = EMPTY2
				newGrid[y][x*2+1] = EMPTY2
			}
		}
	}

	return &State2{
		grid:     newGrid,
		robotPos: newRobotPos,
		moves:    s.moves,
		xLen:     s.xLen * 2,
		yLen:     s.yLen,
	}
}

func (s *State2) ConsumeAllMoves() {
	for len(s.moves) > 0 {
		// s.PrintNextMove()
		// s.Print()
		s.ConsumeMove()
	}
}

func (s *State2) ConsumeMove() bool {
	if len(s.moves) == 0 {
		return false
	}
	move := s.moves[0]
	s.moves = s.moves[1:]

	newRobotPos := s.Move(s.robotPos, move)
	s.robotPos = newRobotPos

	return true
}

func (s State2) IsValid(pos Vect) bool {
	return pos.x >= 0 && pos.x < s.xLen && pos.y >= 0 && pos.y < s.yLen
}

func (s State2) IsMovable(pos Vect) bool {
	if !s.IsValid(pos) {
		return false
	}
	cell := s.grid[pos.y][pos.x]
	return cell != EMPTY2 && cell != WALL2 && cell != BOXEND
}

func (s State2) CanMove(pos Vect, move Vect) bool {
	nextPos := Vect{x: pos.x + move.x, y: pos.y + move.y}
	return s.IsValid(nextPos) && s.grid[nextPos.y][nextPos.x] == EMPTY2
}

func (s State2) SumBoxesGPSCoordinates() int {
	sum := 0
	for y := range s.grid {
		for x := range s.grid[y] {
			// Only count BOXSTART positions
			if s.grid[y][x] == BOXSTART {
				sum += (100 * y) + (x)
			}
		}
	}
	return sum
}

func (s State2) Print() {
	fmt.Println("Debug:")
	for y := range s.grid {
		for x := range s.grid[y] {
			switch s.grid[y][x] {
			case EMPTY2:
				fmt.Printf(".")
			case WALL2:
				fmt.Printf("#")
			case ROBOT2:
				fmt.Printf("@")
			case BOXSTART:
				fmt.Printf("[")
			case BOXEND:
				fmt.Printf("]")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (s State2) PrintNextMove() {
	if len(s.moves) == 0 {
		fmt.Println("No more moves")
		return
	}

	move := s.moves[0]
	var direction string
	switch {
	case move.x == -1:
		direction = "left"
	case move.x == 1:
		direction = "right"
	case move.y == -1:
		direction = "up"
	case move.y == 1:
		direction = "down"
	}

	fmt.Printf("Next move: %s (%+v)\n", direction, move)
	fmt.Printf("Robot currently at: (%d, %d)\n", s.robotPos.x/2, s.robotPos.y)
}

type Box struct {
	start Vect // Left or top position
	end   Vect // Right or bottom position
}

// findBoxAt returns the complete box if pos is part of a box, nil otherwise
func (s State2) findBoxAt(pos Vect) *Box {
	if !s.IsValid(pos) {
		return nil
	}

	cell := s.grid[pos.y][pos.x]
	if cell != BOXSTART && cell != BOXEND {
		return nil
	}

	// If we found a BOXEND, look left for the BOXSTART
	if cell == BOXEND {
		if s.IsValid(Vect{x: pos.x - 1, y: pos.y}) && s.grid[pos.y][pos.x-1] == BOXSTART {
			return &Box{
				start: Vect{x: pos.x - 1, y: pos.y},
				end:   pos,
			}
		}
		return nil
	}

	// If we found a BOXSTART, look right for the BOXEND
	if s.IsValid(Vect{x: pos.x + 1, y: pos.y}) && s.grid[pos.y][pos.x+1] == BOXEND {
		return &Box{
			start: pos,
			end:   Vect{x: pos.x + 1, y: pos.y},
		}
	}
	return nil
}

// findMovableBoxes returns all boxes that would be pushed by moving from pos in direction move
func (s State2) findMovableBoxes(pos Vect, move Vect) []Box {
	var boxes []Box
	var checked = make(map[Vect]bool)
	toCheck := []Vect{pos}

	for len(toCheck) > 0 {
		current := toCheck[0]
		toCheck = toCheck[1:]

		if checked[current] {
			continue
		}
		checked[current] = true

		// If it's part of a box, add the box and the next position to check
		if box := s.findBoxAt(current); box != nil {
			boxes = append(boxes, *box)

			// Add positions adjacent to both ends of the box in the direction of movement
			nextStart := Vect{x: box.start.x + move.x, y: box.start.y + move.y}
			nextEnd := Vect{x: box.end.x + move.x, y: box.end.y + move.y}

			if !checked[nextStart] {
				toCheck = append(toCheck, nextStart)
			}
			if !checked[nextEnd] {
				toCheck = append(toCheck, nextEnd)
			}
		}
	}

	return boxes
}

// canMoveBoxes checks if all boxes can be moved in the given direction
func (s State2) canMoveBoxes(boxes []Box, move Vect) bool {
	// Create a map of current box positions to check for collisions
	occupied := make(map[Vect]bool)
	for _, box := range boxes {
		occupied[box.start] = true
		occupied[box.end] = true
	}

	// Check if each box can move
	for _, box := range boxes {
		newStart := Vect{x: box.start.x + move.x, y: box.start.y + move.y}
		newEnd := Vect{x: box.end.x + move.x, y: box.end.y + move.y}

		// Check if new positions are valid and empty (or part of a moving box)
		if !s.IsValid(newStart) || !s.IsValid(newEnd) {
			return false
		}

		if !occupied[newStart] {
			if s.grid[newStart.y][newStart.x] != EMPTY2 {
				return false
			}
		}
		if !occupied[newEnd] {
			if s.grid[newEnd.y][newEnd.x] != EMPTY2 {
				return false
			}
		}
	}
	return true
}

func (s *State2) Move(pos Vect, move Vect) Vect {
	nextPos := Vect{x: pos.x + move.x, y: pos.y + move.y}

	// If next position is empty, just move the robot
	if s.IsValid(nextPos) && s.grid[nextPos.y][nextPos.x] == EMPTY2 {
		s.grid[nextPos.y][nextPos.x] = ROBOT2
		s.grid[pos.y][pos.x] = EMPTY2
		return nextPos
	}

	// Find all boxes that would be moved
	boxes := s.findMovableBoxes(nextPos, move)

	// If no boxes found or boxes can't be moved, stay in place
	if len(boxes) == 0 || !s.canMoveBoxes(boxes, move) {
		return pos
	}

	// Move all boxes
	for _, box := range boxes {
		// Clear old positions
		s.grid[box.start.y][box.start.x] = EMPTY2
		s.grid[box.end.y][box.end.x] = EMPTY2
	}

	// Place boxes in new positions
	for _, box := range boxes {
		newStart := Vect{x: box.start.x + move.x, y: box.start.y + move.y}
		newEnd := Vect{x: box.end.x + move.x, y: box.end.y + move.y}
		s.grid[newStart.y][newStart.x] = BOXSTART
		s.grid[newEnd.y][newEnd.x] = BOXEND
	}

	// Move robot
	s.grid[nextPos.y][nextPos.x] = ROBOT2
	s.grid[pos.y][pos.x] = EMPTY2

	return nextPos
}
