package kata

func LongestRepetition(text string) Result {
	var res Result
	var cur Result
	for _, c := range text {
		if c == cur.C {
			cur.L++
		} else {
			cur.C = c
			cur.L = 1
		}
		if cur.L > res.L {
			res = cur
		}
	}
	return res
}
