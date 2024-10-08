package sprint

func StrReverse(s string) string {
	runes := []rune(s) //Преобразуем строку в срез символов
	for i, j := 0, len(rune)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes [j], runes[i]
	}
	returnstring(runes)
}