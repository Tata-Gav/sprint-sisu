package sprint

func RemoveDuplicates(arr []int) []int {
    seen := make(map[int]bool)
    result := []int{}

    for _, value := range arr {
        if !seen[value] {
            seen[value] = true
            result = append(result, value)
        }
    }

    return result
}