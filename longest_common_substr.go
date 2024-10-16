package sprint

func LongestCommonSubstr(str1, str2 string) string {
    m := len(str1)
    n := len(str2)
    maxLength := 0
    endingIndex := 0

    lcsuff := make([][]int, m+1)
    for i := range lcsuff {
        lcsuff[i] = make([]int, n+1)
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if str1[i-1] == str2[j-1] {
                lcsuff[i][j] = lcsuff[i-1][j-1] + 1
                if lcsuff[i][j] > maxLength {
                    maxLength = lcsuff[i][j]
                    endingIndex = i
                }
            }
        }
    }

    if maxLength == 0 {
        return ""
    }

    return str1[endingIndex-maxLength : endingIndex]
}