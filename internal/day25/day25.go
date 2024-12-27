package day25

import (
	"regexp"
	"strings"

	"aoc-2024-go/helpers"
)

var allPound = regexp.MustCompile(`^#+\n`)

type lock [5]int
type key [5]int

func CountFittingLockKeyPairs(input string) int {
	blocks := strings.Split(input, "\n\n")
	locks := []lock{}
	keys := []key{}
	for _, block := range blocks {
		if allPound.MatchString(block) {
			locks = append(locks, readLock(block))
		} else {
			keys = append(keys, readKey(block))
		}
	}
	fitting := 0
	for _, l := range locks {
		for _, k := range keys {
			if !haveOverlap(l, k) {
				fitting++
			}
		}
	}
	return fitting
}

func haveOverlap(l lock, k key) bool {
	for i := 0; i < 5; i++ {
		if l[i]+k[i] > 5 {
			return true
		}
	}
	return false
}

func readKey(block string) key {
	newKey := key{}
	grid := helpers.StringToGrid(block)
	for y, row := range grid {
		for x, element := range row {
			if element == '.' {
				newKey[x] = 5 - y
			}
		}
	}
	return newKey
}

func readLock(block string) lock {
	newLock := lock{}
	grid := helpers.StringToGrid(block)
	for y, row := range grid {
		for x, element := range row {
			if element == '#' {
				newLock[x] = y
			}
		}
	}
	return newLock
}
