package day13

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type position struct {
	x, y int
}

func CalculateMinTokensToWinPossiblePrices(input string) int {
	machines := parseInput(input)
	return lo.Sum(lo.Map(machines, func(item machine, _ int) int {
		aPresses, bPresses := calculateRequiredPresses(item)
		if aPresses*item.buttonA.x+bPresses*item.buttonB.x == item.destination.x &&
			aPresses*item.buttonA.y+bPresses*item.buttonB.y == item.destination.y {
			return aPresses*3 + bPresses
		}
		return 0
	}))
}

func CalculateMinTokensToWinPossiblePricesForRealPositions(input string) int {
	machines := parseInput(input)
	machines = lo.Map(machines, func(item machine, _ int) machine {
		return machine{
			buttonA: item.buttonA,
			buttonB: item.buttonB,
			destination: position{
				x: item.destination.x + 10000000000000,
				y: item.destination.y + 10000000000000,
			},
		}
	})
	return lo.Sum(lo.Map(machines, func(item machine, _ int) int {
		aPresses, bPresses := calculateRequiredPresses(item)
		if aPresses*item.buttonA.x+bPresses*item.buttonB.x == item.destination.x &&
			aPresses*item.buttonA.y+bPresses*item.buttonB.y == item.destination.y {
			return aPresses*3 + bPresses
		}
		return 0
	}))
}

type machine struct {
	buttonA, buttonB, destination position
}

func parseInput(input string) []machine {
	machineBlocks := strings.Split(input, "\n\n")
	return lo.Map(machineBlocks, func(block string, _ int) machine {
		lines := strings.Split(block, "\n")
		return machine{
			buttonA:     parseButton(lines[0]),
			buttonB:     parseButton(lines[1]),
			destination: parseDestination(lines[2]),
		}
	})
}

func parseDestination(line string) position {
	regex := regexp.MustCompile("X=([0-9]+), Y=([0-9]+)")
	result := regex.FindStringSubmatch(line)
	x, _ := strconv.Atoi(result[1])
	y, _ := strconv.Atoi(result[2])
	return position{
		x: x,
		y: y,
	}
}

func parseButton(line string) position {
	regex := regexp.MustCompile("X\\+([0-9]+), Y\\+([0-9]+)")
	result := regex.FindStringSubmatch(line)
	x, _ := strconv.Atoi(result[1])
	y, _ := strconv.Atoi(result[2])
	return position{
		x: x,
		y: y,
	}
}

func calculateRequiredPresses(m machine) (int, int) {
	aPresses := ((m.buttonB.y * m.destination.x) - (m.buttonB.x * m.destination.y)) / ((m.buttonB.y * m.buttonA.x) - (m.buttonB.x * m.buttonA.y))
	bPresses := (m.destination.y - m.buttonA.y*aPresses) / m.buttonB.y
	return aPresses, bPresses
}
