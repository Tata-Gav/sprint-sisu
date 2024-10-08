package Sprint

func ToCapitalCase(s string) string {
    var result []rune
    for i, r := range s {
        // Если это первый символ или предыдущий символ - разделитель
        if i == 0  isWordBoundary(s[i-1]) {
            // Преобразуем первую букву в заглавную
            if 'a' <= r && r <= 'z' {
                r = r - ('a' - 'A')
            }
        } else {
            // Преобразуем заглавные буквы в строчные
            if 'A' <= r && r <= 'Z' {
                r = r + ('a' - 'A')
            }
        }

        // Обрабатываем все буквы, кроме первой
        if i > 0 && !isWordBoundary(s[i-1]) {
            if 'A' <= r && r <= 'Z' {
                r = r + ('a' - 'A')
            }
        }

        result = append(result, r)
    }
    return string(result)
}

// Функция для проверки границ слова
func isWordBoundary(c byte) bool {
    return c == ' '  c == '!'  c == '?'  c == '['  c == '{'  c == '(' 
        c == ':'  c == '}'  c == '-'  c == '/'  c == '+'  c == '%'
}
