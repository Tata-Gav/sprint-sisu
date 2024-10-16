package sprint

func LongestClimb(arr []int) []int {
    if len(arr) == 0 {
        return []int{}
    }

    longestStart, longestEnd := 0, 0
    currentStart := 0

    for i := 1; i < len(arr); i++ {
        if arr[i] > arr[i-1] {
            if i - currentStart > longestEnd - longestStart {
                longestStart = currentStart
                longestEnd = i
            }
        } else {
            currentStart = i
        }
    }

    return arr[longestStart:longestEnd+1]
}