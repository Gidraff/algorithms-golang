package numeric

func BaseToBase(value string, base, newBase int) string {
	dec := BaseToDec(value, base)
	return DecToBase(dec, newBase)
}
