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

func TestMakeCounter(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		calls    []int
		expected []int
	}{
		{
			name:     "Counter from 0",
			start:    0,
			calls:    []int{0, 0, 0},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Counter from 100",
			start:    100,
			calls:    []int{0, 0},
			expected: []int{101, 102},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := MakeCounter(tt.start)
			for i, _ := range tt.calls {
				got := counter()
				if got != tt.expected[i] {
					t.Errorf("counter() = %v, want %v", got, tt.expected[i])
				}
			}
		})
	}
}

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name     string
		factor   int
		inputs   []int
		expected []int
	}{
		{
			name:     "Doubler",
			factor:   2,
			inputs:   []int{1, 2, 5},
			expected: []int{2, 4, 10},
		},
		{
			name:     "Tripler",
			factor:   3,
			inputs:   []int{2, 4},
			expected: []int{6, 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mult := MakeMultiplier(tt.factor)
			for i, input := range tt.inputs {
				got := mult(input)
				if got != tt.expected[i] {
					t.Errorf("mult(%v) = %v, want %v", input, got, tt.expected[i])
				}
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	tests := []struct {
		name       string
		initial    int
		operations []struct {
			opType string
			value  int
		}
		expected int
	}{
		{
			name:    "Add and subtract",
			initial: 100,
			operations: []struct {
				opType string
				value  int
			}{
				{"add", 50},
				{"subtract", 30},
			},
			expected: 120,
		},
		{
			name:    "Multiple adds",
			initial: 0,
			operations: []struct {
				opType string
				value  int
			}{
				{"add", 10},
				{"add", 5},
				{"subtract", 3},
			},
			expected: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add, sub, get := MakeAccumulator(tt.initial)
			for _, op := range tt.operations {
				if op.opType == "add" {
					add(op.value)
				} else if op.opType == "subtract" {
					sub(op.value)
				}
			}
			got := get()
			if got != tt.expected {
				t.Errorf("accumulator final value = %v, want %v", got, tt.expected)
			}
		})
	}
}
