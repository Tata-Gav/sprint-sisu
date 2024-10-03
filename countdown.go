package sprint

func Countdown(n int) string {
	result := ""

	for i := n; i >= 0; i -= 2 {
		if result != "" {
			result += ", "
		}
		result += string('0' + i)
	}

	if n > 0 { 
		result += ", "
	}
	result += "0!"

	return result
}