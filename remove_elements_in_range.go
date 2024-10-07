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
