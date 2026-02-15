package main

import (
	"fmt"
)

func identifyType(v any) {
	switch v.(type) {
	case int:
		fmt.Println("тип переменной int")
	case string:
		fmt.Println("иип переменной string")
	case bool:
		fmt.Println("тип переенной bool")
	case chan int:
		fmt.Println("тип переменной chan int")
	case chan string:
		fmt.Println("тип переменной chan string")
	default:
		fmt.Println("тип переменной неизвестен")

	}
}

func main() {
	identifyType(5)
	identifyType("sss")
	identifyType(make(chan int))
	identifyType(true)
	identifyType(make(chan string))
	identifyType(3.15)

}
