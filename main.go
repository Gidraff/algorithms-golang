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

	fmt.Println("DecToBase =>:")
	fmt.Println(numeric.DecToBase(10, 2))

	fmt.Println("BaseToBase =>:")
	fmt.Println(numeric.BaseToBase("10101", 2, 10))

	fmt.Println("FindTwoThatSum =>:")
	fmt.Println(numeric.FindTwoThatSum([]int{1, 2, 3, 4, 5}, 9))

	fmt.Println("Factors of Given number =>:")
	fmt.Println(numeric.Factor([]int{3, 5, 7, 9}, 9))

	fmt.Println("Fibonacci Sequence:")
	fmt.Println(numeric.Fibonacci(6))

	fmt.Println("GCD =>:")
	fmt.Println(numeric.GCD(20, 15))
}
