package sprint

// IsLower проверяет, содержит ли строка только строчные символы
func IsLower(s string) bool {
    for _, c := range s {
        if c < 'a' || c > 'z' { // Если символ не в диапазоне строчных букв
            return false
        }
    }
    return len(s) > 0 // Возвращаем true, если строка не пустая и все символы строчные
}

// IsUpper проверяет, содержит ли строка только заглавные буквы
func IsUpper(s string) bool {
    hasLetter := false
    for _, c := range s {
        if c >= 'A' && c <= 'Z' {
            hasLetter = true
        } else if c >= 'a' && c <= 'z' || c >= '0' && c <= '9' {
            return false
        }
    }
    return hasLetter
}

// IsAlphanumeric проверяет, содержит ли строка только алфавитные и числовые символы
func IsAlphanumeric(s string) bool {
    for _, c := range s {
        if !(('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') || ('0' <= c && c <= '9')) {
            return false
        }
    }
    return len(s) > 0 // Возвращаем true, если строка не пустая и содержит только буквы/цифры
}

// IsNumeric проверяет, содержит ли строка только цифры
func IsNumeric(s string) bool {
    for _, c := range s {
        if !('0' <= c && c <= '9') {
            return false
        }
    }
    return len(s) > 0 // Возвращаем true, если строка не пустая и содержит только цифры
}

// ArrCountIf возвращает количество элементов в массиве строк, которые возвращают true для функции f
func ArrCountIf(f func(string) bool, a []string) int {
    count := 0
    for _, s := range a {
        if f(s) {
            count++ // Увеличиваем счетчик, если функция f вернула true для элемента
        }
    }
    return count
}