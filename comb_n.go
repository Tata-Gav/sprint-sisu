package sprint

import (
	"fmt"
)

func CombN(n int) []string {
	var result []string
	var generate func(prefix string, start int)

	generate = func(prefix string, start int) {
		if len(prefix) == n {
			result = append(result, prefix)
			return
		}
		for i := start; i <= 9; i++ {
			generate(prefix+fmt.Sprintf("%d", i), i+1)	
		}
	}

		generate("", 0)
		return result
}