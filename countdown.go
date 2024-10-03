package sprint
import (
	"fmt"
)
func Countdown(n int) string {
	var result string

	for i :=n; i >= 0; i -= 2 {
		if result != "" {
			result += ", "
		}
		result += fmt.Sprintf("%d", i)
	}
	result += "!"
	return result
}