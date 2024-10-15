package sprint

func IsUpper(s string) bool {
    for _, c := range s {
        if c >= 'a' && c <= 'z' {
            return false
        }
    }
    return true
}

func IsAlphanumeric(s string) bool {
    for _, c := range s {
        if !(('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') || ('0' <= c && c <= '9')) {
            return false
        }
    }
    return true
}

func IsLower(s string) bool {
    for _, c := range s {
        if c >= 'A' && c <= 'Z' {
            return false
        }
    }
    return true
}

func ArrAny(f func(string) bool, a []string) bool {
    for _, s := range a {
        if f(s) {
            return true
        }
    }
    return false
}