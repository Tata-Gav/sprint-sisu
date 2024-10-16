package sprint

import (
    "sort" // Make sure to import the sort package
)
func SortWordArr(a []string) []string {
	sort.String(a)
	return a
}