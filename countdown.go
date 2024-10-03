package sprint

func Countdown(n int) string {
	result := ""

	for i :=n; i >= 0; i -= 2 {
		if result != "" {
			result += ", "
		}
		if i == 0 {
			result += "0!"
		} else {
			result += string('0'+i)
		}
	}
	return result
}