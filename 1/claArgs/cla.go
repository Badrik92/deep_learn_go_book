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
}