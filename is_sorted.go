package sprint

import "strings"

func StrCompare(a, b string) int {
    return strings.Compare(a, b)
}

func IsSorted(f func(a, b string) int, arr []string) bool {
    for i := 0; i < len(arr)-1; i++ {
        if f(arr[i], arr[i+1]) > 0 {
            return false // Not sorted in ascending order
        }
    }
    return true
}