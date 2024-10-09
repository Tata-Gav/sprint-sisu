package sprint

import (
	"math"
	"strings"
)

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	if len(baseFrom) < 2 || len(baseTo) < 2 || strings.ContainsAny(baseFrom, "+-") || strings.ContainsAny(baseTo, "+-") {
			return ""
	}

	fromBaseMap := make(map[rune]int)
	for i, char := range baseFrom {
			fromBaseMap[char] = i
	}

	toBaseMap := make(map[rune]int)
	for i, char := range baseTo {
			toBaseMap[char] = i
	}

	decimalValue := 0
	baseFromLen := len(baseFrom)
	for i, char := range nbr {
			if _, ok := fromBaseMap[char]; !ok {
					return ""
			}
			decimalValue += fromBaseMap[char] * int(math.Pow(float64(baseFromLen), float64(len(nbr)-i-1)))
	}

	result := ""
	baseToLen := len(baseTo)
	for decimalValue > 0 {
			remainder := decimalValue % baseToLen
			result = string(baseTo[toBaseMap[rune(baseTo[remainder])]]) + result
			decimalValue /= baseToLen
	}

	return result
}