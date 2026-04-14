# Challenge 09: Shopping List

## Objective

Build a shopping list manager using arrays, slices, `append`, `make`, `range`, and slice operations — the key building blocks for working with collections in Go.

## What You'll Learn

- Arrays have a fixed size: `[n]T`
- Slices are dynamically-sized views into arrays: `[]T`
- Slices reference underlying arrays — changes to a slice affect the array
- Slice literals, slice defaults (omitting low/high bounds)
- `len()` and `cap()` — length vs capacity
- `nil` slices have length and capacity 0
- Creating slices with `make([]T, len, cap)`
- Growing slices with `append(slice, elements...)`
- Iterating with `range` — gives index and value
- Skipping index or value with `_`

## Tour Reference — Read These First

1. [Arrays](https://go.dev/tour/moretypes/6) — Fixed-size: `var a [10]int`
2. [Slices](https://go.dev/tour/moretypes/7) — Dynamically-sized views: `a[1:4]`
3. [Slices are like references to arrays](https://go.dev/tour/moretypes/8) — Changing a slice changes the underlying array
4. [Slice literals](https://go.dev/tour/moretypes/9) — `[]int{1, 2, 3}` creates a slice
5. [Slice defaults](https://go.dev/tour/moretypes/10) — `a[:5]`, `a[2:]`, `a[:]`
6. [Slice length and capacity](https://go.dev/tour/moretypes/11) — `len(s)` and `cap(s)`
7. [Nil slices](https://go.dev/tour/moretypes/12) — Zero value of a slice is `nil`
8. [Creating a slice with make](https://go.dev/tour/moretypes/13) — `make([]int, 5)` or `make([]int, 0, 5)`
9. [Slices of slices](https://go.dev/tour/moretypes/14) — Slices can contain slices
10. [Appending to a slice](https://go.dev/tour/moretypes/15) — `append(s, 1, 2, 3)` grows the slice
11. [Range](https://go.dev/tour/moretypes/16) — `for i, v := range slice`
12. [Range continued](https://go.dev/tour/moretypes/17) — Skip index with `_`, or omit value

## What to Build

A shopping list program that:
1. Starts with a fixed list of "aisle items" in an array
2. Creates a flexible shopping list as a slice
3. Adds items, prints them, shows how slices grow

## File to Create

```
challenge09/main.go
```

## Requirements

1. Declare an **array** of 5 strings: `[5]string{"Milk", "Eggs", "Bread", "Butter", "Cheese"}`
2. Create a **slice** from the array containing just the first 3 items (use slicing `array[0:3]`)
3. Print the slice and its len/cap
4. Declare a `nil` var `var shoppingList []string` — print it and check if it's nil
5. Use `append` to add `"Apples"`, `"Bananas"`, and `"Oranges"` to `shoppingList`
6. Print the shopping list using `range`, showing the index and item
7. Use `append` to add two more items at once: `"Yogurt"` and `"Juice"`
8. Print the final list with **only the item names** (skip the index using `_`)
9. Print the final `len` and `cap` of `shoppingList`

## Expected Output

```
Store aisle items: [Milk Eggs Bread Butter Cheese]
First 3 aisle items: [Milk Eggs Bread], len=3, cap=5

Shopping list is nil: true

Shopping list:
  0: Apples
  1: Bananas
  2: Oranges

Final shopping list:
  - Apples
  - Bananas
  - Oranges
  - Yogurt
  - Juice

List length: 5, capacity: 6
```

> Note: the final capacity may be 6 or 8 depending on Go's internal growth strategy — any value >= 5 is correct.

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge09/main.go
```

## Hints

- Array declaration: `aisle := [5]string{"Milk", "Eggs", "Bread", "Butter", "Cheese"}`
- Slicing: `firstThree := aisle[0:3]` or `aisle[:3]`
- A `var s []string` with no initialization is `nil`
- `s = append(s, "Apples", "Bananas", "Oranges")` adds multiple items
- `for i, item := range shoppingList` gives index and value
- `for _, item := range shoppingList` skips the index
- `fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))`
