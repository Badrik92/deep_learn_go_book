// Комплексные числа
package main

import "fmt"

func main() {
	// Первый способ создания комплексного числа - напрямую
	c1 := 12 + 1i
	c2 := complex(5, 7)

	// Выводим типы с1 и с2
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of c2: %T\n", c2)
}