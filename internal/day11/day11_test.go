package day11

import (
	"testing"
)

func TestCountStonesAfterBlinks(t *testing.T) {
	tests := []struct {
		name   string
		blinks int
		want   int
	}{
		{"Blink 1", 1, 3},
		{"Blink 2", 2, 4},
		{"Blink 3", 3, 5},
		{"Blink 4", 4, 9},
		{"Blink 5", 5, 13},
		{"Blink 6", 6, 22},
		{"Blink 25", 25, 55312},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountStonesAfterBlinks("125 17", tt.blinks); got != tt.want {
				t.Errorf("CountStonesAfterBlinks() = %v, want %v", got, tt.want)
			}
		})
	}
}
