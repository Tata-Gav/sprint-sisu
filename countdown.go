package sprint

func Countdown(n int) string {
	var result string

	for i := n; i >= 0; i -= 2 {
		if result != "" {
			result += ", " 
		}
		result += string('0' + rune(i)) 
	}

	result += "!"

	return result
}
