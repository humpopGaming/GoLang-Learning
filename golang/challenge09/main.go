package main

import "fmt"

func main() {
	list := [5]string{"Milk", "Eggs", "Bread", "Butter", "Cheese"}
	fmt.Printf("Store aisle items: %v\n", list) // need to learn the difference between %d, %v and other options

	s := list[0:3]
	fmt.Printf("First 3 aisle items: %v, %d, %d\n", s, len(s), cap(s))
	fmt.Println()

	// shoppingList := []string
	var shoppingList []string // this works but not the above, why?
	/* Answer to above question:
	var shoppingList []string  // ✅ declares a nil slice
	shoppingList := []string   // ❌ syntax error — := needs a value on the right
	shoppingList := []string{} // ✅ declares an empty (non-nil) slice

	The := operator is short variable declaration — it declares and initializes. You need something on the right side.
	*/

	fmt.Printf("Shopping list is nil: %v\n", shoppingList == nil)
	fmt.Println()
	// why cant i append to list, but I can to s and shoppingList?
	/* Answer to above question:
	list := [5]string{...}     // ARRAY — fixed size, can't grow
	s := list[0:3]              // SLICE — can grow with append
	shoppingList := []string    // SLICE — can grow with append

	Arrays are fixed. Slices are dynamic. append only works on slices!
	*/

	shoppingList = append(shoppingList, "Apples", "Bananas", "Oranges")
	printShoppingList(shoppingList)

	shoppingList = append(shoppingList, "Yogurt", "Juice")
	printFinalShoppingList(shoppingList)
	printSlice(shoppingList)
}

func printSlice(s []string) {
	fmt.Printf("List length: %d, capacity: %d\n", len(s), cap(s))
}

func printShoppingList(l []string) {
	fmt.Println("Shopping list:")
	for i, item := range l { //how is this for loop working, what is range, what other key words can be used here?
		/* Answer to above question:
		range iterates over a slice (or array, map, string):

		for index, value := range mySlice {
			// index = 0, 1, 2, ...
			// value = the item at that position
		}

		for _, value := range mySlice {  // skip the index
		for index := range mySlice {     // skip the value (just get indices)
		No other keywords work here — range is the Go way to iterate collections.
		*/

		fmt.Printf("    %v: %v\n", i, item)
	}
	fmt.Println()
}

func printFinalShoppingList(l []string) {
	fmt.Println("Final shopping list:")
	for _, item := range l {
		fmt.Printf("    - %v\n", item)
	}
	fmt.Println()
}
