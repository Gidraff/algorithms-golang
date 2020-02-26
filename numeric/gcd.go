package numeric

//func GCD(a, b int) int {
//	if b == 0 {
//		return a
//	}
//	a, b = b, a%b
//	return GCD(a, b)
//}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
