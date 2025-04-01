package day13

import "testing"

func Test_calculateRequiredPresses(t *testing.T) {
	tests := []struct {
		name     string
		m        machine
		aPresses int
		bPresses int
	}{
		{name: "First example",
			m: machine{
				buttonA:     position{94, 34},
				buttonB:     position{22, 67},
				destination: position{8400, 5400},
			},
			aPresses: 80,
			bPresses: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calculateRequiredPresses(tt.m)
			if got != tt.aPresses {
				t.Errorf("calculateRequiredPresses() got = %v, want %v", got, tt.aPresses)
			}
			if got1 != tt.bPresses {
				t.Errorf("calculateRequiredPresses() got1 = %v, want %v", got1, tt.bPresses)
			}
		})
	}
}

func TestCalculateMinTokensToWinPossiblePrices(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "Example",
			input: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`,
			want: 480},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMinTokensToWinPossiblePrices(tt.input); got != tt.want {
				t.Errorf("CalculateMinTokensToWinPossiblePrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
