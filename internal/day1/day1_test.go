package day1

import "testing"

func Test_calculate(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{
			name:  "Example data",
			input: [][]int{{3, 4, 2, 1, 3, 3}, {4, 3, 5, 3, 9, 3}},
			want:  11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.input); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_similarityScore(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{
			name:  "Example data",
			input: [][]int{{3, 4, 2, 1, 3, 3}, {4, 3, 5, 3, 9, 3}},
			want:  31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := similarityScore(tt.input); got != tt.want {
				t.Errorf("similarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
