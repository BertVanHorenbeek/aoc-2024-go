package day6

import (
	"aoc-2024-go/helpers"
	"github.com/samber/lo"
)

type position struct {
	x, y int
}

type direction int

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

type guard struct {
	position
	d direction
}

func (g *guard) inMap(x int, y int) bool {
	return g.x >= 0 && g.x < x && g.y >= 0 && g.y < y
}

func (g *guard) step(obstacles []position) bool {
	var newPosition position
	switch g.d {
	case UP:
		newPosition = position{
			x: g.x,
			y: g.y - 1,
		}
	case RIGHT:
		newPosition = position{
			x: g.x + 1,
			y: g.y,
		}
	case DOWN:
		newPosition = position{
			x: g.x,
			y: g.y + 1,
		}
	case LEFT:
		newPosition = position{
			x: g.x - 1,
			y: g.y,
		}
	}
	if !lo.Contains(obstacles, newPosition) {
		g.position = newPosition
		return true
	} else {
		return false
	}
}

func (g *guard) rotate() {
	g.d = (g.d + 1) % 4
}

func CalculatedVisitedPositions(input string) int {
	obstacles, g, maxX, maxY := loadMap(input)
	visitedFields := map[position]bool{}
	visitedFields[g.position] = true
	for g.inMap(maxX, maxY) {
		if g.step(obstacles) {
			if g.inMap(maxX, maxY) {
				visitedFields[g.position] = true
			}
		} else {
			g.rotate()
		}
	}
	return len(lo.Keys(visitedFields))
}
func NumberOfPossibleLoops(input string) int {
	obstacles, g, maxX, maxY := loadMap(input)
	visitedFields := map[position]bool{}
	startingPosition := g
	for g.inMap(maxX, maxY) {
		if g.step(obstacles) {
			if g.inMap(maxX, maxY) {
				visitedFields[g.position] = true
			}
		} else {
			g.rotate()
		}
	}
	loopCount := 0
	for _, extraObstacle := range lo.Keys(visitedFields) {
		if detectLoop(startingPosition, append(obstacles, extraObstacle), maxX, maxY) {
			loopCount++
		}
	}
	return loopCount
}

func detectLoop(g guard, obstacles []position, x int, y int) bool {
	visitedFields := map[guard]bool{}
	visitedFields[g] = true
	for g.inMap(x, y) {
		if g.step(obstacles) {
			if visitedFields[g] {
				return true
			}
			visitedFields[g] = true
		} else {
			g.rotate()
		}
	}
	return false
}

func loadMap(input string) (positions []position, g guard, maxX int, maxY int) {
	grid := helpers.StringToGrid(input)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '#' {
				positions = append(positions, position{
					x: x,
					y: y,
				})
			} else if grid[y][x] == '^' {
				g = guard{
					position: position{
						x: x,
						y: y,
					},
					d: UP,
				}
			}
		}
	}
	return positions, g, len(grid[0]), len(grid)
}
