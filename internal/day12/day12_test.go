package day12

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "tiny example",
			input: `AAAA
BBCD
BBCC
EEEC`,
			want: 140},
		{name: "Small example",
			input: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			want: 772},
		{name: "Large example",
			input: `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
			want: 1930},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.input); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveWithDiscount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "tiny example",
			input: `AAAA
BBCD
BBCC
EEEC`,
			want: 80},
		{name: "Small example",
			input: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			want: 436},
		{name: "E example",
			input: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			want: 236},
		{name: "Large example",
			input: `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
			want: 1206},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveWithDiscount(tt.input); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
