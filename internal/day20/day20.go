package day20

import (
	"fmt"
	"log"

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

type shortcut struct {
	start, end position
	timeSaved  int
}

func CountCheatsSavingAtLeastXps(input string, minimalTimeSaved int) int {
	track, _ := buildRaceTrack(input)
	shortcuts := findShortCuts(track)
	return lo.CountBy(shortcuts, func(item shortcut) bool {
		return item.timeSaved >= minimalTimeSaved
	})
}

func CountLongCheatsSavingAtLeastXps(input string, minimalTimeSaved int) int {
	track, walls := buildRaceTrack(input)
	shortcuts := findLongShortcuts(track, walls, minimalTimeSaved)
	return lo.CountBy(shortcuts, func(item shortcut) bool {
		return item.timeSaved >= minimalTimeSaved
	})
}

func FindCheatsOverXps(input string, minimalTimeSaved int) map[int]int {
	track, walls := buildRaceTrack(input)
	shortcuts := findLongShortcuts(track, walls, minimalTimeSaved)
	usefullShortcuts := lo.Filter(shortcuts, func(item shortcut, index int) bool {
		return item.timeSaved >= minimalTimeSaved
	})
	return lo.CountValuesBy(usefullShortcuts, func(item shortcut) int {
		return item.timeSaved
	})
}

func findShortCuts(track []position) []shortcut {
	shortcuts := make([]shortcut, 0, len(track))
	for i := 0; i < len(track); i++ {
		for j := i + 3; j < len(track); j++ {
			if isShortCut(track[i], track[j]) {
				shortcuts = append(shortcuts, shortcut{track[i], track[j], j - i - 2})
			}
		}
	}
	return shortcuts
}

func findLongShortcuts(track []position, walls map[position]bool, minimalTimeSaved int) []shortcut {
	shortcuts := make([]shortcut, 0, len(track))
	fmt.Printf("Tracklength: %d\n", len(track))
	for i := 0; i < len(track); i++ {
		//fmt.Printf("Find shortcut for %d position %v\n", i, track[i])
		for j := i + minimalTimeSaved; j < len(track); j++ {
			if cost, ok := isLongShortCut(track[i], track[j], walls); ok {
				shortcuts = append(shortcuts, shortcut{track[i], track[j], j - i - cost})
			}
		}
	}
	return shortcuts
}

func isLongShortCut(start position, end position, walls map[position]bool) (int, bool) {
	if abs(start.x-end.x)+abs(start.y-end.y) > 20 {
		return -1, false
	}
	currentPositions := []position{start}
	pastPositions := []position{}
	for i := 1; i <= 21; i++ {
		neighbours := lo.Uniq(lo.FlatMap(currentPositions, func(item position, _ int) []position {
			return item.neighbours()
		}))
		if lo.Contains(neighbours, end) {
			return i, true
		}
		newNeighbours := lo.Without(neighbours, currentPositions...)
		neverVisited := lo.Without(newNeighbours, pastPositions...)
		pastPositions = lo.Union(pastPositions, neverVisited)
		currentPositions = neverVisited
	}
	return -1, false
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func isShortCut(p position, p2 position) bool {
	return p.x == p2.x && p.y == p2.y+2 ||
		p.x == p2.x && p.y == p2.y-2 ||
		p.x == p2.x+2 && p.y == p2.y ||
		p.x == p2.x-2 && p.y == p2.y
}

func buildRaceTrack(input string) ([]position, map[position]bool) {
	grid := helpers.StringToGrid(input)
	var start position
	var end position
	fields := make(map[position]bool)
	walls := make(map[position]bool)
	for y, row := range grid {
		for x, field := range row {
			if field == 'S' {
				start = position{x, y}
			} else if field == 'E' {
				end = position{x, y}
				fields[position{x, y}] = true
			} else if field == '.' {
				fields[position{x, y}] = true
			} else if field == '#' {
				walls[position{x, y}] = true
			}
		}
	}
	track := []position{start}
	for lo.LastOrEmpty(track) != end {
		possibleNextPositions := track[len(track)-1].neighbours()
		next := lo.Filter(possibleNextPositions, func(item position, _ int) bool {
			return fields[item] && (len(track) <= 1 || item != track[len(track)-2])
		})
		if len(next) != 1 {
			log.Fatal("There should only be one next position.")
		}
		track = append(track, next[0])
	}
	return track, walls
}
