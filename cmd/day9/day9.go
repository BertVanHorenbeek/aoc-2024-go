package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day9"
)

func main() {
	input, err := helpers.ReadInput("day9")
	if err != nil {
		panic("could not read input")
	}
	result := day9.Solve(input)
	result2 := day9.SolvePart2(input)
	fmt.Printf("The results are %d and %d\n", result, result2)
}
