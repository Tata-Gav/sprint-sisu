package sprint

func AlphaNumber(n int) string {
	result := ""
	numStr := ""

	// Проверка на отрицательное число
	if n < 0 {
		numStr = "-" + intToString(-n)
	} else {
		numStr = intToString(n)
	}

	// Преобразование каждой цифры в соответствующую букву
	for _, ch := range numStr {
		if ch == '-' {
			result += "-"
		} else {
			result += string('a' + (ch - '0')) // Преобразуем цифру в букву
		}
	}
	return result
}

// Вспомогательная функция для преобразования числа в строку
func intToString(num int) string {
	if num == 0 {
		return "0"
	}
	
	result := ""
	for num > 0 {
		digit := num % 10
		result = string('0'+digit) + result
		num /= 10
	}
	return result
}