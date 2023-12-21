package main

import "fmt"

func insertionSortOne(arr []int) []int {
	// from left to right
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		lst := i
		// from right to left
		for j := i - 1; j > -1; j-- {
			if arr[j] < key {
				break
			}
			arr[j+1] = arr[j]
			lst = j
		}
		arr[lst] = key

	}
	return arr
}

func insertionSortTwo(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	return arr
}

func main() {
	arr := []int{5, 1, 3, 7, 2, 4, 0, 9, 8, 6, 10}

	fmt.Println("Insertion Sort №1:", insertionSortOne(arr))
	fmt.Println("Insertion Sort №2:", insertionSortTwo(arr))


}
