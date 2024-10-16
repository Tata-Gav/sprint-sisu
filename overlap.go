package sprint

func Overlap(arr1, arr2 []int) []int {
    elementCount := make(map[int]int)
    for _, num := range arr1 {
        elementCount[num]++
    }

    commonElements := []int{}
    for _, num := range arr2 {
        if elementCount[num] > 0 {
            commonElements = append(commonElements, num)
            elementCount[num]--
        }
    }

    for i := 0; i < len(commonElements)-1; i++ {
        for j := 0; j < len(commonElements)-1-i; j++ {
            if commonElements[j] > commonElements[j+1] {
                commonElements[j], commonElements[j+1] = commonElements[j+1], commonElements[j]
            }
        }
    }
	 return commonElements
}