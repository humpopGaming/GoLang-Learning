package main

import "fmt"

// POINTERS IN GO — Quick Reference
//
// Two operators: & (address-of) and * (dereference)
//
//   x := 10       x is an int, holds the value 10
//   &x            gives the memory address of x (e.g., 0xc000...)
//   *x            COMPILE ERROR — can't dereference a plain int
//
//   p := &x       p is a *int (pointer to int), holds x's address
//   *p            gives the value at that address (10)
//   *p = 42       changes the value at that address (x becomes 42)
//
// * is contextual:
//   - In a type (*int, *string):  means "pointer to"
//   - In an expression (*p):      means "get/set value at this address"
//
// Passing pointers to functions:
//   - func f(n int)   → receives a COPY of the value. Changes don't affect the original.
//   - func f(n *int)  → receives the ADDRESS. Use *n to read/write the original.
//   - Call with f(&x) to pass x's address.
//
// Trap: copying a pointer (temp := a) copies the ADDRESS, not the value.
//   Both point to the same memory, so changing *a also changes *temp.
//   To save a value before overwriting: temp := *a (copies the value itself).

func swap(a *int, b *int) {
	tempA := *a
	tempB := *b

	*a = tempB
	*b = tempA
}

func doubleValue(n *int) {
	*n = *n * 2
}

func failedDouble(n int) {
	n = n * 2
}

func main() {
	x, y := 10, 20

	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swap:  x = %d, y = %d\n", x, y)

	doubleValue(&x)
	fmt.Printf("After doubling x: %d\n", x)

	fmt.Printf("Before failedDouble: x = %d\n", x)
	failedDouble(x)
	fmt.Printf("After failedDouble:  x = %d (unchanged!)\n", x)
}
