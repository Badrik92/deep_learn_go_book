package main

import(
	"fmt"
	"sort"
)
// Создаем структуру
type aStructure struct {
	person string
	heigth int
	weight int
}

func main() {
	// Создаем новый срез
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStructure{"Athina", 134, 40})
	fmt.Println("0:", mySlice)

	// Сортируем слайс с помощью анонимной ф-ии 
	sort.Slice(mySlice, func(i, j int) bool {
		// Сортируем по росту снизу вверх
		return mySlice[i].heigth < mySlice[j].heigth
	})
	fmt.Println("<:", mySlice)

	// Сортируем слайс с помощью анонимной ф-ии
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].heigth > mySlice[j].heigth
	})

	fmt.Println(">:", mySlice)
	
}