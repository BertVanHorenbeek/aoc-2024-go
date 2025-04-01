package day20

import (
	"reflect"
	"testing"
)

func TestCountCheatsSavingAtLeastXps(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		minimalTimeSaved int
		want             int
	}{
		{
			name: "example",
			input: `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`,
			minimalTimeSaved: 64,
			want:             1,
		},
		{
			name: "example2",
			input: `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`,
			minimalTimeSaved: 65,
			want:             0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountCheatsSavingAtLeastXps(tt.input, tt.minimalTimeSaved); got != tt.want {
				t.Errorf("CountCheatsSavingAtLeastXps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountLongCheatsSavingAtLeastXps(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		minimalTimeSaved int
		want             int
	}{
		{
			name: "example",
			input: `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`,
			minimalTimeSaved: 76,
			want:             3,
		},
		{
			name: "example",
			input: `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`,
			minimalTimeSaved: 77,
			want:             0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountLongCheatsSavingAtLeastXps(tt.input, tt.minimalTimeSaved); got != tt.want {
				t.Errorf("CountLongCheatsSavingAtLeastXps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindCheatsOverXps(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		minimalTimeSaved int
		want             map[int]int
	}{
		{
			name: "example",
			input: `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`,
			minimalTimeSaved: 50,
			want: map[int]int{
				50: 32,
				52: 31,
				54: 29,
				56: 39,
				58: 25,
				60: 23,
				62: 20,
				64: 19,
				66: 12,
				68: 14,
				70: 12,
				72: 22,
				74: 4,
				76: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCheatsOverXps(tt.input, tt.minimalTimeSaved); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCheatsOverXps() = %v, want %v", got, tt.want)
			}
		})
	}
}
