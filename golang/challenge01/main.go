package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello, Go Playground!")
	fmt.Printf("Pi is: %f \n", math.Pi)
	fmt.Printf("Random number: %d \n", rand.Intn(100))
	fmt.Printf("Square root of 144 is: %g \n", math.Sqrt(144))
}