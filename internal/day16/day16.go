package day16

import (
	"aoc-2024-go/helpers"
	"github.com/samber/lo"
)

type position struct {
	x, y int
}

type reindeer struct {
	p position
	o position
}

type reindeerWithTail struct {
	r    reindeer
	tail []position
}

func (r reindeerWithTail) newReindeer(new reindeer) reindeerWithTail {
	if r.r.p == new.p {
		return reindeerWithTail{r: new, tail: r.tail}
	}
	return reindeerWithTail{
		new,
		lo.Union(r.tail, []position{new.p}),
	}
}

func (r reindeer) forward() reindeer {
	return reindeer{
		p: position{r.p.x + r.o.x, r.p.y + r.o.y},
		o: r.o,
	}
}

func (r reindeer) rotateR() reindeer {
	return reindeer{
		p: r.p,
		o: position{-r.o.y, r.o.x},
	}
}

func (r reindeer) rotateL() reindeer {
	return reindeer{
		p: r.p,
		o: position{r.o.y, -r.o.x},
	}
}

func (r reindeer) rotated() []reindeer {
	return []reindeer{r.rotateL(), r.rotateR()}
}

func FindBestPossibleScore(input string) int {
	costs := calculateBestPath(input)
	return lo.Min(lo.Keys(costs))
}

func calculateBestPath(input string) map[int][]reindeerWithTail {
	blockages := make([]position, 0, 200)
	var start reindeer
	var end position
	grid := helpers.StringToGrid(input)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '#' {
				blockages = append(blockages, position{x: x, y: y})
			} else if grid[y][x] == 'S' {
				start = reindeer{p: position{x, y}, o: position{1, 0}}
			} else if grid[y][x] == 'E' {
				end = position{x, y}
			}
		}
	}

	costMap := map[reindeer]int{start: 0}
	costs := map[int][]reindeerWithTail{}
	current := []reindeerWithTail{{r: start, tail: []position{start.p}}}
	for {
		nextReindeer := make([]reindeerWithTail, 0, len(current)*2)
		for _, r := range current {
			forward := r.r.forward()
			cost := costMap[r.r] + 1
			if oldCost, ok := costMap[forward]; (!ok || oldCost >= cost) && !lo.Contains(blockages, forward.p) {
				costMap[forward] = cost
				if forward.p == end {
					costs[cost] = append(costs[cost], r.newReindeer(forward))
					continue
				} else {
					nextReindeer = append(nextReindeer, r.newReindeer(forward))
				}
			}
			rotationCost := costMap[r.r] + 1000
			for _, rotated := range r.r.rotated() {
				if oldCost, ok := costMap[rotated]; !ok || oldCost >= rotationCost {
					costMap[rotated] = rotationCost
					nextReindeer = append(nextReindeer, r.newReindeer(rotated))
				}
			}
		}
		if len(nextReindeer) == 0 {
			break
		}
		current = nextReindeer
	}
	return costs
}

func FindBestSeats(input string) int {
	costs := calculateBestPath(input)
	bestScore := lo.Min(lo.Keys(costs))
	positions := lo.FlatMap(costs[bestScore], func(withTail reindeerWithTail, _ int) []position {
		return withTail.tail
	})
	return len(lo.Uniq(positions))
}
