package sprint

func GSD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a

}
func LCM(a, b int) int {
	return abs(a * b) / GCD(a, b)

}
 func abc(x int) int {
	if x < 0 {
		return -x
	}
	return x
 }