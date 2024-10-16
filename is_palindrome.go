package sprint

// Function to check if a string is a palindrome
func IsPalindrome(s string) bool {
    n := len(s)
    for i := 0; i < n/2; i++ {
        if s[i] != s[n-1-i] {
            return false
        }
    }
    return true
}