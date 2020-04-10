package sorting

import "fmt"

// Iterates through a list, comparing the first element and the second element.
// if the first element is larger than the second element, it swaps them.
// It performs this action for the entire list. BubbleSort TC is O(N^2)

// BubbleSortInt sorts a list of integers
func BubbleSortInt(list []int) {
	for sweepNum := 0; sweepNum < len(list); sweepNum++ {
		swapped := false
		for i := 0; i < len(list)-1-sweepNum; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Println(list)
}
