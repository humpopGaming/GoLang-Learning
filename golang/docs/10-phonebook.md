# Challenge 10: Phonebook

## Objective

Build a phonebook using maps and explore function values and closures — two powerful Go features for storing and processing data.

## What You'll Learn

- Maps: key-value stores — `map[string]string`
- Creating maps with `make` and with map literals
- Adding, updating, reading, and deleting map entries
- Checking if a key exists with `value, ok := m[key]`
- Functions are values — they can be stored in variables and passed around
- Closures: functions that capture and "remember" variables from their outer scope

## Tour Reference — Read These First

1. [Maps](https://go.dev/tour/moretypes/19) — `map[KeyType]ValueType`; use `make` to create
2. [Map literals](https://go.dev/tour/moretypes/20) — Initialize a map inline with `map[K]V{ key: value, ... }`
3. [Map literals continued](https://go.dev/tour/moretypes/21) — Type name can be omitted in literal elements
4. [Mutating maps](https://go.dev/tour/moretypes/22) — Insert, update, delete, and check existence
5. [Function values](https://go.dev/tour/moretypes/24) — Functions can be passed around like any other value
6. [Function closures](https://go.dev/tour/moretypes/25) — A closure captures variables from its surrounding scope

## What to Build

A phonebook program that:
1. Creates a phonebook using a map
2. Performs lookups (including checking if a contact exists)
3. Adds, updates, and deletes entries
4. Uses a closure to create a "search logger" that tracks how many searches you've done

## File to Create

```
challenge10/main.go
```

## Requirements

1. Create a `phonebook` using a **map literal** with these initial entries:
   - `"Alice"` → `"555-1234"`
   - `"Bob"` → `"555-5678"`
   - `"Charlie"` → `"555-9012"`
2. Write `func lookup(pb map[string]string, name string)` that:
   - Uses the two-value form `number, ok := pb[name]` to check existence
   - Prints the number if found, or `"not found"` if not
3. In `main()`:
   - Look up `"Alice"` (should be found)
   - Look up `"Dave"` (should not be found)
   - **Add** `"Dave"` with number `"555-3456"`
   - Look up `"Dave"` again (now found)
   - **Update** `"Alice"` to `"555-0000"`
   - **Delete** `"Charlie"`
   - Print the final phonebook by iterating with `range`
4. Write `func makeSearchCounter() func(string) int` — a function that **returns a closure**. The closure takes a search term (string), increments an internal counter, and returns the total number of searches performed so far.
5. In `main()`, create a `searchCounter` from `makeSearchCounter()`, call it 3 times with any names, and print the count each time.

## Expected Output

```
Looking up Alice: 555-1234
Looking up Dave: not found
Added Dave.
Looking up Dave: 555-3456
Updated Alice.
Deleted Charlie.

Final phonebook:
  Alice: 555-0000
  Bob: 555-5678
  Dave: 555-3456

Search count after "Alice": 1
Search count after "Bob": 2
Search count after "unknown": 3
```

> Note: Map iteration order is random in Go, so your phonebook entries may print in a different order.

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge10/main.go
```

## Hints

- Map literal:
  ```go
  phonebook := map[string]string{
      "Alice": "555-1234",
      "Bob":   "555-5678",
  }
  ```
- Two-value lookup: `number, ok := phonebook["Alice"]` — `ok` is `true` if key exists
- Delete: `delete(phonebook, "Charlie")`
- Iterate: `for name, number := range phonebook`
- A closure that counts:
  ```go
  func makeSearchCounter() func(string) int {
      count := 0
      return func(name string) int {
          count++
          return count
      }
  }
  ```
  The returned function "remembers" `count` between calls — that's the closure!
