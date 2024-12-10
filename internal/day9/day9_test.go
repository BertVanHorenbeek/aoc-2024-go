package day9

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example",
			input: "2333133121414131402",
			want:  1928,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.input); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expand(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{
			name:  "Short example",
			input: "12345",
			want:  []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expand(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defragmentImproved(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Example",
			input: []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9},
			want:  []int{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defragmentImproved(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defragmentImproved() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "Example",
			input: "2333133121414131402",
			want:  2858,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePart2(tt.input); got != tt.want {
				t.Errorf("SolvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
