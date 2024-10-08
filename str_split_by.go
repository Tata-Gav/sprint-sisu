package sprint

import "strings"

func StrSplitBy(s, sep string) []string {

    result := []string{}
    startIndex := 0

    for {
        endIndex := strings.IndexByte(s[startIndex:], sep[0])
        if endIndex == -1 {
            break
        }

        if match := s[startIndex : startIndex+endIndex+1]; match == sep {
            result = append(result, s[startIndex:startIndex+endIndex])
            startIndex += endIndex + len(sep)
        } else {
            startIndex += endIndex + 1
        }
    }

    if startIndex < len(s) {
        result = append(result, s[startIndex:])
    }

    return result
}