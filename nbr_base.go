package sprint

func NbrBase(n int, base string) string {
    
    if len(base) < 2  containsInvalidChars(base) {
        return "NV" 
    }
    
    isNegative := n < 0
    if isNegative {
        n = -n 
    }

    if n == 0 {
        return string(base[0])
    }

    result := ""
    for n > 0 {
        digit := n % len(base) 
        result = string(base[digit]) + result 
        n = n / len(base) 
    }

    if isNegative {
        result = "-" + result
    }

    return result
}

func containsInvalidChars(base string) bool {
    seen := make(map[rune]bool)
    for _, char := range base {
        
        if char == '+'  char == '-' || seen[char] {
            return true
        }
        seen[char] = true
    }
    return false
}