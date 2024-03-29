package main 

import (
	"fmt"
)
// ОБЪЯСНЕНИЕ:
/*
На каждой итерации внешнего цикла мы перебираем все элементы в слайсе, начинаная с первого и заканчивая n-ым элементом.
Однако на каждой итерации внутреннего цикла мы не рассматриваем последний элемент, поскольку уже знаем, 
что он находится в правильном порядке.
Если бы не использовали -1 в выражении n-1-i, то в последней итерации внешнего цикла внутренний цикл бы продолжался до n-го элемента,
включая последний элемент, который мы уже знаем, что находится в правильном порядке.
Это было бы неэффективным, потому что мы тратили бы время на проверку элемента, который уже отсортирован.
Потому, чтобы увеличить эффективность алгоритма, мы ограничиваем диапозон элементов, который рассматриваем на каждой итерации
внутреннего цикла, от 0 до n-1-i. 
Таким образом, на последней итерации внешнего цикла мы проверяем только первый элемент, который еще не отсортирован, и слайс 
уже отсортирован.
*/
func bubbleSort(slice []int) {
	// Длина нашего слайса
	n := len(slice)
	// Внешний цикл проходит от 0 до n-1
	for i := 0; i < n - 1; i++ {
		// Внутренний цикл проходит от 0 до n-1-i
		for j := 0; j < n - 1 - i; j++ {
			// Если текущий элемент больше следующего, то меняем их местами
			if slice[j] > slice[j + 1] {
				slice[j], slice[j + 1] = slice[j + 1], slice[j]
			}
		}
	} 

}

func main() {
	slice := []int{5, 1, 3, 7, 2, 4, 0, 9, 8, 6}
	bubbleSort(slice)
	fmt.Println(slice)
}




func main_() {
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