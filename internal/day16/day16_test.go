package day16

import (
	"reflect"
	"testing"
)

func TestFindBestPossibleScore(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`,
			want: 7036,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindBestPossibleScore(tt.input); got != tt.want {
				t.Errorf("FindBestPossibleScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reindeer_rotateR(t *testing.T) {
	tests := []struct {
		name        string
		orientation position
		want        position
	}{
		{
			name:        "N->E",
			orientation: position{0, -1},
			want:        position{1, 0},
		},
		{
			name:        "E->S",
			orientation: position{1, 0},
			want:        position{0, 1},
		},
		{
			name:        "S->W",
			orientation: position{0, 1},
			want:        position{-1, 0},
		},
		{
			name:        "W->N",
			orientation: position{-1, 0},
			want:        position{0, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := reindeer{
				p: position{},
				o: tt.orientation,
			}
			if got := r.rotateR(); !reflect.DeepEqual(got.o, tt.want) {
				t.Errorf("rotateR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reindeer_rotateL(t *testing.T) {
	tests := []struct {
		name        string
		orientation position
		want        position
	}{
		{
			name:        "E->N",
			orientation: position{1, 0},
			want:        position{0, -1},
		},
		{
			name:        "S->E",
			orientation: position{0, 1},
			want:        position{1, 0},
		},
		{
			name:        "W->S",
			orientation: position{-1, 0},
			want:        position{0, 1},
		},
		{
			name:        "N->W",
			orientation: position{0, -1},
			want:        position{-1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := reindeer{
				p: position{},
				o: tt.orientation,
			}
			if got := r.rotateL(); !reflect.DeepEqual(got.o, tt.want) {
				t.Errorf("rotateL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindBestSeats(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "first example",
			input: `###############
#.......#....E#
#.#.###.#.###O#
#.....#.#...#O#
#.###.#####.#O#
#.#.#.......#O#
#.#.#####.###O#
#..OOOOOOOOO#O#
###O#O#####O#O#
#OOO#O....#O#O#
#O#O#O###.#O#O#
#OOOOO#...#O#O#
#O###.#.#.#O#O#
#S..#.....#OOO#
###############`,
			want: 45,
		},
		{
			name: "second example",
			input: `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#O#
#.#.#.#...#...#O#
#.#.#.#.###.#.#O#
#OOO#.#.#.....#O#
#O#O#.#.#.#####O#
#O#O..#.#.#OOOOO#
#O#O#####.#O###O#
#O#O#..OOOOO#OOO#
#O#O###O#####O###
#O#O#OOO#..OOO#.#
#O#O#O#####O###.#
#O#O#OOOOOOO..#.#
#O#O#O#########.#
#S#OOO..........#
#################`,
			want: 64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindBestSeats(tt.input); got != tt.want {
				t.Errorf("FindBestSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}
