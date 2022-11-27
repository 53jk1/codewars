package kata

import (
	"strings"
)

func ToJadenCase(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, " ")
}
