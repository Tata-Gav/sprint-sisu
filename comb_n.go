package sprint

import (
	"fmt"
)

func CombN(n int) []string {
	var results []string
	combination := make([]int, n)

	var backtrack func(start int, depth int)
	backtrack = func (start int, depth int)  {
		if depth  == n {
			str := ""
			for _, num := range combination {
				str += fmt.Sprint(num)
			}
			results = append(results, str)
			return
			}

		for i := start; i <= 9; i++ {
			combination[depth] = i
			backtrack(i+1, depth+1)
			}
		backtrack(0, 0)
		return results
		}
		func main()  {
			combination := CombN(3)
			fmt.Println(combinations)	
		}
	}
