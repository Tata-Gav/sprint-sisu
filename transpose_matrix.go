package sprint

func TransposeMatrix(matrix [][]int) [][]int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return [][]int{}
    }
    rows := len(matrix)
    cols := len(matrix[0])
    transposed := make([][]int, cols)
    
    for i := range transposed {
        transposed[i] = make([]int, rows)
        for j := range matrix {
            transposed[i][j] = matrix[j][i]
        }
    }
    return transposed
}