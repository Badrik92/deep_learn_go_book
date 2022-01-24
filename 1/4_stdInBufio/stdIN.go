package main

import (
	"os"
	"bufio" // Для ввода и вывода файлов
	"fmt"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	// bufio.NewScanner передается в стандартный поток ввода os.Stdin
	scanner := bufio.NewScanner(f)
	// bufio.NewScanner возвращает переменную bufio.Scanner
	// которая пользуется ф-ей Scan() для построчного чтения из os.Stdin
	// каждая прочитанная строка выводится на экран, после чего считывается следующая строка
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}