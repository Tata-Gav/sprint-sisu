package sprint

func Countdown(n int) string {
	 var result string

	for i := n; i >= 1; i -= 2 {
		result += string('0' + rune(i))


	if i >= 1 { 
		result += ", "
	}
	return result + "0!"
	}
}
