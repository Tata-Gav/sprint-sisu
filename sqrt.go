package sprint

func Sqrt(n int) int {
    sqrt := int(n / 2) // Initial guess for the square root

    for sqrt*sqrt > n {
        sqrt--
    }

    for sqrt*sqrt < n {
        sqrt++
    }

    if sqrt*sqrt == n {
        return sqrt
    }
    return 0
}