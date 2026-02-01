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

			got, err := Factorial(tt.input)

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

func TestApply(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	tests := []struct {
		name string
		op   func(int) int
		want []int
	}{
		{"square", func(x int) int { return x * x }, []int{1, 4, 9, 16, 25}},
		{"double", func(x int) int { return x * 2 }, []int{2, 4, 6, 8, 10}},
		{"negate", func(x int) int { return -x }, []int{-1, -2, -3, -4, -5}},
		{"add 10", func(x int) int { return x + 10 }, []int{11, 12, 13, 14, 15}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(nums, tt.op)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Apply() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 12, 15}

	tests := []struct {
		name string
		pred func(int) bool
		want []int
	}{
		{"even numbers", func(x int) bool { return x%2 == 0 }, []int{2, 4, 6, 12}},
		{"greater than 10", func(x int) bool { return x > 10 }, []int{12, 15}},
		{"divisible by 3", func(x int) bool { return x%3 == 0 }, []int{3, 6, 12, 15}},
		{"odd numbers", func(x int) bool { return x%2 != 0 }, []int{1, 3, 5, 15}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(nums, tt.pred)
			if len(got) != len(tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Filter() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	tests := []struct {
		name string
		init int
		op   func(int, int) int
		want int
	}{
		{"sum", 0, func(acc, cur int) int { return acc + cur }, 15},
		{"product", 1, func(acc, cur int) int { return acc * cur }, 120},
		{"max", nums[0], func(acc, cur int) int {
			if cur > acc {
				return cur
			}
			return acc
		}, 5},
		{"min", nums[0], func(acc, cur int) int {
			if cur < acc {
				return cur
			}
			return acc
		}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(nums, tt.init, tt.op)
			if got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	tests := []struct {
		name  string
		f     func(int) int
		g     func(int) int
		input int
		want  int
	}{
		{"double then addTwo", func(x int) int { return x + 2 }, func(x int) int { return x * 2 }, 5, 12},
		{"addTwo then double", func(x int) int { return x * 2 }, func(x int) int { return x + 2 }, 5, 14},
		{"square then negate", func(x int) int { return -x }, func(x int) int { return x * x }, 3, -9},
		{"negate then square", func(x int) int { return x * x }, func(x int) int { return -x }, 3, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compose(tt.f, tt.g)(tt.input)
			if got != tt.want {
				t.Errorf("Compose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwapValues(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		wantA int
		wantB int
	}{
		{"positive numbers", 5, 10, 10, 5},
		{"zero and positive", 0, 7, 7, 0},
		{"negative numbers", -3, -8, -8, -3},
		{"mixed signs", -4, 6, 6, -4},
		{"same values", 9, 9, 9, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotA, gotB := SwapValues(tt.a, tt.b)

			if gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf(
					"SwapValues(%d, %d) = (%d, %d), want (%d, %d)",
					tt.a, tt.b,
					gotA, gotB,
					tt.wantA, tt.wantB,
				)
			}
		})
	}
}

func TestSwapPointers(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		wantA int
		wantB int
	}{
		{"positive numbers", 5, 10, 10, 5},
		{"zero and positive", 0, 7, 7, 0},
		{"negative numbers", -3, -8, -8, -3},
		{"mixed signs", -4, 6, 6, -4},
		{"same values", 9, 9, 9, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			a := tt.a
			b := tt.b

			SwapPointers(&a, &b)

			if a != tt.wantA || b != tt.wantB {
				t.Errorf(
					"SwapPointers(%d, %d) = (%d, %d), want (%d, %d)",
					tt.a, tt.b,
					a, b,
					tt.wantA, tt.wantB,
				)
			}
		})
	}
}
