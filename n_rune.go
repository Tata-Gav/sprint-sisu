package sprint

func NRune(s string, i int) rune {
	if i < 0 || i >= len(s) {
		return 0 // return 0 for invalid indises
	}
	return rune(s[i])
}