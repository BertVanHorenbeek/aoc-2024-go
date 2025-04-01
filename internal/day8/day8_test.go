package day8

import (
	"reflect"
	"testing"
)

func Test_findAntiNodes(t *testing.T) {
	tests := []struct {
		name     string
		antennas pair
		want     []position
	}{
		{
			name: "first example",
			antennas: pair{
				first:  position{x: 4, y: 3},
				second: position{x: 5, y: 5},
			},
			want: []position{
				{x: 3, y: 1},
				{x: 6, y: 7},
			},
		}, {
			name: "opposite example",
			antennas: pair{
				second: position{x: 4, y: 3},
				first:  position{x: 5, y: 5},
			},
			want: []position{
				{x: 6, y: 7},
				{x: 3, y: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAntiNodes(tt.antennas, 0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAntiNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountAntiNodes(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "Example",
			input: `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`,
			want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountAntiNodes(tt.input); got != tt.want {
				t.Errorf("CountAntiNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountAntiNodesUsingHarmonics(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "Example",
			input: `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`,
			want: 34},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountAntiNodesUsingHarmonics(tt.input); got != tt.want {
				t.Errorf("CountAntiNodesUsingHarmonics() = %v, want %v", got, tt.want)
			}
		})
	}
}
