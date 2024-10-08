package sprint

func ToCapitalCase(s string) string {
    var result []rune
    capitalizeNext := true

    for i, r := range s {
        if isWordBoundary(byte(r)) {
            result = append(result, r)
            capitalizeNext = true
        } else {
            if capitalizeNext {
                if 'a' <= r && r <= 'z' {
                    r = r - ('a' - 'A')
                }
                capitalizeNext = false
            } else {
                if 'A' <= r && r <= 'Z' {
                    r = r + ('a' - 'A')
                }
            }
            result = append(result, r)
        }
    }
    return string(result)
}

func isWordBoundary(c byte) bool {
    return c == ' ' || c == '!' || c == '?' || c == '[' || c == '{' || 
           c == '(' || c == ':' || c == '}' || c == '-' || c == '/' || 
           c == '+' || c == '%'
}
