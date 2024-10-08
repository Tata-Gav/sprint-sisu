package sprint

func ToUpperCase(s string) string {
    result := ""
    for _, char := range s {
        if char >= 'a' && char <= 'z' {
            // Преобразуем строчную букву в заглавную
            char = char - 32
        }
        result += string(char) // Добавляем символ в результат
    }
    return result
}