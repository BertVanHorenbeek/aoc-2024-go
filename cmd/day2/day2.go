package main

import (
	"fmt"

	"aoc-2024-go/helpers"
	"aoc-2024-go/internal/day2"
)

func main() {
	input, err := helpers.ReadInput("day2")
	if err != nil {
		panic("could not read input")
	}
	validReports := day2.ValidateReports(input)
	validWithDamper := day2.ValidateReportsWithDamper(input)
	fmt.Printf("The amount of valid reports is %d and after damper %d \n", validReports, validWithDamper)
}
