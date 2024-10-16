package sprint

func FromRoman(s string) int {
    romanToInt := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    total := 0
    length := len(s)
    
    for i := 0; i < length; i++ {
        if i < length-1 && romanToInt[s[i]] < romanToInt[s[i+1]] {
            total -= romanToInt[s[i]]
        } else {
            total += romanToInt[s[i]]
        }
    }
	return total
}