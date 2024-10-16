package sprint

import (
    "fmt"
    "strconv"
)

func StrCompress(input string) string {
    if len(input) == 0 {
        return ""
    }

    var result string
    count := 1

    for i := 1; i < len(input); i++ {
        if input[i] == input[i-1] {
            count++
        } else {
            if count > 1 {
                result += strconv.Itoa(count)
            }
            result += string(input[i-1])
            count = 1
        }
    }

    if count > 1 {
        result += strconv.Itoa(count)
    }
    result += string(input[len(input)-1])
	return result
}