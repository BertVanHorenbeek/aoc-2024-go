package day12

import (
	"aoc-2024-go/helpers"
	"github.com/samber/lo"
)

type position struct {
	x, y int
}

func (p position) neighbours() []position {
	return []position{
		{p.x + 1, p.y},
		{p.x - 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}

type fence struct {
	pos  position
	side position
}

func (p position) isNeighbour(pos position) bool {
	return lo.Contains(p.neighbours(), pos)
}

func (f fence) isNeighbour(f2 fence) bool {
	return f.pos.isNeighbour(f2.pos) && f.side.isNeighbour(f2.side)
}

func Solve(input string) int {
	cropMap := map[rune][]position{}
	grid := helpers.StringToGrid(input)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			cropMap[grid[x][y]] = append(cropMap[grid[x][y]], position{x, y})
		}
	}
	return lo.Sum(lo.Map(lo.Values(cropMap), func(field []position, _ int) int {
		return calculateFenceForField(field)
	}))
}

func calculateFenceForField(patches []position) int {
	fields := splitInFields(patches)
	fenceCost := 0
	for _, field := range fields {
		fenceLength := 0
		for _, patch := range field {
			fenceLength += 4 - lo.CountBy(field, func(otherPatch position) bool {
				return otherPatch.isNeighbour(patch)
			})
		}
		fenceCost += fenceLength * len(field)
	}
	return fenceCost
}

func calculateDiscountedFenceForField(patches []position) int {
	fields := splitInFields(patches)
	fenceCost := 0
	for _, field := range fields {
		fences := []fence{}
		for _, patch := range field {
			edges := lo.Reject(patch.neighbours(), func(otherPatch position, _ int) bool {
				return lo.Contains(field, otherPatch)
			})
			fences = append(fences, lo.Map(edges, func(edge position, _ int) fence {
				return fence{pos: patch, side: edge}
			})...)
		}
		fenceCost += len(calculateDiscounted(fences)) * len(field)
	}
	return fenceCost
}

func calculateDiscounted(fences []fence) [][]fence {
	discounted := [][]fence{}
	remainingFences := fences[1:]
	discounted = append(discounted, []fence{fences[0]})
	for len(remainingFences) > 0 {
		fenceToGrow := discounted[len(discounted)-1]
		remainingAfterThisRound := []fence{}
		for _, patch := range remainingFences {
			partOfTheField := lo.SomeBy(fenceToGrow, func(candidate fence) bool {
				return candidate.isNeighbour(patch)
			})
			if partOfTheField {
				fenceToGrow = append(fenceToGrow, patch)
			} else {
				remainingAfterThisRound = append(remainingAfterThisRound, patch)
			}
		}
		if len(remainingFences) == len(remainingAfterThisRound) {
			discounted = append(discounted, []fence{remainingFences[0]})
			remainingAfterThisRound = remainingAfterThisRound[1:]
		} else {
			discounted[len(discounted)-1] = fenceToGrow
		}
		remainingFences = remainingAfterThisRound
	}
	return discounted
}

func splitInFields(patches []position) [][]position {
	fields := [][]position{}
	remainingPatches := patches[1:]
	fields = append(fields, []position{patches[0]})
	for len(remainingPatches) > 0 {
		fieldToGrow := fields[len(fields)-1]
		remainingAfterThisRound := []position{}
		for _, patch := range remainingPatches {
			partOfTheField := lo.SomeBy(fieldToGrow, func(candidate position) bool {
				return candidate.isNeighbour(patch)
			})
			if partOfTheField {
				fieldToGrow = append(fieldToGrow, patch)
			} else {
				remainingAfterThisRound = append(remainingAfterThisRound, patch)
			}
		}
		if len(remainingPatches) == len(remainingAfterThisRound) {
			fields = append(fields, []position{remainingPatches[0]})
			remainingAfterThisRound = remainingAfterThisRound[1:]
		} else {
			fields[len(fields)-1] = fieldToGrow
		}
		remainingPatches = remainingAfterThisRound
	}
	return fields
}

func SolveWithDiscount(input string) int {
	cropMap := map[rune][]position{}
	grid := helpers.StringToGrid(input)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			cropMap[grid[x][y]] = append(cropMap[grid[x][y]], position{x, y})
		}
	}
	return lo.Sum(lo.Map(lo.Values(cropMap), func(field []position, _ int) int {
		return calculateDiscountedFenceForField(field)
	}))
}
