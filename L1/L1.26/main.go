package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	str := strings.ToLower(s)
	m := make(map[rune]struct{})
	for _, v := range str {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abcA"))
}
