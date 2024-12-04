package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day3"
)

func main() {
	input, err := helpers.ReadInput("day3")
	if err != nil {
		panic("could not read input")
	}
	result := day3.Solve(input)
	result2 := day3.SolveWithActivation(input)
	fmt.Printf("The results are %d and %d \n", result, result2)
}
