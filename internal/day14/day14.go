package day14

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type position struct {
	x, y int
}

type robot struct {
	pos      position
	velocity position
}

func CalculateSafetyFactor(input string, width, height int) int {
	robots := readRobots(input)
	botsAfter100s := lo.Map(robots, func(bot robot, _ int) robot {
		return robot{
			pos: position{
				x: (bot.pos.x + 100*(width+bot.velocity.x)) % width,
				y: (bot.pos.y + 100*(height+bot.velocity.y)) % height},
			velocity: bot.velocity,
		}
	})
	return safetyFactor(botsAfter100s, width, height)
}

func FindChristmastree(input string, width, height int) int {
	robots := readRobots(input)
	output, _ := os.Create("output2.txt")
	defer output.Close()

	for s := 0; s < 10000; s++ {
		robots = lo.Map(robots, func(bot robot, _ int) robot {
			return robot{
				pos: position{
					x: (bot.pos.x + width + bot.velocity.x) % width,
					y: (bot.pos.y + height + bot.velocity.y) % height,
				},
				velocity: bot.velocity,
			}
		})
		//byPosition := lo.GroupBy(robots, func(bot robot) position {
		//	return bot.pos
		//})
		//keys := lo.Keys(byPosition)
		//if lo.Every(keys, expectedTree(width, 4)) {
		if possibleTree(robots) {
			_, _ = fmt.Fprintf(output, "%d\n", s)
			printBots(robots, width, height, output)
		}
		//return s
		//}
	}
	return -1
}

func possibleTree(robots []robot) bool {
	byPosition := lo.GroupBy(robots, func(bot robot) position {
		return bot.pos
	})
	for _, bot := range robots {
		if botIsPartOfTop(bot, byPosition) {
			return true
		}
	}
	return false
}

func botIsPartOfTop(bot robot, byPosition map[position][]robot) bool {
	for i := 0; i < 3; i++ {
		possibleBotLeft := position{
			x: bot.pos.x - i,
			y: bot.pos.y + i,
		}
		possibleBotRight := position{
			x: bot.pos.x + i,
			y: bot.pos.y + i,
		}
		if _, ok := byPosition[possibleBotLeft]; !ok {
			return false
		}
		if _, ok := byPosition[possibleBotRight]; !ok {
			return false
		}
	}
	return true
}

func expectedTree(width int, height int) []position {
	positions := []position{}
	for y := 0; y < height; y++ {
		center := width / 2
		positions = append(positions, position{x: center + y, y: y})
		if y > 0 {
			positions = append(positions, position{x: center - y, y: y})
		}
	}
	return positions
}

func printBots(robots []robot, width, height int, output io.Writer) {
	byPosition := lo.GroupBy(robots, func(bot robot) position {
		return bot.pos
	})
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if _, ok := byPosition[position{x, y}]; ok {
				fmt.Fprint(output, "X")
			} else {
				fmt.Fprint(output, ".")
			}
		}
		fmt.Fprintln(output)
	}
	fmt.Fprintln(output)
}

func safetyFactor(bots []robot, width int, height int) int {
	grouped := lo.GroupBy(bots, func(bot robot) int {
		switch {
		case bot.pos.x < width/2 && bot.pos.y < height/2:
			return 0
		case bot.pos.x > width/2 && bot.pos.y < height/2:
			return 1
		case bot.pos.x > width/2 && bot.pos.y > height/2:
			return 2
		case bot.pos.x < width/2 && bot.pos.y > height/2:
			return 3
		default:
			return -1
		}
	})
	return len(grouped[0]) * len(grouped[1]) * len(grouped[2]) * len(grouped[3])
}

func readRobots(input string) []robot {
	lines := strings.Split(input, "\n")

	return lo.Map(lines, func(line string, _ int) robot {
		var pos, vel string
		fmt.Sscanf(line, "p=%s v=%s", &pos, &vel)
		posSplit := strings.Split(pos, ",")
		velSplit := strings.Split(vel, ",")
		posx, _ := strconv.Atoi(posSplit[0])
		posy, _ := strconv.Atoi(posSplit[1])
		vx, _ := strconv.Atoi(velSplit[0])
		vy, _ := strconv.Atoi(velSplit[1])
		return robot{
			pos:      position{x: posx, y: posy},
			velocity: position{x: vx, y: vy},
		}
	})
}
