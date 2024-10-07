package sprint
 
func SortIntegerTable(table []int) []int {
    
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] >= table[j] {

				table[i], table[j] = table[j], table[i]
			}
		}
	}
    return table
}