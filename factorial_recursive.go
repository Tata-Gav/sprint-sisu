package sprint

func FactorialRecursive(n int) int {
    if n < 0 {
        return 0 // Handle negative numbers
    }

    if n == 0 || n == 1 {
        return 1
    }

    // Check for overflow before multiplying
    result := FactorialRecursive(n-1)
    if n*result < 0 {
        return 0 // Return 0 for overflow
    }

    return n * result
}