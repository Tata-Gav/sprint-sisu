package sprint

func Countdown(n int) string {
	result := ""

	for i := n; i >= 0; i -= 2 {
		if result != "" {
			result += ", "
		}

		if i == 0 {
			result += "0!"
		} else {
			result += fmt.Sprintf("%d", i)
		}
	}

	return result
}
