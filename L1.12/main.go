package main

import "fmt"

func set(str []string) []string {
	res := []string{}
	set := make(map[string]struct{})
	for _, v := range str {
		if _, ok := set[v]; ok {
			continue
		}
		res = append(res, v)
		set[v] = struct{}{}
	}
	return res
}

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(set(str))
}
