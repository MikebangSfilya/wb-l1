package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("10000000000000000000000000", 10)
	b := new(big.Int)
	b.SetString("5000000000000000000000000", 10)
	mul := new(big.Int).Mul(a, b)
	fmt.Println("Умножение: ", mul)
	if b.Sign() == 0 {
		fmt.Println("деление на ноль невозможно.")
	} else {
		div := new(big.Int).Div(a, b)
		fmt.Println("Деление: ", div)
	}
	plus := new(big.Int).Add(a, b)
	fmt.Println("Сложение: ", plus)
	sub := new(big.Int).Sub(a, b)
	fmt.Println("Вычитание :", sub)

}
