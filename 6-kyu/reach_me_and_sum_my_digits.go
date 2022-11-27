package kata

func SumDigNthTerm(initVal int, pattern []int, nth int) int {
	if nth == 1 {
		return initVal
	}
	sum := initVal
	for i := 1; i < nth; i++ {
		sum = sum + pattern[(i-1)%len(pattern)]
	}

	var sumResult int

	for sum != 0 {
		reminder := sum % 10
		sumResult += reminder
		sum = sum / 10
	}

	return sumResult
}
