package main

import (
	"fmt"
	"strconv"
)

/*
- The zero value of a map is nil. A nil map has no keys, nor can keys be added. why would this be useful? wouldn't me using the make function returns a map of the given type, initialized and ready for use.
Why Nil Maps Matter
1. Reading is safe, writing panics
var m map[string]int  // nil map
value := m["key"]      // ✓ OK! Returns 0 (zero value)
m["key"] = 42          // ✗ PANIC! Can't write to nil map

2. Detecting "never initialized" vs "empty"
var config map[string]string  // nil = "no config loaded yet"
config = make(map[string]string)  // empty = "config loaded, just has no entries"

if config == nil {
    fmt.Println("Config never loaded")
} else if len(config) == 0 {
    fmt.Println("Config loaded but empty")
}

3. Optional data without allocation
type User struct {
    Name     string
    Metadata map[string]string  // nil by default, only allocate if needed
}

u := User{Name: "Alice"}  // Metadata is nil - no memory wasted
// Later, if needed:
u.Metadata = make(map[string]string)
u.Metadata["role"] = "admin"

4. Safe iteration
var m map[string]int  // nil
for k, v := range m {  // ✓ OK! Just doesn't loop (0 iterations)
    fmt.Println(k, v)
}
*/

/*
NOTES
- Map literals are like struct literals, but the keys are required.
- Mutating maps:
	Insert or update an element in map m: m[key] = elem
	Retrieve an element: elem = m[key]
	Delete an element: delete(m, key)
	Test that a key is present with a two-value assignment: elem, ok = m[key]
- Go functions may be closures. A closure is a function value that references variables from outside its body.
*/

var s = makeSearchCounter()

func main() {
	var phonebook = map[string]string{
		"Alice":   "555-1234",
		"Bob":     "555-678",
		"Charlie": "555-9012",
	}

	searchAlice := lookup(phonebook, "Alice")
	searchUnknown := lookup(phonebook, "Dave")

	phonebook["Dave"] = "555-3456"
	fmt.Println("Added Dave")

	searchDave := lookup(phonebook, "Dave")
	phonebook["Alice"] = "555-0000"
	fmt.Println("Updated Alice")

	delete(phonebook, "Charlie")
	fmt.Println("Deleted Charlie")

	printPhonebook(phonebook)

	fmt.Println("Search count after \"Alice\": " + strconv.Itoa(searchAlice))
	fmt.Println("Search count after \"unknown\": " + strconv.Itoa(searchUnknown))
	fmt.Println("Search count after \"Dave\": " + strconv.Itoa(searchDave))
}

func lookup(pb map[string]string, name string) int {
	number, ok := pb[name]

	if ok {
		fmt.Println("Looking up " + name + ": " + number)
	} else {
		fmt.Println("Looking up " + name + ": " + "Not found")
	}

	return s(1)
}

func makeSearchCounter() func(int) int {
	searchCount := 0
	return func(x int) int {
		searchCount += x
		return searchCount
	}
}

func printPhonebook(pb map[string]string) {
	fmt.Println()
	fmt.Println("Final phonebook:")

	for key, value := range pb {
		fmt.Println("    " + key + ": " + value)
	}
	fmt.Println()
}
