package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day14"
)

func main() {
	input, err := helpers.ReadInput("day14")
	if err != nil {
		panic("could not read input")
	}
	result := day14.CalculateSafetyFactor(input, 101, 103)
	result2 := day14.FindChristmastree(input, 101, 103)
	fmt.Printf("The results are %d and %d\n", result, result2)
}

// 226118750 too high

// 10000 too high
// 7082 too low
