package main

import "fmt"

func quickSort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	pivotIdx := len(nums) / 2
	pivot := nums[pivotIdx]
	low := []int{}
	middle := []int{}
	high := []int{}

	for i := range nums {
		if nums[i] > pivot {
			high = append(high, nums[i])
		} else if nums[i] < pivot {
			low = append(low, nums[i])
		} else {
			middle = append(middle, nums[i])
		}
	}
	return append(append(quickSort(low), middle...), quickSort(high)...)
}

func main() {
	nums := []int{3, 52, -255, 322, 12, 137, 1, 1337, 2, -1, 0, 2}
	fmt.Println(quickSort(nums))
}
