# Challenge 01: Hello Playground

## Objective

Write your first Go program that imports packages, calls functions from the standard library, and prints formatted output.

## What You'll Learn

- Every Go program starts in `package main`
- How to import packages using a factored import statement
- Exported names start with a capital letter (e.g. `math.Pi` not `math.pi`)
- Using `fmt.Println` and `fmt.Printf` to print output
- Using functions from the `math` and `math/rand` packages

## Tour Reference — Read These First

1. [Packages](https://go.dev/tour/basics/1) — Programs are made of packages; `main` is the entry point
2. [Imports](https://go.dev/tour/basics/2) — Use factored (parenthesized) import statements
3. [Exported names](https://go.dev/tour/basics/3) — Capital letter = exported = accessible from outside the package

## What to Build

A program that:
1. Prints a greeting message
2. Prints the value of Pi from the `math` package
3. Generates and prints a random number using `math/rand`
4. Prints the square root of a number using `math.Sqrt`

## File to Create

```
challenge01/main.go
```

## Requirements

1. Use `package main` and a `func main()`
2. Use a **factored import statement** (parenthesized) to import `fmt`, `math`, and `math/rand`
3. Print `Hello, Go Playground!` on the first line
4. Print `Pi is: ` followed by the value of `math.Pi`
5. Print `Random number: ` followed by `rand.Intn(100)` (a random int from 0-99)
6. Print `Square root of 144 is: ` followed by `math.Sqrt(144)`

## Expected Output

```
Hello, Go Playground!
Pi is: 3.141592653589793
Random number: 81
Square root of 144 is: 12
```

> Note: Your random number will differ — that's fine! The point is that it prints _some_ number.

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge01/main.go
```

## Hints

- Your import block should look like:
  ```go
  import (
      "fmt"
      "math"
      "math/rand"
  )
  ```
- Use `fmt.Println(...)` for simple output
- Remember: it's `math.Pi` (capital P), not `math.pi` — exported names start with uppercase
- `rand.Intn(100)` returns a random integer between 0 and 99
