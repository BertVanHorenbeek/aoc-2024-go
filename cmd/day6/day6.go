package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day6"
)

func main() {
	input, err := helpers.ReadInput("day6")
	if err != nil {
		panic("could not read input")
	}
	result := day6.CalculatedVisitedPositions(input)
	result3 := day6.NumberOfPossibleLoops(input)
	fmt.Printf("The results are %d and %d\n", result, result3)
}
