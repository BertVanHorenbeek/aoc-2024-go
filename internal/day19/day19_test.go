package day19

import "testing"

func TestCountPossibleDesigns(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountPossibleDesigns(tt.input); got != tt.want {
				t.Errorf("CountPossibleDesigns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountDesignOptions(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`,
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountDesignOptions(tt.input); got != tt.want {
				t.Errorf("CountDesignOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
