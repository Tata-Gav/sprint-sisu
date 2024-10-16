package sprint

func ToRoman(num int) string {
    if num < 1 || num > 3999 {
        return "Invalid input"
    }

    val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

    roman := ""
    for i := 0; i < len(val); i++ {
        for num >= val[i] {
            num -= val[i]
            roman += syms[i]
        }
    }
	return roman
}