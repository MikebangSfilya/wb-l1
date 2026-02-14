package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]struct{})
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			m[i] = struct{}{}

		}()
	}
	wg.Wait()
	fmt.Println(len(m))
}
