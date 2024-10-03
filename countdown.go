package sprint

func Countdown(n int) string {
	result := ""

	for i := n; i >= 0; i -= 2 {
		if result != "" {
			result += ", "
		}

		if i == 0 {
			result += "0!" 
		} else if i < 10 {
			result += string('0' + i) 
		} else {
			result += string(i + '0') 
		}
	}
	return result
}