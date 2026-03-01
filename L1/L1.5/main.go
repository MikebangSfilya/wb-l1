package main

import (
	"fmt"
	"math/rand"
	"time"
)

const baseTime = 10 * time.Second

func main() {

	timer := time.After(baseTime)
	pipe := make(chan int)

	go func() {
		for v := range pipe {
			if v == 322 {
				fmt.Println("деньги не проблема", v)
			} else {
				fmt.Println("Канал доставил новое число: ", v)
			}

		}
	}()

	for {
		select {
		case <-timer:
			close(pipe)
			fmt.Println("А таймер то кончился")
			return
		case pipe <- rand.Intn(323):
			fmt.Println("Число отправилось в канал")
			time.Sleep(322 * time.Millisecond)
		}
	}

}
