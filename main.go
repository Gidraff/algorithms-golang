package main

import (
	"fmt"

	"github.com/Gidraff/algorithms-golang/numeric"
)

func main() {
	numList := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(numeric.NumInList(numList, 5))
}
