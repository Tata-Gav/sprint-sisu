package sprint

func ReverseAlphabet(step int) string {

	alphabet := "zyxwvutsrqponmlkjihgfedcba"
	if step <= 0 {
	step = 1
	}
	result := ""

	for i := 0; i < len(alphabet); i += step {
	result += string(alphabet[i])
	}
	
	return result
}