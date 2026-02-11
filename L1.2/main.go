package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for i := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nums[i] *= nums[i]
		}()
	}
	wg.Wait()
	fmt.Println("nums:", nums)
}
