package sprint

import (
	"unicode"
)

func ArrAny(f func(string) bool, a []string) bool {
	for _, s := range a {
			if f(s) {
					return true
			}
	}
	return false
}

func IsUpper(s string) bool {
	for _, c := range s {
			if !unicode.IsUpper(c) {
					return false
			}
	}
	return true
}

func IsAlphanumeric(s string) bool {
	for _, c := range s {
			if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
					return false
			}
	}
	return true
}