package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type Input struct {
	Rules   [][]int
	Updates [][]int
}

func Run(input string) (int, error) {
	parsedInput, err := Parse(input)
	if err != nil {
		return 0, err
	}

	correctUpdates := FindCorrectUpdates(parsedInput)

	var sum int
	for _, update := range correctUpdates {
		mid := update[len(update)/2]
		sum += mid
	}
	return sum, nil
}

func Run2(input string) (int, error) {
	parsedInput, err := Parse(input)
	if err != nil {
		return 0, err
	}

	var pagesBefore map[int][]int = make(map[int][]int)
	var incorrectUpdates [][]int

	for _, rule := range parsedInput.Rules {
		pagesBefore[rule[1]] = append(pagesBefore[rule[1]], rule[0])
	}

	for _, update := range parsedInput.Updates {
		if !IsCorrectUpdate(pagesBefore, update) {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	var sum int
	for i := range incorrectUpdates {
		update := incorrectUpdates[i]
		// fmt.Printf("before sort: %v\n", update)
		Sort(pagesBefore, update)
		// fmt.Printf("after  sort: %v\n", update)

		mid := update[len(update)/2]
		sum += mid
	}
	return sum, nil
}

func Parse(input string) (*Input, error) {
	var rules, updates [][]int

	parseRules := true
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			parseRules = false
			continue
		}
		if parseRules {
			rule := strings.Split(line, "|")
			if len(rule) != 2 {
				return nil, fmt.Errorf("Rule longer than 2: %v\n", rule)
			}
			before, err := strconv.Atoi(rule[0])
			if err != nil {
				return nil, fmt.Errorf("Rule not int: %w", err)
			}
			after, err := strconv.Atoi(rule[1])
			if err != nil {
				return nil, fmt.Errorf("Rule not int: %w", err)
			}
			rules = append(rules, []int{before, after})
		} else {
			updateStrs := strings.Split(line, ",")
			var update []int
			for _, page := range updateStrs {
				pageI, err := strconv.Atoi(page)
				if err != nil {
					return nil, fmt.Errorf("Update Page not int: %w", err)
				}
				update = append(update, pageI)
			}
			updates = append(updates, update)
		}
	}

	return &Input{Rules: rules, Updates: updates}, nil
}

func FindCorrectUpdates(input *Input) [][]int {
	var pagesBefore map[int][]int = make(map[int][]int)
	var correctUpdates [][]int

	for _, rule := range input.Rules {
		pagesBefore[rule[1]] = append(pagesBefore[rule[1]], rule[0])
	}

	for _, update := range input.Updates {
		if IsCorrectUpdate(pagesBefore, update) {
			correctUpdates = append(correctUpdates, update)
		}
	}

	return correctUpdates
}

func IsCorrectUpdate(pagesBeforeRules map[int][]int, update []int) bool {
	for i, page := range update {
		pagesNeededBefore := pagesBeforeRules[page]
		if ContainsAny(update[i+1:], pagesNeededBefore) {
			return false
		}
	}
	return true
}

func ContainsAny(slice []int, subslice []int) bool {
	var elems map[int]bool = make(map[int]bool)

	for _, e := range slice {
		elems[e] = true
	}

	for _, x := range subslice {
		if elems[x] {
			return true
		}
	}
	return false
}

func Contains(slice []int, elem int) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}

func Sort(pagesBeforeRules map[int][]int, update []int) {
	quicksort(pagesBeforeRules, update, 0, len(update)-1)
}

func isABeforeB(pagesBeforeRules map[int][]int, update []int, a, b int) bool {
	before := pagesBeforeRules[b]
	// fmt.Printf("a %v <= b %v beforeB = %v\n", a, b, before)
	if a == b || Contains(before, a) {
		return true
	} else {
		for _, beforeB := range pagesBeforeRules[b] {
			if Contains(update, beforeB) && isABeforeB(pagesBeforeRules, update, a, beforeB) {
				return true
			}
		}
	}
	return false
}

func quicksort(pagesBeforeRules map[int][]int, update []int, lo, hi int) {
	if lo >= hi || lo < 0 {
		return
	}

	pivotIdx := partition(pagesBeforeRules, update, lo, hi)

	quicksort(pagesBeforeRules, update, lo, pivotIdx-1)
	quicksort(pagesBeforeRules, update, pivotIdx+1, hi)
}

func partition(pagesBeforeRules map[int][]int, update []int, lo, hi int) int {
	pivot := update[hi]
	i := lo

	for j := lo; j < hi; j++ {
		isBefore := isABeforeB(pagesBeforeRules, update, update[j], pivot)
		// fmt.Printf("Is %v before %v = %v\n", update[j], pivot, isBefore)
		if isBefore {
			// fmt.Printf("swapping %v and %v pivot = %v\n", update[j], update[i], pivot)
			update[j], update[i] = update[i], update[j]
			i++
		}
	}

	update[i], update[hi] = update[hi], update[i]
	return i
}
