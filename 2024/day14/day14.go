package day14

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Robot struct {
	Pos Vect
	Vel Vect
}

type Vect struct {
	X int
	Y int
}

var patternRegex = regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

func Parse(input string) []Robot {
	var robots []Robot
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		var robot Robot
		var err error

		matches := patternRegex.FindStringSubmatch(line)
		if len(matches) != 5 { // Full match + capture groups
			log.Fatalf("invalid input format. expected = 5, got = %v", len(matches))
		}

		robot.Pos.X, err = strconv.Atoi(matches[1])
		if err != nil {
			log.Fatalf("invalid robot.Pos.X: %v", err)
		}
		robot.Pos.Y, err = strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalf("invalid robot.Pos.Y: %v", err)
		}
		robot.Vel.X, err = strconv.Atoi(matches[3])
		if err != nil {
			log.Fatalf("invalid robot.Vel.X: %v", err)
		}
		robot.Vel.Y, err = strconv.Atoi(matches[4])
		if err != nil {
			log.Fatalf("invalid robot.Vel.Y: %v", err)
		}

		robots = append(robots, robot)
	}

	return robots
}

func Tick(times int, robots []Robot, xLen, yLen int) {
	// var variances []float64
	for n := 0; n < times; n++ {
		// for i := range robots {
		// 	robots[i].Pos.X = robots[i].Pos.X + robots[i].Vel.X
		// 	for robots[i].Pos.X >= xLen {
		// 		robots[i].Pos.X = robots[i].Pos.X - xLen
		// 	}
		// 	for robots[i].Pos.X < 0 {
		// 		robots[i].Pos.X = robots[i].Pos.X + xLen
		// 	}
		// 	robots[i].Pos.Y = robots[i].Pos.Y + robots[i].Vel.Y
		// 	for robots[i].Pos.Y >= yLen {
		// 		robots[i].Pos.Y = robots[i].Pos.Y - yLen
		// 	}
		// 	for robots[i].Pos.Y < 0 {
		// 		robots[i].Pos.Y = robots[i].Pos.Y + yLen
		// 	}
		// }
		for i := range robots {
			robots[i].Pos.X = ((robots[i].Pos.X+robots[i].Vel.X)%xLen + xLen) % xLen
			robots[i].Pos.Y = ((robots[i].Pos.Y+robots[i].Vel.Y)%yLen + yLen) % yLen
		}
		// variance := calculateVariance(robots)
		// variances = append(variances, variance)
		// fmt.Printf("variance: %v\n", variance)
		if hasLine(9, robots, xLen, yLen) {
			Print(n, robots, xLen, yLen)
			break
		}
	}
	// var lowestVariance float64 = variances[0]
	// var lowestVarianceIdx int = 0
	//
	// for i := range variances {
	// 	if variances[i] < lowestVariance {
	// 		lowestVariance = variances[i]
	// 		lowestVarianceIdx = i
	// 	}
	// }
	// fmt.Printf("Lowest variance occurs at %v with %v\n", lowestVarianceIdx, lowestVariance)
	// printTopTenLowestVariances(variances)

}

func Print(iteration int, robots []Robot, xLen, yLen int) {
	screen := make([][]int, yLen)

	for i := range yLen {
		screen[i] = make([]int, xLen)
	}

	for i := range robots {
		screen[robots[i].Pos.Y][robots[i].Pos.X] += 1
	}

	fmt.Println("Iteration: " + strconv.Itoa(iteration+1))
	for y := range screen {
		for x := range screen[y] {
			if screen[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("a")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func CalculateSafetyFactor(robots []Robot, xLen, yLen int) int {
	var q1, q2, q3, q4 int

	for i := range robots {
		if robots[i].Pos.X < xLen/2 && robots[i].Pos.Y < yLen/2 {
			q1 += 1
		}
		if robots[i].Pos.X > xLen/2 && robots[i].Pos.Y < yLen/2 {
			q2 += 1
		}
		if robots[i].Pos.X < xLen/2 && robots[i].Pos.Y > yLen/2 {
			q3 += 1
		}
		if robots[i].Pos.X > xLen/2 && robots[i].Pos.Y > yLen/2 {
			q4 += 1
		}
	}
	log.Printf("q1: %v, q2: %v, q3: %v, q4: %v", q1, q2, q3, q4)

	return q1 * q2 * q3 * q4
}

func calculateVariance(robots []Robot) float64 {
	if len(robots) <= 1 {
		return 0
	}

	// Calculate sum of minimum distances
	var sumMinDistances float64
	for i, robot := range robots {
		// Initialize minDist to a large number
		minDist := math.MaxFloat64

		// Find closest point
		for j, other := range robots {
			if i == j {
				continue // Skip self
			}

			// Calculate Euclidean distance
			dx := float64(robot.Pos.X - other.Pos.X)
			dy := float64(robot.Pos.Y - other.Pos.Y)
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist < minDist {
				minDist = dist
			}
		}

		sumMinDistances += minDist
	}

	// Calculate average minimum distance
	variance := sumMinDistances / float64(len(robots))
	return variance
}

type VarianceItem struct {
	Index    int
	Variance float64
}

func printTopTenLowestVariances(variances []float64) {
	// Create slice of VarianceItems
	items := make([]VarianceItem, len(variances))
	for i, v := range variances {
		items[i] = VarianceItem{
			Index:    i,
			Variance: v,
		}
	}

	// Sort by variance
	sort.Slice(items, func(i, j int) bool {
		return items[i].Variance < items[j].Variance
	})

	// Print top 10 (or less if array is smaller)
	count := 10
	if len(items) < 10 {
		count = len(items)
	}

	fmt.Println("Top", count, "lowest variances:")
	fmt.Println("Rank\tIndex\tVariance")
	fmt.Println("------------------------")
	for i := 0; i < count; i++ {
		fmt.Printf("%d\t%d\t%.4f\n", i+1, items[i].Index, items[i].Variance)
	}
}

func hasLine(minLength int, robots []Robot, xLen, yLen int) bool {
	// Create the screen
	screen := make([][]int, yLen)
	for i := range screen {
		screen[i] = make([]int, xLen)
	}

	// Fill the screen with robot positions
	for i := range robots {
		screen[robots[i].Pos.Y][robots[i].Pos.X] += 1
	}

	// Check horizontal lines
	for y := 0; y < yLen; y++ {
		count := 0
		for x := 0; x < xLen; x++ {
			if screen[y][x] > 0 {
				count++
				if count >= minLength {
					return true
				}
			} else {
				count = 0
			}
		}
	}

	// Check vertical lines
	for x := 0; x < xLen; x++ {
		count := 0
		for y := 0; y < yLen; y++ {
			if screen[y][x] > 0 {
				count++
				if count >= minLength {
					return true
				}
			} else {
				count = 0
			}
		}
	}

	return false
}
