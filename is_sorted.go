package sprint

func StrCompare(a, b string) int {
    // Handle leading zeros
    a = strings.TrimLeft(a, "0")
    b = strings.TrimLeft(b, "0")

    // Compare strings character by character
    for i := 0; i < len(a) && i < len(b); i++ {
        if a[i] < b[i] {
            return -1
        } else if a[i] > b[i] {
            return 1
        }
    }

    // Compare lengths if strings are equal up to this point
    if len(a) < len(b) {
        return -1
    } else if len(a) > len(b) {
        return 1
    }

    return 0
}