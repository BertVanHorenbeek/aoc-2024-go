package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day10"
)

func main() {
	input, err := helpers.ReadInput("day10")
	if err != nil {
		panic("could not read input")
	}
	result := day10.CountTrails(input)
	result2 := day10.CountTrails2(input)
	fmt.Printf("The results are %d and %d\n", result, result2)
}
