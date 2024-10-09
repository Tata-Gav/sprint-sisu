package sprint

func NbrBase(n int, base string) string {
    if base == "" || len(base) < 2 {
        panic("Invalid base")
    }

    isNegative := n < 0
    if isNegative {
        n = -n
    }

    result := ""
    baseMap := make(map[rune]int)
    for i, char := range base {
        baseMap[char] = i
    }

    for n > 0 {
        remainder := n % len(base)
        result = string(base[baseMap[rune(base[remainder])]]) + result
        n /= len(base)
    }

    if isNegative {
        result = "-" + result
    }

    return result
}