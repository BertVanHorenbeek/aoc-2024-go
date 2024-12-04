package day4

import (
	"aoc-2024-go/helpers"
	"github.com/samber/lo"
)

type position struct {
	x, y int
}

func (p position) adjacentPositions() []position {
	adjacentPositions := make([]position, 0, 8)
	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			if x != p.x || y != p.y {
				adjacentPositions = append(adjacentPositions, position{
					x: x,
					y: y,
				})
			}
		}
	}
	return adjacentPositions
}

func CountXmas(input string) int {
	grid := helpers.StringToGrid(input)
	legend := buildLegend(grid)
	hiddenWord := []rune("XMAS")
	return findWords(legend, hiddenWord[1:], lo.Map(legend[hiddenWord[0]], func(item position, index int) []position {
		return []position{item}
	}))
}

func CountMasInXShape(input string) int {
	grid := helpers.StringToGrid(input)
	legend := buildLegend(grid)
	centerPositions := legend['A']
	return lo.CountBy(centerPositions, func(center position) bool {
		if center.x < 1 || center.y < 1 || center.x >= len(grid[0])-1 || center.y >= len(grid)-1 {
			return false
		}
		topLeft := position{
			x: center.x - 1,
			y: center.y - 1,
		}
		topRight := position{
			x: center.x + 1,
			y: center.y - 1,
		}
		bottomLeft := position{
			x: center.x - 1,
			y: center.y + 1,
		}
		bottomRight := position{
			x: center.x + 1,
			y: center.y + 1,
		}
		return (grid[topLeft.y][topLeft.x] == 'M' && grid[bottomRight.y][bottomRight.x] == 'S' ||
			grid[topLeft.y][topLeft.x] == 'S' && grid[bottomRight.y][bottomRight.x] == 'M') &&
			(grid[topRight.y][topRight.x] == 'M' && grid[bottomLeft.y][bottomLeft.x] == 'S' ||
				grid[topRight.y][topRight.x] == 'S' && grid[bottomLeft.y][bottomLeft.x] == 'M')
	})
}

func findWords(legend map[rune][]position, restOfWord []rune, positions [][]position) int {
	if len(restOfWord) == 0 {
		return len(positions)
	}
	possibleNextPositions := lo.FlatMap(positions, calculateNextPositions)
	positionsOfNextLetter := legend[restOfWord[0]]

	return findWords(legend, restOfWord[1:], lo.Filter(possibleNextPositions, func(possibleSnippet []position, _ int) bool {
		return lo.Contains(positionsOfNextLetter, lo.LastOrEmpty(possibleSnippet))
	}))
}

func calculateNextPositions(knownWordSnippet []position, _ int) [][]position {
	if len(knownWordSnippet) == 1 {
		return lo.Map(knownWordSnippet[0].adjacentPositions(), func(adjacentPosition position, _ int) []position {
			return []position{knownWordSnippet[0], adjacentPosition}
		})
	}
	lastPosition := len(knownWordSnippet) - 1
	nextPosition := position{
		x: knownWordSnippet[lastPosition].x + knownWordSnippet[lastPosition].x - knownWordSnippet[lastPosition-1].x,
		y: knownWordSnippet[lastPosition].y + knownWordSnippet[lastPosition].y - knownWordSnippet[lastPosition-1].y,
	}
	nextSnippet := append(knownWordSnippet, nextPosition)
	return [][]position{nextSnippet}
}

func buildLegend(grid [][]rune) map[rune][]position {
	legend := make(map[rune][]position)
	for y, column := range grid {
		for x, char := range column {
			legend[char] = append(legend[char], position{
				x: x,
				y: y,
			})
		}
	}
	return legend
}
