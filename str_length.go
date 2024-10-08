package sprint

func StrLength(s string) []int {
	
	runeCount := len([]rune(s))
	byteCount := len(s)
	return []int{runeCount, byteCount}
}