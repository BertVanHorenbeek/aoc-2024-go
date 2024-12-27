package day24

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"

	"github.com/samber/lo"
)

type instruction struct {
	operants  []string
	operation string
	result    string
}

type instructionList []instruction

func (l instructionList) findByResult(result string) (instruction, bool) {
	for _, inst := range l {
		if inst.result == result {
			return inst, true
		}
	}
	return instruction{}, false
}

func (l instructionList) findByOperantsAndOperation(operants []string, operation string) (instruction, bool) {
	for _, inst := range l {
		if lo.Every(inst.operants, operants) && operation == inst.operation {
			return inst, true
		}
	}
	return instruction{}, false
}

func Solve(input string) int {
	parts := strings.Split(input, "\n\n")
	stateMap := readInitialState(parts[0])
	instructions := readInstructions(parts[1])
	for actionPerformed := true; actionPerformed; {
		actionPerformed = false
		for _, inst := range instructions {
			if _, ok := stateMap[inst.operants[0]]; !ok {
				continue
			}
			if _, ok := stateMap[inst.operants[1]]; !ok {
				continue
			}
			if _, ok := stateMap[inst.result]; ok {
				continue
			}
			actionPerformed = true
			switch inst.operation {
			case "AND":
				stateMap[inst.result] = stateMap[inst.operants[0]] && stateMap[inst.operants[1]]
			case "OR":
				stateMap[inst.result] = stateMap[inst.operants[0]] || stateMap[inst.operants[1]]
			case "XOR":
				stateMap[inst.result] = stateMap[inst.operants[0]] != stateMap[inst.operants[1]]
			}
		}
	}
	outputMap := lo.PickBy(stateMap, func(key string, value bool) bool {
		return strings.HasPrefix(key, "z")
	})
	return readOutput(outputMap)
}

func FindSwappedPairs(input string) string {
	parts := strings.Split(input, "\n\n")
	stateMap := readInitialState(parts[0])
	instructions := readInstructions(parts[1])
	xMap := lo.PickBy(stateMap, func(key string, value bool) bool {
		return strings.HasPrefix(key, "x")
	})
	x := readOutput(xMap)
	yMap := lo.PickBy(stateMap, func(key string, value bool) bool {
		return strings.HasPrefix(key, "y")
	})
	y := readOutput(yMap)
	expectedOutput := x + y

	realOutput := performInstructions(instructions, stateMap)
	difference := realOutput ^ expectedOutput

	wrongNumbers := listBinOnes(difference)

	wrongOutputs := findUpstreamOutputs(wrongNumbers, instructions)

	pairs := make([][2]int, 0, len(instructions)*len(instructions))
	for i := 0; i < len(instructions); i++ {
		for j := i + 1; j < len(instructions); j++ {
			pairs = append(pairs, [2]int{i, j})
		}
	}

	for i := 0; i < len(pairs); i++ {
		for j := i + 1; j < len(pairs); j++ {
			for k := j + 1; k < len(pairs); k++ {
				for l := k + 1; l < len(pairs); l++ {
					swappedInstructions := swapInstructions(instructions, pairs[i], pairs[j], pairs[k], pairs[l])
					swappedOutput := performInstructions(swappedInstructions, stateMap)
					if swappedOutput == expectedOutput {
						fmt.Printf("We found a winner, %d, %d, %d, %d\n", i, j, k, l)
					}
				}
			}
		}
	}

	_ = wrongOutputs

	fmt.Printf("The real output is: %v but the expected was %v\n", realOutput, expectedOutput)
	return ""
}

