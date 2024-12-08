package day6

import "testing"

func TestCalculatedVisitedPositions(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "Example",
			input: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			want: 41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculatedVisitedPositions(tt.input); got != tt.want {
				t.Errorf("CalculatedVisitedPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberOfPossibleLoops(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "Example",
			input: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfPossibleLoops(tt.input); got != tt.want {
				t.Errorf("NumberOfPossibleLoops() = %v, want %v", got, tt.want)
			}
		})
	}
}
