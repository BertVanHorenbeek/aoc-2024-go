package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day17"
)

func main() {
	input, err := helpers.ReadInput("day17")
	if err != nil {
		panic("could not read input")
	}
	result := day17.RunProgram(input)
	result2 := day17.FindStartToProduceItSelf(input)
	fmt.Printf("The results are \"%s\" and %d\n", result, result2)
}

// 10000000000 too low

// 1000024978 too low

// 39846323390557074 too high
