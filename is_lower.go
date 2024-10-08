package sprint

import "unicode"

func IsLower(s string) bool {
    for _, char := range s {
        if !unicode.IsLower(char) {
            return false
        }
    }
    return true
}