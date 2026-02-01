package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func Factorial(n int) (int, error) {

	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}

	result := 1

	for i := 1; i <= n; i++ {
		result = result * i
	}

	return result, nil
}

func IsPrime(n int) (bool, error) {

	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}

	limit := int(math.Sqrt(float64(n)))

	for i := 2; i <= limit; i++ {

		if n%i == 0 {

			return false, nil

		}
	}

	return true, nil

}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}

	return result, nil
}

func MakeCounter(start int) func() int {
	counter := start
	return func() int {
		counter++
		return counter
	}

}

func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	total := initial

	add = func(x int) {
		total += x
	}

	subtract = func(x int) {
		total -= x
	}

	get = func() int {
		return total
	}

	return add, subtract, get
}

func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = operation(v)
	}
	return result
}

func Filter(nums []int, predicate func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
	acc := initial
	for _, v := range nums {
		acc = operation(acc, v)
	}
	return acc
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func ExploreProcess() {
	fmt.Println("=== Process Information ===")

	pid := os.Getpid()
	fmt.Printf("Current Process ID: %d\n", pid)

	ppid := os.Getppid()
	fmt.Printf("Parent Process ID: %d\n", ppid)

	data := []int{1, 2, 3, 4, 5}

	fmt.Printf("Memory address of slice: %p\n", &data)

	fmt.Printf("Memory address of first element: %p\n", &data[0])

	fmt.Println("Note: Other processes cannot access these memory addresses due to process isolation.")
	fmt.Println("Process IDs identify each running program, and each process has its own memory space.")
	fmt.Println("The slice header stores metadata (length, capacity, pointer).")
	fmt.Println("The element addresses are actual data locations in memory.")
}

func DoubleValue(x int) int {
	x = x * 2
	return x
}

// Comment: This function works on a copy of x, so the original variable stays unchanged

func DoublePointer(x *int) {
	*x = *x * 2
}

// Comment: This modifies the original variable because we are using a pointer

func CreateOnStack() int {
	x := 42
	return x
}

// Comment: This variable stays on the stack, no escape

func CreateOnHeap() *int {
	x := 42
	return &x
}

// Comment: This variable "escapes" to the heap because we are returning a pointer to it

func SwapValues(a, b int) (int, int) {
	return b, a
}

// Comment: Original variables in main remain unchanged

func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}

// Comment: Original variables are modified because pointers are used

func AnalyzeEscape() {
	stackVar := CreateOnStack()
	heapVar := CreateOnHeap()

	fmt.Println("Stack variable value:", stackVar)
	fmt.Println("Heap variable value:", *heapVar)
}

/*
Escape Analysis Comments:

- stackVar does NOT escape to the heap. It's created and used inside the function.
- heapVar escapes to the heap because we return a pointer to it, which could live beyond the function's lifetime.
- "Escapes to heap" means Go allocates the variable on the heap instead of stack, so it persists beyond the function call.
*/
func main() {

	// ===============================
	// Part 4: Process Information
	// ===============================
	fmt.Println("=== Process Information ===")
	ExploreProcess()

	// ===============================
	// Part 1: Math Operations
	// ===============================
	fmt.Println("\n=== Math Operations ===")

	// Factorial
	f0, _ := Factorial(0)
	f5, _ := Factorial(5)
	f10, _ := Factorial(10)

	fmt.Println("Factorial(0) =", f0)
	fmt.Println("Factorial(5) =", f5)
	fmt.Println("Factorial(10) =", f10)

	// IsPrime
	p17, _ := IsPrime(17)
	p20, _ := IsPrime(20)
	p25, _ := IsPrime(25)

	fmt.Println("IsPrime(17) =", p17)
	fmt.Println("IsPrime(20) =", p20)
	fmt.Println("IsPrime(25) =", p25)

	// Power
	pow1, _ := Power(2, 8)
	pow2, _ := Power(5, 3)

	fmt.Println("Power(2,8) =", pow1)
	fmt.Println("Power(5,3) =", pow2)

	// ===============================
	// Part 2: Closures
	// ===============================
	fmt.Println("\n=== Closure Demonstration ===")

	counter1 := MakeCounter(0)
	counter2 := MakeCounter(100)

	fmt.Println("Counter1:", counter1())
	fmt.Println("Counter1:", counter1())
	fmt.Println("Counter2:", counter2())
	fmt.Println("Counter1:", counter1())

	doubler := MakeMultiplier(2)
	tripler := MakeMultiplier(3)

	fmt.Println("Doubler(5):", doubler(5))
	fmt.Println("Tripler(5):", tripler(5))

	add, sub, get := MakeAccumulator(100)

	add(50)
	fmt.Println("Accumulator:", get())

	sub(30)
	fmt.Println("Accumulator:", get())

	// ===============================
	// Part 3: Higher-Order Functions
	// ===============================
	fmt.Println("\n=== Higher-Order Functions ===")

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Original:", nums)

	squared := Apply(nums, func(x int) int {
		return x * x
	})
	fmt.Println("Squared:", squared)

	evens := Filter(nums, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Evens:", evens)

	sum := Reduce(nums, 0, func(a, b int) int {
		return a + b
	})
	fmt.Println("Sum:", sum)

	double := func(x int) int {
		return x * 2
	}

	addTen := func(x int) int {
		return x + 10
	}

	composed := Compose(addTen, double)

	fmt.Println("Compose (double then add 10):", composed(5))

	// ===============================
	// Part 5: Pointer Demonstration
	// ===============================
	fmt.Println("\n=== Pointer Demonstration ===")

	num := 10

	fmt.Println("Before DoubleValue:", num)
	DoubleValue(num)
	fmt.Println("After DoubleValue:", num, "(unchanged)")

	fmt.Println("Before DoublePointer:", num)
	DoublePointer(&num)
	fmt.Println("After DoublePointer:", num, "(changed)")

	a := 5
	b := 10

	fmt.Println("Before SwapValues: a =", a, "b =", b)
	x, y := SwapValues(a, b)
	fmt.Println("After SwapValues: a =", a, "b =", b, "(originals unchanged)")
	fmt.Println("Returned:", x, y)

	fmt.Println("Before SwapPointers: a =", a, "b =", b)
	SwapPointers(&a, &b)
	fmt.Println("After SwapPointers: a =", a, "b =", b)

	// ===============================
	// Escape Analysis Demo
	// ===============================
	fmt.Println("\n=== Escape Analysis ===")
	AnalyzeEscape()

	fmt.Println("\n=== Program Finished ===")
}
