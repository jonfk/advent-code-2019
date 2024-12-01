package day1

import (
	"fmt"
	"strconv"
	"strings"
)

type Input struct {
	l1 []int
	l2 []int
}

func Parse(input string) (*Input, error) {
	res := new(Input)
	for _, str := range strings.Split(input, "\n") {
		inputs := strings.Split(str, "   ")
		if len(inputs) != 2 {
			return nil, fmt.Errorf("Input row did not contain right number of columns")
		}
		l1, err := strconv.Atoi(inputs[0])
		if err != nil {
			return nil, err
		}
		res.l1 = append(res.l1, l1)

		l2, err := strconv.Atoi(inputs[1])
		if err != nil {
			return nil, err
		}
		res.l2 = append(res.l2, l2)
	}
	return res, nil
}

func SumDistancesBetweenSortedLists(input *Input) (int, error) {
	Sort(input.l1)
	Sort(input.l2)

	if len(input.l1) != len(input.l2) {
		return 0, fmt.Errorf("l1 and l2 are not same lengths")
	}

	sumDist := 0
	for i := range input.l1 {
		sumDist += abs(input.l1[i] - input.l2[i])
	}
	return sumDist, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Sort(list []int) {
	quicksort(list, 0, len(list)-1)
}

func quicksort(a []int, lo, hi int) {
	if lo >= hi || lo < 0 {
		return
	}

	pivotIdx := partition(a, lo, hi)

	quicksort(a, lo, pivotIdx-1)
	quicksort(a, pivotIdx+1, hi)
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo

	for j := lo; j < hi; j++ {
		if a[j] <= pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}

	a[i], a[hi] = a[hi], a[i]
	return i
}

func SumSimilarityScore(input *Input) int {
	occurrances := make(map[int]int)

	for _, x := range input.l2 {
		occurrances[x] += 1
	}

	sumSimilarityScores := 0
	for _, x := range input.l1 {
		sumSimilarityScores += (occurrances[x] * x)
	}

	return sumSimilarityScores
}
