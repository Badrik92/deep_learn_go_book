// Разбираем работу ф-ии copy()
package main

import "fmt"

func main() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	// копируем a6 и a4
	copy(a6, a4)
	fmt.Println("a6:", a6) // [-1 -2 -3 -4 4 5]
	fmt.Println("a4:", a4) // [-1 -2 -3 -4]
	fmt.Println()

	// Вторая часть кода
	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	copy(b4, b6)
	fmt.Println("b6:", b6) // [-10 1 2 3 4 5]
	fmt.Println("b4:", b4) // [-10 1 2 3]

	// Третий фрагмент кода
	fmt.Println()
	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, 1, -1, -1, 5, -5}
	// Преобразовываем массив в срез с помощью нотации [:]
	copy(s6, array4[0:])
	fmt.Println("array4:", array4[0:]) // [4 -4 4 -4]
	fmt.Println("s6:", s6)	// [4 -4 4 -4 5 -5]
	fmt.Println()

	// Последняя часть кода
	array5 := [5]int{5, -5, 5, -5, 5}
	s7 := []int{7,7, -7, -7, 7, -7, 7}
	// Преобразовываем массив в срез
	// Если не преобразовать массив в срез, то будет ошибка
	copy(array5[0:], s7)
	fmt.Println("array5:", array5) // [7 7 -7 -7 7]
	fmt.Println("s7:", s7)		// [7 7 -7 -7 7 -7 7]
	



}