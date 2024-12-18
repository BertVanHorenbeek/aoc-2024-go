package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day18"
)

func main() {
	input, err := helpers.ReadInput("day18")
	if err != nil {
		panic("could not read input")
	}
	result := day18.FindShortestPath(input, day18.NewPosition(70, 70), 1024)
	result2 := day18.FindFirstBlockingByte(input, day18.NewPosition(70, 70))
	fmt.Printf("The results are %d and %v\n", result, result2)
}

// 258 too high
