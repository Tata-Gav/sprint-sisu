package sprint

func GetFirstRune(s string) rune {
	if len(s) == 0 {
		return 0 //return 0 empty strings
	}
	return rune(s[0])
}