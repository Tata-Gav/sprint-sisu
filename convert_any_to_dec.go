package sprint

import (
	"strings"
)

func ConvertAnyToDec(s string, base string) int {
	if len(base) < 2 || strings.ContainsAny(base, "+-") {
			return 0
	}

	baseMap := make(map[rune]int)
	for i, char := range base {
			baseMap[char] = i
	}

	result := 0
	baseLen := len(base)
	for _, char := range s {
			if _, ok := baseMap[char]; !ok {
					return 0
			}
			result = result*baseLen + baseMap[char]
	}

	return result
}