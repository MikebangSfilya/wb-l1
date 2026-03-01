package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	sync.Mutex
	cnt int64
}

func (c *counter) incAtomic() {
	atomic.AddInt64(&c.cnt, 1)
}

func (c *counter) incMutex() {
	c.Lock()
	defer c.Unlock()
	c.cnt++
}

func (c *counter) res() int64 {
	return atomic.LoadInt64(&c.cnt)
}

func main() {

	c := counter{}
	wg := sync.WaitGroup{}

	numGoroutines := 100
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			c.incAtomic()
		}()

	}

	wg.Wait()
	fmt.Println("С помощью атомиков получилось: ", c.res())

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {

		go func() {
			defer wg.Done()
			c.incMutex()
		}()
	}

	wg.Wait()
	fmt.Println("С помощью мьютекса получилось: ", c.cnt)

}
