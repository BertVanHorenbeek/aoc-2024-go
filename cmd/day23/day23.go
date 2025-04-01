package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day23"
)

func main() {
	input, err := helpers.ReadInput("day23")
	if err != nil {
		panic("could not read input")
	}
	result := day23.Solve(input)
	result2 := day23.Solve2(input)
	fmt.Printf("The results are %d and %v\n", result, result2)
}
