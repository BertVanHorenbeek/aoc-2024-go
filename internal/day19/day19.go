package day19

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/samber/lo"
)

func CountPossibleDesigns(input string) int {
	towels, designs := parseInput(input)

	return lo.CountBy(designs, func(design string) bool {
		return isDesignPossible(towels, design)
	})
}

func CountDesignOptions(input string) int {
	towels, designs := parseInput(input)
	possibleDesigns := lo.Filter(designs, func(design string, _ int) bool {
		return isDesignPossible(towels, design)
	})
	designCache := map[string]int{}
	return lo.Sum(lo.Map(possibleDesigns, func(design string, _ int) int {
		return countOptions(towels, design, &designCache)
	}))
}

func countOptions(towels []string, design string, designCache *map[string]int) int {
	if count, ok := (*designCache)[design]; ok {
		return count
	}
	count := 0
	for _, towel := range towels {
		if tail, ok := strings.CutPrefix(design, towel); ok {
			if len(tail) == 0 {
				count += 1
			} else {
				count += countOptions(towels, tail, designCache)
			}
		}
	}
	(*designCache)[design] = count
	return count
}

func parseInput(input string) ([]string, []string) {
	parts := strings.Split(input, "\n\n")
	towels := lo.Map(strings.Split(parts[0], ","), func(s string, _ int) string {
		return strings.TrimSpace(s)
	})
	designs := strings.Split(parts[1], "\n")
	return towels, designs
}

func isDesignPossible(towels []string, design string) bool {
	regex := regexp.MustCompile(fmt.Sprintf("^(%s)+$", strings.Join(towels, "|")))
	return regex.MatchString(design)
}
