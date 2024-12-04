package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day1"
)

func main() {
	input, err := helpers.ReadInput("day1")
	if err != nil {
		panic("could not read input")
	}
	distance, similarity := day1.Solve(input)
	fmt.Printf("The distance is %d and similarity is %d\n", distance, similarity)
}
