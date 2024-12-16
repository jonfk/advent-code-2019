package day13

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type PrizeMachine struct {
	A     Coord
	B     Coord
	Prize Coord
}

type Coord struct {
	X int
	Y int
}

type PrizeMachineResult struct {
	A int
	B int
}

func Parse(input string) []PrizeMachine {
	var res []PrizeMachine
	for _, prizeMachineStr := range strings.Split(input, "\n\n") {
		prizeMachine, err := ParsePrizeMachine(prizeMachineStr)
		if err != nil {
			log.Fatalf("error parsing prize machine: %s", err)
		}
		res = append(res, prizeMachine)
	}
	return res
}

var patternRegex = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)

func ParsePrizeMachine(input string) (PrizeMachine, error) {
	matches := patternRegex.FindStringSubmatch(input)
	if len(matches) != 7 { // Full match + 6 capture groups
		return PrizeMachine{}, fmt.Errorf("invalid input format")
	}

	// Convert string matches to integers
	aX, err := strconv.Atoi(matches[1])
	if err != nil {
		return PrizeMachine{}, fmt.Errorf("invalid A.X value: %v", err)
	}
	aY, err := strconv.Atoi(matches[2])
	if err != nil {
		return PrizeMachine{}, fmt.Errorf("invalid A.Y value: %v", err)
	}
	bX, err := strconv.Atoi(matches[3])
	if err != nil {
		return PrizeMachine{}, fmt.Errorf("invalid B.X value: %v", err)
	}
	bY, err := strconv.Atoi(matches[4])
	if err != nil {
		return PrizeMachine{}, fmt.Errorf("invalid B.Y value: %v", err)
	}
	prizeX, err := strconv.Atoi(matches[5])
	if err != nil {
		return PrizeMachine{}, fmt.Errorf("invalid Prize.X value: %v", err)
	}
	prizeY, err := strconv.Atoi(matches[6])
	if err != nil {
		return PrizeMachine{}, fmt.Errorf("invalid Prize.Y value: %v", err)
	}

	return PrizeMachine{
		A:     Coord{X: aX, Y: aY},
		B:     Coord{X: bX, Y: bY},
		Prize: Coord{X: prizeX, Y: prizeY},
	}, nil
}

func FindPrize(a, b, prize int) (PrizeMachineResult, bool) {
	bN := prize / b
	for bN >= 0 {
		remaining := prize - (b * bN)
		if remaining%a == 0 {
			aN := remaining / a
			// if aN+bN > 100 {
			// 	return PrizeMachineResult{}, false
			// }
			return PrizeMachineResult{A: aN, B: bN}, true
		}
		bN -= 1
	}
	return PrizeMachineResult{}, false
}

func FindMinTokenForAllPossiblePrizes(machines []PrizeMachine) int {
	tokens := 0
	for _, machine := range machines {
		countA, countB, possible := SolvePrizeSystem(machine.Prize.X, machine.Prize.Y, machine.A.X, machine.B.X, machine.A.Y, machine.B.Y)
		if possible {
			log.Printf("countA: %v, countB: %v\n", countA, countB)
			tokens += (countA * 3) + countB
		}
	}
	return tokens
}

func FindMinTokenForAllPossiblePrizesCorrected(machines []PrizeMachine) int {
	tokens := 0
	for _, machine := range machines {
		countA, countB, possible := SolvePrizeSystem2(machine.Prize.X+10000000000000, machine.Prize.Y+10000000000000, machine.A.X, machine.B.X, machine.A.Y, machine.B.Y)
		if possible {
			log.Printf("countA: %v, countB: %v\n", countA, countB)
			tokens += (countA * 3) + countB
		}
	}
	return tokens
}

func SolvePrizeSystem(xPrize, yPrize, Ax, Bx, Ay, By int) (countA int, countB int, possible bool) {
	// Calculate initial maximum countB
	maxCountB := xPrize / Bx

	// Try each possible countB value from highest to lowest
	for countB = maxCountB; countB >= 0; countB-- {
		// Calculate potential countA
		remainder := (xPrize - countB*Bx) % Ax

		if remainder != 0 {
			continue
		}

		countA = (xPrize - countB*Bx) / Ax
		if countA < 0 {
			continue
		}

		// Verify yPrize equation
		calculatedYPrize := countA*Ay + countB*By
		if calculatedYPrize == yPrize {
			return countA, countB, true
		}
	}
	return 0, 0, false
}

func SolvePrizeSystem2(xPrize, yPrize, Ax, Bx, Ay, By int) (countA int, countB int, possible bool) {
	fmt.Printf("\nSolving system of linear equations:\n")
	fmt.Printf("%d = %d*countA + %d*countB\n", xPrize, Ax, Bx)
	fmt.Printf("%d = %d*countA + %d*countB\n", yPrize, Ay, By)

	// Using Cramer's rule to solve the system
	// The system can be written as:
	// Ax*countA + Bx*countB = xPrize
	// Ay*countA + By*countB = yPrize

	determinant := Ax*By - Ay*Bx
	if determinant == 0 {
		fmt.Printf("System has no unique solution (determinant = 0)\n")
		return 0, 0, false
	}

	// Calculate numerators for countA and countB
	countANumerator := xPrize*By - yPrize*Bx
	countBNumerator := Ax*yPrize - Ay*xPrize

	fmt.Printf("\nCalculation steps:\n")
	fmt.Printf("1. Determinant = %d*%d - %d*%d = %d\n", Ax, By, Ay, Bx, determinant)
	fmt.Printf("2. countA numerator = %d*%d - %d*%d = %d\n", xPrize, By, yPrize, Bx, countANumerator)
	fmt.Printf("3. countB numerator = %d*%d - %d*%d = %d\n", Ax, yPrize, Ay, xPrize, countBNumerator)

	// Check if we'll get integer solutions
	if countANumerator%determinant != 0 || countBNumerator%determinant != 0 {
		fmt.Printf("\nNo integer solution exists:\n")
		fmt.Printf("countA = %d/%d (not divisible)\n", countANumerator, determinant)
		fmt.Printf("countB = %d/%d (not divisible)\n", countBNumerator, determinant)
		return 0, 0, false
	}

	// Calculate final values
	countA = countANumerator / determinant
	countB = countBNumerator / determinant

	fmt.Printf("\nSolution found:\n")
	fmt.Printf("countA = %d/%d = %d\n", countANumerator, determinant, countA)
	fmt.Printf("countB = %d/%d = %d\n", countBNumerator, determinant, countB)

	return countA, countB, true
}
