# Challenge 07: Pointer Swap

## Objective

Write a program that swaps two variables using pointers, demonstrating how pointers let you modify values from outside the scope where they were declared.

## What You'll Learn

- A pointer holds the memory address of a value
- `&` gives you the address of a variable (creates a pointer)
- `*` dereferences a pointer (reads/writes the value at that address)
- The zero value of a pointer is `nil`
- Go has no pointer arithmetic (unlike C)

## Tour Reference — Read These First

1. [Pointers](https://go.dev/tour/moretypes/1) — `*T` is a pointer to a `T` value; `&` creates a pointer; `*` dereferences it

## What to Build

A program that:
1. Swaps two integers using a function that takes pointers
2. Doubles a number using a function that takes a pointer
3. Shows that without pointers, changes don't affect the original variable

## File to Create

```
challenge07/main.go
```

## Requirements

1. Write `func swap(a *int, b *int)` that swaps the values of the two integers pointed to by `a` and `b`
2. Write `func doubleValue(n *int)` that doubles the value pointed to by `n`
3. Write `func failedDouble(n int)` that tries to double `n` but takes it by value (to show it doesn't work)
4. In `main()`:
   - Declare `x := 10` and `y := 20`
   - Print `x` and `y` before the swap
   - Call `swap(&x, &y)` and print `x` and `y` after
   - Call `doubleValue(&x)` and print `x` after
   - Print the value of `x` before calling `failedDouble(x)`, then print `x` after to show it didn't change

## Expected Output

```
Before swap: x = 10, y = 20
After swap:  x = 20, y = 10
After doubling x: 40
Before failedDouble: x = 40
After failedDouble:  x = 40 (unchanged!)
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge07/main.go
```

## Hints

- To swap via pointers:
  ```go
  temp := *a
  *a = *b
  *b = temp
  ```
- `*n = *n * 2` doubles the value that `n` points to
- `failedDouble` receives a _copy_ of the value — changes to `n` inside the function don't affect the original
- This is one of the most important concepts in Go: understanding when you're working with a value vs. a pointer to a value
