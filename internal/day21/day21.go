package day21

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type position struct {
	x, y int
}

type keypad struct {
	keyMap           map[string]position
	fingerPosition   position
	forbidden        position
	allMovesAreEqual bool
}

func (pad *keypad) moveFingerTo(key string) position {
	destination := pad.keyMap[key]
	difference := position{destination.x - pad.fingerPosition.x, destination.y - pad.fingerPosition.y}
	pad.fingerPosition = destination
	return difference
}

func (pad *keypad) moveFingerAndReturnMovementOptions(key string) []string {
	destination := pad.keyMap[key]
	start := pad.fingerPosition
	difference := pad.moveFingerTo(key)
	moveOptions := make([]string, 0, 2)
	if !(destination.x == pad.forbidden.x && start.y == pad.forbidden.y) {
		move2 := ""
		if difference.x > 0 {
			for x := 0; x < difference.x; x++ {
				move2 += ">"
			}
		}
		if difference.x < 0 {
			for x := 0; x > difference.x; x-- {
				move2 += "<"
			}
		}
		if difference.y > 0 {
			for y := 0; y < difference.y; y++ {
				move2 += "v"
			}
		}
		if difference.y < 0 {
			for y := 0; y > difference.y; y-- {
				move2 += "^"
			}
		}
		move2 += "A"
		moveOptions = append(moveOptions, move2)
	}
	if !(destination.y == pad.forbidden.y && start.x == pad.forbidden.x) {
		move1 := ""
		if difference.y > 0 {
			for y := 0; y < difference.y; y++ {
				move1 += "v"
			}
		}
		if difference.y < 0 {
			for y := 0; y > difference.y; y-- {
				move1 += "^"
			}
		}
		if difference.x > 0 {
			for x := 0; x < difference.x; x++ {
				move1 += ">"
			}
		}
		if difference.x < 0 {
			for x := 0; x > difference.x; x-- {
				move1 += "<"
			}
		}
		move1 += "A"
		moveOptions = append(moveOptions, move1)
	}
	if pad.allMovesAreEqual {
		return []string{moveOptions[0]}
	}
	return moveOptions
}

func CalculateComplexity(input string, numbots int) int {
	codes := strings.Split(input, "\n")
	//calculationMap := make(map[string]int)
	//for _, arrow := range strings.Split("><^vA", "") {
	//	calculationMap[arrow] = calculateArrow(arrow, 2)
	//}
	return lo.Sum(lo.Map(codes, func(code string, _ int) int {
		codeAsInt, err := strconv.Atoi(code[:len(code)-1])
		if err != nil {
			panic(err)
		}

		sequence := calculateShortestSequence(code, numbots)
		return len(sequence) * codeAsInt
	}))
}

func CalculateComplexityForX(input string, calculationDepth int) int {
	codes := strings.Split(input, "\n")
	return lo.Sum(lo.Map(codes, func(code string, _ int) int {
		codeAsInt, err := strconv.Atoi(code[:len(code)-1])
		if err != nil {
			panic(err)
		}
		return calculateShortestSequenceWithSplitRecursive(code, calculationDepth) * codeAsInt
	}))
}

func calculateShortestSequenceWithSplit(code string, calculationDepth int) int {
	numericKeypad := newNumericKeypad()
	directionalKeypad := newDirectionalKeypad()

	inputForDirectional := lo.Uniq(handleCodeForKeypadImproved(numericKeypad, code))

	results := lo.Map(inputForDirectional, func(code string, _ int) int {
		currentValues := stringToValueMap(code)
		for i := 0; i < calculationDepth+1; i++ {
			nextValues := map[string]int{}
			for key, val := range currentValues {
				resultingString := handleCodeForKeypadImproved(directionalKeypad, key)[0]
				resultingMap := stringToValueMap(resultingString)
				for k, v := range resultingMap {
					nextValues[k] += val * v
				}
			}
			currentValues = nextValues
		}
		score := 0
		for _, val := range currentValues {
			score += val
		}
		return score
	})
	return lo.Min(results)
}

func calculateShortestSequenceWithSplitRecursive(code string, calculationDepth int) int {
	numericKeypad := newNumericKeypad()

	inputForDirectional := lo.Uniq(handleCodeForKeypadImproved(numericKeypad, code))
	return lo.Min(lo.Map(inputForDirectional, func(code string, _ int) int {
		return calculateForKeypads(code, calculationDepth)
	}))
}

type cacheKey struct {
	code  string
	depth int
}

var calculationCache = map[cacheKey]int{}

func calculateForKeypads(input string, calculationDepth int) int {
	if val, ok := calculationCache[cacheKey{input, calculationDepth}]; ok {
		return val
	}
	if calculationDepth == 0 {
		return len(input)
	}
	directionalKeypad := newDirectionalKeypad()
	neededOnKeypadList := handleCodeForKeypadImproved(directionalKeypad, input)
	bestScore := 0
	for _, neededOnKeypad := range neededOnKeypadList {
		score := 0
		currentValues := stringToValueMap(neededOnKeypad)
		for key, val := range currentValues {
			score += val * calculateForKeypads(key, calculationDepth-1)
		}
		if bestScore > score || bestScore == 0 {
			bestScore = score
		}
	}
	calculationCache[cacheKey{input, calculationDepth}] = bestScore
	return bestScore
}

