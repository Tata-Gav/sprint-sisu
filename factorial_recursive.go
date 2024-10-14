package sprint

import (
    "math"
)
func FactorialRecursive(n int) int {
    if n < 0 {
        return 0 // non-possible value
    }
    if n == 0 || n == 1 {
        return 1
    }
    result := n * FactorialRecursive(n-1)
    if result < 0 || result > math.MaxInt { // Check for overflow
        return 0
    }
    return result
}
