package day2

import (
	"strconv"
	"strings"
)

type Input struct {
	reports [][]int
}

func Parse(input string) (*Input, error) {
	var res Input
	for _, report := range strings.Split(input, "\n") {
		var levels []int
		for _, level := range strings.Split(report, " ") {
			ilevel, err := strconv.Atoi(level)
			if err != nil {
				return nil, err
			}
			levels = append(levels, ilevel)
		}
		res.reports = append(res.reports, levels)
	}
	return &res, nil
}

func CountSafe(input *Input) int {
	safeCount := 0
	for _, report := range input.reports {
		if IsSafe(report) {
			safeCount += 1
		}
	}
	return safeCount
}

func IsSafe(report []int) bool {
	var direction int = 0
	for i, lvl := range report {
		if i != 0 && lvl == report[i-1] {
			return false
		}
		if i != 0 && direction == 0 {
			if lvl > report[i-1] {
				direction = 1
			}
			if lvl < report[i-1] {
				direction = -1
			}
		}

		if i != 0 && direction == 1 && (lvl > (report[i-1]+3) || lvl < report[i-1]) {
			return false
		}
		if i != 0 && direction == -1 && (lvl < (report[i-1]-3) || lvl > report[i-1]) {
			return false
		}
	}
	return true
}

func IsSafeWithDampenerBruteForce(report []int) bool {
	if IsSafe(report) {
		return true
	} else {
		for i := range report {
			var dampenedReport []int = make([]int, len(report))
			copy(dampenedReport, report)
			dampenedReport = append(dampenedReport[:i], dampenedReport[i+1:]...)

			if IsSafe(dampenedReport) {
				return true
			}
		}
		return false
	}
}

func CountSafeWithDampenerBruteForce(input *Input) int {
	safeCount := 0
	for _, report := range input.reports {
		if IsSafeWithDampenerBruteForce(report) {
			safeCount += 1
		}
	}
	return safeCount
}

// WARN: Doesn't account for all edge cases with is why it doesn't give the right answer
func CountSafeWithDampener(input *Input) int {
	safeCount := 0
	for _, report := range input.reports {
		if IsSafeWithDampener(report) {
			safeCount += 1
		}
	}
	return safeCount
}

func IsSafeWithDampener(report []int) bool {
	var direction int = 0
	var viewedLvls []int
	var levelRemoved bool

	for i := 0; i < len(report); {
		lvl := report[i]

		if len(viewedLvls) == 0 {
			viewedLvls = append(viewedLvls, lvl)
			i++
			continue
		}
		prev := viewedLvls[len(viewedLvls)-1]

		if lvl == prev {
			if !levelRemoved {
				levelRemoved = true
				viewedLvls = viewedLvls[:len(viewedLvls)-1]
				continue
			} else {
				return false
			}
		}
		if direction == 0 {
			if lvl > prev {
				direction = 1
			}
			if lvl < prev {
				direction = -1
			}
		}

		if direction == 1 && (lvl > (prev+3) || lvl < prev) {
			if !levelRemoved {
				levelRemoved = true
				viewedLvls = viewedLvls[:len(viewedLvls)-1]
				continue
			} else {
				return false
			}
		}
		if i != 0 && direction == -1 && (lvl < (prev-3) || lvl > prev) {
			if !levelRemoved {
				levelRemoved = true
				viewedLvls = viewedLvls[:len(viewedLvls)-1]
				continue
			} else {
				return false
			}
		}

		viewedLvls = append(viewedLvls, lvl)
		i++
	}
	return true
}
