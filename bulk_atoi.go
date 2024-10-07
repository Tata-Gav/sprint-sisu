package sprint

func ConvertStrToInt(s string) int {
	if s == "" {
        return 0
    }
    
    isNegative := false
    if s[0] == '-' {
        isNegative = true
        s = s[1:]
    }
    result := 0
    for _, char := s {
    if char < '0' || char > '9' {
        return 0
    }
    result = result*10 + int(char-'0')
}
    if isNegative{
    return -result
}
return result
}

func BulkAtoi(arr []string) []int {
    result := []int{}

    for _, str := range arr {
        result = append(result, ConvertStrToInt(str))
    }
    return result 
}