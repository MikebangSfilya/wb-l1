package main

import "fmt"

func reverseString(s string) string {
	srune := []rune(s)

	for i, j := 0, len(srune)-1; i < j; i, j = i+1, j-1 {
		srune[i], srune[j] = srune[j], srune[i]
	}
	return string(srune)
}

func main() {
	fmt.Println(reverseString("главрыба"))
	fmt.Println("go")
}