func FindSwappedPairs2(input string) string {
	parts := strings.Split(input, "\n\n")
	instructions := readInstructions(parts[1])

	cary := "mkf"
	for i := 1; i < 45; i++ {
		inputs := []string{fmt.Sprintf("x%02d", i), fmt.Sprintf("y%02d", i)}

		firstXorInstruction, xorFound := lo.Find(instructions, func(item instruction) bool {
			return lo.Every(item.operants, inputs) && item.operation == "XOR"
		})
		if !xorFound {
			panic(fmt.Sprintf("xor instruction not found for %s and %s", inputs[0], inputs[1]))
		}
		secondXorInstruction, xorFound := lo.Find(instructions, func(item instruction) bool {
			return lo.Every(item.operants, []string{cary, firstXorInstruction.result}) && item.operation == "XOR"
		})
		if !xorFound {
			panic(fmt.Sprintf("second xor instruction not found for %s and %s", cary, firstXorInstruction.result))
		}
		if secondXorInstruction.result != fmt.Sprintf("z%02d", i) {
			panic(fmt.Sprintf("output for %d is not correct", i))
		}
		firstAndInstruction, andFound := lo.Find(instructions, func(item instruction) bool {
			return lo.Every(item.operants, []string{cary, firstXorInstruction.result}) && item.operation == "AND"
		})
		if !andFound {
			panic(fmt.Sprintf("firts and instruction not found for %s and %s", cary, firstXorInstruction.result))
		}
		secondAndInstruction, andFound := lo.Find(instructions, func(item instruction) bool {
			return lo.Every(item.operants, inputs) && item.operation == "AND"
		})
		if !andFound {
			panic(fmt.Sprintf("second and instruction not found for %s and %s", inputs[0], inputs[1]))
		}
		orInstruction, orFound := lo.Find(instructions, func(item instruction) bool {
			return lo.Every(item.operants, []string{firstAndInstruction.result, secondAndInstruction.result}) && item.operation == "OR"
		})
		if !orFound {
			panic(fmt.Sprintf("or instruction not found for %s and %s", firstAndInstruction.result, secondAndInstruction.result))
		}
		cary = orInstruction.result
	}
	return ""
}

func FindSwappedPairs3(input string) string {
	parts := strings.Split(input, "\n\n")
	instructions := readInstructions(parts[1])
	swappableFields := []string{}
	orGates := lo.Filter(instructions, func(item instruction, _ int) bool {
		return item.operation == "OR"
	})
	andGates := lo.Filter(instructions, func(item instruction, _ int) bool {
		return item.operation == "AND"
	})
	xorGates := lo.Filter(instructions, func(item instruction, _ int) bool {
		return item.operation == "XOR"
	})
	_ = xorGates

	swappableFields = append(swappableFields, allAndsGoToOr(andGates, orGates)...)
	swappableFields = append(swappableFields, allAndsGoToOr(andGates, orGates)...)

	//fmt.Printf("The diffs are %v and %v\n", diff1, diff2)
	return ""
}

func FindSwappedPairs4(input string) string {
	parts := strings.Split(input, "\n\n")
	instructions := readInstructions(parts[1])
	for i := 44; i >= 0; i-- {
		inputs := []string{fmt.Sprintf("x%02d", i), fmt.Sprintf("y%02d", i)}
		_ = fmt.Sprintf("x%02d", i)
		_, xorFound := lo.Find(instructions, func(item instruction) bool {
			return lo.Every(item.operants, inputs) && item.operation == "XOR"
		})
		if !xorFound {
			panic(fmt.Sprintf("xor instruction not found for %s and %s", inputs[0], inputs[1]))
		}

	}
	return ""
}

