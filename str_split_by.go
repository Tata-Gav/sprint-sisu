package sprint

func StrSplitBy(s, sep string) []string {
    result := []string{}
    startIndex := 0
    endIndex := strings.Index(s[startIndex:], sep)

    for endIndex != -1 {
        result = append(result, s[startIndex:startIndex+endIndex])
        startIndex = startIndex + endIndex + len(sep)
        endIndex = strings.Index(s[startIndex:], sep)
    }

    if startIndex < len(s) {
        result = append(result, s[startIndex:])
    }

    return result
}