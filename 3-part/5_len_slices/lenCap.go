package main

import "fmt"

// Ф-я, помогает вывести на экран одномерный срез
// без необходимости постоянно повторять один и тот же код
func printSlice(x []int) { // принимает слайс в качестве аргумента
	for _, number := range x {
		fmt.Println(number, " ")
	}
	fmt.Println()
}

func main() {
	// Инициализируем слайс
	aSlice := []int{-1, 0, 4}
	fmt.Println("aSlice: ", aSlice)
	// Передаем в ф-ю наш слайс
	printSlice(aSlice)
	// Выводим длину и емкость слайса
	fmt.Printf("Cap: %d, length: %d\n", cap(aSlice), len(aSlice))
	// Добавим в слайс элементы
	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, length: %d\n", cap(aSlice), len(aSlice))

	// Добавим в слайс элементы
	aSlice = append(aSlice, -2)
	aSlice = append(aSlice, -3)
	aSlice = append(aSlice, -4)
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

}