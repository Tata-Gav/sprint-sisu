package sprint

import (
    "sort" // Make sure to import the sort package
)
func SortWordArr(a []string) []string {
	sort.Strings(a)
	return a
}