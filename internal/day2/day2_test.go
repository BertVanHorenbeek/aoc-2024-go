package day2

import "testing"

func TestValidateReports(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Safe",
			input: "7 6 4 2 1",
			want:  1,
		},
		{
			name: "example",
			input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateReports(tt.input); got != tt.want {
				t.Errorf("ValidateReports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateReportsWithDamper(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Safe",
			input: "7 6 4 2 1",
			want:  1,
		}, {
			name:  "Safe after damper",
			input: "1 3 2 4 5",
			want:  1,
		},
		{
			name: "example",
			input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateReportsWithDamper(tt.input); got != tt.want {
				t.Errorf("ValidateReports() = %v, want %v", got, tt.want)
			}
		})
	}
}
