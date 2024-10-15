package main

import "strings"

func StrCompare(a, b string) int {
    return strings.Compare(a, b)
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
    // Implementation using bubble sort for simplicity
    for i := 0; i < len(a)-1; i++ {
        for j := i + 1; j < len(a); j++ {
            if f(a[i], a[j]) > 0 {
                a[i], a[j] = a[j], a[i]
            }
        }
    }
    return a
}