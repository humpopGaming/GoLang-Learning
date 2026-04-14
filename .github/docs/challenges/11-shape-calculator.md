# Challenge 11: Shape Calculator

## Objective

Build a shape calculator with methods on struct types, practising both value receivers and pointer receivers — and understanding when to use each.

## What You'll Learn

- Methods: functions with a receiver argument — `func (v Vertex) Abs() float64`
- Methods are just functions with a special receiver parameter
- You can define methods on non-struct types too (e.g. `type MyFloat float64`)
- Pointer receivers (`*T`) can modify the receiver; value receivers cannot
- Go auto-dereferences: `v.Scale(5)` works even when Scale has a `*Vertex` receiver
- Choosing between value and pointer receivers

## Tour Reference — Read These First

1. [Methods](https://go.dev/tour/methods/1) — A method has a receiver between `func` and the method name
2. [Methods are functions](https://go.dev/tour/methods/2) — A method is just a function with a receiver argument
3. [Methods continued](https://go.dev/tour/methods/3) — Methods on non-struct types
4. [Pointer receivers](https://go.dev/tour/methods/4) — Use `*T` to modify the receiver
5. [Pointers and functions](https://go.dev/tour/methods/5) — Compare methods vs equivalent plain functions
6. [Methods and pointer indirection](https://go.dev/tour/methods/6) — `v.Scale(5)` auto-becomes `(&v).Scale(5)`
7. [Methods and pointer indirection (2)](https://go.dev/tour/methods/7) — `p.Abs()` auto-becomes `(*p).Abs()`
8. [Choosing a value or pointer receiver](https://go.dev/tour/methods/8) — Use pointer receivers to modify or avoid copying

## What to Build

A shape calculator with two shapes: `Rectangle` and `Circle`. Each has methods for area and perimeter. `Rectangle` also has a `Scale` method that uses a pointer receiver to modify the shape in-place.

## File to Create

```
challenge11/main.go
```

## Requirements

1. Define `type Rectangle struct` with `Width, Height float64`
2. Define `type Circle struct` with `Radius float64`
3. Add **value receiver** methods on `Rectangle`:
   - `Area() float64` — returns `Width * Height`
   - `Perimeter() float64` — returns `2 * (Width + Height)`
4. Add a **pointer receiver** method on `Rectangle`:
   - `Scale(factor float64)` — multiplies both Width and Height by factor
5. Add **value receiver** methods on `Circle`:
   - `Area() float64` — returns `math.Pi * Radius * Radius`
   - `Perimeter() float64` — returns `2 * math.Pi * Radius`
6. In `main()`:
   - Create `rect := Rectangle{Width: 10, Height: 5}`
   - Print rect's area and perimeter
   - Call `rect.Scale(2)` and print the new area to show the pointer receiver modified rect
   - Create `circle := Circle{Radius: 7}`
   - Print circle's area and perimeter

## Expected Output

```
Rectangle (10.0 x 5.0):
  Area:      50.00
  Perimeter: 30.00

After Scale(2) — Rectangle (20.0 x 10.0):
  Area:      200.00
  Perimeter: 60.00

Circle (radius 7.0):
  Area:      153.94
  Perimeter: 43.98
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge11/main.go
```

## Hints

- Value receiver method: `func (r Rectangle) Area() float64 { return r.Width * r.Height }`
- Pointer receiver method: `func (r *Rectangle) Scale(f float64) { r.Width *= f; r.Height *= f }`
- Import `"math"` for `math.Pi`
- Use `fmt.Printf("  Area:      %.2f\n", rect.Area())` for two-decimal formatting
- `rect.Scale(2)` works even though `rect` is not a pointer — Go automatically takes `&rect`
