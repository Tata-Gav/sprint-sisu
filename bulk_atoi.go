package sprint

func BulkAtoi(arr []string) []int {
	
	result := 0
    sign := 1
    i := 0

    // Handle leading whitespace
    for i < len(s) && s[i] == ' ' {
        i++
    }

    // Handle sign
    if i < len(s) && (s[i] == '+' || s[i] == '-') {
        if s[i] == '-' {
            sign = -1
        }
        i++
    }

    // Parse digits
    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        digit := int(s[i] - '0')
        result = result*10 + digit
        i++
    }

    return result * sign
}

func BulkAtoi(arr []string) []int {
    result := make([]int, len(arr))
    for i, s := range arr {
        result[i] = StrToInt(s)
    }
    return result
}
