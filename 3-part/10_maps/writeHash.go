package main

import "fmt"

func main(){
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println(aMap) // map[test:1]

	// aMap2 := map[string]int{}
	aMap = nil
	aMap["test2"] = 2
}