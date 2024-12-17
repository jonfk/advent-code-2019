package day16

import (
	"container/heap"
)

// PriorityQueue implementation for A* search
type PathNode struct {
	pos      Vect
	dir      Direction
	cost     int
	turns    int
	priority int
	index    int
	previous *PathNode
}

type PriorityQueue []*PathNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PathNode)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Manhattan distance heuristic
func (p Puzzle) heuristic(pos Vect) int {
	return abs(pos.x-p.end.x) + abs(pos.y-p.end.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Direction to vector mapping
var dirVectors = map[Direction]Vect{
	NORTH: {0, -1},
	SOUTH: {0, 1},
	EAST:  {1, 0},
	WEST:  {-1, 0},
}

func (p Puzzle) FindMinPathCostAStar() int {
	// Initialize visited map using direction as part of the key
	visited := make(map[Vect]map[Direction]bool)

	// Initialize priority queue with starting position
	pq := &PriorityQueue{}
	heap.Init(pq)

	startNode := &PathNode{
		pos:      p.start,
		dir:      EAST,
		cost:     0,
		turns:    0,
		priority: p.heuristic(p.start),
	}
	heap.Push(pq, startNode)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*PathNode)

		// Check if we reached the end
		if current.pos == p.end {
			return current.cost
		}

		// Skip if we've visited this position with this direction
		if visited[current.pos] != nil && visited[current.pos][current.dir] {
			continue
		}

		// Mark as visited
		if visited[current.pos] == nil {
			visited[current.pos] = make(map[Direction]bool)
		}
		visited[current.pos][current.dir] = true

		// Try all possible directions
		for newDir := range dirVectors {
			// Skip reverse direction
			if (newDir == NORTH && current.dir == SOUTH) ||
				(newDir == SOUTH && current.dir == NORTH) ||
				(newDir == EAST && current.dir == WEST) ||
				(newDir == WEST && current.dir == EAST) {
				continue
			}

			dv := dirVectors[newDir]
			newPos := Vect{current.pos.x + dv.x, current.pos.y + dv.y}

			// Skip invalid moves
			if cell := p.get(newPos); cell == INVALID || cell == WALL {
				continue
			}

			// Calculate new cost including turn penalty
			turnCost := 0
			if newDir != current.dir {
				turnCost = 1000
			}

			newNode := &PathNode{
				pos:      newPos,
				dir:      newDir,
				cost:     current.cost + 1 + turnCost,
				turns:    current.turns + 1,
				previous: current,
			}
			newNode.priority = newNode.cost + p.heuristic(newPos)

			heap.Push(pq, newNode)
		}
	}

	return -1 // No path found
}

// PathSet stores all paths that reach the end with minimum cost
type PathSet struct {
	paths []*PathNode
	cost  int
}

// FindAllMinCostPaths finds all paths with minimum cost and counts unique spots using Dijkstra
func (p Puzzle) FindAllMinCostPaths() (int, int) {
	// Initialize priority queue with starting position
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Track minimum cost to each position+direction combination
	costs := make(map[Vect]map[Direction]int)

	// Track all paths that reach the end with minimum cost
	var endPaths []*PathNode
	var minEndCost = -1

	startNode := &PathNode{
		pos:      p.start,
		dir:      EAST,
		cost:     0,
		priority: 0, // No heuristic for Dijkstra
	}
	heap.Push(pq, startNode)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*PathNode)

		// If we've found end paths and current cost is higher than min cost, we can stop
		if minEndCost != -1 && current.cost > minEndCost {
			break
		}

		// Check if we reached the end
		if current.pos == p.end {
			if minEndCost == -1 || current.cost == minEndCost {
				endPaths = append(endPaths, current)
				minEndCost = current.cost
			}
			continue
		}

		// Check if we've found a better or equal cost path to this position+direction
		if costs[current.pos] != nil {
			if prevCost, exists := costs[current.pos][current.dir]; exists {
				if current.cost > prevCost {
					continue
				}
			}
		}

		// Update cost for this position+direction
		if costs[current.pos] == nil {
			costs[current.pos] = make(map[Direction]int)
		}
		costs[current.pos][current.dir] = current.cost

		// Try all possible directions
		for newDir := range dirVectors {
			// Skip reverse direction
			if (newDir == NORTH && current.dir == SOUTH) ||
				(newDir == SOUTH && current.dir == NORTH) ||
				(newDir == EAST && current.dir == WEST) ||
				(newDir == WEST && current.dir == EAST) {
				continue
			}

			dv := dirVectors[newDir]
			newPos := Vect{current.pos.x + dv.x, current.pos.y + dv.y}

			// Skip invalid moves
			if cell := p.get(newPos); cell == INVALID || cell == WALL {
				continue
			}

			// Calculate new cost including turn penalty
			turnCost := 0
			if newDir != current.dir {
				turnCost = 1000
			}

			newCost := current.cost + 1 + turnCost

			// Only add if this cost is not higher than any known end path
			if minEndCost == -1 || newCost <= minEndCost {
				newNode := &PathNode{
					pos:      newPos,
					dir:      newDir,
					cost:     newCost,
					priority: newCost, // Priority is just the cost for Dijkstra
					previous: current,
				}
				heap.Push(pq, newNode)
			}
		}
	}

	// Count unique spots from all minimum cost paths
	uniqueSpots := make(map[Vect]bool)
	for _, path := range endPaths {
		// Traverse back through the path and mark all spots
		current := path
		for current != nil {
			uniqueSpots[current.pos] = true
			current = current.previous
		}
	}

	if minEndCost == -1 {
		return -1, 0
	}
	return minEndCost, len(uniqueSpots)
}
