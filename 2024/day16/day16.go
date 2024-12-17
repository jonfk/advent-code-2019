package day16

import (
	"fmt"
	"log"
	"slices"
	"strings"
)

type Direction int

const (
	NORTH Direction = iota
	SOUTH
	EAST
	WEST
)

var directions = []struct {
	dir Direction
	dx  int
	dy  int
}{
	{NORTH, 0, -1},
	{SOUTH, 0, 1},
	{EAST, 1, 0},
	{WEST, -1, 0},
}

type Cell int

const (
	WALL Cell = iota
	EMPTY
	START
	END
	INVALID
)

type Puzzle struct {
	start Vect
	end   Vect
	grid  [][]Cell
}

type Vect struct {
	x, y int
}

type Path struct {
	points     []Vect
	directions []Direction
}

func (p Path) Clone() Path {
	points := make([]Vect, len(p.points))
	copy(points, p.points)
	directions := make([]Direction, len(p.directions))
	copy(directions, p.directions)
	return Path{points, directions}
}

func (p Path) Cost() int {
	if len(p.points) == 0 {
		return 0
	}
	var cost int
	var lastDir Direction = p.directions[0]
	for i := range p.directions {
		if p.directions[i] != lastDir {
			cost += 1000
			lastDir = p.directions[i]
		}
	}
	return cost + len(p.points) - 1
}

type VectDir struct {
	v Vect
	d Direction
}

func Parse(input string) Puzzle {
	var puzzle Puzzle

	for y, line := range strings.Split(input, "\n") {
		var row []Cell
		for x, ch := range line {
			switch ch {
			case '#':
				row = append(row, WALL)
			case '.':
				row = append(row, EMPTY)
			case 'S':
				row = append(row, START)
				puzzle.start = Vect{x, y}
			case 'E':
				row = append(row, END)
				puzzle.end = Vect{x, y}
			}
		}
		puzzle.grid = append(puzzle.grid, row)
	}
	if puzzle.get(puzzle.start) != START || puzzle.get(puzzle.end) != END {
		log.Fatal("Incorrect grid")
	}
	return puzzle
}

func (p Puzzle) get(v Vect) Cell {
	if v.x >= 0 && v.x < len(p.grid[0]) && v.y >= 0 && v.y < len(p.grid) {
		return p.grid[v.y][v.x]
	} else {
		return INVALID
	}
}

// State represents a position and direction
type State struct {
	pos       Vect
	dir       Direction
	cost      int
	prevState *State
}

func (p Puzzle) FindMinPathCostBFS() int {

	// Initialize queue with starting state towards the east
	queue := make([]*State, 0, len(directions))
	queue = append(queue, &State{
		pos:  p.start,
		dir:  EAST,
		cost: 0,
	})

	// Use a map for visited states
	visited := make(map[Vect]map[Direction]bool)
	minCost := -1

	// Process queue
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Skip if we've seen this state before
		if visited[current.pos] != nil && visited[current.pos][current.dir] {
			continue
		}

		// Mark as visited
		if visited[current.pos] == nil {
			visited[current.pos] = make(map[Direction]bool)
		}
		visited[current.pos][current.dir] = true

		// Check if we reached the end
		if p.get(current.pos) == END {
			if minCost == -1 || current.cost < minCost {
				minCost = current.cost
			}
			continue
		}

		// Try each direction
		for _, d := range directions {
			// Skip reverse direction
			if (d.dir == NORTH && current.dir == SOUTH) ||
				(d.dir == SOUTH && current.dir == NORTH) ||
				(d.dir == EAST && current.dir == WEST) ||
				(d.dir == WEST && current.dir == EAST) {
				continue
			}

			newPos := Vect{current.pos.x + d.dx, current.pos.y + d.dy}

			// Skip invalid moves
			if cell := p.get(newPos); cell == INVALID || cell == WALL {
				continue
			}

			// Calculate new cost including turn penalty
			newCost := current.cost + 1
			if d.dir != current.dir {
				newCost += 1000
			}

			// If we've found a solution and this path is already more expensive, skip it
			if minCost != -1 && newCost >= minCost {
				continue
			}

			newState := &State{
				pos:       newPos,
				dir:       d.dir,
				cost:      newCost,
				prevState: current,
			}
			queue = append(queue, newState)
		}
	}

	return minCost
}