func stringToValueMap(input string) map[string]int {
	parts := strings.Split(input, "A")
	withA := lo.Map(parts[:len(parts)-1], func(code string, _ int) string {
		return code + "A"
	})
	return lo.CountValues(withA)
}

func buildArrowMap(numbots int) map[string]int {
	arrowMap := map[string]int{}
	for _, item := range []string{"v<<A", ">>^A", "v<A", "<A", "vA", "^<A", ">A", "^A", "^>A", "v>A", "A", ">^^A", "vvvA", ">^A"} {
		arrowMap[item] = calculateArrow(item, numbots)
	}
	return arrowMap
}

func newNumericKeypad() keypad {
	return keypad{
		keyMap: map[string]position{
			"7": {0, 0}, "8": {1, 0}, "9": {2, 0},
			"4": {0, 1}, "5": {1, 1}, "6": {2, 1},
			"1": {0, 2}, "2": {1, 2}, "3": {2, 2},
			"0": {1, 3}, "A": {2, 3}},
		fingerPosition: position{2, 3},
		forbidden:      position{0, 3},
	}
}

func newDirectionalKeypad() keypad {
	return keypad{
		keyMap: map[string]position{
			"^": {1, 0}, "A": {2, 0},
			"<": {0, 1}, "v": {1, 1}, ">": {2, 1}},
		fingerPosition: position{2, 0},
		forbidden:      position{0, 0},
	}
}

func calculateArrow(code string, numberOfBots int) int {
	directionalKeypad := newDirectionalKeypad()
	inputForDirectional := []string{code}
	for i := 0; i < numberOfBots; i++ {
		inputForDirectional = lo.Uniq(lo.FlatMap(inputForDirectional, func(input string, _ int) []string {
			return handleCodeForKeypadImproved(directionalKeypad, input)
		}))
		inputForDirectional = onlyKeepBestOptions(inputForDirectional)
	}
	shortest := lo.MinBy(inputForDirectional, func(a string, b string) bool {
		return len(a) < len(b)
	})
	return len(shortest)
}

func calculateShortestSequence(code string, numberOfDirectionalKeypads int) string {
	numericKeypad := newNumericKeypad()
	directionalKeypad := newDirectionalKeypad()

	inputForDirectional := handleCodeForKeypadImproved(numericKeypad, code)

	for i := 0; i < numberOfDirectionalKeypads; i++ {
		inputForDirectional = lo.Uniq(lo.FlatMap(inputForDirectional, func(input string, _ int) []string {
			return handleCodeForKeypadImproved(directionalKeypad, input)
		}))
		inputForDirectional = onlyKeepBestOptions(inputForDirectional)
	}

	return lo.MinBy(inputForDirectional, func(a string, b string) bool {
		return len(a) < len(b)
	})
}

func calculateShortestSequenceBetterPerformance(code string, numberOfBots int, calculationMap map[string]int) int {
	numericKeypad := newNumericKeypad()
	directionalKeypad := newDirectionalKeypad()

	inputForDirectional := handleCodeForKeypadImproved(numericKeypad, code)

	for i := 0; i < numberOfBots; i++ {
		inputForDirectional = lo.Uniq(lo.FlatMap(inputForDirectional, func(input string, _ int) []string {
			return handleCodeForKeypadImproved(directionalKeypad, input)
		}))
		inputForDirectional = onlyKeepBestOptions(inputForDirectional)
	}
	bestScore := 1000000000000000000
	for _, t := range inputForDirectional {
		score := 0
		parts := strings.Split(t, "A")
		for _, arrow := range parts[:len(parts)-1] {
			i, ok := calculationMap[arrow+"A"]
			if !ok {
				fmt.Printf("%s is not found\n", arrow)
				panic(arrow + " is not found")
			}
			score += i
		}
		if score < bestScore {
			bestScore = score
		}
	}
	return bestScore
}

func onlyKeepBestOptions(movements []string) []string {
	bestScore := lo.MinBy(movements, func(a string, b string) bool {
		return len(a) < len(b)
	})
	filtered := lo.Filter(movements, func(item string, _ int) bool {
		return len(item) < (len(bestScore) + 5)
	})
	//if len(filtered) > 10000000 {
	//	return filtered[:10000000]
	//}
	return filtered
}

func handleCodeForKeypad(pad keypad, code string) []position {
	moves := []position{}
	for _, number := range strings.Split(code, "") {
		moves = append(moves, pad.moveFingerTo(number))
	}
	return moves
}

func handleCodeForKeypadImproved(pad keypad, code string) []string {
	options := []string{""}
	for _, number := range strings.Split(code, "") {
		movementOptions := pad.moveFingerAndReturnMovementOptions(number)
		options = lo.FlatMap(movementOptions, func(newMovementOption string, _ int) []string {
			return lo.Map(options, func(existingOption string, _ int) string {
				return existingOption + newMovementOption
			})
		})
		options = lo.Uniq(options)
	}
	return options
}
