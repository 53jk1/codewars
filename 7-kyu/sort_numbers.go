package kata

func SortNumbers(numbers []int) []int {
	return sort(numbers)
}

func sort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	middle := len(numbers) / 2
	left := numbers[:middle]
	right := numbers[middle:]

	return merge(sort(left), sort(right))
}

func merge(ints []int, ints2 []int) []int {
	result := make([]int, len(ints)+len(ints2))
	i := 0
	j := 0
	k := 0
	for i < len(ints) && j < len(ints2) {
		if ints[i] < ints2[j] {
			result[k] = ints[i]
			i++
		} else {
			result[k] = ints2[j]
			j++
		}
		k++
	}

	for i < len(ints) {
		result[k] = ints[i]
		i++
		k++
	}

	for j < len(ints2) {
		result[k] = ints2[j]
		j++
		k++
	}

	return result
}
