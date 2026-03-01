package main

import "fmt"

func setBit(num int64, pos uint, bit int) int64 {
	mask := int64(1) << pos
	if bit == 1 {
		return num | mask
	} else {
		return num &^ mask
	}
}

func main() {
	//4
	fmt.Println(setBit(5, 0, 0))
	//1
	fmt.Println(setBit(5, 2, 0))
	//7
	fmt.Println(setBit(5, 1, 1))
}
