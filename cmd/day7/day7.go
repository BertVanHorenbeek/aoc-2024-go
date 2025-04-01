package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day7"
)

func main() {
	input, err := helpers.ReadInput("day7")
	if err != nil {
		panic("could not read input")
	}
	result := day7.TotalCalibrationResult(input)
	result2 := day7.TotalCalibrationResultWithConcat(input)
	fmt.Printf("The results are %d and %d\n", result, result2)
}

//1298300128588 to low
