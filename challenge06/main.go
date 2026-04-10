package main

import "fmt"

func greet(name string) {
	defer fmt.Println("Goodbye, " + name + "!")
	fmt.Println("Hello, " + name + "?")
	fmt.Println("How are you, " + name + "!")
}

func countdown(from int) {
	fmt.Println("Starting countdown...")
	for i := 1; i <= from; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Go!")
}

func main() {
	greet("Gopher")
	fmt.Println()
	countdown(5)
}
