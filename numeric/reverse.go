package numeric

// Reverse will return the provided word in reverse
func Reverse(word string) string {
	var res string
	for i := len(word) - 1; i >= 0; i-- {
		res = res + string(word[i])
	}
	return res
}
