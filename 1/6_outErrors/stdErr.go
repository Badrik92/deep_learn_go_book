package main

import (
	"io"
	"os"
)


func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument"
	} else {
		myString = arguments[1]
	}

	// Вызываем ф-ю 2 раза io.WriteString для записи в стандартный поток ошибок(os.Stderr)
	// и один раз для записи в стандартный поток вывода (os.Stdout)
	io.WriteString(os.Stdout, "This is Standart output\n")
	io.WriteString(os.Stderr, myString) 
	io.WriteString(os.Stderr, "\n")
}