package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day24"
)

func main() {
	input, err := helpers.ReadInput("day24")
	if err != nil {
		panic("could not read input")
	}
	result := day24.Solve(input)
	result2 := day24.FindSwappedPairs5(input)
	fmt.Printf("The results are %d and %s\n", result, result2)
}
