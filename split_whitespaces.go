package sprint

func SplitWhitespaces(s string) []string {
    words := []string{}
    word := ""

    for _, char := range s {
        if unicode.IsSpace(char) || char == '\t' || char == '\n' {
            if word != "" {
                words = append(words, word)
                word = ""
            }
        } else {
            word += string(char)
        }
    }

    if word != "" {
        words = append(words, word)
    }

    return words
}