package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day5"
)

func main() {
	input, err := helpers.ReadInput("day5")
	if err != nil {
		panic("could not read input")
	}
	result := day5.CountValidQueues(input)
	result2 := day5.FixInvalidQueues(input)
	fmt.Printf("The results are %d and %d\n", result, result2)
}
