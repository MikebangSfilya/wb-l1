package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 32, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target := 6
	copy(nums[target:], nums[target+1:])
	nums = nums[:len(nums)-1]
	fmt.Println(nums)
}
