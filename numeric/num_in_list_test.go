package numeric

import "testing"

func TestNumInList(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5}
	result := NumInList(list1, 5)
	if result == !true {
		t.Errorf("NumInList(%p, %d) = %t", list1, 5, result)
	}
}