func FindSwappedPairs5(input string) string {
	parts := strings.Split(input, "\n\n")
	instructions := readInstructions(parts[1])
	instructions = swap(instructions, "z05", "bpf")
	instructions = swap(instructions, "z11", "hcc")
	instructions = swap(instructions, "z35", "fdw")
	instructions = swap(instructions, "hqc", "qcw")

	swappableFields := []string{}
	for i := 0; i < 44; i++ {
		if _, ok := checkKepingRecords(buildEmptyMap(), instructions, []string{fmt.Sprintf("x%02d", i)}, []string{fmt.Sprintf("z%02d", i)}); !ok {
			swappableFields = append(swappableFields, fmt.Sprintf("z%02d", i))
			fmt.Printf("x%02d = 1 y%02d = 0 expected z%02d = 1\n", i, i, i)
		}
		if _, ok := checkKepingRecords(buildEmptyMap(), instructions, []string{fmt.Sprintf("y%02d", i)}, []string{fmt.Sprintf("z%02d", i)}); !ok {
			swappableFields = append(swappableFields, fmt.Sprintf("z%02d", i))
			fmt.Printf("x%02d = 0 y%02d = 1 expected z%02d = 1\n", i, i, i)
		}
		if _, ok := checkKepingRecords(buildEmptyMap(), instructions, []string{fmt.Sprintf("y%02d", i), fmt.Sprintf("x%02d", i)}, []string{fmt.Sprintf("z%02d", i+1)}); !ok {
			swappableFields = append(swappableFields, fmt.Sprintf("z%02d", i))
			fmt.Printf("x%02d = 1 y%02d = 1 expected z%02d = 0 and z%02d = 1\n", i, i, i, i+1)
		}
	}
	return strings.Join(lo.Uniq(swappableFields), ",")
}

func swap(instructions instructionList, result1 string, result2 string) instructionList {
	swapped := make([]instruction, len(instructions))
	for index, inst := range instructions {
		if inst.result == result1 {
			swapped[index] = instruction{
				operants:  inst.operants,
				operation: inst.operation,
				result:    result2,
			}
		} else if inst.result == result2 {
			swapped[index] = instruction{
				operants:  inst.operants,
				operation: inst.operation,
				result:    result1,
			}
		} else {
			swapped[index] = inst
		}
	}
	return swapped
}

var inputField = regexp.MustCompile("^x|y[0-9]+")
var outputField = regexp.MustCompile("^z[0-9]+")

func ignoreInputFields(item string, _ int) bool {
	return !inputField.MatchString(item)
}

func checkKepingRecords(stateMap map[string]bool, instructions instructionList, onesInInput []string, expectedOnesInOutput []string) (map[string]bool, bool) {
	for _, one := range onesInInput {
		stateMap[one] = true
	}
	resultMap := onlyPerform(instructions, stateMap)
	for key, value := range resultMap {
		if !outputField.MatchString(key) {
			continue
		}
		if (lo.Contains(expectedOnesInOutput, key) && !value) ||
			((!lo.Contains(expectedOnesInOutput, key)) && value) {
			return resultMap, false
		}
	}
	return resultMap, true
}

func buildEmptyMap() map[string]bool {
	emptyMap := make(map[string]bool)
	for i := 0; i <= 44; i++ {
		emptyMap[fmt.Sprintf("x%02d", i)] = false
		emptyMap[fmt.Sprintf("y%02d", i)] = false
	}
	return emptyMap
}

func validateAdder(result string, instructions instructionList) []string {
	node, _ := instructions.findByResult(result)
	if !(node.operation == "XOR") {
		println("")
	}
	return []string{}
}
func allAndsGoToOr(andGates []instruction, orGates []instruction) []string {
	andExits := lo.Map(andGates, func(item instruction, _ int) string {
		return item.result
	})
	orInputs := lo.FlatMap(orGates, func(item instruction, _ int) []string {
		return item.operants
	})
	diff1, diff2 := lo.Difference(orInputs, andExits)
	return lo.Union(diff1, diff2)
}

