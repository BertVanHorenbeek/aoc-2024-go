package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day13"
)

func main() {
	input, err := helpers.ReadInput("day13")
	if err != nil {
		panic("could not read input")
	}
	result := day13.CalculateMinTokensToWinPossiblePrices(input)
	result2 := day13.CalculateMinTokensToWinPossiblePricesForRealPositions(input)
	fmt.Printf("The results are %d and %d\n", result, result2)
}
