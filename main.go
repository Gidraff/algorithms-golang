package main

import (
	"fmt"

	"github.com/Gidraff/algorithms-golang/numeric"
)

func main() {
	numList := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(numeric.NumInList(numList, 5))

	fmt.Print("Fizz Buzz Algo")
	fmt.Println()
	numeric.FizzBuzz(15)

	fmt.Println("DecToBase==>")
	fmt.Print(numeric.DecToBase(10, 2))
}
