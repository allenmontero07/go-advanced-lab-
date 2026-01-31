package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {

	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{"0 factorial", 0, 1, false},         // 0!
		{"1 factorial", 1, 1, false},         // 1!
		{"5 factorial", 5, 120, false},       // positive
		{"7 factorial", 7, 5040, false},      // positive
		{"10 factorial", 10, 3628800, false}, // bigger
		{"negative", -4, 0, true},            // negative
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got, err := factorial(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {

	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{"prime 2", 2, true, false},       // required
		{"small prime 3", 3, true, false}, // small prime
		{"small prime 5", 5, true, false}, // small prime
		{"composite 4", 4, false, false},  // composite
		{"composite 9", 9, false, false},  // composite
		{"edge case 1", 1, false, true},   // edge
		{"edge case 0", 0, false, true},   // edge
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool
	}{
		{"positive base, positive exponent", 2, 3, 8, false},
		{"positive base, zero exponent", 5, 0, 1, false},
		{"negative base, even exponent", -2, 4, 16, false},
		{"negative base, odd exponent", -2, 3, -8, false},
		{"base zero", 0, 5, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Power() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}
