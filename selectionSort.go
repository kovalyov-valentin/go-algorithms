// package main 

// import (
// 	"fmt"
// )

// func main() {
// 	arrays := []int{5, 1, 3, 7, 2, 4, 0, 9, 8, 6}
// 	for i := 0; i < len(arrays)-1; i++ {
// 		min := i 
// 		for j := i + 1; j < len(arrays); j++ {
// 			if arrays[j] < arrays[min] {
// 				arrays[j], arrays[min] = arrays[min], arrays[j]
// 			}
// 		}
// 	}
// 	fmt.Println(arrays)
// }

package main 

import (
	"fmt"
)

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				arr[j], arr[min] = arr[min], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
func main() {
	arrays := []int{5, 1, 3, 7, 2, 4, 0, 9, 8, 6}
	selectionSort(arrays)
}


