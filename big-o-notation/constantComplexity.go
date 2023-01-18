// O(1) — константная сложность

package main 

import (
	"fmt"
)

func GetCount(items []int) int {
	return len(items)
}

func main() {
	i := []int{1,2,3,4,5,6,7,8,9,0,10,11}
	fmt.Println(GetCount(i))
}