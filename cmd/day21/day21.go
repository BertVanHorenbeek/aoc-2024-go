package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day21"
)

func main() {
	input, err := helpers.ReadInput("day21")
	if err != nil {
		panic("could not read input")
	}
	//result := day21.CalculateComplexityForX(input, 2)
	result2 := day21.CalculateComplexityForX(input, 25)
	fmt.Printf("The results are %d and %v\n", result2, result2)
}

// 308275027809642 too high
// 120702950410039 too low
// 187276527880227 too low
// 400827838993176
// 156944425344350
// 247965001155352 wrong
