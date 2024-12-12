package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day12"
)

func main() {
	input, err := helpers.ReadInput("day12")
	if err != nil {
		panic("could not read input")
	}
	result := day12.Solve(input)
	result2 := day12.SolveWithDiscount(input)
	fmt.Printf("The results are %d and %d\n", result, result2)
}
