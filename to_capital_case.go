package Sprint

func ToCapitalCase(s string) string {
    var result []rune
    capitalizeNext := true // Flag to determine if the next character should be capitalized

    for i, r := range s {
        if isWordBoundary(byte(r)) {
            result = append(result, r) // Append the word boundary character
            capitalizeNext = true      // Next alphanumeric character should be capitalized
        } else {
            if capitalizeNext {
                // Capitalize the first letter of the word
                if 'a' <= r && r <= 'z' {
                    r = r - ('a' - 'A')
                }
                capitalizeNext = false // Reset flag after capitalizing
            } else {
                // Convert uppercase letters to lowercase for subsequent characters
                if 'A' <= r && r <= 'Z' {
                    r = r + ('a' - 'A')
                }
            }
            result = append(result, r) // Append the modified character
        }
    }
    return string(result)
}

// Function to check for word boundaries
func isWordBoundary(c byte) bool {
    return c == ' ' || c == '!' || c == '?' || c == '[' || c == '{' || 
           c == '(' || c == ':' || c == '}' || c == '-' || c == '/' || 
           c == '+' || c == '%'
}