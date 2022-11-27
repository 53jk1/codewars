package kata

func ProperFractions(n int) int {
	if n == 1 {
		return 0
	}
	result := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			result -= result / i
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}
