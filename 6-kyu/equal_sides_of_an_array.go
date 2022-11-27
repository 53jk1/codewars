package kata

func FindEvenIndex(arr []int) int {
	for i := 0; i < len(arr); i++ {
		if sum(arr[:i]) == sum(arr[i+1:]) {
			return i
		}
	}
	return -1
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
