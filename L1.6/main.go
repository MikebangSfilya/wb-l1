package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	exitByAtomic()
	exitByDoneChan()
	closeByContext()
	exitByStopChan()
	exitByTimeout()
	exitByGoexit()
}

// Атомарный флаг
func exitByAtomic() {
	var stop atomic.Bool
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		for !stop.Load() {
			fmt.Println("реально тяжелые вычисления c atomic bool: ", rand.Intn(13)*rand.Intn(37))
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("👺👺👺👺 end work 👺👺👺👺")
	}()

	time.Sleep(1 * time.Second)
	stop.Store(true)
	wg.Wait()
	fmt.Println("Горутина с atomic bool отменена")
}

// Канал заврешения
func exitByDoneChan() {
	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {

			select {
			case <-done:
				fmt.Println("👺👺👺👺 end work 👺👺👺👺")
				return
			default:
				fmt.Println("реально тяжелые вычисления у горутины с каналом завершения: ", rand.Intn(100)*rand.Intn(100))
				time.Sleep(300 * time.Millisecond)
			}
		}

	}()

	time.Sleep(1 * time.Second)
	done <- struct{}{}
	wg.Wait()
	fmt.Println("Горутина с done chan отменена")
}

// Отмена через контекст
func closeByContext() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {

			select {
			case <-ctx.Done():
				fmt.Println("👺👺👺👺 end work 👺👺👺👺")
				return
			default:
				fmt.Println("реально тяжелые вычисления у горутины что отменяется контекстом: ", rand.Intn(100)*rand.Intn(100))
				time.Sleep(300 * time.Millisecond)
			}
		}

	}()

	time.Sleep(1 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("Горутина c ctx отменена")
}

// Канал отмены
func exitByStopChan() {
	stop := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {

			select {
			case <-stop:
				fmt.Println("👺👺👺👺 end work 👺👺👺👺")
				return
			default:
				fmt.Println("Я еще живу и напоминаю об этом")
				time.Sleep(300 * time.Millisecond)
			}
		}

	}()

	time.Sleep(1 * time.Second)
	close(stop)
	wg.Wait()
	fmt.Println("Горутина с stop chan отменена")
}

// Goexit
func exitByGoexit() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Горутина готовится снова выполнить свои тяжелые действия")
		runtime.Goexit()
	}()

	wg.Wait()
	fmt.Println("Горутина с runtime.Goexit отменена👺👺👺👺")
}

// Таймаут
func exitByTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("👺👺👺👺 таймаут 👺👺👺👺")
				return
			default:
				fmt.Println("Горутина делает очень важную работу")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Горутина с таймаут контекстом отменена")
}
