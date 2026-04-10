# Challenge 15: Generic Helpers

## Objective

Write generic utility functions that work on slices of any type, learning Go's type parameter syntax and the `comparable` constraint.

## What You'll Learn

- Generic functions with type parameters: `func Name[T constraint](params) returnType`
- The `comparable` constraint — allows `==` and `!=` on the type parameter
- The `any` constraint — allows any type at all
- How generics eliminate the need to write duplicate code for different types

## Tour Reference — Read These First

1. [Type parameters](http://127.0.0.1:3999/tour/generics/1) — Functions can have type parameters in `[brackets]`

## What to Build

Three generic helper functions:
1. `Contains` — checks if a slice contains a given element
2. `Filter` — returns a new slice with only elements that pass a test
3. `Map` — transforms each element of a slice using a function

## File to Create

```
challenge15/main.go
```

## Requirements

1. Write `func Contains[T comparable](slice []T, target T) bool` that returns `true` if `target` is found in `slice`
2. Write `func Filter[T any](slice []T, test func(T) bool) []T` that returns a new slice containing only elements where `test` returns `true`
3. Write `func Map[T any, U any](slice []T, transform func(T) U) []U` that returns a new slice where each element has been transformed
4. In `main()`, demonstrate each with **two different types**:
   - `Contains` with `[]int` and `[]string`
   - `Filter` with `[]int` (keep even numbers) and `[]string` (keep strings longer than 3 chars)
   - `Map` on `[]int` (double each number) and `[]string` (convert each to uppercase)

## Expected Output

```
Contains([]int{1,2,3,4,5}, 3): true
Contains([]int{1,2,3,4,5}, 6): false
Contains([]string{"go","rust","python"}, "go"): true
Contains([]string{"go","rust","python"}, "java"): false

Filter even numbers from [1 2 3 4 5 6 7 8 9 10]: [2 4 6 8 10]
Filter long strings from [Go Rust Python JS]: [Rust Python]

Map double [1 2 3 4 5]: [2 4 6 8 10]
Map uppercase [hello world go]: [HELLO WORLD GO]
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge15/main.go
```

## Hints

- Generic function syntax:
  ```go
  func Contains[T comparable](slice []T, target T) bool {
      for _, v := range slice {
          if v == target {
              return true
          }
      }
      return false
  }
  ```
- `comparable` is needed for `Contains` because it uses `==`
- `Filter` and `Map` only need `any` because they don't compare elements directly
- For `Map` with two type parameters: `func Map[T any, U any](...)` — `T` is input type, `U` is output type
- Use `strings.ToUpper` from the `"strings"` package for the uppercase transformation
- Even number test: `func(n int) bool { return n%2 == 0 }`
