package numeric

import "fmt"

// FizzBuzz prints FizzBuzz
func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("Fizz Buzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		case i%5 == 0:
			fmt.Print("Buzz")
		default:
			fmt.Print(n)
		}
		if i != n {
			fmt.Print(". ")
		}
	}
	fmt.Println()
}
