package sprint

func GetLastRune(s string) rune {
	 if len(s) == 0 {
		return 0 //return 0 for empty string
	 }
	 return rune(s[len(s)-1])
}