package sorting

import "fmt"

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
