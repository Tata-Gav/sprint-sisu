package sprint

func NextPrime(n int) int {
    if n <= 1 {
        n = 2 // Start with 2 if n is less than or equal to 1
    }

    for !isPrime(n) {
        n++
    }
    return n
}

func isPrime(n int) bool {
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