package main

import "fmt"

func main() {
	aSlice := []int{1,2,3,4,5}
	fmt.Println(aSlice)
	integers := make([]int, 2)
	integers = nil
	fmt.Println(integers) // []

	anArray := [5]int{-1, -2, -3, -4, -5}
	// Создаем новый срез с помощью нотации [:]
	// Помним, что мы создаем не копию массива, а лишь ссылку на него
	refAnArray := anArray[:]

	fmt.Println(anArray) // [-1 -2 -3 -4 -5]
	fmt.Println(refAnArray) // [-1 -2 -3 -4 -5]
	anArray[4] = -100
	fmt.Println(refAnArray) // [-1 -2 -3 -4 -100]

	// Создаем два среза, одномерный и двумерный
	s := make([]byte, 5)
	fmt.Println(s) // [0 0 0 0 0]
	twoD := make([][]int, 3)
	// Элементы многомерного среза тоже являются срезами
	fmt.Println(twoD) // [[] [] []]
	fmt.Println()
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 2; j++ {
			// Для добавления элементов, нужно явно использвать ф-ю append()
			twoD[i] = append(twoD[i], i*j)
		}
	}
	fmt.Println(twoD) // [[0 0] [0 1] [0 2]]

	// Использвание range для вывода на экран всех элементов среза
	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i:", i, "value", y)
			// i: 0 value 0
			// i: 1 value 0
			// i: 0 value 0
			// i: 1 value 1
			// i: 0 value 0
			// i: 1 value 2
		}
		fmt.Println()
	}
}