func swapInstructions(instructions []instruction, swaps ...[2]int) []instruction {
	swapped := make([]instruction, len(instructions))
	for index, inst := range instructions {
		if swap, found := lo.Find(swaps, func(swap [2]int) bool {
			return swap[0] == index
		}); found {
			swapped[index] = instruction{
				operants:  []string{inst.operants[1], inst.operants[0]},
				operation: inst.operation,
				result:    instructions[swap[1]].result,
			}
		} else if swap2, found2 := lo.Find(swaps, func(swap [2]int) bool {
			return swap[1] == index
		}); found2 {
			swapped[index] = instruction{
				operants:  []string{inst.operants[1], inst.operants[0]},
				operation: inst.operation,
				result:    instructions[swap2[0]].result,
			}
		} else {
			swapped[index] = inst
		}
	}
	return swapped
}

func findUpstreamOutputs(numbers []string, instructions []instruction) []string {
	return lo.Uniq(lo.FlatMap(numbers, func(output string, _ int) []string {
		inputs := lo.FlatMap(lo.Filter(instructions, func(item instruction, _ int) bool {
			return item.result == output
		}), func(item instruction, _ int) []string {
			return item.operants
		})
		return append(findUpstreamOutputs(inputs, instructions), output)
	}))
}

func listBinOnes(difference int) []string {
	ones := make([]string, 0, 45)
	remaining := difference
	for i := 44; i >= 0; i-- {
		if remaining >= int(math.Pow(2, float64(i))) {
			remaining -= int(math.Pow(2, float64(i)))
			ones = append(ones, fmt.Sprintf("z%02d", i))
		}
	}
	return ones
}

func performInstructions(instructions []instruction, stateMap map[string]bool) int {
	stateMap = onlyPerform(instructions, stateMap)
	outputMap := lo.PickBy(stateMap, func(key string, value bool) bool {
		return strings.HasPrefix(key, "z")
	})
	return readOutput(outputMap)
}

func onlyPerform(instructions []instruction, stateMap map[string]bool) map[string]bool {
	for actionPerformed := true; actionPerformed; {
		actionPerformed = false
		for _, inst := range instructions {
			if _, ok := stateMap[inst.operants[0]]; !ok {
				continue
			}
			if _, ok := stateMap[inst.operants[1]]; !ok {
				continue
			}
			if _, ok := stateMap[inst.result]; ok {
				continue
			}
			actionPerformed = true
			switch inst.operation {
			case "AND":
				stateMap[inst.result] = stateMap[inst.operants[0]] && stateMap[inst.operants[1]]
			case "OR":
				stateMap[inst.result] = stateMap[inst.operants[0]] || stateMap[inst.operants[1]]
			case "XOR":
				stateMap[inst.result] = stateMap[inst.operants[0]] != stateMap[inst.operants[1]]
			}
		}
	}

	return stateMap
}

func readOutput(outputMap map[string]bool) int {
	keys := lo.Keys(outputMap)
	slices.Sort(keys)
	number := 0
	for i, key := range keys {
		if outputMap[key] {
			number += int(math.Pow(2, float64(i)))
		}
	}
	return number
}

func readInstructions(input string) instructionList {
	lines := strings.Split(input, "\n")
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		instructions[i] = instruction{
			operants:  []string{parts[0], parts[2]},
			operation: parts[1],
			result:    parts[4],
		}
	}
	return instructions
}

func readInitialState(input string) map[string]bool {
	stateMap := make(map[string]bool)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		stateMap[parts[0]] = parts[1] == "1"
	}
	return stateMap
}

func WriteOutFormulas(input string) string {
	parts := strings.Split(input, "\n\n")
	instructions := readInstructions(parts[1])
	for i := 0; i <= 45; i++ {
		result := fmt.Sprintf("z%02d", i)
		expanded := expandInstruction(result, instructions)
		fmt.Printf("%s = %s\n", expanded, result)
	}
	return ""
}

func expandInstruction(result string, instructions instructionList) string {
	if found, ok := instructions.findByResult(result); ok {
		return fmt.Sprintf("(%s) %s (%s)", expandInstruction(found.operants[0], instructions), found.operation, expandInstruction(found.operants[1], instructions))
	} else {
		return result
	}
}
