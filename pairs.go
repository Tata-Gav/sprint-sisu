package sprint

import (
	"fmt"
	"strings"
)

func Pairs() string {
	var result strings.Builder

	for i := 0; i < 100; i++ {

		for j := i + 1; j < 100; j++ {
			
			pair := fmt.Sprintf("%d %d", i, j)

			if result.Len() > 0 {
				result.WriteString(", ")
			}
			result.WriteString(pair)
		}
	}
	return result.String()
}