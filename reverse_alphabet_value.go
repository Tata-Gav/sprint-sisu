package sprint

func ReverseAlphabetValue(ch rune) rune {

	if ch < 'a' || ch > 'z' {
		return 0
	}
	return 'z' - (ch - 'a')
}