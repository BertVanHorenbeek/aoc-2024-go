package day5

import (
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type rule struct {
	first  int
	second int
}

func CountValidQueues(input string) int {
	parts := strings.Split(input, "\n\n")
	ruleBlock := parts[0]
	queuesInput := parts[1]

	rules := parseRules(ruleBlock)
	validQueues := lo.Filter(lo.Map(strings.Split(queuesInput, "\n"), parseQueue), validateQueue(rules))
	return lo.Sum(lo.Map(validQueues, func(queue []int, index int) int {
		return queue[len(queue)/2]
	}))
}

func FixInvalidQueues(input string) int {
	parts := strings.Split(input, "\n\n")
	ruleBlock := parts[0]
	queuesInput := parts[1]

	rules := parseRules(ruleBlock)
	inValidQueues := lo.Reject(lo.Map(strings.Split(queuesInput, "\n"), parseQueue), validateQueue(rules))
	fixedQueues := lo.Map(inValidQueues, fixQueue(rules))
	return lo.Sum(lo.Map(fixedQueues, func(queue []int, index int) int {
		return queue[len(queue)/2]
	}))
}

func fixQueue(rules []rule) func([]int, int) []int {
	return func(queue []int, _ int) []int {
		sortedList := make([]int, 0, len(queue))
		remainingElements := queue
		for i := 0; i < len(queue); i++ {
			if len(remainingElements) == 1 {
				sortedList = append(sortedList, remainingElements[0])
			} else {
				relevantRules := filterRelevantRules(rules, remainingElements)
				smallest := smallestElement(relevantRules)
				sortedList = append(sortedList, smallest)
				remainingElements = lo.Without(remainingElements, smallest)
			}
		}
		return sortedList
	}
}

func filterRelevantRules(rules []rule, queue []int) []rule {
	return lo.Filter(rules, func(r rule, _ int) bool {
		return lo.Every(queue, []int{r.first, r.second})
	})
}

func smallestElement(rules []rule) int {
	firstElements := lo.Map(rules, func(r rule, _ int) int {
		return r.first
	})
	secondElements := lo.Map(rules, func(r rule, _ int) int {
		return r.second
	})
	return lo.FirstOrEmpty(lo.Without(firstElements, secondElements...))
}

func parseQueue(queueAsString string, _ int) []int {
	parts := strings.Split(queueAsString, ",")
	return lo.Map(parts, func(item string, _ int) int {
		page, _ := strconv.Atoi(item)
		return page
	})
}

func validateQueue(rules []rule) func([]int, int) bool {
	return func(queue []int, _ int) bool {
		for i, page := range queue {
			pagesLargerThenCurrentPage := pageIsSmallerThen(page, rules)
			if len(lo.Intersect(pagesLargerThenCurrentPage, queue[:i])) > 0 {
				return false
			}
		}
		return true
	}
}

func parseRules(block string) []rule {
	return lo.Map(strings.Split(block, "\n"), func(line string, _ int) rule {
		parts := strings.Split(line, "|")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		return rule{
			first:  first,
			second: second,
		}
	})
}

func pageIsSmallerThen(page int, rules []rule) []int {
	return lo.Map(lo.Filter(rules, func(item rule, _ int) bool {
		return item.first == page
	}), func(r rule, _ int) int {
		return r.second
	})
}
