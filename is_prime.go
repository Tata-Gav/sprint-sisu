package sprint

func IsPrime(n int) bool {
    if n <= 1 {
        return false // Handle numbers less than or equal to 1
    }

    // Check divisibility up to the square root of n
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false // Not prime if divisible
        }
    }

    return true // Prime if no divisors found
}