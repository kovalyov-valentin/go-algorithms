package main 

import (
	"fmt"
)

func main() {
	arrays := []int{5, 1, 3, 7, 2, 4, 0, 9, 8, 6}
	for i := 0; i < len(arrays) - 1; i++ {
		for j := 0; j <  len(arrays) - 1 - i; j++ {
			if arrays[j] > arrays[j + 1] {
				tmp := arrays[j]
				arrays[j] = arrays[j + 1]
				arrays[j + 1] = tmp
			}
		}
	}
	fmt.Println(arrays)
}