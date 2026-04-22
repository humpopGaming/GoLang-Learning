# Challenge 08: Contact Card

## Objective

Build a small program that creates and manipulates contact cards using structs — Go's way of grouping related data together.

## What You'll Learn

- Declaring a struct type with `type Name struct { ... }`
- Accessing struct fields with `.` (dot notation)
- Pointers to structs — `p.X` works the same as `(*p).X`
- Struct literals: creating structs by field name or by position
- The `&` prefix on a struct literal returns a pointer

## Tour Reference — Read These First

1. [Structs](tour/moretypes/02-structs.html) — A struct is a collection of fields
2. [Struct Fields](tour/moretypes/03-struct-fields.html) — Accessed using a dot
3. [Pointers to Structs](tour/moretypes/04-pointers-to-structs.html) — `p.X` instead of `(*p).X`
4. [Struct Literals](tour/moretypes/05-struct-literals.html) — Create by listing field values, or by `Name: value` syntax

## What to Build

A contact card system that:
1. Defines a `Contact` struct with name, email, and age
2. Creates contacts using different struct literal styles
3. Modifies a contact through a pointer
4. Has a function that prints a formatted contact card

## File to Create

```
challenge08/main.go
```

## Requirements

1. Define a `Contact` struct with fields: `Name string`, `Email string`, `Age int`
2. Write `func printCard(c Contact)` that prints a formatted card like:
   ```
   ┌──────────────────────────┐
   │ Name:  Alice Smith       │
   │ Email: alice@example.com │
   │ Age:   30                │
   └──────────────────────────┘
   ```
   (Don't worry about exact box alignment — just use `fmt.Println` with a simple format)
3. In `main()`:
   - Create `alice` using a **positional struct literal**: `Contact{"Alice Smith", "alice@example.com", 30}`
   - Create `bob` using a **named-field struct literal**: `Contact{Name: "Bob Jones", Email: "bob@example.com"}` (Age omitted — will be zero value)
   - Print both cards
   - Create a pointer to `alice` and **change her email** through the pointer
   - Print Alice's card again to show the change
   - Print Bob's age to show the zero value for int

## Expected Output

```
--- Alice's Card ---
Name:  Alice Smith
Email: alice@example.com
Age:   30

--- Bob's Card ---
Name:  Bob Jones
Email: bob@example.com
Age:   0

--- Alice's Updated Card ---
Name:  Alice Smith
Email: alice.smith@newjob.com
Age:   30

Bob's age (zero value): 0
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge08/main.go
```

## Hints

- Positional literal: `Contact{"Alice", "alice@example.com", 30}` — must supply all fields in order
- Named literal: `Contact{Name: "Bob", Email: "bob@example.com"}` — can omit fields (they get zero values)
- Pointer to struct: `p := &alice` then `p.Email = "new@email.com"` modifies `alice`
- `%d` for integers, `%s` for strings in `fmt.Printf`
- When you omit `Age` in Bob's struct, it will be `0` (the zero value for `int`)
