package sprint

func StrCompare(a, b string) int {
    for i := 0; i < len(a) && i < len(b); i++ {
        if a[i] < b[i] {
            return -1
        } else if a[i] > b[i] {
            return 1
        }
    }

    if len(a) < len(b) {
        return -1
    } else if len(a) > len(b) {
        return 1
    }

    return 0
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
    // Implementation using bubble sort for simplicity
    for i := 0; i < len(a)-1; i++ {
        for j := i + 1; j < len(a); j++ {
            if f(a[i], a[j]) > 0 {
                a[i], a[j] = a[j], a[i]
            }
        }
    }
    return a
}