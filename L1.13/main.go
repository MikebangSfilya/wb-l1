package main

import "fmt"

func main() {
	a, b := 1, 2
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("math:", a, b)

	c, d := 5, 6
	c, d = d, c
	fmt.Println("XOR:", c, d)
}
