package sprint

func ShiftBy(r rune, step int) rune {

	if r < 'a' || r > 'z' {
		return r
	}
	return 'a' + (r-'a'+rune(step)+26)%26
}