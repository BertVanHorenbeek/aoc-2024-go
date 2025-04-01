package day14

import "testing"

func Test_calculateSafetyFactor(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`,
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateSafetyFactor(tt.input, 11, 7); got != tt.want {
				t.Errorf("calculateSafetyFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_FindChristmastree(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`,
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindChristmastree(tt.input, 11, 7); got != tt.want {
				t.Errorf("FindChristmastree() = %v, want %v", got, tt.want)
			}
		})
	}
}
