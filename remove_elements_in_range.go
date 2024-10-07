package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
    // Handle negative indices
    if from < 0 {
        from = 0
    }
    if to < 0 {
        to = 0
    }

    // Ensure from <= to
    if from > to {
        from, to = to, from
    }

    // Handle indices larger than array length
    if from >= len(arr) {
        return arr
    }
    if to > len(arr) {
        to = len(arr)
    }

    // Remove elements within the specified range
    result := append(arr[:from], arr[to:]...)
    return result
}
func main() {
    arr := []float64{10.0, 0.8, -0.4, 20.0, 7.7, 3.0}
    result := RemoveElementsInRange(arr, 4, 1)
    fmt.Println(result) // Output: [10 7.7 3]
}