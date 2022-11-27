package kata

import "regexp"

func alphanumeric(str string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return re.MatchString(str)
}
