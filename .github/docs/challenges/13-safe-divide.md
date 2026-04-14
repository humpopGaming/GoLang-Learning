# Challenge 13: Safe Divide

## Objective

Build a "safe divide" function that returns errors for invalid inputs, practising Go's error handling pattern, custom error types, type assertions, and type switches.

## What You'll Learn

- The `error` interface: `type error interface { Error() string }`
- Creating custom error types that implement the `error` interface
- Returning `(result, error)` from functions — Go's standard error pattern
- Checking for errors: `if err != nil { ... }`
- Type assertions: `t := i.(T)` and `t, ok := i.(T)` (safe form)
- Type switches: `switch v := i.(type) { case T: ... }`

## Tour Reference — Read These First

1. [Type assertions](https://go.dev/tour/methods/15) — Extract concrete type from an interface value
2. [Type switches](https://go.dev/tour/methods/16) — Switch on the type of an interface value
3. [Errors](https://go.dev/tour/methods/19) — The `error` interface and how Go handles errors
4. [Exercise: Errors](https://go.dev/tour/methods/20) — Example of custom error types

## What to Build

A division function that safely handles division by zero and negative inputs, with custom error types for each case. Then use type assertions and type switches to react differently to each error type.

## File to Create

```
challenge13/main.go
```

## Requirements

1. Define `type DivByZeroError struct{}` with an `Error() string` method returning `"cannot divide by zero"`
2. Define `type NegativeInputError struct{ Value float64 }` with an `Error() string` method returning `"negative input not allowed: <value>"`
3. Write `func safeDivide(a, b float64) (float64, error)` that:
   - Returns `DivByZeroError{}` if `b == 0`
   - Returns `NegativeInputError{a}` if `a < 0`
   - Returns `NegativeInputError{b}` if `b < 0`
   - Otherwise returns `a / b, nil`
4. Write `func describeError(err error)` that uses a **type switch** on `err` to print:
   - For `DivByZeroError`: `"Error type: DivByZeroError — <message>"`
   - For `NegativeInputError`: `"Error type: NegativeInputError — <message> (value was: <Value>)"`
   - Default: `"Unknown error: <message>"`
5. In `main()`, test with these calls:
   - `safeDivide(10, 3)` — should succeed
   - `safeDivide(10, 0)` — should return DivByZeroError
   - `safeDivide(-5, 2)` — should return NegativeInputError
   - `safeDivide(10, -3)` — should return NegativeInputError
   - For each, print the result or call `describeError`

## Expected Output

```
10 / 3 = 3.3333333333333335

10 / 0:
  Error type: DivByZeroError — cannot divide by zero

-5 / 2:
  Error type: NegativeInputError — negative input not allowed: -5 (value was: -5)

10 / -3:
  Error type: NegativeInputError — negative input not allowed: -3 (value was: -3)
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge13/main.go
```

## Hints

- Custom error types just need an `Error() string` method:
  ```go
  func (e DivByZeroError) Error() string { return "cannot divide by zero" }
  ```
- Return errors like: `return 0, DivByZeroError{}`
- Check errors like: `result, err := safeDivide(10, 3); if err != nil { ... }`
- Type switch:
  ```go
  switch e := err.(type) {
  case DivByZeroError:
      fmt.Println("division by zero!")
  case NegativeInputError:
      fmt.Printf("negative value: %v\n", e.Value)
  }
  ```
- When implementing `Error()` for `NegativeInputError`, use `fmt.Sprintf` — but be careful not to call `fmt.Sprint(e)` directly (this causes infinite recursion). Use `fmt.Sprintf("negative input not allowed: %g", float64(e.Value))`
