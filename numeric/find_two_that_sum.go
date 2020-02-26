package numeric

// FindTwoThatSum will look for two numbers in the provided numbers list that
// sum up to the sum argument. I will then return the indices of each of these
// numbers
func FindTwoThatSum(numbers []int, sum int) (int, int) {
	for i, val := range numbers {
		for j, val2 := range numbers {
			if i == j {
				continue
			}
			if val+val2 == sum {
				return i, j
			}
		}
	}
	// base case, will return if no answer
	return -1, -1
}
