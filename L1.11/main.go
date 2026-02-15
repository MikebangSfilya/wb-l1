package main

import "fmt"

func set(nums1, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	set := make(map[int]struct{})
	for _, v := range nums1 {
		set[v] = struct{}{}
	}
	res := []int{}
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}

func main() {
	a, b := []int{1, 2, 3}, []int{2, 3, 4}
	fmt.Println(set(a, b))
}
