package sprint

func Payout(amount int, denominations []int) (payout []int) {
    // Bubble sort to order denominations in descending order
    n := len(denominations)
    for i := 0; i < n; i++ {
        for j := 0; j < n-i-1; j++ {
            if denominations[j] < denominations[j+1] {
                denominations[j], denominations[j+1] = denominations[j+1], denominations[j]
            }
        }
    }

    for _, coin := range denominations {
        for amount >= coin {
            amount -= coin
            payout = append(payout, coin)
        }
    }

    if amount != 0 {
        return []int{}
    }
}