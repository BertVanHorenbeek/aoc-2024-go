package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day22"
)

func main() {
	input, err := helpers.ReadInput("day22")
	if err != nil {
		panic("could not read input")
	}
	result := day22.CalculateSecretNumberSum(input, 2000)
	result2 := day22.CalculateMostBananasYouCanEarn(input)
	fmt.Printf("The results are %d and %v\n", result, result2)
}

// 14887539874 too low
