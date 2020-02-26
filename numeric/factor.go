package numeric

// Factor takes in a list of primes and a number
// and factors that number with the provided primes.
func Factor(primes []int, number int) []int {
	var res []int
	for _, prime := range primes {
		for number%prime == 0 {
			res = append(res, prime)
			number = number / prime
		}
	}
	if number > 1 {
		res = append(res, number)
	}
	return res
}
