package main

import (
	"io" // Пакет для доступа к ф-ям ввода-вывода
	
	"os" // Пакет для чтения программой аргументов командной строки
)

func main() {
	myString := "" // Инициализируем пустую переменную
	arguments := os.Args // пер-я принимает на вход аргументы командной строки
	if len(arguments) == 1 { // Проверяем, переданы ли аргументы
		// Если длина Args 1 (т.е 0 аргументов передано) выводим сообщение
		myString = "Please give me one argument!"
	} else {
		// если передано 1 и больше аргументов
		myString = arguments[1]
	}
	// Выводим аргументы
	// То же самое, что и Print, только имеется файл, 
	// в который будут записаны аргументы
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}