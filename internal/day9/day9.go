package day9

import (
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func Solve(input string) int {
	expanded := expand(input)
	defragmented := defragment(expanded)
	return calculateCheckSum(defragmented)
}

func SolvePart2(input string) int {
	expanded := expand(input)
	defragmented := defragmentImproved(expanded)
	return calculateCheckSum(defragmented)
}

func defragmentImproved(input []int) []int {
	result := make([]int, len(input))
	copy(result, input)
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == -1 {
			continue
		}
		j := i
		for ; j >= 0 && input[j] == input[i]; j-- {
		}
		blockLength := i - j

		indexOfEmptyBlock := findFirstEmptyBlock(result[:i], blockLength)
		if indexOfEmptyBlock == -1 {
			for ; i > j; i-- {
				result[i] = input[i]
			}
			i++
			continue
		}
		endOfBlock := indexOfEmptyBlock + blockLength
		for ; indexOfEmptyBlock < endOfBlock; indexOfEmptyBlock++ {
			result[indexOfEmptyBlock] = input[i]
		}
		for ; i > j; i-- {
			result[i] = -1
		}
		i++
	}
	return result
}

func findFirstEmptyBlock(input []int, length int) int {
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input) && input[j] == -1; j++ {
			if j-i+1 >= length {
				return i
			}
		}
	}
	return -1
}

func calculateCheckSum(defragmented []int) int {
	return lo.Sum(lo.Map(defragmented, func(item int, index int) int {
		if item == -1 {
			return 0
		}
		return item * index
	}))
}

func defragment(expanded []int) []int {
	end := len(expanded) - 1
	defragmented := make([]int, 0, len(expanded))
	for i := 0; i <= end; i++ {
		if expanded[i] == -1 {
			for expanded[end] == -1 {
				end--
			}
			if end <= i {
				continue
			}
			defragmented = append(defragmented, expanded[end])
			end--
		} else {
			defragmented = append(defragmented, expanded[i])
		}
	}
	return defragmented
}

func expand(input string) []int {
	expanded := []int{}
	for index, numberAsString := range strings.Split(input, "") {
		n, _ := strconv.Atoi(numberAsString)
		if index%2 == 0 {
			fileID := index / 2
			for i := 0; i < n; i++ {
				expanded = append(expanded, fileID)
			}
		} else {
			for i := 0; i < n; i++ {
				expanded = append(expanded, -1)
			}
		}
	}
	return expanded
}
