package main

import (
	"fmt"
)

func main() {
	// Инициализируем мапу с помощью make
	iMap := make(map[string]int)
	// Передаем ключ и значение
	iMap["k1"] = 12
	iMap["k2"] = 13
	fmt.Println("iMap:", iMap) // iMap: map[k1:12 k2:13]

	// Инициализируем мапу с помощью литерала хеш-таблицы
	anotherMap := map[string]int {
		"k1": 12,
		"k2": 13,
	}

	fmt.Println("anotherMap:", anotherMap ) // anotherMap: map[k1:12 k2:13]
	
	// Удаление элемента из хеш-таблицы
	// Многократный вызов одного и того же оператора delete ни на что не влияет
	// и не генерирует предупреждающих сообщений об ошибке
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	fmt.Println("anotherMap:", anotherMap) // anotherMap: map[k2:13]

	// Узнаем, есть ли в таблице данный ключ
	// Пользуемся записью _, ok
	_, ok := iMap["doesItExist"]
	if ok {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Does NOT exist")
	}

	//  Выводим элементы с мопощью цикла
	for key, value := range iMap {
		fmt.Println(key, value)
	}

}