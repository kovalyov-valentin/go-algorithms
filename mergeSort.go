package main 

import (
	"fmt"
)

func merge(left []int, right []int) []int {

	// Merge two lists in ascending order.
	lst := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			lst = append(lst, left[0])
			left = left[1:]
		} else {
			lst = append(lst, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		lst = append(lst, left...)
	}

	if len(right) > 0 {
		lst = append(lst, right...)
	}

	return lst 
}

func mergeSort(lst []int) []int {

	// Sort the list by merging O(n * log n).
	lenght := len(lst)
	if lenght >= 2 {
		mid := lenght / 2
		lst = merge(mergeSort(lst[:mid]), mergeSort(lst[mid:]))
	}
	return lst 
}

func main() {
	lst := []int{2, 123, 443, 223, 3, 5, 6, 432, 1}
	lst = mergeSort(lst)
	fmt.Println(lst)
}