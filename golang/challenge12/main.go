package main

import (
	"fmt"
	"math"
)

type Rectangle struct { // are all types structs?
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimiter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Print() {
	fmt.Printf("Rectangle (%.1f x %.1f):\n", r.Width, r.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimiter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Print() {
	fmt.Printf("Circle (radius %.1f):\n", c.Radius)
}

// can i use an interface to mkae this better?
func PrintSize(a float64, p float64) {
	fmt.Printf("    Area:         %.2f\n", a)
	fmt.Printf("    Perimiter:    %.2f\n", p)
	fmt.Println("")
}

func (r *Rectangle) Scale(f float64) {
	r.Width *= f
	r.Height *= f
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	rect.Print()
	PrintSize(rect.Area(), rect.Perimiter())
	rect.Scale(2)
	rect.Print()
	PrintSize(rect.Area(), rect.Perimiter())

	circle := Circle{Radius: 7}
	circle.Print()
	PrintSize(circle.Area(), circle.Perimiter())
}
