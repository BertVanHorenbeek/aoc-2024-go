package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func CountStonesAfterBlinks(input string, blinks int) int {
	fields := strings.Fields(input)
	stones := lo.Map(fields, func(item string, _ int) int {
		atoi, _ := strconv.Atoi(item)
		return atoi
	})
	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}
	return len(stones)
}

func CountStonesAfter75Blinks(input string) int {
	fields := strings.Fields(input)
	stones := lo.Map(fields, func(item string, _ int) int {
		atoi, _ := strconv.Atoi(item)
		return atoi
	})
	totalAmountOfStones := 0
	solutionMap := map[int][]int{}
	for _, stone := range stones {
		fmt.Printf("Starting with stone %d\n", stone)
		solution := []int{stone}
		for i := 0; i < 2; i++ {
			fmt.Printf("Step %d: %d\n", i, stone)
			fmt.Printf("Solution is of length: %d\n", len(solution))
			solution = do25steps(solution, &solutionMap)
		}
		for _, solutionStone := range solution {
			totalAmountOfStones += len(solutionFromSolutionMap(solutionStone, &solutionMap))
		}
	}
	return totalAmountOfStones
}

func do25steps(stones []int, solutionMap *map[int][]int) []int {
	solution := make([]int, 0, len(stones))
	for _, stone := range stones {
		solution = append(solution, solutionFromSolutionMap(stone, solutionMap)...)
	}
	return solution
}

func solutionFromSolutionMap(stone int, solutionMap *map[int][]int) []int {
	solution, ok := (*solutionMap)[stone]
	if ok {
		return solution
	}
	result := []int{stone}
	for i := 0; i < 25; i++ {
		result = blink(result)
	}
	(*solutionMap)[stone] = result
	return result
}

func blink(stones []int) []int {
	output := make([]int, 0, len(stones)*2)
	for i := 0; i < len(stones); i++ {
		if stones[i] == 0 {
			output = append(output, 1)
		} else if asString := fmt.Sprintf("%d", stones[i]); len(asString)%2 == 0 {
			firstHalf, _ := strconv.Atoi(asString[:len(asString)/2])
			secondHalf, _ := strconv.Atoi(asString[len(asString)/2:])
			output = append(output, firstHalf, secondHalf)
		} else {
			output = append(output, stones[i]*2024)
		}
	}
	return output
}
