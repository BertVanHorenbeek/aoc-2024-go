package day8

import (
	"fmt"

	"aoc-2024-go/helpers"
	"github.com/samber/lo"
)

type position struct {
	x, y int
}

func (p position) plus(p2 position) position {
	return position{
		x: p.x + p2.x,
		y: p.y + p2.y,
	}
}

type pair struct {
	first, second position
}

func (p position) diff(p2 position) position {
	return position{
		x: p.x - p2.x,
		y: p.y - p2.y,
	}
}

func CountAntiNodes(input string) int {
	grid := helpers.StringToGrid(input)
	maxX := len(grid[0])
	maxY := len(grid)

	antennaMap := map[rune][]position{}

	for y, row := range grid {
		for x, field := range row {
			if field != '.' {
				antennaMap[field] = append(antennaMap[field], position{
					x: x,
					y: y,
				})
			}
		}
	}
	antiNodes := lo.FlatMap(lo.Values(antennaMap), func(positions []position, _ int) []position {
		return lo.FlatMap(pairUp(positions), findAntiNodes)
	})
	validAntiNodes := lo.Filter(lo.Uniq(antiNodes), inMap(maxX, maxY))
	fmt.Printf("The antinodes are: %v\n", validAntiNodes)
	return len(validAntiNodes)
}

func CountAntiNodesUsingHarmonics(input string) int {
	grid := helpers.StringToGrid(input)
	maxX := len(grid[0])
	maxY := len(grid)

	antennaMap := map[rune][]position{}

	for y, row := range grid {
		for x, field := range row {
			if field != '.' {
				antennaMap[field] = append(antennaMap[field], position{
					x: x,
					y: y,
				})
			}
		}
	}
	antiNodes := lo.FlatMap(lo.Values(antennaMap), func(positions []position, _ int) []position {
		return lo.FlatMap(pairUp(positions), findAntiNodesUsingHarmonics(maxX, maxY))
	})
	validAntiNodes := lo.Filter(lo.Uniq(antiNodes), inMap(maxX, maxY))
	fmt.Printf("The antinodes are: %v\n", validAntiNodes)
	return len(validAntiNodes)
}

func findAntiNodesUsingHarmonics(x int, y int) func(pair, int) []position {
	return func(antennas pair, _ int) []position {
		antiNodes := []position{}
		diff := antennas.first.diff(antennas.second)
		plusPosition := antennas.first
		for inMap(x, y)(plusPosition, 0) {
			antiNodes = append(antiNodes, plusPosition)
			plusPosition = plusPosition.plus(diff)
		}
		minPosition := antennas.second
		for inMap(x, y)(minPosition, 0) {
			antiNodes = append(antiNodes, minPosition)
			minPosition = minPosition.diff(diff)
		}

		return antiNodes
	}
}

func inMap(x int, y int) func(position, int) bool {
	return func(item position, _ int) bool {
		return item.x >= 0 && item.x < x && item.y >= 0 && item.y < y
	}
}

func pairUp(antennas []position) []pair {
	pairs := []pair{}
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			pairs = append(pairs, pair{
				first:  antennas[i],
				second: antennas[j],
			})
		}
	}
	return pairs
}

func findAntiNodes(antennas pair, _ int) []position {
	diff := antennas.first.diff(antennas.second)
	return []position{
		antennas.first.plus(diff),
		antennas.second.diff(diff),
	}
}
