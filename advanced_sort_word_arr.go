package sprint

// StrCompare function should be defined beforehand
func StrCompare(a, b string) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    }
    return 0
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
    sort.Slice(a, func(i, j int) bool {
        return f(a[i], a[j]) < 0
    })
    return a
}