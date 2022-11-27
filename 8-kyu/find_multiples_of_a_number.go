package kata

func FindMultiples(integer, limit int) []int {
	var result []int
	for i := integer; i <= limit; i += integer {
		result = append(result, i)
	}
	return result
}
