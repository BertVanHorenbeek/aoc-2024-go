package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day25"
)

func main() {
	input, err := helpers.ReadInput("day25")
	if err != nil {
		panic("could not read input")
	}
	result := day25.CountFittingLockKeyPairs(input)
	fmt.Printf("The results are %d and %d\n", result, result)
}
