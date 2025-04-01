package day18

import "testing"

func TestFindShortestPath(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		destination Position
		bytes       int
		want        int
	}{
		{
			name: "example",
			input: `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`,
			destination: Position{6, 6},
			bytes:       12,
			want:        22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindShortestPath(tt.input, tt.destination, tt.bytes); got != tt.want {
				t.Errorf("FindShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
