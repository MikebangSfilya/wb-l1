package main

import (
	"fmt"
)

func write(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range in {
			out <- v * 2
		}
	}()

	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	ch := square(write(nums))
	for v := range ch {
		fmt.Println(v)
	}

}
