package sprint

import ( 
"strings"
"math" )

func NbrBase(n int, base string) string {
	if len(base) < 2 || strings.ContainsAny(base, "+-") {
		return "NV"
	}

	isNegative := n < 0
	n = int(math.Abs(float64(n)))

	baseMap := make(map[int]rune)
	for i, char := range base {
		baseMap[i] = char
	}

	result := ""
	for n > 0 {
		remainder := n % len(base)
		result = string(baseMap[remainder]) + result
		n /= len(base)
	}

	if isNegative {
		result = "-" + result
	}

	return result
}