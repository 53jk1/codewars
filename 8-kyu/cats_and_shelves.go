package kata

import (
	"math"
)

func Cats(start, finish int) int {
	diff := finish - start
	return int(math.Floor(float64(diff/3 + diff%3)))
}
