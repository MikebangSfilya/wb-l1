package main

import "fmt"

func binarySerach(nums []int, target int) int {

	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i + 1
	}
	target := 732
	fmt.Println(binarySerach(nums, target))
}
