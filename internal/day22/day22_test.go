package day22

import "testing"

func Test_calculateSecretNumberIterations(t *testing.T) {
	tests := []struct {
		name   string
		number int
		n      int
		want   int
	}{
		{
			name:   "test 1",
			number: 123,
			n:      1,
			want:   15887950,
		}, {
			name:   "test 2",
			number: 123,
			n:      2,
			want:   16495136,
		}, {
			name:   "test 3",
			number: 2024,
			n:      2000,
			want:   8667524,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSecretNumberIterations(tt.number, tt.n); got != tt.want {
				t.Errorf("calculateSecretNumberIterations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mix(t *testing.T) {
	tests := []struct {
		name   string
		n1, n2 int
		want   int
	}{
		{name: "example",
			n1:   15,
			n2:   42,
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mix(tt.n1, tt.n2); got != tt.want {
				t.Errorf("mix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateSecretNumberSum(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `1
10
100
2024`,
			want: 37327623,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateSecretNumberSum(tt.input, 2000); got != tt.want {
				t.Errorf("CalculateSecretNumberSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateMostBananasYouCanEarn(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `1
2
3
2024`,
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMostBananasYouCanEarn(tt.input); got != tt.want {
				t.Errorf("CalculateMostBananasYouCanEarn() = %v, want %v", got, tt.want)
			}
		})
	}
}
