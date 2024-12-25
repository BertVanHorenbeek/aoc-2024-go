package day23

import (
	"fmt"
	"slices"
	"strings"

	"github.com/samber/lo"
)

func Solve(input string) int {
	connectionMap := buildConnectionMap(input)
	startingWithT := lo.PickBy(connectionMap, func(key string, value []string) bool {
		return strings.HasPrefix(key, "t")
	})

	networks := [][]string{}
	for key, values := range startingWithT {
		for _, value := range values {
			others := lo.Intersect(connectionMap[key], connectionMap[value])
			for _, other := range others {
				networks = append(networks, []string{key, value, other})
			}
		}
	}
	fmt.Printf("The networks are: %v\n", networks)
	return len(lo.UniqBy(networks, func(item []string) string {
		slices.Sort(item)
		return strings.Join(item, ",")
	}))
}

func buildConnectionMap(input string) map[string][]string {
	lines := strings.Split(input, "\n")
	entries := lo.FlatMap(lines, func(line string, _ int) [][]string {
		split := strings.Split(line, "-")
		return [][]string{split, {split[1], split[0]}}
	})
	entryMap := lo.GroupBy(entries, func(item []string) string {
		return item[0]
	})
	connectionMap := lo.MapValues(entryMap, func(value [][]string, key string) []string {
		return lo.Map(value, func(item []string, index int) string {
			return item[1]
		})
	})
	return connectionMap
}

func Solve2(input string) string {
	connectionMap := buildConnectionMap(input)
	networks := [][]string{}
	for key, values := range connectionMap {
		for _, value := range values {
			network := []string{key, value}
			networkCandidates := lo.Intersect(connectionMap[key], connectionMap[value])
			for _, networkCandidate := range networkCandidates {
				if lo.Every(connectionMap[networkCandidate], network) {
					network = append(network, networkCandidate)
				}
			}
			networks = append(networks, network)
		}
	}
	longest := lo.MaxBy(networks, func(a []string, b []string) bool {
		return len(a) > len(b)
	})
	slices.Sort(longest)
	return strings.Join(longest, ",")
}
