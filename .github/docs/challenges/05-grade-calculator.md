# Challenge 05: Grade Calculator

## Objective

Build a grade calculator that converts a numeric score to a letter grade using `switch` statements, practising all three forms of switch in Go.

## What You'll Learn

- Standard `switch` with a value to match against
- Switch with a short statement (like `if` short statements)
- Switch evaluation order — cases are checked top to bottom
- Conditionless switch (`switch { ... }`) as a cleaner if/else chain
- Go's switch doesn't need `break` — only the matching case runs

## Tour Reference — Read These First

1. [Switch](http://127.0.0.1:3999/tour/flowcontrol/9) — Go's switch; only runs the matching case (no fallthrough by default)
2. [Switch evaluation order](http://127.0.0.1:3999/tour/flowcontrol/10) — Cases are evaluated top to bottom
3. [Switch with no condition](http://127.0.0.1:3999/tour/flowcontrol/11) — `switch {}` is a clean way to write if/else chains

## What to Build

A program with two grading functions:
1. One that uses a **conditionless switch** to assign letter grades based on score ranges
2. One that uses a **standard switch** to describe what each letter grade means

## File to Create

```
challenge05/main.go
```

## Requirements

1. Write `func letterGrade(score int) string` that uses a **conditionless switch** (`switch { ... }`) to return:
   - `"A"` for scores 90–100
   - `"B"` for scores 80–89
   - `"C"` for scores 70–79
   - `"D"` for scores 60–69
   - `"F"` for scores below 60
2. Write `func gradeDescription(grade string) string` that uses a **standard switch** on `grade` to return:
   - `"A"` → `"Excellent"`
   - `"B"` → `"Good"`
   - `"C"` → `"Average"`
   - `"D"` → `"Below Average"`
   - `"F"` → `"Failing"`
   - default → `"Invalid grade"`
3. In `main()`, test with these scores: `95, 82, 74, 65, 43`
   - For each score print: `Score: XX → Grade: Y (Description)`

## Expected Output

```
Score: 95 → Grade: A (Excellent)
Score: 82 → Grade: B (Good)
Score: 74 → Grade: C (Average)
Score: 65 → Grade: D (Below Average)
Score: 43 → Grade: F (Failing)
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge05/main.go
```

## Hints

- A conditionless switch looks like:
  ```go
  switch {
  case score >= 90:
      return "A"
  case score >= 80:
      return "B"
  ...
  }
  ```
- A standard switch looks like:
  ```go
  switch grade {
  case "A":
      return "Excellent"
  ...
  }
  ```
- Order matters in the conditionless switch — check the highest range first
- The `→` character can be typed in Go source as `\u2192` or just use `->` if you prefer
- Use `fmt.Printf("Score: %d → Grade: %s (%s)\n", ...)` for formatting
