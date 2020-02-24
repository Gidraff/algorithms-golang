package numeric

import "fmt"

// Sum returns sum of integers in a list
// func Sum(numbers []int) int {
// 	total := 0
// 	for _, val := range numbers {
// 		total += val
// 	}
// 	return total
// }

// Sum returns a sum of integers
// used recursion
func Sum(numbers []int) int {
	fmt.Println(numbers)
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}
