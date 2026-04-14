# Challenge 02: Temperature Converter

## Objective

Build a program with multiple functions that convert temperatures between Celsius and Fahrenheit, practising function declarations, parameters, return values, and type conversions.

## What You'll Learn

- Declaring functions with typed parameters
- Shortening parameter lists when consecutive params share a type (`x, y int`)
- Returning multiple values from a function
- Named return values and naked returns
- Go's basic types (`float64`, `int`, `string`)
- Zero values for different types
- Explicit type conversions (`float64(x)`, `int(f)`)
- Type inference with `:=`

## Tour Reference — Read These First

1. [Functions](https://go.dev/tour/basics/4) — Functions take typed parameters; type comes after the name
2. [Functions continued](https://go.dev/tour/basics/5) — Consecutive params of the same type can be shortened
3. [Multiple results](https://go.dev/tour/basics/6) — A function can return multiple values
4. [Named return values](https://go.dev/tour/basics/7) — Return values can be named; naked return returns them
5. [Basic types](https://go.dev/tour/basics/11) — Go's built-in types: `int`, `float64`, `bool`, `string`, etc.
6. [Zero values](https://go.dev/tour/basics/12) — Variables without initial values get zero values
7. [Type conversions](https://go.dev/tour/basics/13) — Use `T(v)` to convert value `v` to type `T`
8. [Type inference](https://go.dev/tour/basics/14) — `:=` infers the type from the right-hand side

## What to Build

A temperature converter that has:
- A function `celsiusToFahrenheit` that takes a `float64` and returns a `float64`
- A function `fahrenheitToCelsius` that takes a `float64` and returns a `float64`
- A function `boilingAndFreezing` that returns **two** `float64` values (boiling point and freezing point in Fahrenheit)
- Demonstration of type conversion by converting a `float64` result to `int` for display

## File to Create

```
challenge02/main.go
```

## Requirements

1. Write `func celsiusToFahrenheit(c float64) float64` that returns `c*9/5 + 32`
2. Write `func fahrenheitToCelsius(f float64) float64` that returns `(f - 32) * 5 / 9`
3. Write `func boilingAndFreezing() (float64, float64)` that returns the boiling (100°C) and freezing (0°C) points converted to Fahrenheit — this function must call `celsiusToFahrenheit` twice and return both results
4. In `main()`:
   - Convert 100°C to °F and print the result
   - Convert 72°F to °C and print the result
   - Call `boilingAndFreezing()` and print both returned values
   - Show type conversion: convert the Celsius result of 72°F to an `int` and print it
   - Declare a `float64` variable without initializing it and print its zero value

## Expected Output

```
100°C = 212°F
72°F = 22.22222222222222°C
Boiling: 212°F, Freezing: 32°F
72°F as whole number Celsius: 22
Zero value of float64: 0
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge02/main.go
```

## Hints

- The formula is: `F = C × 9/5 + 32` and `C = (F - 32) × 5/9`
- To return two values: `func name() (float64, float64) { return a, b }`
- To receive two values: `boil, freeze := boilingAndFreezing()`
- Type conversion: `int(22.7)` gives `22` (truncates, doesn't round)
- `fmt.Printf` with `%v` or `%g` is useful for formatting numbers
- A `var f float64` without `= value` will be `0` (the zero value)
