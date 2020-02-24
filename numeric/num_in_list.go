package numeric

// Determine if a number is in a list
// Given a list of numbers, determine if a specific number is in that list

// NumInList will return true if num is in the
// list slice, and false otherwise
func NumInList(list []int, num int) bool {
	for _, i := range list {
		if i == num {
			return true
		}
	}
	return false
}
