package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()
	for v := range jobs {
		fmt.Printf("Воркер %d выполнил свою тяжелую работу %d\n", id, v)
	}
}

func main() {
	jobs := make(chan int)
	wg := &sync.WaitGroup{}

	workers := flag.Int("workers", 5, "Количество воркеров для выполнения задачи")
	flag.Parse()

	for i := 1; i <= *workers; i++ {
		wg.Add(1)
		go worker(i, wg, jobs)
	}

	for {
		jobs <- rand.Int()
		time.Sleep(50 * time.Millisecond)
	}

}
