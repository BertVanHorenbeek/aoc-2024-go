package day10

import (
	"aoc-2024-go/helpers"
	"github.com/samber/lo"
)

type position struct {
	x, y int
}

func (p position) inMap(x int, y int) bool {
	return p.x >= 0 && p.x < x && p.y >= 0 && p.y < y
}

func (p position) nextPositions() []position {
	return []position{
		{p.x + 1, p.y},
		{p.x - 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}

func CountTrails(input string) int {
	grid := helpers.StringToGrid(input)

	trailHeads := []position{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '0' {
				trailHeads = append(trailHeads, position{x, y})
			}
		}
	}
	return lo.Sum(lo.Map(trailHeads, func(trailHead position, _ int) int {
		tops := findTopsReachable(trailHead, grid)
		uniqueTops := lo.Uniq(tops)
		return len(uniqueTops)
	}))
}

func findTopsReachable(pos position, grid [][]rune) []position {
	if grid[pos.y][pos.x] == '9' {
		return []position{pos}
	}
	currentElevation := grid[pos.y][pos.x]
	nextPositions := lo.Filter(pos.nextPositions(), func(item position, _ int) bool {
		return item.inMap(len(grid[0]), len(grid)) && grid[item.y][item.x] == currentElevation+1
	})
	if len(nextPositions) == 0 {
		return []position{}
	}
	return lo.FlatMap(nextPositions, func(next position, _ int) []position {
		return findTopsReachable(next, grid)
	})
}

func CountTrails2(input string) int {
	grid := helpers.StringToGrid(input)

	trailHeads := []position{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '0' {
				trailHeads = append(trailHeads, position{x, y})
			}
		}
	}
	return lo.Sum(lo.Map(trailHeads, func(trailHead position, _ int) int {
		paths := countPathsToTop(trailHead, grid)
		return paths
	}))
}

func countPathsToTop(pos position, grid [][]rune) int {
	if grid[pos.y][pos.x] == '9' {
		return 1
	}
	currentElevation := grid[pos.y][pos.x]
	nextPositions := lo.Filter(pos.nextPositions(), func(item position, _ int) bool {
		return item.inMap(len(grid[0]), len(grid)) && grid[item.y][item.x] == currentElevation+1
	})
	if len(nextPositions) == 0 {
		return 0
	}
	return lo.Sum(lo.Map(nextPositions, func(next position, _ int) int {
		return countPathsToTop(next, grid)
	}))
}
