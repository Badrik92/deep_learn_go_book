package main

import(
	"fmt"
	// Загрузка пакета с помощбю команды go get -v
	// Временные файлы загруженного пакета можно удалить
	// go clean -i -v -x  fullFIleName
	// rm -rf filename  удаления всего пакета
	"github.com/mactsouk/go/simpleGitHub"
)

func main() {
	fmt.Println(simpleGitHub.AddTwo(5,6))
}