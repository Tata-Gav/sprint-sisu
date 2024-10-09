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
    for n > 0 {
        remainder := n % len(base)
        result = string(base[remainder]) + result
        n /= len(base)
    }

    if isNegative {
        result = "-" + result
    }

    return result
}