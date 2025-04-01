package day7

import (
	"testing"

	"github.com/samber/lo"
)

func TestTotalCalibrationResult(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example",
			input: `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`,
			want: 3749,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalCalibrationUsingReduce(tt.input); got != tt.want {
				t.Errorf("TotalCalibrationResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalCalibrationResultWithConcat(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example",
			input: `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`,
			want: 11387,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalCalibrationResultWithConcat(tt.input); got != tt.want {
				t.Errorf("TotalCalibrationResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listPossibleSolutions(t *testing.T) {
	tests := []struct {
		name     string
		operants []int
		want     []int
	}{
		{
			name:     "190: 10 19",
			operants: []int{10, 19},
			want:     []int{190, 29},
		},
		{
			name:     "3267: 81 40 27",
			operants: []int{81, 40, 27},
			want: []int{
				(81 + 40) + 27,
				(81 + 40) * 27,
				(81 * 40) + 27,
				(81 * 40) * 27},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listPossibleSolutions(tt.operants, operations)
			diff1, diff2 := lo.Difference(got, tt.want)
			if len(diff1) > 0 || len(diff2) > 0 {
				t.Errorf("listPossibleSolutions() = %v, want %v", got, tt.want)
			}
		})
	}
}
