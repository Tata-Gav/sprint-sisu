package sprint

func FactorialRecursive(n int) int {
    if n < 0 {
        return 0 
    }

    if n == 0 || n == 1 {
        return 1
    }

    result := FactorialRecursive(n-1)
    if n*result < 0 {
        return 0
    }
    return n * result
}
