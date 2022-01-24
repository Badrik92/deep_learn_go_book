package main

import(
	"fmt"
	// Добавляем пакет, что бы была возможность 
	// преобразовывать аргументы командной строки
	"strconv" 
	"os"
)
func main() {
	if len(os.Args) == 1 { // ППроверка
		fmt.Println("Please give one or more floats")
		os.Exit(1) // Выходим с ошибкой
		// Нулевой статус Exit  о
		// Exit немедленно завершает работу, даже другие ф-ии не выполняются
	}

	arguments := os.Args
	// strconv.ParseFloat возвращает аргумент командной строки и тип error
	// в данном примере мы игнорируем значение error нижним подчеркиванием
	// т.е пустым идентификатором
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	// Цикл for позволяет перебрать все элементы среза os.Args
	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		
		if n < min { // если аргумент меньше
			min = n	 // в n кладется этот аргумент	
		}
		if n > max {
			max = n
		}

	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}