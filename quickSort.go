// package main 

// import (
// 	"fmt"
// )

// func quickSort(lst []int, left, right int) []int {

// 	// Создаем копии пришедших переменных, с которым будет в дальнейшем работать 
// 	l := left 
// 	r := right 

// 	// Вычисляем ключевой, центральный элемент, на который будем опираться
// 	// Берем значение центральной ячейки массива 

// 	center := lst[(left + right) / 2]

// 	fmt.Println(l, r, (left + right) / 2)

// 	// Цикл, начинающий сортировку
// 	for l <= r {

// 		// Ищем значения больше центра
// 		for lst[r] > center {
// 			r--
// 		}

// 		// Ищем значения меньше центра
// 		for lst[l] < center {
// 			l++
// 		}

// 		// После прохода циклов проверяем счетчики циклов
// 		if l <= r {

// 			// И если условие true, то меняем ячейки друг с другом
// 			lst[r], lst[l] = lst[l], lst[r]
// 			l++ 
// 			r--
// 		}
// 	}

// 	if r > left {
// 		quickSort(lst, left, r)
// 	}

// 	if l < right {
// 		quickSort(lst, l, right)
// 	}

// 	return lst 
// }

// func main() {
// 	lst := []int{5, 4, 1, 2, 0, 123, 1234, 32, 12, 2345, 99}
// 	lst = quickSort(lst, 0, len(lst) - 1)
// 	fmt.Println(lst)
// }


package main 

import (
	"fmt"
)

func Sort(input []int) []int {
    l := len(input)
    if l < 2 {
        return input
    }
    less := make([]int, 0)
    bigger := make([]int, 0)
    pivot := input[0]
    for _, v := range input[1:] {
        if v > pivot {
            bigger = append(bigger, v)
        } else {
            less = append(less, v)
        }
    }
    input = append(Sort(less), pivot)
    input = append(input, Sort(bigger)...)
    return input
}

func main() {
	array := []int{5, 4, 1, 2, 0, 123, 1234, 32, 12, 2345, 99}
	fmt.Println(Sort(array))
}

