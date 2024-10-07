package sprint

func SimpleStrToInt(arr []string) []int {
	
	result := 0
    i := 0
    for {
        if i >= len(s) {
       break
    }
    
    char := s[i]
    if char >=  '0' && char <= '9' {
        digit := int (char - '0')
        result = result*10 + digit
        }
        i++
    }
    
    if len(s) > 0 && s[0] = '-' {
        return -result
    }
    return result

}
func BulkAtoi(arr []string) []int {
    result := []int{}
    for i := 0, i < len(arr); i++ {
    value := SimpleStrToInt(arr[i])
    result = append(result)
    }
    return result 
}