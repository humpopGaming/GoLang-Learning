package main

import "fmt"

type Contact struct {
	Name  string
	Email string
	Age   int
}

func printCard(c Contact) {
	fmt.Printf("Name: %v\n", c.Name)
	fmt.Printf("Email: %v\n", c.Email)
	fmt.Printf("Age: %v\n", c.Age)
}

func main() {

	alice := Contact{"Alice Smith", "alice@example.com", 30}

	fmt.Println("--- Alice's Card ---")
	printCard(alice)
	fmt.Println()

	bob := Contact{Name: "Bob Jones", Email: "bob@example.com"}

	fmt.Println("--- Bob's Card ---")
	printCard(bob)
	fmt.Println()

	p := &alice
	p.Email = "alice@newExample.com"
	/*
		the above can also be achieved with
		alice.Email = "alice@newExample.com"

		This demonstrates Go's handy feature: p.Email automatically means (*p).Email — you don't need to dereference manually!

		This pointer practice here sets up for Challenge 11
	*/

	fmt.Println("--- Alice's Updated Card ---")
	printCard(alice)
	fmt.Println()

	fmt.Printf("Bob's age (zero value): %v\n", bob.Age)
}
