package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	if duration <= 0 {
		return
	}
	<-time.After(duration)
}

func main() {
	fmt.Println("Засыпаем")
	sleep(2 * time.Second)
	fmt.Println("Просыпаемся")
}
