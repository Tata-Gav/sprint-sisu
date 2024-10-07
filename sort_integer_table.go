package sprint

func SortIntegerTable(table []int) []int {
    // Create a copy of the original table to avoid modifying it in place
    sortedTable := make([]int, len(table))
    copy(sortedTable, table)

    // Sort the copied table using the built-in sort function
    sort.Ints(sortedTable)

    return sortedTable
}