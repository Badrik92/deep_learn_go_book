package main

import "fmt"

func main() {
	// Демонстрация работы цикла for
	for i := 0; i < 100; i++ {
		if i % 20 == 0 {
			continue
		}
		if i == 95 {
			break 
		}
		fmt.Print(i, " ")
	}

	fmt.Println()
	i := 10
	// for работает как while
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()

	// for работает как do...while
	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}
	fmt.Print(i, " ")
	i++
	}
	fmt.Println()

	// Использование range
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		// range возвращает индекс и значение
		fmt.Println("Index: ", i, "value:", value)
	}
}
