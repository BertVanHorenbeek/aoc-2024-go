package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day11"
)

func main() {
	input, err := helpers.ReadInput("day11")
	if err != nil {
		panic("could not read input")
	}
	//result := day11.CountStonesAfterBlinks(input, 25)
	result2 := day11.CountStonesAfter75Blinks(input)
	fmt.Printf("The results are %d and %d\n", result2, result2)
}

// 239420053552058 to high
// 239413123020116
