package day17

import "testing"

func TestRunProgram(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "Example",
			input: `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`,
			want: "4,6,3,5,6,3,5,2,1,0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunProgram(tt.input); got != tt.want {
				t.Errorf("RunProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindStartToProduceItSelf(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`,
			want: 117440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindStartToProduceItSelf(tt.input); got != tt.want {
				t.Errorf("FindStartToProduceItSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
