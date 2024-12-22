package day21

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/samber/lo"
)

func Test_handleCodeForKeypad(t *testing.T) {
	tests := []struct {
		code string
		want []position
	}{
		{
			code: "029A",
			want: []position{{-1, 0}, {0, -1}, {1, -2}, {0, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			if got := handleCodeForKeypad(newNumericKeypad(), tt.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleCodeForKeypad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateShortestSequence(t *testing.T) {
	tests := []struct {
		code string
		want string
	}{
		{
			code: "029A",
			want: "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A",
		}, {
			code: "980A",
			want: "<v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A",
		}, {
			code: "179A",
			want: "<v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A",
		}, {
			code: "456A",
			want: "<v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A",
		}, {
			code: "379A",
			want: "<v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			if got := calculateShortestSequence(tt.code, 2); len(got) != len(tt.want) {
				t.Errorf("calculateLengthOfShortestSequence() = %v(%d), want %v(%d)", got, len(got), tt.want, len(tt.want))
			}
		})
	}
}

func TestCalculateComplexity(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `029A
980A
179A
456A
379A`,
			want: 126384,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateComplexity(tt.input, 2); got != tt.want {
				t.Errorf("CalculateComplexity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keypad_moveFingerAndReturnMovementOptions(t *testing.T) {
	tests := []struct {
		name        string
		pad         keypad
		destination string
		want        []string
	}{
		{
			name:        "example",
			pad:         newNumericKeypad(),
			destination: "7",
			want:        []string{"^^^<<A"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pad.moveFingerAndReturnMovementOptions(tt.destination); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveFingerAndReturnMovementOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateArrow(t *testing.T) {
	tests := []struct {
		name         string
		code         string
		numberOfBots int
		want         int
	}{
		{
			name:         "<A^A>^^AvvvA * 2",
			code:         "<A^A>^^AvvvA",
			numberOfBots: 2,
			want:         68,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateArrow(tt.code, tt.numberOfBots); got != tt.want {
				t.Errorf("calculateArrow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hypotheses(t *testing.T) {
	arrowMap := map[string]int{}
	for _, item := range []string{"v<<A", ">>^A", "v<A", "<A", "vA", "^<A", ">A", "^A", "^>A", "v>A", "A", ">^^A", "vvvA"} {
		arrowMap[item] = calculateArrow(item, 5)
	}

	want := calculateArrow("<A^A>^^AvvvA", 5)
	split := strings.Split("<A^A>^^AvvvA", "A")
	got := lo.Sum(lo.Map(split[:len(split)-1], func(item string, index int) int {
		i := arrowMap[item+"A"]
		fmt.Printf("Value for %s is %d\n", item, i)
		return i
	}))
	if got != want {
		t.Errorf("calculateArrow() = %v, want %v", got, want)
	}
}

//func Test_calculateShortestSequenceBetterPerformance(t *testing.T) {
//	type args struct {
//		code           string
//		numberOfBots   int
//		calculationMap map[string]int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		{
//			name: "example",
//			args: args{
//				code:           "029A",
//				numberOfBots:   1,
//				calculationMap: buildArrowMap(1),
//			},
//			want: 68,
//		},
//		{
//			name: "example",
//			args: args{
//				code:           "980A",
//				numberOfBots:   1,
//				calculationMap: buildArrowMap(1),
//			},
//			want: 60,
//		},
//		{
//			name: "example",
//			args: args{
//				code:           "179A",
//				numberOfBots:   1,
//				calculationMap: buildArrowMap(1),
//			},
//			want: 68,
//		},
//		{
//			name: "example",
//			args: args{
//				code:           "789A",
//				numberOfBots:   1,
//				calculationMap: buildArrowMap(1),
//			},
//			want: len(calculateShortestSequence("789A", 2)),
//		},
//		{
//			name: "example",
//			args: args{
//				code:           "968A",
//				numberOfBots:   1,
//				calculationMap: buildArrowMap(1),
//			},
//			want: len(calculateShortestSequence("968A", 2)),
//		},
//		{
//			name: "example",
//			args: args{
//				code:           "286A",
//				numberOfBots:   7,
//				calculationMap: buildArrowMap(5),
//			},
//			want: len(calculateShortestSequence("286A", 12)),
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := calculateShortestSequenceBetterPerformance(tt.args.code, tt.args.numberOfBots, tt.args.calculationMap); got != tt.want {
//				t.Errorf("calculateShortestSequenceBetterPerformance() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_CompareOptimisation(t *testing.T) {
	input := `789A
968A
286A
349A
170A`
	want := CalculateComplexity(input, 3)
	got := CalculateComplexityForX(input, 3)
	if got != want {
		t.Errorf("calculateShortestSequenceBetterPerformance() = %v, want %v", got, want)
	}
}
