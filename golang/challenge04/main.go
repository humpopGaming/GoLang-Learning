// copilot inline suggestions have been disabled!!!
package main

func isDivisibleBy(i int) (bool, bool) {
	return i%3 == 0, i%5 == 0
}

func findSpecialNumber() int {
	n := 101
	for n%21 != 0 {
		n++
	}
	return n
}

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if three, five := isDivisibleBy(i); three && five {
			println("FizzBuzz")
		} else if five {
			println("Buzz")
		} else if three {
			println("Fizz")
		} else {
			println(i)
		}
	}
}

func main() {
	fizzBuzz(30)
	println(findSpecialNumber())
}
