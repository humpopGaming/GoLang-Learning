# Challenge 04: FizzBuzz

## Objective

Implement the classic FizzBuzz problem to practise Go's `for` loop (all three forms) and `if`/`else` statements.

## What You'll Learn

- The standard `for` loop with init, condition, and post: `for i := 0; i < n; i++`
- Dropping the init and post to make a "while" loop: `for condition {}`
- Using `if`, `else if`, and `else`
- `if` with a short statement (declaring a variable in the `if` itself)
- The modulo operator `%` to check divisibility

## Tour Reference — Read These First

1. [For](http://127.0.0.1:3999/tour/flowcontrol/1) — Go's only loop construct; three components separated by semicolons
2. [For continued](http://127.0.0.1:3999/tour/flowcontrol/2) — Init and post statements are optional
3. [For is Go's "while"](http://127.0.0.1:3999/tour/flowcontrol/3) — Drop the semicolons for a while-style loop
4. [If](http://127.0.0.1:3999/tour/flowcontrol/5) — No parentheses around condition, braces are required
5. [If with a short statement](http://127.0.0.1:3999/tour/flowcontrol/6) — Execute a short statement before the condition
6. [If and else](http://127.0.0.1:3999/tour/flowcontrol/7) — Variables from `if` short statements are available in `else` blocks

## What to Build

A FizzBuzz program:
- Print numbers 1 through 30
- For multiples of 3, print "Fizz" instead
- For multiples of 5, print "Buzz" instead
- For multiples of both 3 and 5, print "FizzBuzz" instead

Then, as a second part, use a while-style loop to find the first number above 100 that is divisible by both 7 and 3.

## File to Create

```
challenge04/main.go
```

## Requirements

1. Write a `func fizzBuzz(n int)` that uses a standard **`for` loop** (with init/condition/post) to iterate from 1 to `n` and prints the appropriate FizzBuzz output for each number
2. Write a `func findSpecialNumber() int` that uses a **while-style `for` loop** (condition only, no semicolons) starting at 101 and returns the first number divisible by both 7 and 3. Use an **`if` with a short statement** to check divisibility (e.g. `if remainder := n % 21; remainder == 0`)
3. In `main()`:
   - Call `fizzBuzz(30)`
   - Print a blank line separator
   - Call `findSpecialNumber()` and print the result

## Expected Output

```
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
Fizz
22
23
Fizz
Buzz
26
Fizz
28
29
FizzBuzz

First number above 100 divisible by both 7 and 3: 105
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge04/main.go
```

## Hints

- Check "divisible by both" first (`n%15 == 0`), then by 3, then by 5, then the default
- A while-style loop looks like: `for n < 1000 { n++ }`
- An if-with-short-statement looks like: `if r := n % 21; r == 0 { ... }`
- The variable declared in the `if` short statement is scoped to the `if`/`else` block
- Use `fmt.Println` to print each value
