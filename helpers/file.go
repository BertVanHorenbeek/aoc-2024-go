package helpers

import (
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo"
)

func ReadInput(day string) (string, error) {
	file, err := os.ReadFile(fmt.Sprintf("assets/%s/input.txt", day))
	if err != nil {
		return "", err
	}
	fileString := string(file)
	return fileString, nil
}

func StringToGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")
	return lo.Map(lines, lineToFields)
}

func lineToFields(line string, _ int) []rune {
	return []rune(line)
}
