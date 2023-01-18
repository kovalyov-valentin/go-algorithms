package main 

import (
	"fmt"
)


func main() {
	arr := []int{5, 1, 3, 7, 2, 4, 0, 9, 8, 6, 10}
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
		fmt.Println(arr)
	}

}