package main

import (
	"fmt"
	"strings"
)

/*
Pattern counting Problem: Count the number of times that a pattern appears
in a longer test
	Input: Strings text and pattern.
	Output: The number of times that pattern appears as substring of text
*/
func CountPattern(text, pattern string) int {
	count := 0
	n := len(text)
	k := len(pattern)

	for i := 0; i <= n-k; i++ {
		if pattern == text[i:i+k] {
			count++
		}
	}
	return count
}

func main() {
	//common string operations
	fmt.Println(strings.Count("banana", "an"))
	fmt.Println(strings.Contains("johndoe@gmail.com", "@"))

	fmt.Println("Counter: ", CountPattern("CGATATAGGATAC", "ATA"))

	fmt.Println("Hello " + "there")

	s := "Hi there!"
	fmt.Println("The length of a is", len(s))
	fmt.Println("The first symbol of s is", string(s[0]))
	fmt.Println("The last symbol of s is", string(s[len(s)-1]))

	if s[1] == 'i' {
		fmt.Println("Second symbol is e!")
	}

	// substrings
	fmt.Println(s[3:5])
	fmt.Println(s[4 : len(s)-1])
	fmt.Println(s[:4])

	substr := s[3:8]
	fmt.Println(substr)
	fmt.Println(string(substr[2]))
}
