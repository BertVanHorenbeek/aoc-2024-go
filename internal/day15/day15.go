package day15

import (
	"strings"

	"aoc-2024-go/helpers"
	"github.com/samber/lo"
)

type position struct {
	x, y int
}

type movable interface {
	doMovement(move position, world map[position]movable) (map[position]movable, bool)
	canMove(move position, world map[position]movable) bool
	calculateCoordinateScore() int
	toString() string
}

type box struct {
	position position
}

func (b *box) doMovement(move position, world map[position]movable) (map[position]movable, bool) {
	newPosition := position{b.position.x + move.x, b.position.y + move.y}
	m := world[newPosition]
	ok := true
	if m != nil {
		world, ok = m.doMovement(move, world)
	}
	if ok {
		delete(world, b.position)
		b.position = newPosition
		world[newPosition] = b
		return world, ok
	}
	return world, false
}

func (b *box) canMove(move position, world map[position]movable) bool {
	newPosition := position{b.position.x + move.x, b.position.y + move.y}
	m := world[newPosition]
	ok := true
	if m != nil {
		ok = m.canMove(move, world)
	}
	return ok
}

func (b *box) calculateCoordinateScore() int {
	return 100*b.position.y + b.position.x
}

func (b *box) toString() string {
	return "O"
}

type wall struct {
	position position
}

func (w *wall) doMovement(move position, world map[position]movable) (map[position]movable, bool) {
	return world, false
}

func (w *wall) canMove(move position, world map[position]movable) bool {
	return false
}

func (w *wall) calculateCoordinateScore() int {
	return 0
}

func (w *wall) toString() string {
	return "#"
}

type robot struct {
	position position
}

func (bot *robot) doMovement(move position, world map[position]movable) (map[position]movable, bool) {
	newPosition := position{bot.position.x + move.x, bot.position.y + move.y}
	m := world[newPosition]
	ok := true
	if m != nil {
		world, ok = m.doMovement(move, world)
	}
	if ok {
		bot.position = newPosition
	}
	return world, false
}

func (bot *robot) canMove(move position, world map[position]movable) bool {
	newPosition := position{bot.position.x + move.x, bot.position.y + move.y}
	m := world[newPosition]
	ok := true
	if m != nil {
		world, ok = m.doMovement(move, world)
	}
	return ok
}

func (bot *robot) calculateCoordinateScore() int {
	return 0
}

func (bot *robot) toString() string {
	return "@"
}

type wideBox struct {
	positions []position
}

func (w *wideBox) doMovement(move position, world map[position]movable) (map[position]movable, bool) {
	newPositions := lo.Map(w.positions, func(pos position, _ int) position {
		return position{pos.x + move.x, pos.y + move.y}
	})
	obstacles := lo.Uniq(lo.Filter(lo.Map(newPositions, func(pos position, _ int) movable {
		return world[pos]
	}), func(item movable, _ int) bool {
		return item != nil && item != w
	}))
	canMove := lo.EveryBy(obstacles, func(item movable) bool {
		return item.canMove(move, world)
	})
	if !canMove {
		return world, false
	}
	for _, obstacle := range obstacles {
		var ok bool
		world, ok = obstacle.doMovement(move, world)
		if !ok {
			panic("We checked but fucked up!")
		}
	}
	for _, pos := range w.positions {
		delete(world, pos)
	}
	w.positions = newPositions
	for _, pos := range newPositions {
		world[pos] = w
	}
	return world, true
}

func (w *wideBox) canMove(move position, world map[position]movable) bool {
	newPositions := lo.Map(w.positions, func(pos position, _ int) position {
		return position{pos.x + move.x, pos.y + move.y}
	})
	obstacles := lo.Uniq(lo.Filter(lo.Map(newPositions, func(pos position, _ int) movable {
		return world[pos]
	}), func(item movable, _ int) bool {
		return item != nil && item != w
	}))
	canMove := lo.EveryBy(obstacles, func(item movable) bool {
		return item.canMove(move, world)
	})
	return canMove
}

func (w *wideBox) calculateCoordinateScore() int {
	return 100*w.positions[0].y + w.positions[0].x
}

func (w *wideBox) toString() string {
	return "O"
}

var instructionMap = map[string]position{
	"^": position{x: 0, y: -1},
	">": position{x: 1, y: 0},
	"v": position{x: 0, y: 1},
	"<": position{x: -1, y: 0},
}

func Solve(input string) int {
	parts := strings.Split(input, "\n\n")
	world, bot := readWorld(parts[0])
	instructions := strings.Replace(parts[1], "\n", "", -1)
	for _, instruction := range strings.Split(instructions, "") {
		move := instructionMap[instruction]
		world, _ = bot.doMovement(move, world)
	}
	return scoreOfBoxes(world)
}

func SolvePart2(input string) int {
	parts := strings.Split(input, "\n\n")
	world, bot := readWiderWorld(parts[0])
	instructions := strings.Replace(parts[1], "\n", "", -1)
	for _, instruction := range strings.Split(instructions, "") {
		move := instructionMap[instruction]
		world, _ = bot.doMovement(move, world)
	}
	return scoreOfBoxes(world)
}

func readWiderWorld(input string) (world map[position]movable, bot robot) {
	world = make(map[position]movable)
	grid := helpers.StringToGrid(input)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			p := position{x * 2, y}
			switch grid[y][x] {
			case '@':
				bot = robot{p}
			case 'O':
				p2 := position{p.x + 1, y}
				w := wideBox{positions: []position{p, p2}}
				world[p] = &w
				world[p2] = &w
			case '#':
				p2 := position{p.x + 1, y}
				world[p] = &wall{p}
				world[p2] = &wall{p2}
			}
		}
	}
	return
}

func scoreOfBoxes(world map[position]movable) int {
	return lo.Sum(lo.Map(lo.Uniq(lo.Values(world)), func(m movable, _ int) int {
		return m.calculateCoordinateScore()
	}))
}

func readWorld(input string) (world map[position]movable, bot robot) {
	world = make(map[position]movable)
	grid := helpers.StringToGrid(input)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			p := position{x, y}
			switch grid[y][x] {
			case '@':
				bot = robot{p}
			case 'O':
				world[p] = &box{p}

			case '#':
				world[p] = &wall{p}
			}
		}
	}
	return
}

func printWorld(world map[position]movable, bot robot) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 20; x++ {
			p := position{x, y}
			if bot.position == p {
				print(bot.toString())
			} else {
				field := world[p]
				if field != nil {
					print(field.toString())
				} else {
					print(".")
				}
			}
		}
		println()
	}
	println()
}
