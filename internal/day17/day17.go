package day17

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func RunProgram(input string) string {
	parts := strings.Split(input, "\n\n")
	registers := loadInitialState(parts[0])
	program := loadProgram(parts[1])
	out := run(program, registers, "")
	return strings.Join(out, ",")
}

func run(program []int, registers map[string]int, expected string) []string {
	aregister := registers["A"]
	out := []string{}
	instructionPointer := 0
	for instructionPointer < len(program) {
		//printState(registers, program, instructionPointer)
		if !strings.HasPrefix(expected, strings.Join(out, ",")) && expected != "" {
			if len(out) > 12 {
				fmt.Printf("A register with %d is 8 long\n", aregister)
			}
			return out
		}
		switch program[instructionPointer] {
		case 0:
			registers["A"] = registers["A"] / int(math.Pow(2, float64(combo(program[instructionPointer+1], registers))))
		case 1:
			registers["B"] = registers["B"] ^ program[instructionPointer+1]
		case 2:
			registers["B"] = combo(program[instructionPointer+1], registers) % 8
		case 3:
			if registers["A"] != 0 {
				instructionPointer = program[instructionPointer+1] - 2
			}
		case 4:
			registers["B"] = registers["B"] ^ registers["C"]
		case 5:
			out = append(out, strconv.Itoa(combo(program[instructionPointer+1], registers)%8))
		case 6:
			registers["B"] = registers["A"] / int(math.Pow(2, float64(combo(program[instructionPointer+1], registers))))
		case 7:
			registers["C"] = registers["A"] / int(math.Pow(2, float64(combo(program[instructionPointer+1], registers))))
		}
		instructionPointer += 2
	}
	return out
}

func FindStartToProduceItSelf(input string) int {
	parts := strings.Split(input, "\n\n")
	registers := loadInitialState(parts[0])
	program := loadProgram(parts[1])
	expected := lo.Map(program, func(item int, _ int) string {
		return strconv.Itoa(item)
	})
	expectedString := strings.Join(expected, ",")
	i := 1447123790738
	for ; i < 39846323390557074; i += 268435456 {
		runRegisters := map[string]int{"A": i, "B": registers["B"], "C": registers["C"]}
		out := run(program, runRegisters, expectedString)
		if len(out) == len(expected) && out[len(expected)-1] == expected[len(expected)-1] {
			break
		}
		//fmt.Printf("We match the first %d characters after %d\n", n, i)
	}
	return i
	//for i := 9917604604818; i < 100000000000000000; i += 748129615872 {
	//	runRegisters := map[string]int{"A": i, "B": registers["B"], "C": registers["C"]}
	//	expectedString := strings.Join(expected, ",")
	//	out := run(program, runRegisters, expectedString)
	//	if len(expected) == len(out) {
	//		return i
	//	}
	//}
}

func printState(registers map[string]int, program []int, pointer int) {
	for key, value := range registers {
		fmt.Printf("Register %s = %d\n", key, value)
	}
	fmt.Printf("instruction %d, operand %d\n\n", program[pointer], program[pointer+1])
}

func combo(operant int, registers map[string]int) int {
	switch operant % 8 {
	case 0, 1, 2, 3:
		return operant
	case 4:
		return registers["A"]
	case 5:
		return registers["B"]
	case 6:
		return registers["C"]
	case 7:
		log.Fatal("RESERVED!!!!")
	}
	log.Fatal("invalid operant")
	return 0
}

func loadProgram(input string) []int {
	parts := strings.Split(input, ":")
	program := strings.Split(strings.TrimSpace(parts[1]), ",")
	return lo.Map(program, func(stringInstruction string, _ int) int {
		instruction, _ := strconv.Atoi(stringInstruction)
		return instruction
	})
}

func loadInitialState(input string) map[string]int {
	state := make(map[string]int)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var key string
		var val int
		_, err := fmt.Sscanf(line, "Register %1s: %d", &key, &val)
		if err != nil {
			log.Fatal("error loading register " + line)
		}
		state[key] = val
	}
	return state
}
