package sprint

func StrCompress(input string) string {
    var compressed string
    count := 1

    for i := 1; i < len(input); i++ {
        if input[i] == input[i-1] {
            count++
        } else {
            if count > 1 {
                compressed += string(count) + string(input[i-1])
            } else {
                compressed += string(input[i-1])
            }
            count = 1
        }
    }

    if count > 1 {
        compressed += string(count) + string(input[len(input)-1])
    } else {
        compressed += string(input[len(input)-1])
    }

    return compressed
}