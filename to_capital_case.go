package sprint

func ToCapitalCase(s string) string {
    var result strings.Builder
    capitalizeNext := true

    for _, r := range s {
        if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
            if capitalizeNext {
                r = r - 32
            } else {
                r = r + 32
            }
            result.WriteRune(r)
            capitalizeNext = false
        } else if r == ' ' || r == '-' || r == '_' {
            result.WriteRune(r)
            capitalizeNext = true
        } else {
            result.WriteRune(r)
        }
    }

    return result.String()
}