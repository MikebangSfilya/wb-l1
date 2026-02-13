package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return
		case v, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("Воркер %d выполнил свою тяжелую работу %d\n", id, v)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	jobs := make(chan int)
	wg := &sync.WaitGroup{}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	workers := flag.Int("workers", 5, "Количество воркеров для выполнения задачи")
	flag.Parse()

	for i := 1; i <= *workers; i++ {
		wg.Add(1)
		go worker(ctx, i, wg, jobs)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("writer stopped by context\n")
				return
			case jobs <- rand.Int():
				time.Sleep(50 * time.Millisecond)
			}

		}
	}()

	wg.Wait()
	cancel()
	fmt.Printf("Stopped gracefully\n")

}
