package day18

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type Position struct {
	x, y int
}

func NewPosition(x, y int) Position {
	return Position{x, y}
}

func (pos Position) neighbours(mapSize Position) []Position {
	neighbours := []Position{
		{pos.x + 1, pos.y},
		{pos.x - 1, pos.y},
		{pos.x, pos.y - 1},
		{pos.x, pos.y + 1},
	}
	return lo.Filter(neighbours, func(p Position, _ int) bool {
		return p.inGrid(mapSize)
	})
}

func (pos Position) inGrid(size Position) bool {
	return pos.x >= 0 &&
		pos.x <= size.x &&
		pos.y >= 0 &&
		pos.y <= size.y
}

func FindShortestPath(input string, destination Position, bytes int) int {
	lines := strings.Split(input, "\n")
	blockages := lo.Map(lines, func(line string, _ int) (p Position) {
		_, err := fmt.Sscanf(line, "%d,%d", &p.x, &p.y)
		if err != nil {
			panic(err)
		}
		return p
	})[:bytes]
	startPos := Position{0, 0}
	positions := []Position{startPos}
	pathMap := map[Position]int{startPos: 0}
	for {
		nextPositions := make([]Position, 0, len(positions)*3)
		for _, pos := range positions {
			newNeighbours := lo.Reject(pos.neighbours(destination), func(p Position, _ int) bool {
				_, ok := pathMap[p]
				return ok || lo.Contains(blockages, p)
			})
			pathCost := pathMap[pos]
			for _, neighbour := range newNeighbours {
				if neighbour == destination {
					return pathCost + 1
				}
				pathMap[neighbour] = pathCost + 1
				nextPositions = append(nextPositions, neighbour)
			}
		}
		positions = nextPositions
	}
}

func FindFirstBlockingByte(input string, destination Position) Position {
	lines := strings.Split(input, "\n")
	blockages := lo.Map(lines, func(line string, _ int) (p Position) {
		_, err := fmt.Sscanf(line, "%d,%d", &p.x, &p.y)
		if err != nil {
			panic(err)
		}
		return p
	})
	for i := 1500; i <= 3450; i++ {
		blockagesBeingTested := blockages[:i]
		if !hasAPath(destination, blockagesBeingTested) {
			return lo.LastOrEmpty(blockagesBeingTested)
		}
	}
	return Position{}
}

func hasAPath(destination Position, blockages []Position) bool {
	startPos := Position{0, 0}
	positions := []Position{startPos}
	pathMap := map[Position]int{startPos: 0}
	for {
		nextPositions := make([]Position, 0, len(positions)*3)
		for _, pos := range positions {
			newNeighbours := lo.Reject(pos.neighbours(destination), func(p Position, _ int) bool {
				_, ok := pathMap[p]
				return ok || lo.Contains(blockages, p)
			})
			pathCost := pathMap[pos]
			for _, neighbour := range newNeighbours {
				if neighbour == destination {
					return true
				}
				pathMap[neighbour] = pathCost + 1
				nextPositions = append(nextPositions, neighbour)
			}
		}
		if len(nextPositions) == 0 {
			return false
		}
		positions = nextPositions
	}
}