func (p Puzzle) nextSteps(path Path) []VectDir {
	var next []VectDir
	last := path.points[len(path.points)-1]

	west := Vect{x: last.x - 1, y: last.y}
	east := Vect{x: last.x + 1, y: last.y}
	north := Vect{x: last.x, y: last.y - 1}
	south := Vect{x: last.x, y: last.y + 1}

	westCell := p.get(Vect{x: last.x - 1, y: last.y})
	eastCell := p.get(Vect{x: last.x + 1, y: last.y})
	northCell := p.get(Vect{x: last.x, y: last.y - 1})
	southCell := p.get(Vect{x: last.x, y: last.y + 1})

	if westCell != INVALID && westCell != WALL && !slices.Contains(path.points, west) {
		next = append(next, VectDir{west, WEST})
	}
	if eastCell != INVALID && eastCell != WALL && !slices.Contains(path.points, east) {
		next = append(next, VectDir{east, EAST})
	}
	if northCell != INVALID && northCell != WALL && !slices.Contains(path.points, north) {
		next = append(next, VectDir{north, NORTH})
	}
	if southCell != INVALID && southCell != WALL && !slices.Contains(path.points, south) {
		next = append(next, VectDir{south, SOUTH})
	}

	return next
}

// Reconstructs the path from a final state
func reconstructPath(state *State) []Vect {
	var path []Vect
	for current := state; current != nil; current = current.prevState {
		path = append([]Vect{current.pos}, path...)
	}
	return path
}

func (p Puzzle) MinPaths() [][]Vect {
	queue := make([]*State, 0, len(directions))
	for _, d := range directions {
		queue = append(queue, &State{
			pos:  p.start,
			dir:  d.dir,
			cost: 0,
		})
	}

	visited := make(map[Vect]map[Direction]bool)
	var minCost int = -1
	var minPaths [][]Vect

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.pos] != nil && visited[current.pos][current.dir] {
			continue
		}

		if visited[current.pos] == nil {
			visited[current.pos] = make(map[Direction]bool)
		}
		visited[current.pos][current.dir] = true

		if p.get(current.pos) == END {
			if minCost == -1 || current.cost < minCost {
				minCost = current.cost
				minPaths = [][]Vect{reconstructPath(current)}
			} else if current.cost == minCost {
				minPaths = append(minPaths, reconstructPath(current))
			}
			continue
		}

		for _, d := range directions {
			if (d.dir == NORTH && current.dir == SOUTH) ||
				(d.dir == SOUTH && current.dir == NORTH) ||
				(d.dir == EAST && current.dir == WEST) ||
				(d.dir == WEST && current.dir == EAST) {
				continue
			}

			newPos := Vect{current.pos.x + d.dx, current.pos.y + d.dy}

			if cell := p.get(newPos); cell == INVALID || cell == WALL {
				continue
			}

			newCost := current.cost + 1
			if d.dir != current.dir {
				newCost += 1000
			}

			if minCost != -1 && newCost > minCost {
				continue
			}

			newState := &State{
				pos:       newPos,
				dir:       d.dir,
				cost:      newCost,
				prevState: current,
			}
			queue = append(queue, newState)
		}
	}

	return minPaths
}

func (p Puzzle) MinPathsSpots() int {
	minPaths := p.MinPaths()
	spots := make(map[Vect]bool)

	// Add all points from all minimum paths
	for _, path := range minPaths {
		for _, point := range path {
			spots[point] = true
		}
	}

	return len(spots)
}

func (p Puzzle) PrintMinPaths() {
	minPaths := p.MinPaths()
	spots := make(map[Vect]bool)

	// Mark all spots in minimum paths
	for _, path := range minPaths {
		for _, point := range path {
			spots[point] = true
		}
	}

	// Print the grid
	for y := 0; y < len(p.grid); y++ {
		for x := 0; x < len(p.grid[0]); x++ {
			pos := Vect{x, y}
			if spots[pos] {
				fmt.Print("0")
			} else if p.grid[y][x] == WALL {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
