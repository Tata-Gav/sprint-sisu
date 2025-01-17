package sprint

func FactorialIterative(n int) int {
    if n < 0 {
        return 0 // Handle negative numbers
    }

    result := 1
    for i := 1; i <= n; i++ {
        // Check for overflow before multiplying
        if result*i < 0 {
            return 0 // Return 0 for overflow
        }
        result *= i
    }
    return result
}