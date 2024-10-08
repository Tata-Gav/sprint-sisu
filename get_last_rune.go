package sprint

import "slices"

func GetLastRune(s string) rune {
	 if len(s) == 0 {
		return 0 //return 0 for empty string
	 }
	 return rune(slices[len(s)-1])
}