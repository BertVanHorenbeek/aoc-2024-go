package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day19"
)

func main() {
	input, err := helpers.ReadInput("day19")
	if err != nil {
		panic("could not read input")
	}
	result := day19.CountPossibleDesigns(input)
	result2 := day19.CountDesignOptions(input)
	fmt.Printf("The results are %d and %v\n", result, result2)
}
