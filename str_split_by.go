package sprint

func StrSplitBy(s, sep string) []string {
    result := []string{}
    startIndex := 0
    sepLen := len(sep)

    for {
        endIndex := indexSubString(s[startIndex:], sep)

        if endIndex == -1 {
            break
        }

        result = append(result, s[startIndex:startIndex+endIndex])
        startIndex += endIndex + sepLen
    }

    if startIndex < len(s) {
        result = append(result, s[startIndex:])
    }

    return result
}

func indexSubString(s, sub string) int {
    for i := 0; i <= len(s)-len(sub); i++ {
        if s[i:i+len(sub)] == sub {
            return i
        }
    }
    return -1
}