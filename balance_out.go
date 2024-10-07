package sprint

func BalanceOut(arr []bool) []bool {
    trueCount, falseCount := 0, 0
    for _, b := range arr {
        if b {
            trueCount++
        } else {
            falseCount++
        }
    }

    diff := trueCount - falseCount
    if diff == 0 {
        return arr
    }

    // Determine which value needs to be added
    var valueToAdd bool
    if diff > 0 {
        valueToAdd = false
    } else {
        valueToAdd = true
    }

    // Add the necessary number of values
    for i := 0; i < abs(diff); i++ {
        arr = append(arr, valueToAdd)
    }

    return arr
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}