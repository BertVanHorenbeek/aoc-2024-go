package day10

import "testing"

func Test_countTrails(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example",
			input: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			want: 36,
		},
		{
			name: "Simple Example",
			input: `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`,
			want: 2,
		},
		{
			name: "Next Example",
			input: `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountTrails(tt.input); got != tt.want {
				t.Errorf("CountTrails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countTrails2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example",
			input: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			want: 81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountTrails2(tt.input); got != tt.want {
				t.Errorf("CountTrails2() = %v, want %v", got, tt.want)
			}
		})
	}
}
