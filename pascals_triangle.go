package sprint

func PascalsTriangle(n int) [][]int {
    if n <= 0 {
        return [][]int{}
    }
    
    triangle := make([][]int, n)
    
    for i := 0; i < n; i++ {
        // Initialize the row
        triangle[i] = make([]int, i+1)
        triangle[i][0], triangle[i][i] = 1, 1
        
        // Fill in the interior values
        for j := 1; j < i; j++ {
            triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
        }
    }
    return triangle
}