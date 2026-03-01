package main

import "fmt"

func reverseRune(s []rune, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reversePhrase(s string) string {
	runes := []rune(s)
	reverseRune(runes, 0, len(runes)-1)

	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverseRune(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func main() {
	fmt.Println(reversePhrase("sun dog snow"))

}
