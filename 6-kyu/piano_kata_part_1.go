package kata

func BlackOrWhiteKey(keyPressCount int) string {
	switch (keyPressCount - 1) % 88 % 12 {
	case 0, 2, 3, 5, 7, 8, 10:
		return "white"
	default:
		return "black"
	}
}
