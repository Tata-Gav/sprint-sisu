package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
    result := [][]int{}
    for _, subarr := range arr {
        sum := 0
        for _, num := range subarr {
            sum += num
        }
        if sum >= limit {
            result = append(result, subarr)
        }
    }
    return result
}