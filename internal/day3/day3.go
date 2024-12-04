package day3

import (
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

func Solve(input string) int {
	pattern := regexp.MustCompile("mul\\(([0-9]+),([0-9]+)\\)")
	matches := pattern.FindAllStringSubmatch(input, -1)
	return lo.Sum(lo.Map(matches, func(command []string, index int) int {
		first, _ := strconv.Atoi(command[1])
		second, _ := strconv.Atoi(command[2])
		return first * second
	}))
}

func SolveWithActivation(input string) int {
	pattern := regexp.MustCompile("mul\\(([0-9]+),([0-9]+)\\)|do\\(\\)|don't\\(\\)")
	matches := pattern.FindAllStringSubmatch(input, -1)
	active := true
	sum := 0
	for _, command := range matches {
		switch command[0] {
		case "don't()":
			active = false
		case "do()":
			active = true
		default:
			if active {
				first, _ := strconv.Atoi(command[1])
				second, _ := strconv.Atoi(command[2])
				sum += first * second
			}
		}
	}
	return sum
}
