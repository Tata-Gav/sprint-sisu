package sprint

func IsLower(s string) bool {
	for _, char := range s {
		if !unicode.IsLower(char) {
			return false
		}
	}
	return true
}