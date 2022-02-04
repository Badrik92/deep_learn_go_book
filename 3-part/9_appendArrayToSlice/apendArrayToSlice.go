package main

import (
	"fmt"
)

func main() {
	// Инициализация слайса
	s := []int{1,2,3}
	// Инициализация массива
	a := [3]int{4,5,6}

	// СОздаем ссылку на существующий массив
	ref := a[:]
	fmt.Println("Existing array:\t", ref) // [4 5 6]
	// Объединяем в cлайс t массив ref с помощью ... (троеточие)
	t := append(s, ref...)
	fmt.Println("New slice:\t", t) // [1 2 3 4 5 6]
	// Переопределяем слайс s объединяя массив ref и слайс s 
	s = append(s, ref...)
	fmt.Println("Existing slice:\t", s)
	// Обединяем элементы нового среза s
	s = append(s, s...)
	fmt.Println("s + s: \t\t", s) // [1 2 3 4 5 6 1 2 3 4 5 6]

}