package sprint

func Countdown(n int) string {
	var result string

	for i := n; i >= 0; i -= 2 {
		if result != "" {
			result += ", " 
		}
		result += string('0' + rune(i)) 
	}

	if result[len(result)-1] != '0' { 
		result += ", 0!"
	} else {
		result += "!" 
	}

	return result
}