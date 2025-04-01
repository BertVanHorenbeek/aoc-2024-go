package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day16"
)

func main() {
	input, err := helpers.ReadInput("day16")
	if err != nil {
		panic("could not read input")
	}
	result := day16.FindBestPossibleScore(input)
	result2 := day16.FindBestSeats(input)
	fmt.Printf("The results are %d and %d\n", result, result2)
}
