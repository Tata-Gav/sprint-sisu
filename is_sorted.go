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

func IsSorted(f func(a, b string) int, arr []string) bool {
    for i := 0; i < len(arr)-1; i++ {
        if f(arr[i], arr[i+1]) > 0 {
            return false // Not sorted in ascending order
        }
    }
    return true
}