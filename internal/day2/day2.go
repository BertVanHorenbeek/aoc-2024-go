package day2

import (
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func ValidateReports(input string) int {
	lines := strings.Split(input, "\n")
	return lo.CountBy(lines, validateReport)
}

func ValidateReportsWithDamper(input string) int {
	lines := strings.Split(input, "\n")
	return lo.CountBy(lines, validateReportWithDamper)
}

func validateReport(report string) bool {
	values := lineToIntSlice(report)

	allIncreasing, allDecreasing := true, true

	for i := 1; i < len(values); i++ {
		allIncreasing = allIncreasing && values[i-1] < values[i]
		allDecreasing = allDecreasing && values[i-1] > values[i]
		if abs(values[i-1]-values[i]) > 3 {
			return false
		}
	}
	return allDecreasing || allIncreasing
}

func validateReportWithDamper(report string) bool {
	values := lineToIntSlice(report)
	for damper := 0; damper < len(values); damper++ {
		dampened := make([]int, len(values))
		copy(dampened, values)
		dampened = slices.Delete(dampened, damper, damper+1)
		if validate(dampened) {
			return true
		}
	}
	return false
}

func validate(values []int) bool {
	allIncreasing, allDecreasing := true, true
	for i := 1; i < len(values); i++ {
		increasing := values[i-1] < values[i]
		allIncreasing = allIncreasing && increasing
		allDecreasing = allDecreasing && values[i-1] > values[i]
		if abs(values[i-1]-values[i]) > 3 {
			return false
		}
	}
	return allDecreasing || allIncreasing
}

func lineToIntSlice(report string) []int {
	valuesAsStrings := strings.Fields(report)
	values := lo.Map(valuesAsStrings, func(val string, _ int) int {
		atoi, err := strconv.Atoi(val)
		if err != nil {
			panic("could not convert string to number")
		}
		return atoi
	})
	return values
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
