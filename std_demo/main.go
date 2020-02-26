package main

import "fmt"

func main() {
	var n int
	fmt.Println("How many names")
	fmt.Scanf("%d", &n)
	var names []string
	for i := 0; i < n; i++ {
		fmt.Println("Enter a names:")
		var name string
		fmt.Scanf("%s", &name)
		names = append(names, name)
	}
	fmt.Println("Names:", names)
}
