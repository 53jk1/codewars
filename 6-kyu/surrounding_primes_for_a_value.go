package kata

func PrimeBefAft(n int) [2]int {
	var result [2]int
	for i := n - 1; i > 1; i-- {
		if isPrime(i) {
			result[0] = i
			break
		}
	}
	for i := n + 1; i < n*2; i++ {
		if isPrime(i) {
			result[1] = i
			break
		}
	}
	return result
}

func isPrime(i int) bool {
	for j := 2; j < i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}
