package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day20"
)

func main() {
	input, err := helpers.ReadInput("day20")
	if err != nil {
		panic("could not read input")
	}
	result := day20.CountCheatsSavingAtLeastXps(input, 100)
	result2 := day20.CountLongCheatsSavingAtLeastXps(input, 100)
	fmt.Printf("The results are %d and %v\n", result, result2)
}

// 212716 too low
// 231211 too low
