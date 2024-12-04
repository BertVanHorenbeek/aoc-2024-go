package day1

import (
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func Solve(input string) (int, int) {
	lines := strings.Split(input, "\n")
	var lists = make([][]int, 2)
	for _, line := range lines {
		fields := strings.Fields(line)
		element1, err := strconv.Atoi(fields[0])
		if err != nil {
			panic("Could not convert to int")
		}
		lists[0] = append(lists[0], element1)
		element2, err := strconv.Atoi(fields[1])
		if err != nil {
			panic("Could not convert to int")
		}
		lists[1] = append(lists[1], element2)
	}
	return calculate(lists), similarityScore(lists)
}

func calculate(lists [][]int) int {
	for _, list := range lists {
		sort.Ints(list)
	}
	totalDistance := 0
	for i := 0; i < len(lists[0]); i++ {
		totalDistance += abs(lists[0][i] - lists[1][i])
	}
	return totalDistance
}

func similarityScore(lists [][]int) int {
	score := 0
	for _, number := range lists[0] {
		score += number * lo.Count(lists[1], number)
	}
	return score
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
