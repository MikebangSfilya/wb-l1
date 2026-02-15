package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(i int) string {
	return strings.Repeat("large", i)
}

func someFunc() {
	//Срез justString ссылается на массив строки v, блокируя его очистку сборщиком мусора
	v := createHugeString(1 << 10)
	//Создание копии строки v. Это гарантирует создание копии строки в новом выделенном пространстве.
	justString = strings.Clone(v[:100])
	fmt.Println(justString)
}

func main() {
	someFunc()
}
