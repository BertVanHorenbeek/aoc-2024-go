package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day4"
)

func main() {
	input, err := helpers.ReadInput("day4")
	if err != nil {
		panic("could not read input")
	}
	result := day4.CountXmas(input)
	result2 := day4.CountMasInXShape(input)
	fmt.Printf("The results are %d and %d\n", result, result2)
}
