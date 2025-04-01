package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day15"
)

func main() {
	input, err := helpers.ReadInput("day15")
	if err != nil {
		panic("could not read input")
	}
	result := day15.Solve(input)
	result2 := day15.SolvePart2(input)
	fmt.Printf("The results are %d and %d\n", result, result2)
}
