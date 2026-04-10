# Challenge 03: Currency Exchange

## Objective

Build a simple currency exchange calculator to practise variables, variable declarations, constants, and short variable declarations.

## What You'll Learn

- Declaring variables with `var` at package and function level
- Initializing variables with values (type becomes optional)
- Short variable declarations with `:=` (inside functions only)
- Constants with `const` — cannot use `:=` for constants
- Numeric constants and high-precision constant arithmetic

## Tour Reference — Read These First

1. [Variables](http://127.0.0.1:3999/tour/basics/8) — `var` declares variables; type comes last
2. [Variables with initializers](http://127.0.0.1:3999/tour/basics/9) — With an initializer, the type can be omitted
3. [Short variable declarations](http://127.0.0.1:3999/tour/basics/10) — `:=` for implicit-type declaration inside functions
4. [Constants](http://127.0.0.1:3999/tour/basics/15) — `const` keyword; cannot use `:=`
5. [Numeric constants](http://127.0.0.1:3999/tour/basics/16) — High-precision values that take the type their context needs

## What to Build

A currency exchange calculator that:
- Defines exchange rates as constants
- Uses different variable declaration styles throughout
- Converts USD amounts to EUR, GBP, and JPY
- Demonstrates that constants cannot be changed

## File to Create

```
challenge03/main.go
```

## Requirements

1. Declare a **factored `const` block** (parenthesized) with these exchange rates:
   - `USDToEUR = 0.92`
   - `USDToGBP = 0.79`
   - `USDToJPY = 154.50`
2. Declare a **package-level variable** (using `var`) called `bankName` of type `string` with value `"Go Exchange"`
3. In `main()`:
   - Use a **short variable declaration** (`:=`) for `amountUSD` set to `100.0`
   - Use a `var` **with initializer** (no explicit type) for `euros` set to `amountUSD * USDToEUR`
   - Use a `var` **with explicit type** for `pounds` as `float64` set to `amountUSD * USDToGBP`
   - Use `:=` for `yen` set to `amountUSD * USDToJPY`
   - Print the bank name
   - Print each converted amount on its own line
   - Print the sum of all three converted values as "Total value in foreign currencies"

## Expected Output

```
Welcome to Go Exchange
$100.00 = €92.00
$100.00 = £79.00
$100.00 = ¥15450.00
Total value in foreign currencies: 15621.00
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge03/main.go
```

## Hints

- A factored const block looks like:
  ```go
  const (
      A = 1
      B = 2
  )
  ```
- Package-level variables use `var name type = value` outside any function
- Short declarations `:=` only work inside functions
- Use `fmt.Printf("$%.2f = €%.2f\n", amountUSD, euros)` for two-decimal formatting
- Constants in Go are just replaced at compile time — they don't have a type until they're used, which lets them be high precision
