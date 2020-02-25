package numeric

func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"
	res := 0
	multiplier := 1
	for i := len(value) - 1; i >= 0; i-- {
		index := -1
		for j, char := range charset {
			if char == rune(value[i]) {
				index = j
				break
			}
		}
		if index < 0 {
			// Shouldn't happen
			panic("Something went wrong!")
		}
		res = res + index*multiplier
		multiplier = multiplier * base
	}

	return res
}
