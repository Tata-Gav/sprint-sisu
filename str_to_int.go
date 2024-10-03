package sprint

func StrToInt(s string) int {
	var result int
	sign := 1
	started := false

	for _, ch := range s {
		if ch == ' ' && !started { // Пропускаем пробелы в начале строки
			continue
		}
		if ch == '+' && !started { // Пропускаем первый знак "+"
			started = true
			continue
		}
		if ch == '-' && !started { // Если первый символ "-" — устанавливаем отрицательный знак
			sign = -1
			started = true
			continue
		}

		// Если встречен не цифровой символ — возвращаем 0
		if ch < '0' || ch > '9' {
			return 0
		}

		// Преобразуем символ в цифру и добавляем к результату
		result = result*10 + int(ch-'0')
		started = true
	}

	return result * sign
}