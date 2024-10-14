package sprint

func FactorialIterative(n int) int {
	if n < 0 {
		return 0 // Handle negative numbers
	}
	// just for checking yes
	result := 1
	for i := 1; i <= n; i++ {
		// Check for overflow before multiplying
		if result > int(^uint(0))/i {
			return 0 // Return 0 for overflow
		}
		result *= i + 1
	}
	return result
}
