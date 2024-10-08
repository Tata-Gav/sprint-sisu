package sprint

func ToCapitalCase(s string) string {
	result := ""
	capitalizeNext := true // Флаг для определения, нужно ли капитализировать следующую букву

	for _, char := range s {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			if capitalizeNext {
				if char >= 'a' && char <= 'z' {
					char -= 32 // Преобразуем строчную букву в заглавную
				}
				result += string(char) // Добавляем заглавную букву
				capitalizeNext = false  // Сбрасываем флаг
			} else {
				if char >= 'A' && char <= 'Z' {
					char += 32 // Преобразуем заглавную букву в строчную
				}
				result += string(char) // Добавляем строчную букву
			}
		} else {
			// Если символ не является буквой или цифрой, добавляем его без изменений
			result += string(char)
			if char == ' ' || char == '-' || char == '_' {
				capitalizeNext = true // Устанавливаем флаг для следующего слова
			}
		}
	}
	return result // Возвращаем результат только после обработки всех символов
}
