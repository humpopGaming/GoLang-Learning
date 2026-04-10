# Challenge 12: Animal Sounds

## Objective

Build an "animal sounds" program to learn interfaces — Go's way of defining behaviour. You'll see how types implement interfaces implicitly and how the `Stringer` interface controls how values are printed.

## What You'll Learn

- An interface is a set of method signatures
- Types implement interfaces **implicitly** — no "implements" keyword needed
- Interface values hold a (value, type) pair under the hood
- Nil interface values cause panics; nil concrete values inside an interface don't
- The empty interface `interface{}` can hold any value
- `fmt.Stringer` — implement `String() string` to control how your type is printed

## Tour Reference — Read These First

1. [Interfaces](http://127.0.0.1:3999/tour/methods/9) — An interface type is a set of method signatures
2. [Interfaces are implemented implicitly](http://127.0.0.1:3999/tour/methods/10) — No "implements" keyword
3. [Interface values](http://127.0.0.1:3999/tour/methods/11) — Under the hood: (value, type) tuple
4. [Interface values with nil underlying values](http://127.0.0.1:3999/tour/methods/12) — Methods can handle nil receivers
5. [Nil interface values](http://127.0.0.1:3999/tour/methods/13) — Calling a method on a nil interface panics
6. [The empty interface](http://127.0.0.1:3999/tour/methods/14) — `interface{}` holds any type
7. [Stringers](http://127.0.0.1:3999/tour/methods/17) — `fmt.Stringer` lets you control how a type prints

## What to Build

An animal sounds program where different animal types all implement a `Speaker` interface, plus custom string representations via `Stringer`.

## File to Create

```
challenge12/main.go
```

## Requirements

1. Define an interface `Speaker` with one method: `Speak() string`
2. Define three struct types: `Dog` (with `Name string`), `Cat` (with `Name string`), `Fish` (with `Name string`)
3. Implement `Speak()` on each:
   - `Dog` returns `"Woof! Woof!"`
   - `Cat` returns `"Meow!"`
   - `Fish` returns `"..."`
4. Implement `String() string` (the `fmt.Stringer` interface) on each:
   - `Dog` returns `"Dog(<Name>)"`
   - `Cat` returns `"Cat(<Name>)"`
   - `Fish` returns `"Fish(<Name>)"`
5. Write `func introduce(s Speaker)` that prints: `"<s> says: <s.Speak()>"`
   (When you pass a `Dog` to `fmt.Printf("%v", s)`, it will use your `String()` method)
6. In `main()`:
   - Create a `Dog{"Rex"}`, `Cat{"Whiskers"}`, `Fish{"Nemo"}`
   - Put them all in a `[]Speaker` slice
   - Loop over the slice and call `introduce` for each
   - Demonstrate the empty interface: create a `[]interface{}` (or `[]any`) containing a string, an int, and one of the animals, then print each value's type and value using `fmt.Printf("(%v, %T)\n", ...)`

## Expected Output

```
Dog(Rex) says: Woof! Woof!
Cat(Whiskers) says: Meow!
Fish(Nemo) says: ...

Values of any type:
  ("hello", string)
  (42, int)
  (Dog(Rex), main.Dog)
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge12/main.go
```

## Hints

- Interface definition:
  ```go
  type Speaker interface {
      Speak() string
  }
  ```
- Implementing the interface — just write the method on the type:
  ```go
  func (d Dog) Speak() string { return "Woof! Woof!" }
  ```
  No "implements" keyword needed. If Dog has a `Speak() string` method, it _is_ a Speaker.
- Stringer:
  ```go
  func (d Dog) String() string { return fmt.Sprintf("Dog(%s)", d.Name) }
  ```
- `introduce` uses the `Speaker` interface, but `fmt.Printf("%v", s)` uses the `Stringer` interface — your types implement both!
- Empty interface: `[]interface{}{"hello", 42, rex}` — or in modern Go: `[]any{"hello", 42, rex}`
