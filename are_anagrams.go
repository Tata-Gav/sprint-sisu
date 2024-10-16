package sprint

import (
    "sort"
    "strings"
)

func AreAnagrams(str1, str2 string) bool {
    // Удаляем пробелы и приводим строки к нижнему регистру
    str1 = strings.ReplaceAll(str1, " ", "")
    str2 = strings.ReplaceAll(str2, " ", "")
    str1 = strings.ToLower(str1)
    str2 = strings.ToLower(str2)
    
    // Если длины строк не равны, они не могут быть анаграммами
    if len(str1) != len(str2) {
        return false
    }

    // Превращаем строки в срезы рун (символов)
    s1 := []rune(str1)
    s2 := []rune(str2)

    // Сортируем срезы рун
    sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
    sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })

    // Сравниваем отсортированные строки
    return string(s1) == string(s2)
}