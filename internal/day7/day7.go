package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type calibration struct {
	expectedResult int
	operants       []int
}

var operations = []func(int, int) int{
	func(n1, n2 int) int {
		return n1 + n2
	},
	func(n1, n2 int) int {
		return n1 * n2
	},
}

func TotalCalibrationResult(input string) int {
	lines := strings.Split(input, "\n")
	calibrations := lo.Map(lines, parseCalibrationLine)

	total := 0

	for _, c := range calibrations {
		solutions := listPossibleSolutions(c.operants, operations)
		if lo.Contains(solutions, c.expectedResult) {
			total += c.expectedResult
		}
	}
	return total
}

func TotalCalibrationUsingReduce(input string) int {
	lines := strings.Split(input, "\n")
	calibrations := lo.Map(lines, parseCalibrationLine)

	total := 0

	for _, c := range calibrations {
		solutions := lo.Reduce(c.operants, func(agg []int, nextValue int, _ int) []int {
			if len(agg) == 0 {
				return []int{nextValue}
			}
			return lo.FlatMap(agg, func(subTotal int, _ int) []int {
				return lo.Map(operations, func(op func(int, int) int, _ int) int {
					return op(subTotal, nextValue)
				})
			})
		}, []int{})
		if lo.Contains(solutions, c.expectedResult) {
			total += c.expectedResult
		}
	}
	return total
}

func TotalCalibrationResultWithConcat(input string) int {
	lines := strings.Split(input, "\n")
	calibrations := lo.Map(lines, parseCalibrationLine)

	total := 0

	for _, c := range calibrations {
		solutions := listPossibleSolutions(c.operants, append(operations, func(i int, i2 int) int {
			atoi, _ := strconv.Atoi(fmt.Sprintf("%d%d", i2, i))
			return atoi
		}))
		if lo.Contains(solutions, c.expectedResult) {
			total += c.expectedResult
		}
	}
	return total
}

func parseCalibrationLine(line string, _ int) calibration {
	parts := strings.Split(line, ":")
	expectedResult, _ := strconv.Atoi(parts[0])
	operants := lo.Map(strings.Fields(parts[1]), func(item string, _ int) int {
		n, _ := strconv.Atoi(item)
		return n
	})
	return calibration{
		expectedResult: expectedResult,
		operants:       operants,
	}
}

func listPossibleSolutions(operants []int, ops []func(int, int) int) []int {
	if len(operants) == 1 {
		return operants
	}
	lastIndex := len(operants) - 1
	firstNumber := operants[lastIndex]
	rest := operants[:lastIndex]
	possibleSolutions := []int{}
	for _, secondNumber := range listPossibleSolutions(rest, ops) {
		for _, operation := range ops {
			possibleSolutions = append(possibleSolutions, operation(firstNumber, secondNumber))
		}
	}
	return possibleSolutions
}
