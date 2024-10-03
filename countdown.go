package sprint

func Countdown(n int) string {
	result := ""

	for i := n; i >= 0; i -= 2 {
		if result != "" {
			result += ", "
		}
		// Добавляем текущее число в строку
		result += string('0' + i)
	}

	result += "!" // Добавляем "!" в конце после цикла
	return result
}