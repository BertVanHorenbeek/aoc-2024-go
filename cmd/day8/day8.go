package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day8"
)

func main() {
	input, err := helpers.ReadInput("day8")
	if err != nil {
		panic("could not read input")
	}
	result := day8.CountAntiNodes(input)
	result2 := day8.CountAntiNodesUsingHarmonics(input)
	fmt.Printf("The results are %d and %d\n", result, result2)
}

// 294 is too high

//part 2
// 901 to Low
